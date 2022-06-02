package models

import "github.com/astaxie/beego/orm"

type AddressBalanceDaily struct {
	Address string `json:"address" orm:"size(128)"`
	Date    string `json:"date"`
	Value   string `json:"value" orm:"null"`
	BaseModel
}

func (a *AddressBalanceDaily) TableUnique() [][]string {
	return [][]string{
		{"address", "date"},
	}
}

func init() {
	orm.RegisterModel(new(AddressBalanceDaily))
}
