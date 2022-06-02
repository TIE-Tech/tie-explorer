package models

import "github.com/astaxie/beego/orm"

type BridgeHistory struct {
	Hash             string           `json:"Hash" orm:"size(128);unique"`
	From             string           `json:"User" orm:"size(128);"`
	FromChainId      int              `json:"SrcChainId"`
	To               string           `json:"DstUser" orm:"size(128);"`
	ToChainId        int              `json:"DstChainId"`
	TransferAmount   string           `json:"TransferAmount" orm:"size(128)"`
	ContractHash     string           `json:"ContractHash" orm:"size(128)"`
	State            int64            `json:"State"`
	Time             int64            `json:"Time" orm:"-"`
	TransactionState []*BridgeProcess `json:"TransactionState" orm:"-"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(BridgeHistory))
}
