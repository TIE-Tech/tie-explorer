package models

import "github.com/astaxie/beego/orm"

type AddressName struct {
	Address  string `json:"address" orm:"size(128)"`
	Name     string `json:"name" orm:"size(128)"`
	Primary  bool   `json:"primary"`
	Metadata string `json:"metadata"`
	BaseModel
}

func (a *AddressName) TableUnique() [][]string {
	return [][]string{
		{"address", "name"},
	}
}

func init() {
	orm.RegisterModel(new(AddressName))
}
