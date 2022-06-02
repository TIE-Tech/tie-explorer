package controllers

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"tie-explorer-go/common"
	"tie-explorer-go/models"
	"time"
)

type AddressController struct {
	BaseController
}

func (t AddressController) List() {
	pageSize, _ := t.GetInt64("ps", 20)
	page, _ := t.GetInt64("p", 1)
	if page <= 0 {
		page = 1
	}

	var blocks []models.Block
	o := orm.NewOrm()
	qt := o.QueryTable("block")
	total, _ := qt.Count()
	pagination := common.GetPagination(total, page, pageSize)
	qt.Offset(pagination.Offset).Limit(pageSize).OrderBy("-number").All(&blocks)

	dataRes := make(map[string]interface{})
	dataRes["blocks"] = blocks
	dataRes["pagination"] = pagination

	t.Data["json"] = Response{0, "success", dataRes}

	t.ServeJSON()
}

func (t AddressController) Info() {
	address := t.GetString(":address")

	addressInfo := models.Address{}

	o := orm.NewOrm()
	o.QueryTable("address").Filter("address__iexact", address).One(&addressInfo)

	maxLimit := 25

	var transactionList []*models.Transaction
	transactionMap := make(map[string]uint64)
	cond := orm.NewCondition()
	baseCond := cond.And("from__iexact", address).Or("to__iexact", address)
	conditions := cond.OrCond(baseCond)
	conditions1 := cond.OrCond(baseCond.Or("contract_address__iexact", address))
	// Txs - limit 25
	o.QueryTable("transaction").SetCond(conditions1).OrderBy("-timestamp").Limit(maxLimit).All(&transactionList)
	txnsTotal, _ := o.QueryTable("transaction").SetCond(conditions1).Count()

	for _, item := range transactionList {
		transactionMap[item.Hash] = item.Timestamp
		if strings.ToLower(item.From) == strings.ToLower(item.To) {
			item.Direction = "Self"
		} else if strings.ToLower(item.To) == strings.ToLower(address) {
			item.Direction = "In"
		} else {
			item.Direction = "Out"
		}

		if item.To == "" {
			item.Genre = "<span class='badge bg-cyan'>Contract Creation</span>"
			if item.ContractAddress != "" {
				item.To = item.ContractAddress
			}
		} else {
			var addressModel models.Address
			o.QueryTable("address").Filter("address__iexact", item.To).One(&addressModel)
			if addressModel.IsContract {
				item.Genre = "<span class='badge bg-purple'>Contract Call</span>"
			} else {
				item.Genre = "<span class='badge bg-blue'>Transaction</span>"
			}
		}
	}

	contractMap := make(map[string]string)
	var erc20TransferList []*models.TokenTransfer
	// Token transfer erc-20
	o.QueryTable("token_transfer").SetCond(conditions.And("token_type__iexact", "erc-20")).OrderBy("-block_number").Limit(maxLimit).All(&erc20TransferList)
	erc20Total, _ := o.QueryTable("token_transfer").SetCond(conditions.And("token_type__iexact", "erc-20")).Count()
	for _, item := range erc20TransferList {
		if strings.ToLower(item.From) == strings.ToLower(item.To) {
			item.Direction = "Self"
		} else if strings.ToLower(item.To) == strings.ToLower(address) {
			item.Direction = "In"
		} else {
			item.Direction = "Out"
		}

		if transactionMap[item.TransactionHash] <= 0 {
			var tx models.Transaction
			o.QueryTable("transaction").Filter("hash__iexact", item.TransactionHash).One(&tx)
			transactionMap[item.TransactionHash] = tx.Timestamp
		}
		item.Timestamp = transactionMap[item.TransactionHash]
		item.Datetime = time.Unix(int64(item.Timestamp), 0).Format("2006-01-02 15:04:05")

		if contractMap[item.TokenContractAddress] == "" {
			var addressName models.AddressName
			o.QueryTable("address_name").Filter("address__iexact", item.TokenContractAddress).One(&addressName, "name")
			contractMap[item.TokenContractAddress] = addressName.Name
		}
		item.TokenName = contractMap[item.TokenContractAddress]
	}

	var erc721TransferList []*models.TokenTransfer
	// Token transfer erc-721
	o.QueryTable("token_transfer").SetCond(conditions.And("token_type__iexact", "erc-721")).OrderBy("-block_number").Limit(maxLimit).All(&erc721TransferList)
	erc721Total, _ := o.QueryTable("token_transfer").SetCond(conditions.And("token_type__iexact", "erc-721")).Count()
	for _, item := range erc721TransferList {
		if strings.ToLower(item.From) == strings.ToLower(item.To) {
			item.Direction = "Self"
		} else if strings.ToLower(item.To) == strings.ToLower(address) {
			item.Direction = "In"
		} else {
			item.Direction = "Out"
		}

		if transactionMap[item.TransactionHash] <= 0 {
			var tx models.Transaction
			o.QueryTable("transaction").Filter("hash__iexact", item.TransactionHash).One(&tx)
			transactionMap[item.TransactionHash] = tx.Timestamp
		}
		item.Timestamp = transactionMap[item.TransactionHash]
		item.Datetime = time.Unix(int64(item.Timestamp), 0).Format("2006-01-02 15:04:05")

		if contractMap[item.TokenContractAddress] == "" {
			var addressName models.AddressName
			o.QueryTable("address_name").Filter("address__iexact", item.TokenContractAddress).One(&addressName, "name")
			contractMap[item.TokenContractAddress] = addressName.Name
		}
		item.TokenName = contractMap[item.TokenContractAddress]
	}

	var tokenList []*models.AddressTokenBalance
	o.QueryTable("address_token_balance").Filter("address__iexact", address).All(&tokenList)

	var addressDaily []models.AddressBalanceDaily
	o.QueryTable("address_balance_daily").Filter("address__iexact", address).Limit(30).All(&addressDaily)

	var logList []models.Log
	if addressInfo.IsContract {
		// Log - limit 25
		o.QueryTable("log").Filter("address__iexact", address).Limit(maxLimit).OrderBy("-block_number").All(&logList)

		var creator models.Transaction
		o.QueryTable("transaction").Filter("to", "").Filter("contract_address__iexact", address).One(&creator)
		addressInfo.Creator = creator.From
		addressInfo.CreateAt = creator.Hash
	}

	var tokens []common.BaseToken
	for _, item := range tokenList {
		if contractMap[item.ContractAddress] == "" {
			var addressName models.AddressName
			o.QueryTable("address_name").Filter("address__iexact", item.ContractAddress).One(&addressName, "name")
			contractMap[item.ContractAddress] = addressName.Name
		}
		tokens = append(tokens, common.BaseToken{
			TokenId:       item.TokenId,
			TokenType:     item.TokenType,
			TokenValue:    item.Value,
			TokenName:     contractMap[item.ContractAddress],
			TokenContract: item.ContractAddress,
		})
	}

	dataRes := make(map[string]interface{})
	dataRes["info"] = addressInfo
	dataRes["txns"] = transactionList
	dataRes["txns_total"] = txnsTotal
	dataRes["erc20"] = erc20TransferList
	dataRes["erc20Total"] = erc20Total
	dataRes["erc721"] = erc721TransferList
	dataRes["erc721Total"] = erc721Total
	dataRes["logs"] = logList
	dataRes["daily"] = addressDaily
	dataRes["tokens"] = tokens

	t.Data["json"] = Response{0, "success", dataRes}

	t.ServeJSON()
}
