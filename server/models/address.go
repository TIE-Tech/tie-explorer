package models

import "github.com/astaxie/beego/orm"

type Address struct {
	Address      string `json:"address" orm:"unique;size(128)"`
	IsContract   bool   `json:"isContract"`
	ContractCode string `json:"contractCode" orm:"null"`
	Decompiled   bool   `json:"decompiled"`
	Verified     bool   `json:"verified"`
	Balance      string `json:"balance" orm:"size(128)"`
	Creator      string `json:"creator" orm:"-"`
	CreateAt     string `json:"createAt" orm:"-"`
	Senders      int64  `json:"senders" orm:"-"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(Address))
}
