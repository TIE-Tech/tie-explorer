package models

import "github.com/astaxie/beego/orm"

type Log struct {
	Address          string `json:"address" orm:"size(128)"`
	Topics0          string `json:"topics0" orm:"size(128)"`
	Topics1          string `json:"topics1" orm:"size(128)"`
	Topics2          string `json:"topics2" orm:"size(128)"`
	Topics3          string `json:"topics3" orm:"size(128)"`
	Data             string `json:"data"`
	BlockNumber      uint64 `json:"blockNumber" orm:"column(block_number)"`
	TransactionHash  string `json:"transactionHash" orm:"column(transaction_hash);size(128)"`
	TransactionIndex uint64 `json:"transactionIndex" orm:"column(transaction_index)"`
	BlockHash        string `json:"blockHash" orm:"column(block_hash);size(128)"`
	LogIndex         uint64 `json:"logIndex" orm:"column(log_index)"`
	Removed          bool   `json:"removed"`
	BaseModel
}

func (l *Log) TableUnique() [][]string {
	return [][]string{
		{"transaction_hash", "log_index"},
	}
}

func init() {
	orm.RegisterModel(new(Log))
}
