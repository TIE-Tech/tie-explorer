package models

import "github.com/astaxie/beego/orm"

type AddressTokenBalance struct {
	Address         string `json:"address" orm:"size(128)"`
	BlockNumber     uint64 `json:"blockNumber"`
	ContractAddress string `json:"contractAddress" orm:"size(128)"`
	Value           string `json:"value" orm:"size(128)"`
	TokenId         string `json:"tokenId" orm:"size(128)"`
	TokenType       string `json:"tokenType" orm:"size(128)"`
	BaseModel
}

func (a *AddressTokenBalance) TableUnique() [][]string {
	return [][]string{
		{"address", "contract_address"},
	}
}

func init() {
	orm.RegisterModel(new(AddressTokenBalance))
}
