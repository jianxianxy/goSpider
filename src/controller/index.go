package controller

import (
	"config"
	"fmt"
	"lib/spider"
	"strings"
	"time"
)

//页面分析结果
type pageSeg struct {
	Title string
	Url   string
	Seg   map[string]int
}

//搜索
func Search() {
	urlHttp := "http://chengde.esf.fang.com"

	que := &spider.Queue{}
	que.StacksA = make(map[string]string)
	que.StacksB = make(map[string]string)
	urlMap := map[string]string{spider.HashKey("/house-a013051/"): "/house-a013051/"}
	que.QueueInsert(urlMap)
	//开始循环
	for {
		turl := que.QueueShift()
		if turl != "" {
			curUrl := urlHttp + turl
			fangCom(curUrl, que)
		} else {
			break
		}
	}
	fmt.Println("结束")
}

func fangCom(url string, que *spider.Queue) {
	b_Time := time.Now()

	html, _ := spider.GetHtmlByUrl(url)
	html = spider.ConvertToString(html, "gbk", "utf-8")

	//分页信息
	page := spider.FindByIC(html, "div", ".page_al")
	phref := spider.GetUrlFromString(page)
	hrefMap := spider.HashMap(phref)
	que.QueueInsert(hrefMap)

	//页面数据
	ulHt := spider.FindByIC(html, "div", ".shop_list shop_list_4")
	liArr := make([]string, 0)
	spider.FindByICA(ulHt, "dl", "", &liArr)
	for _, v := range liArr {
		colTb := make(map[string]interface{})
		colTb["href"] = spider.StripQuotation(url)

		h4 := spider.FindByIC(v, "h4", "")
		colTb["title"] = spider.TrimSpace(spider.GetText(h4))

		info := spider.FindByIC(v, "p", ".tel_shop")
		infoArr := strings.Split(spider.TrimSpace(spider.GetText(info)), "|")
		infoLen := len(infoArr)
		infoEnd := infoLen - 1
		var extra string
		for ik, iv := range infoArr {
			if ik == 0 {
				colTb["style"] = iv
			} else if ik == 1 {
				colTb["area"] = iv
			} else if ik == 2 {
				colTb["layer"] = iv
			} else if ik > 2 && ik < infoEnd {
				extra += iv
			}
		}
		colTb["extra"] = extra

		posHt := spider.FindByIC(v, "p", ".add_shop")
		name := spider.FindByIC(posHt, "a", "")
		colTb["name"] = spider.TrimSpace(spider.GetText(name))
		rute := spider.FindByIC(posHt, "span", "")
		ruteArr := strings.Split(spider.TrimSpace(spider.GetText(rute)), "-")
		for rk, rv := range ruteArr {
			if rk == 0 {
				colTb["pos_1"] = spider.TrimSpace(rv)
			} else if rk == 1 {
				colTb["pos_2"] = spider.TrimSpace(rv)
			}
		}

		priceAll := spider.FindByIC(v, "dd", ".price_right")
		price := spider.FindByIC(priceAll, "b", "")
		colTb["price"] = spider.TrimSpace(spider.GetText(price))
		pricem2Hm := make([]string, 0)
		var pricem2 string
		spider.FindByICA(priceAll, "span", "", &pricem2Hm)
		if len(pricem2Hm) > 1 {
			pricem2 = pricem2Hm[1]
		}
		colTb["price_m2"] = spider.PickInt(spider.GetText(pricem2), 0)
		//插入数据库
		iin := insertRow("realty", colTb)
		fmt.Println(iin)
	}

	u_Time := time.Since(b_Time)
	fmt.Println("用时:", u_Time, "地址:", url)
}

//插入数据库
func insertRow(table string, data map[string]interface{}) int {
	mysql := config.DbSpider()
	return mysql.Insert(table, data)
}
