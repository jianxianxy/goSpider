package main

import (
	"config"
	"fmt"
	"lib/spider"
	"strings"
)

func main() {
	analy()
}

func analy() {
	sel := "SELECT `day`,COUNT(`day`) AS num FROM (SELECT DATE_FORMAT(`create`,'%Y-%m-%d') AS `day` FROM `realty`) tbday GROUP BY `day` ORDER BY `day` ASC"
	mysql := config.DbSpider()
	cols := mysql.GetRow(sel)
	var preDay string
	if len(cols) > 0 {
		for _, val := range cols {
			curDay := val["day"]
			if analyNeed(curDay) && preDay != "" {
				analyData(preDay, curDay)
			}
			preDay = curDay
		}
	}
}
func analyNeed(day string) bool {
	sel := "SELECT `id` FROM `data_chart` WHERE `anday` = '" + day + "'"
	mysql := config.DbSpider()
	row := mysql.GetRow(sel)
	if len(row) < 1 {
		return true
	} else {
		return false
	}
}
func analyData(preDay, curDay string) {
	preData := analyDay(preDay)
	curData := analyDay(curDay)
	for key, _ := range curData {
		if _, ok := preData[key]; ok {
			delete(preData, key)
			delete(curData, key)
		}
	}
	//售出
	var saleInfo string
	for sk, sv := range preData {
		if analySale(sk, curDay) == false {
			delete(preData, sk)
		} else {
			saleInfo = saleInfo + "," + sv
		}
	}
	//新增
	for ak, _ := range curData {
		if analyNew(ak, preDay) == false {
			delete(curData, ak)
		}
	}
	inData := make(map[string]interface{})
	inData["anday"] = curDay
	inData["add"] = len(curData)
	inData["reduce"] = len(preData)
	inData["bare"] = len(curData) - len(preData)
	inData["info"] = strings.Trim(saleInfo, ",")
	fmt.Println(inData)
	mysql := config.DbSpider()
	res := mysql.Insert("data_chart", inData)
	fmt.Println(curDay, res)
}

func analyDay(day string) map[string]string {
	sel := "SELECT `id`,`signkey` FROM `realty` WHERE DATE_FORMAT(`create`,'%Y-%m-%d') = '" + day + "'"
	mysql := config.DbSpider()
	all := mysql.GetRow(sel)
	dmp := make(map[string]string)
	if len(all) > 0 {
		for _, val := range all {
			dmp[val["signkey"]] = val["id"]
		}
	}
	return dmp
}

func analySale(signkey, day string) bool {
	sel := "SELECT * FROM realty WHERE signkey = '" + signkey + "' AND DATE_FORMAT(`create`,'%Y-%m-%d') > '" + day + "'"
	mysql := config.DbSpider()
	row := mysql.GetRow(sel)
	if len(row) > 0 {
		return false
	} else {
		return true
	}
}
func analyNew(signkey, day string) bool {
	sel := "SELECT * FROM realty WHERE signkey = '" + signkey + "' AND DATE_FORMAT(`create`,'%Y-%m-%d') < '" + day + "'"
	mysql := config.DbSpider()
	row := mysql.GetRow(sel)
	if len(row) > 0 {
		return false
	} else {
		return true
	}
}

func pefData() {
	sel := "SELECT `id`,`title`,`area`,DATE_FORMAT(`create`,'%Y-%m-%d') AS `create`  FROM `realty` WHERE `signkey` = '' ORDER BY `id` ASC LIMIT 0,30000"
	mysql := config.DbSpider()
	cols := mysql.GetRow(sel)
	if len(cols) > 0 {
		for _, val := range cols {
			pefCol(val)
		}
	}
}
func pefCol(data map[string]string) {
	where := make(map[string]interface{})
	where["id"] = data["id"]
	signKey := spider.HashKey(string(data["title"]) + string(data["area"]))
	upcol := make(map[string]interface{})
	upcol["signkey"] = signKey
	sel := "SELECT `id` FROM `realty` WHERE `signkey` = '" + signKey + "' AND DATE_FORMAT(`create`,'%Y-%m-%d') = '" + string(data["create"]) + "'"
	mysql := config.DbSpider()
	col := mysql.GetRow(sel)
	if len(col) > 0 {
		mysql.Remove("realty", where)
		fmt.Println("删除:", data["id"])
	} else {
		mysql.Update("realty", upcol, where)
		fmt.Println("更新:", data["id"])
	}
}
