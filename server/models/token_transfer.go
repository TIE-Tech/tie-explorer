package models

import "github.com/astaxie/beego/orm"

type TokenTransfer struct {
	BlockNumber          uint64 `json:"blockNumber"`
	BlockHash            string `json:"blockHash" orm:"size(128);"`
	TransactionHash      string `json:"transactionHash" orm:"index;size(128);"`
	From                 string `json:"from" orm:"size(128)"`
	To                   string `json:"to" orm:"size(128);"`
	Direction            string `json:"direction" orm:"-"`
	Timestamp            uint64 `json:"timestamp" orm:"-"`
	Datetime             string `json:"datetime" orm:"-"`
	Value                string `json:"value" orm:"size(128);"`
	LogIndex             uint64 `json:"logIndex"`
	TokenContractAddress string `json:"tokenContractAddress" orm:"size(128);"`
	TokenId              string `json:"tokenId" orm:"null"`
	TokenType            string `json:"tokenType"`
	TokenName            string `json:"tokenName" orm:"-"`
	BaseModel
}

func (t *TokenTransfer) TableUnique() [][]string {
	return [][]string{
		{"transaction_hash", "log_index"},
	}
}

func init() {
	orm.RegisterModel(new(TokenTransfer))
}
