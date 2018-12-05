package main

import (
	"config"
	"fmt"
	"lib/spider"
	"strings"
)

func main() {
	analy()
	//analyData("2018-12-03","2018-12-04")
}

func analy() {
	limno := 31
	sel := "SELECT `day`,COUNT(`day`) AS num FROM (SELECT DATE_FORMAT(`create`,'%Y-%m-%d') AS `day` FROM `realty`) tbday GROUP BY `day` ORDER BY `day` DESC LIMIT 0," + limno
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

//是否需要处理
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

//获取后天的数据
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

//是否售出
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

//是否新增
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
