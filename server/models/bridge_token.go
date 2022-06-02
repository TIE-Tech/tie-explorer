package models

import "github.com/astaxie/beego/orm"

type BridgeToken struct {
	Hash           string `json:"Hash" orm:"size(128);unique"`
	ChainId        int    `json:"ChainId"`
	Name           string `json:"Name" orm:"size(128)"`
	TokenBasicName string `json:"TokenBasicName" orm:"size(128)"`
	Property       int    `json:"Property"`
	Precision      int    `json:"Precision"`
	State          bool   `json:"State"`
	ResourceId     string `json:"ResourceId" orm:"size(128)"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(BridgeToken))
}
