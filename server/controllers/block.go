package controllers

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"tie-explorer-go/common"
	"tie-explorer-go/models"
)

type BlockController struct {
	BaseController
}

func (b BlockController) List() {
	pageSize, _ := b.GetInt64("ps", 20)
	page, _ := b.GetInt64("p", 1)

	var blocks []*models.Block
	o := orm.NewOrm()
	qt := o.QueryTable("block")
	total, _ := qt.Count()
	pagination := common.GetPagination(total, page, pageSize)
	qt.Offset(pagination.Offset).Limit(pageSize).OrderBy("-number").All(&blocks)

	dataRes := make(map[string]interface{})
	dataRes["blocks"] = blocks
	dataRes["pagination"] = pagination
	b.Data["json"] = Response{0, "success", dataRes}

	b.ServeJSON()
}

func (b BlockController) Info() {
	number := b.GetString(":number")
	numberUint64, _ := strconv.ParseUint(number, 0, 64)

	block := models.Block{}

	o := orm.NewOrm()
	o.QueryTable("block").Filter("number", numberUint64).One(&block)

	b.Data["json"] = Response{0, "success", block}

	b.ServeJSON()
}
