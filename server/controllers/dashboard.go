package controllers

import (
	"github.com/astaxie/beego/orm"
	"regexp"
	"strings"
	"time"
)

type DashboardController struct {
	BaseController
}

func (d DashboardController) Info() {
	o := orm.NewOrm()

	days := 13

	now := time.Now()
	maxDate := now.Format("2006-01-02")
	maxTimes, _ := time.Parse("2006-01-02 15:04:05", maxDate+" 23:59:59")
	maxTimestamp := maxTimes.Unix()
	minDate := now.AddDate(0, 0, -days).Format("2006-01-02")
	minTimes, _ := time.Parse("2006-01-02 15:04:05", minDate+" 00:00:00")
	minTimestamp := minTimes.Unix()

	type TxnCount struct {
		Count int64  `json:"count"`
		Date  string `json:"date"`
	}
	var txnCounts []TxnCount
	o.Raw("SELECT COUNT(\"id\") as count, to_char(to_timestamp(\"timestamp\"), 'MM-DD') AS date FROM transaction WHERE timestamp >= ? AND timestamp <= ? GROUP BY \"date\"", minTimestamp, maxTimestamp).QueryRows(&txnCounts)
	txnCountMap := make(map[string]int64)
	for _, item := range txnCounts {
		txnCountMap[item.Date] = item.Count
	}

	var txnCountArr []TxnCount
	for i := days; i >= 0; i-- {
		nowDate := now.AddDate(0, 0, -i).Format("01-02")
		var count int64 = 0
		if txnCountMap[nowDate] > 0 {
			count = txnCountMap[nowDate]
		}
		txnCountArr = append(txnCountArr, TxnCount{
			Count: count,
			Date:  nowDate,
		})
	}

	txns, _ := o.QueryTable("transaction").Count()
	blocks, _ := o.QueryTable("block").Count()
	addresses, _ := o.QueryTable("address").Count()

	dataRes := make(map[string]interface{})
	dataRes["txnCountArr"] = txnCountArr
	dataRes["txns"] = txns
	dataRes["blocks"] = blocks
	dataRes["addresses"] = addresses
	dataRes["avg"] = "2"

	d.Data["json"] = Response{0, "success", dataRes}

	d.ServeJSON()
}

func (d DashboardController) Search() {
	s := d.GetString("s", "")
	t := d.GetString("t", "")

	pattern := "^(\\d+)$"
	isNumber, _ := regexp.MatchString(pattern, s)

	dataRes := make(map[string]interface{})
	var dataType string
	var exists = false
	o := orm.NewOrm()

	if t != "" {
		t = strings.ToLower(t)
		if t == "address" {
			dataType = "address"
			if len(s) == 42 {
				exists = o.QueryTable("address").Filter("address__iexact", s).Exist()
			}
		} else if t == "token" {
			dataType = "token"
			if len(s) == 42 {
				exists = o.QueryTable("token").Filter("contract_address__iexact", s).Exist()
			}
		}
	} else {
		if isNumber {
			dataType = "block"
			exists = o.QueryTable("block").Filter("number", s).Exist()
		} else {
			if len(s) == 42 {
				exists = o.QueryTable("address").Filter("address__iexact", s).Exist()
				if exists {
					dataType = "address"
				} else {
					dataType = "token"
					exists = o.QueryTable("token").Filter("contract_address__iexact", s).Exist()
				}
			} else if len(s) == 66 {
				dataType = "transaction"
				exists = o.QueryTable("transaction").Filter("hash__iexact", s).Exist()
			}
		}
	}

	dataRes["type"] = dataType
	dataRes["exists"] = exists

	d.Data["json"] = Response{0, "success", dataRes}

	d.ServeJSON()
}
