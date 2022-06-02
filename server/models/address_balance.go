package models

import "github.com/astaxie/beego/orm"

type AddressBalance struct {
	Address   string `json:"address" orm:"size(128)"`
	Value     string `json:"value" orm:"null"`
	FetchedAt uint64 `json:"fetchedAt"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(AddressBalance))
}
