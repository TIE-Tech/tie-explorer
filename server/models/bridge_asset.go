package models

import "github.com/astaxie/beego/orm"

type BridgeAsset struct {
	Name      string        `json:"Name" orm:"size(128);unique"`
	Precision int           `json:"Precision"`
	Meta      string        `json:"Meta"`
	State     bool          `json:"-"`
	Tokens    []BridgeToken `json:"Tokens" orm:"-"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(BridgeAsset))
}
