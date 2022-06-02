package models

import "github.com/astaxie/beego/orm"

type BridgeChain struct {
	ChainId int    `json:"ChainId" orm:"unique"`
	Name    string `json:"Name" orm:"size(128)"`
	Rpc     string `json:"Rpc" orm:"size(128)"`
	Topic   string `json:"Topic" orm:"size(128);null"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(BridgeChain))
}
