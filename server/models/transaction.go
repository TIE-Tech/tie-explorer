package models

import (
	"github.com/astaxie/beego/orm"
)

type Transaction struct {
	Hash              string `json:"transactionHash" orm:"unique;index;size(128)"`
	Index             uint64 `json:"transactionIndex" orm:"default(0)"`
	BlockHash         string `json:"blockHash" orm:"index;size(128)"`
	BlockNumber       uint64 `json:"blockNumber" orm:"index"`
	Timestamp         uint64 `json:"timestamp"`
	DateTime          string `json:"dateTime" orm:"size(32)"`
	From              string `json:"from" orm:"index;size(128)"`
	To                string `json:"to" orm:"index;size(128)"`
	ToStr             string `json:"toStr" orm:"-"`
	Direction         string `json:"direction" orm:"-"`
	BlockConfirms     uint64 `json:"blockConfirms" orm:"-"`
	Value             string `json:"value"`
	Nonce             uint64 `json:"nonce"`
	GasPrice          uint64 `json:"gasPrice"`
	Gas               uint64 `json:"gas"`
	GasUsed           uint64 `json:"gasUsed"`
	CumulativeGasUsed uint64 `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	V                 string `json:"v" orm:"size(128)"`
	R                 string `json:"r" orm:"size(128)"`
	S                 string `json:"s" orm:"size(128)"`
	Type              uint8  `json:"type"`
	Genre             string `json:"genre" orm:"-"`
	Root              string `json:"root" orm:"size(128)"`
	IsFake            bool   `json:"isFake"`
	Status            uint64 `json:"status"`
	LogsBoom          string `json:"logsBloom"`
	ContractAddress   string `json:"contractAddress" orm:"size(128)"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(Transaction))
}
