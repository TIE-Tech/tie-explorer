package controllers

import (
	"github.com/astaxie/beego/orm"
	"tie-explorer-go/common"
	"tie-explorer-go/models"
)

type TransactionController struct {
	BaseController
}

func (t TransactionController) List() {
	pageSize, _ := t.GetInt64("ps", 20)
	page, _ := t.GetInt64("p", 1)
	block, _ := t.GetInt64("block", 0)

	var transactions []*models.Transaction
	o := orm.NewOrm()
	qt := o.QueryTable("transaction")
	if block > 0 {
		qt = qt.Filter("block_number", block)
	}
	total, _ := qt.Count()

	pagination := common.GetPagination(total, page, pageSize)

	qt.Offset(pagination.Offset).Limit(pageSize).OrderBy("-timestamp").All(&transactions)

	for _, item := range transactions {
		if item.From == item.To {
			item.Direction = "Self"
		} else {
			item.Direction = "->"
		}

		if item.To == "" {
			item.Genre = "<span class='badge bg-cyan'>Contract Creation</span>"
			if item.ContractAddress != "" {
				item.To = item.ContractAddress
			}
		} else {
			var addressModel models.Address
			o.QueryTable("address").Filter("address", item.To).One(&addressModel)
			if addressModel.IsContract {
				item.Genre = "<span class='badge bg-purple'>Contract Call</span>"
			} else {
				item.Genre = "<span class='badge bg-blue'>Transaction</span>"
			}
		}
	}

	dataRes := make(map[string]interface{})
	dataRes["transactions"] = transactions
	dataRes["pagination"] = pagination

	t.Data["json"] = Response{0, "success", dataRes}

	t.ServeJSON()
}

func (t TransactionController) Info() {
	hash := t.GetString(":hash")
	transaction := &models.Transaction{}
	var tokenTransfer []*models.TokenTransfer
	var log []models.Log

	o := orm.NewOrm()
	o.QueryTable("transaction").Filter("hash", hash).One(transaction)

	if transaction.Hash != "" {
		o.QueryTable("token_transfer").Filter("transaction_hash", hash).All(&tokenTransfer)
		o.QueryTable("log").Filter("transaction_hash", hash).All(&log)

		var block models.Block
		o.QueryTable("block").OrderBy("-number").One(&block)

		blockConfirms := transaction.BlockNumber
		if block.Number-blockConfirms < 0 {
			blockConfirms = 1
		} else {
			blockConfirms = block.Number - blockConfirms + 1
		}

		transaction.BlockConfirms = blockConfirms

		for _, item := range tokenTransfer {
			var addrName models.AddressName
			err := o.QueryTable("address_name").Filter("address", item.TokenContractAddress).One(&addrName)
			if err != nil {
				item.TokenName = "Undefined"
				continue
			}
			item.TokenName = addrName.Name
		}
		if transaction.To == "" {
			transaction.Genre = "<span class='badge bg-cyan'>Contract Creation</span>"
			if transaction.ContractAddress != "" {
				transaction.To = transaction.ContractAddress
				transaction.ToStr = "[Contract <a href='/address/" + transaction.To + "'>" + transaction.ContractAddress + "</a> created]"
			}
		} else {
			transaction.ToStr = "<a href='/address/" + transaction.To + "'>" + transaction.To + "</a>"
			var addressModel models.Address
			o.QueryTable("address").Filter("address", transaction.To).One(&addressModel)
			if addressModel.IsContract {
				transaction.Genre = "<span class='badge bg-purple'>Contract Call</span>"
				transaction.ToStr = "[Contract <a href='/address/" + transaction.To + "'>" + transaction.To + "</a> ]"
			} else {
				transaction.Genre = "<span class='badge bg-blue'>Transaction</span>"
			}
		}
	}

	dataRes := make(map[string]interface{})
	dataRes["transaction"] = transaction
	dataRes["log"] = log
	dataRes["tokenTransfer"] = tokenTransfer

	t.Data["json"] = Response{0, "success", dataRes}

	t.ServeJSON()
}
