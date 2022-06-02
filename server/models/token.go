package models

import "github.com/astaxie/beego/orm"

type Token struct {
	ContractAddress string `json:"contractAddress" orm:"unique;size(128)"`
	Name            string `json:"name" orm:"size(128);"`
	Symbol          string `json:"symbol" orm:"size(128);"`
	TotalSupply     string `json:"totalSupply" orm:"size(128);"`
	Decimals        string `json:"decimals" orm:"size(128);"`
	Type            string `json:"type" orm:"size(32);"`
	Bridged         bool   `json:"bridged"`
	Cataloged       bool   `json:"cataloged"`
	SkipMetadata    bool   `json:"skipMetadata"`
	Holders         int64  `json:"holders" orm:"-"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(Token))
}
