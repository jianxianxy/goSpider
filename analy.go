package main

import (
	"bufio"
	"config"
	"fmt"
	"os"
	"strings"
	"time"
)

type Fang struct {
	Name  string
	Area  string
	Price string
	Day   string
}

func main() {
Loop:
	tips := "选择模式:price[价格分析],hot[热度分析],sale[销售分析]"
	fmt.Println(tips)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		switch input.Text() {
		case "price":
			fmt.Print("输入时间（例:2018-12）:")
			day := bufio.NewScanner(os.Stdin)
			for day.Scan() {
				if day.Text() != "exit" {
					analyPrice(day.Text())
				} else {
					goto Loop
				}
				fmt.Print("输入时间（例:2018-12）:")
			}
		case "hot":
			fmt.Print("输入时间（例:2018-12-01）:")
			day := bufio.NewScanner(os.Stdin)
			for day.Scan() {
				if day.Text() != "exit" {
					analyData(day.Text())
				} else {
					goto Loop
				}
				fmt.Print("输入时间（例:2018-12-01）:")
			}
		case "sale":
			fmt.Print("输入时间（例:2018-12-01）:")
			day := bufio.NewScanner(os.Stdin)
			for day.Scan() {
				if day.Text() != "exit" {
					analySaleInfo(day.Text()) //销售分析
				} else {
					goto Loop
				}
				fmt.Print("输入时间（例:2018-12-01）:")
			}
		}
	}
}

//全部分析
func analy() {
	sel := "SELECT `day`,COUNT(`day`) AS num FROM (SELECT DATE_FORMAT(`create`,'%Y-%m-%d') AS `day` FROM `realty`) tbday GROUP BY `day` ORDER BY `day` ASC"
	mysql := config.DbSpider()
	cols := mysql.GetRow(sel)
	if len(cols) > 0 {
		for _, val := range cols {
			curDay := val["day"]
			if analyNeed(curDay) {
				analyData(curDay)
			}
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

//获取前一天的日期
func getPreDay(curDay string) string {
	sel := "SELECT DATE_FORMAT(`create`,'%Y-%m-%d') AS `day` FROM `realty` WHERE DATE_FORMAT(`create`,'%Y-%m-%d') < '" + curDay + "' ORDER BY `create` DESC LIMIT 0,1"
	mysql := config.DbSpider()
	row := mysql.GetRow(sel)
	preDay := ""
	if len(row) > 0 {
		for _, val := range row {
			return val["day"]
		}
	}
	return preDay
}

//热度分析
func analyData(curDay string) {
	curData := dataByDay(curDay)
	//售出 && 新增
	var saleInfo string
	var saleInt, addInt int
	for sk, sv := range curData {
		if analySale(sk, curDay) {
			saleInt = saleInt + 1
			saleInfo = saleInfo + "," + sv
		} else if analyNew(sk, curDay) {
			addInt = addInt + 1
		}
	}
	//保存
	inData := make(map[string]interface{})
	inData["anday"] = curDay
	inData["add"] = addInt
	inData["reduce"] = saleInt
	inData["bare"] = addInt - saleInt
	inData["info"] = strings.Trim(saleInfo, ",")
	fmt.Println(inData)
	mysql := config.DbSpider()
	res := mysql.Insert("data_chart", inData)
	fmt.Println(curDay, res)
}

//获取某天的数据
func dataByDay(day string) map[string]string {
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

//分析销售详情
func analySaleInfo(day string) {
	curData := dataByDay(day)
	for key, val := range curData {
		if analySale(key, day) {
			fmt.Println(key, day)
			onSale := saleStart(key)
			dayInt := timeSub(onSale, day)
			info := rowInfo(val)
			info["showday"] = dayInt
			info["anday"] = day
			info["rea_id"] = val
			info["onday"] = onSale
			info["offday"] = day
			mysql := config.DbSpider()
			mysql.Insert("data_sale", info)
		}
	}
	fmt.Println("处理完成：", day)
}

//价格分析
func analyPrice(month string) {
	sel := "SELECT DISTINCT signkey FROM `realty` WHERE DATE_FORMAT(`create`,'%Y-%m') = '" + month + "' LIMIT 0,10000"
	mysql := config.DbSpider()
	all := mysql.GetRow(sel)
	if len(all) > 0 {
		for _, val := range all {
			getChangePrice(val["signkey"])
		}
	}
	fmt.Println(".")
	fmt.Println("处理完成：", month)
}

//获取价格
func getChangePrice(signkey string) {
	prcMap := make(map[string]interface{})
	sel := "SELECT `price`,`name`,`area`,DATE_FORMAT(`create`,'%Y-%m-%d') AS `day` FROM `realty` WHERE signkey = '" + signkey + "' ORDER BY `day` ASC"
	mysql := config.DbSpider()
	all := mysql.GetRow(sel)
	if len(all) > 0 {
		for _, val := range all {
			if _, has := prcMap[val["price"]]; !has {
				prcMap[val["price"]] = Fang{Name: val["name"], Area: val["area"], Price: val["price"], Day: val["day"]}
			}
		}
	}
	if len(prcMap) > 1 {
		for _, info := range prcMap {
			data := make(map[string]interface{})
			data["signkey"] = signkey
			data["name"] = info.(Fang).Name
			data["area"] = info.(Fang).Area
			data["price"] = info.(Fang).Price
			data["day"] = info.(Fang).Day
			if needUpPrice(signkey, info.(Fang).Day) {
				mysql := config.DbSpider()
				mysql.Insert("data_price", data)
			}
		}
	}
	fmt.Print(".")
}

//是否需要更新价格流水
func needUpPrice(signkey, day string) bool {
	sel := "SELECT `id` FROM `data_price` WHERE `signkey` = '" + signkey + "' AND `day` = '" + day + "'"
	mysql := config.DbSpider()
	row := mysql.GetRow(sel)
	if len(row) < 1 {
		return true
	} else {
		return false
	}
}

//上架时间
func saleStart(signkey string) string {
	sel := "SELECT DATE_FORMAT(`create`,'%Y-%m-%d') AS `day` FROM `realty` WHERE `signkey`='" + signkey + "' ORDER BY `create` ASC LIMIT 0,1"
	mysql := config.DbSpider()
	row := mysql.GetRow(sel)
	if len(row) > 0 {
		for _, val := range row {
			return val["day"]
		}
	}
	return ""
}

//记录详情
func rowInfo(id string) map[string]interface{} {
	info := make(map[string]interface{})
	sel := "SELECT `name`,`area`,`price` FROM `realty` WHERE id = " + id
	mysql := config.DbSpider()
	row := mysql.GetRow(sel)
	if len(row) > 0 {
		for _, val := range row {
			info["name"] = val["name"]
			info["area"] = val["area"]
			info["price"] = val["price"]
		}
	}
	return info
}

//计算两个日期相差几天
func timeSub(dm, dx string) int {
	layout := "2006-01-02 15:04:05"
	dtm, _ := time.Parse(layout, dm+" 00:00:00")
	dtx, _ := time.Parse(layout, dx+" 00:00:00")
	return int(dtx.Sub(dtm).Hours() / 24)
}
