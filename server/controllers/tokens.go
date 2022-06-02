package controllers

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"tie-explorer-go/common"
	"tie-explorer-go/models"
	"time"
)

type TokenController struct {
	BaseController
}

func (t TokenController) List() {
	tp := t.GetString("t", "")
	pageSize, _ := t.GetInt64("ps", 20)
	page, _ := t.GetInt64("p", 1)
	if page <= 0 {
		page = 1
	}

	dataRes := make(map[string]interface{})
	o := orm.NewOrm()

	if tp == "tie" {
		total, _ := o.QueryTable("address").Filter("balance__gt", 0).Count()
		pagination := common.GetPagination(total, page, pageSize)

		var tokens []*models.Address
		o.Raw("SELECT * FROM address WHERE balance > '0' ORDER BY lpad(balance, 30, '0') DESC OFFSET ? LIMIT ?", pagination.Offset, pageSize).QueryRows(&tokens)

		for _, item := range tokens {
			txns, _ := o.QueryTable("transaction").Distinct().Filter("from__iexact", item.Address).Count()
			item.Senders = txns
		}
		dataRes["pagination"] = pagination
		dataRes["tokens"] = tokens
	} else {
		var tokens []*models.Token
		qt := o.QueryTable("token").Filter("type__iexact", tp).FilterRaw("name", "!= ''")
		total, _ := qt.Count()
		pagination := common.GetPagination(total, page, pageSize)

		qt.Offset(pagination.Offset).Limit(pageSize).OrderBy("-created_at").All(&tokens)

		for _, item := range tokens {
			holders, _ := o.QueryTable("token_transfer").Distinct().Filter("token_type__iexact", tp).Count()
			item.Holders = holders
		}
		dataRes["pagination"] = pagination
		dataRes["tokens"] = tokens
	}

	t.Data["json"] = Response{0, "success", dataRes}

	t.ServeJSON()
}

func (t TokenController) Info() {
	address := t.GetString(":address")

	transactionMap := make(map[string]uint64)
	contractMap := make(map[string]string)
	var tokenInfo models.Token
	var transfer []*models.TokenTransfer
	var holder []models.AddressTokenBalance

	o := orm.NewOrm()
	o.QueryTable("token").Filter("contract_address", address).One(&tokenInfo)

	if tokenInfo.ContractAddress != "" {
		o.QueryTable("token_transfer").Filter("token_contract_address", tokenInfo.ContractAddress).OrderBy("-block_number").All(&transfer)
		for _, item := range transfer {
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
		o.QueryTable("address_token_balance").Filter("contract_address", tokenInfo.ContractAddress).OrderBy("-value").All(&holder)
	}

	dataRes := make(map[string]interface{})
	dataRes["info"] = tokenInfo
	dataRes["transfer"] = transfer
	dataRes["holder"] = holder

	t.Data["json"] = Response{0, "success", dataRes}

	t.ServeJSON()
}
