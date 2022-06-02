package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"tie-explorer-go/models"
	"tie-explorer-go/task"
	"time"
)

type BridgeController struct {
	BaseController
}

func (b BridgeController) ChainHealth() {
	type request struct {
		ChainIds []int64
	}

	var data request
	jsonBytes := b.Ctx.Input.RequestBody
	json.Unmarshal(jsonBytes, &data)

	result := make(map[int64]bool)
	for _, id := range data.ChainIds {
		result[id] = true
	}

	dataRes := make(map[string]interface{})
	dataRes["Result"] = result

	b.Data["json"] = dataRes

	b.ServeJSON()
}

func (b BridgeController) TokenBasics() {
	var assets []*models.BridgeAsset
	o := orm.NewOrm()
	o.QueryTable("bridge_asset").Filter("state", true).All(&assets)

	for i, item := range assets {
		var tokens []models.BridgeToken
		o.QueryTable("bridge_token").Filter("token_basic_name__iexact", item.Name).
			Filter("property", 1).Filter("state", true).All(&tokens)
		if len(assets) > 0 {
			item.Tokens = tokens
		} else {
			assets = append(assets[:i], assets[i+1:]...)
		}
	}

	dataRes := make(map[string]interface{})
	dataRes["TokenBasics"] = assets

	b.Data["json"] = dataRes

	b.ServeJSON()
}

func (b BridgeController) TokenMap() {
	type reqTokenMap struct {
		ChainId int
		Hash    string
	}

	var req reqTokenMap

	json.Unmarshal(b.Ctx.Input.RequestBody, &req)

	var token models.BridgeToken

	o := orm.NewOrm()
	o.QueryTable("bridge_token").Filter("chain_id", req.ChainId).
		Filter("hash__iexact", req.Hash).One(&token)

	type tokenMapStruct struct {
		SrcTokenHash string
		SrcToken     models.BridgeToken
		DstTokenHash string
		DstToken     models.BridgeToken
		Property     int
	}
	var tokenMaps []tokenMapStruct
	if token.Id > 0 {
		var tokens []models.BridgeToken

		o.Raw("SELECT * FROM bridge_token WHERE UPPER (token_basic_name::TEXT)=UPPER ('" +
			token.TokenBasicName + "') AND property = 1 AND state = TRUE AND id !=" +
			strconv.FormatInt(token.Id, 10)).QueryRows(&tokens)

		var tokenMap tokenMapStruct
		for _, item := range tokens {
			tokenMap.SrcTokenHash = token.Hash
			tokenMap.SrcToken = token
			tokenMap.DstTokenHash = item.Hash
			tokenMap.DstToken = item
			tokenMap.Property = token.Property
			tokenMaps = append(tokenMaps, tokenMap)
		}
	}

	dataRes := make(map[string]interface{})
	dataRes["TokenMaps"] = tokenMaps

	b.Data["json"] = dataRes

	b.ServeJSON()
}

func (b BridgeController) GetFee() {
	//{SrcChainId: , SwapTokenHash: '', Hash: '', : 97}
	// TODO 暂时无手续费，假数据
	type reqInfo struct {
		SrcChainId    int
		DstChainId    int
		SwapTokenHash string
		Hash          string
	}

	var req reqInfo

	json.Unmarshal(b.Ctx.Input.RequestBody, &req)

	type resInfo struct {
		SrcChainId               int
		Hash                     string
		DstChainId               int
		UsdtAmount               string
		TokenAmount              string
		TokenAmountWithPrecision string
		SwapTokenHash            string
		Balance                  string
		BalanceWithPrecision     string
	}

	b.Data["json"] = resInfo{
		SrcChainId:               req.SrcChainId,
		Hash:                     req.Hash,
		DstChainId:               req.DstChainId,
		UsdtAmount:               "0",
		TokenAmount:              "0",
		TokenAmountWithPrecision: "",
		SwapTokenHash:            req.SwapTokenHash,
		Balance:                  "99999999999999999999999999999999",
		BalanceWithPrecision:     "0",
	}

	b.ServeJSON()
}

func (b BridgeController) ExpectTime() {
	type reqInfo struct {
		SrcChainId int
		DstChainId int
	}

	var req reqInfo

	json.Unmarshal(b.Ctx.Input.RequestBody, &req)

	type resInfo struct {
		SrcChainId int
		DstChainId int
		Time       int
	}

	rand.Seed(time.Now().UnixNano())
	b.Data["json"] = resInfo{
		SrcChainId: req.SrcChainId,
		DstChainId: req.DstChainId,
		Time:       rand.Intn(100) + 100,
	}

	b.ServeJSON()
}

func (b BridgeController) Transaction() {
	type reqInfo struct {
		Hash string
	}

	var req reqInfo

	json.Unmarshal(b.Ctx.Input.RequestBody, &req)

	var historyInfo models.BridgeHistory

	o := orm.NewOrm()
	o.QueryTable("bridge_history").Filter("hash__iexact", req.Hash).One(&historyInfo)

	if historyInfo.Id > 0 {
		var process []*models.BridgeProcess
		o.QueryTable("bridge_process").Filter("history_id", historyInfo.Id).All(&process)
		if process[0].ChainId != historyInfo.FromChainId {
			historyInfo.TransactionState = []*models.BridgeProcess{
				process[1], process[0],
			}
		} else {
			historyInfo.TransactionState = process
		}

		if historyInfo.State != task.Finished && historyInfo.State != task.Failed {
			if historyInfo.TransactionState[0].State == task.Failed {
				historyInfo.State = task.Failed
				UpdateHistoryState(historyInfo.Id, task.Failed)
			} else if historyInfo.TransactionState[0].State == task.Pending {
				historyInfo.State = task.Pending
			} else {
				historyInfo.TransactionState[0].Blocks = 1
				historyInfo.TransactionState[0].NeedBlocks = 1
				historyInfo.State = task.Pending
				if historyInfo.TransactionState[1].State == task.Confirmed {
					historyInfo.State = task.Finished
					historyInfo.TransactionState[1].Blocks = 1
					historyInfo.TransactionState[1].NeedBlocks = 1
					UpdateHistoryState(historyInfo.Id, task.Finished)
				} else if historyInfo.TransactionState[1].State == task.Failed {
					historyInfo.State = task.Failed
					UpdateHistoryState(historyInfo.Id, task.Failed)
				}
			}
		} else {
			historyInfo.TransactionState[0].Blocks = 1
			historyInfo.TransactionState[0].NeedBlocks = 1
			historyInfo.TransactionState[1].Blocks = 1
			historyInfo.TransactionState[1].NeedBlocks = 1
		}
	}

	b.Data["json"] = historyInfo

	b.ServeJSON()
}

func (b BridgeController) Transactions() {
	type reqInfo struct {
		Addresses []string
		PageNo    int64
		PageSize  int64
	}

	var req reqInfo

	json.Unmarshal(b.Ctx.Input.RequestBody, &req)

	type total struct {
		Count int
	}

	var result total

	addrMap := make(map[string]string)
	var history []*models.BridgeHistory
	var address []string
	o := orm.NewOrm()
	if len(req.Addresses) > 0 {
		sql := "SELECT * FROM \"bridge_history\" WHERE ("
		countSql := "SELECT COUNT(\"id\") AS total FROM \"bridge_history\" WHERE ("
		for _, addr := range req.Addresses {
			if addrMap["0x"+addr] == "" {
				addrMap["0x"+addr] = "0x" + addr
				address = append(address, "UPPER(\"from\"::text) = UPPER('0x"+addr+"')")
			}
		}
		sql += strings.Join(address, " OR ") + ")"
		countSql += strings.Join(address, " OR ") + ")"

		if req.PageNo <= 0 {
			req.PageNo = 1
		}

		if req.PageSize <= 0 {
			req.PageSize = 10
		}

		offset := (req.PageNo - 1) * req.PageSize

		sql += " ORDER BY \"created_at\" DESC OFFSET " + strconv.FormatInt(offset, 10) + " LIMIT " +
			strconv.FormatInt(req.PageSize, 10)
		o.Raw(sql).QueryRows(&history)
		o.Raw(countSql).QueryRow(&result)
	}

	for _, item := range history {
		item.Time = item.CreatedAt.Unix()
		var process []*models.BridgeProcess
		o.QueryTable("bridge_process").Filter("history_id", item.Id).All(&process)
		if process[0].ChainId != item.FromChainId {
			item.TransactionState = []*models.BridgeProcess{
				process[1], process[0],
			}
		} else {
			item.TransactionState = process
		}
		if item.State != task.Finished && item.State != task.Failed {
			if item.TransactionState[0].State == task.Failed {
				item.State = task.Failed
				UpdateHistoryState(item.Id, task.Failed)
			} else if item.TransactionState[0].State == task.Pending {
				item.State = task.Pending
			} else {
				item.TransactionState[0].Blocks = 1
				item.TransactionState[0].NeedBlocks = 1
				item.State = task.Pending
				if item.TransactionState[1].State == task.Confirmed {
					item.State = task.Finished
					item.TransactionState[1].Blocks = 1
					item.TransactionState[1].NeedBlocks = 1
					UpdateHistoryState(item.Id, task.Finished)
				} else if item.TransactionState[1].State == task.Failed {
					item.State = task.Failed
					UpdateHistoryState(item.Id, task.Failed)
				}
			}
		} else {
			item.TransactionState[0].Blocks = 1
			item.TransactionState[0].NeedBlocks = 1
			item.TransactionState[1].Blocks = 1
			item.TransactionState[1].NeedBlocks = 1
		}
	}

	dataRes := make(map[string]interface{})
	dataRes["PageSize"] = req.PageSize
	dataRes["PageNo"] = req.PageNo
	dataRes["TotalPage"] = math.Ceil(float64(result.Count) / float64(req.PageSize))
	dataRes["TotalCount"] = result.Count
	dataRes["Transactions"] = history
	b.Data["json"] = dataRes
	b.ServeJSON()
}

func (b BridgeController) SetExchange() {
	type reqInfo struct {
		FromChainId     int
		FromAddress     string
		ToChainId       int
		ToAddress       string
		ToTokenHash     string
		Amount          string
		TransactionHash string
		Status          int64
	}

	var req reqInfo

	json.Unmarshal(b.Ctx.Input.RequestBody, &req)

	history := models.BridgeHistory{
		Hash:           req.TransactionHash,
		From:           req.FromAddress,
		FromChainId:    req.FromChainId,
		To:             req.ToAddress,
		ToChainId:      req.ToChainId,
		TransferAmount: req.Amount,
		State:          task.Pending,
		ContractHash:   req.ToTokenHash,
	}

	o := orm.NewOrm()
	historyId, err := o.Insert(&history)
	if err == nil {
		processStep1 := models.BridgeProcess{
			HistoryId:   int(historyId),
			Hash:        history.Hash,
			ChainId:     history.FromChainId,
			BlockNumber: 0,
			State:       req.Status,
		}

		var fromChainInfo models.BridgeChain
		o.QueryTable("bridge_chain").Filter("chain_id", history.FromChainId).One(&fromChainInfo)

		var toChainInfo models.BridgeChain
		o.QueryTable("bridge_chain").Filter("chain_id", history.ToChainId).One(&toChainInfo)
		blockNumber := task.GetBlockHeight(toChainInfo.Rpc)

		processStep2 := models.BridgeProcess{
			HistoryId:   int(historyId),
			Hash:        "",
			ChainId:     history.ToChainId,
			State:       task.Pending,
			BlockNumber: blockNumber,
		}
		processId1, _ := o.Insert(&processStep1)
		processId2, _ := o.Insert(&processStep2)
		task.BridgeTxMap[processId1] = task.BridgeTxRelation{
			RPC:       fromChainInfo.Rpc,
			Hash:      history.Hash,
			ProcessId: processId1,
		}

		task.BridgeLogMap[processId2] = task.BridgeProcessRelation{
			RPC:          toChainInfo.Rpc,
			Topic:        toChainInfo.Topic,
			ContractHash: history.ContractHash,
			BlockNumber:  blockNumber,
			ToAddress:    req.ToAddress,
			ProcessId:    processId2,
			QueryCount:   0,
		}
	}

	b.Ctx.WriteString("OK")
}

func UpdateHistoryState(id, state int64) {
	o := orm.NewOrm()
	o.Raw("UPDATE bridge_history SET \"state\" = " +
		strconv.FormatInt(state, 10) +
		" WHERE \"id\" = " + strconv.FormatInt(id, 10)).Exec()
}
