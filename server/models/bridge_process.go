package models

import "github.com/astaxie/beego/orm"

type BridgeProcess struct {
	HistoryId   int    `json:"-"`
	Hash        string `json:"Hash" orm:"size(128);null"`
	ChainId     int    `json:"ChainId"`
	State       int64  `json:"State"`
	BlockNumber int64  `json:"BlockNumber"`
	Blocks      int64  `json:"Blocks" orm:"-"`
	NeedBlocks  int64  `json:"NeedBlocks" orm:"-"`
	BaseModel
}

func init() {
	orm.RegisterModel(new(BridgeProcess))
}
