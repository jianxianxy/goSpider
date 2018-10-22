package controller

import (
	"fmt"
	"lib/spider"
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
	b_Time := time.Now()
	url := "http://chengde.esf.fang.com/house-a013051/"
	html, _ := spider.GetHtmlByUrl(url)
	html = spider.ConvertToString(html, "gbk", "utf-8")
	ulHt := spider.FindByIC(html, "div", ".shop_list shop_list_4")
	var liArr []string
	spider.FindByICA(ulHt, "dl", "", &liArr)
	for k, v := range liArr {
		fmt.Println(k+1, "----------")
		h4 := spider.FindByIC(v, "h4", "")
		fmt.Println(spider.TrimSpace(spider.GetText(h4)))
		info := spider.FindByIC(v, "p", ".tel_shop")
		posHt := spider.FindByIC(v, "p", ".add_shop")
		name := spider.FindByIC(posHt, "a", "")
		rute := spider.FindByIC(posHt, "span", "")
		price := spider.FindByIC(v, "dd", ".price_right")
		fmt.Println(spider.TrimSpace(spider.GetText(price)))
		fmt.Println(spider.TrimSpace(spider.GetText(name)), ":", spider.GetText(rute))
		fmt.Println(spider.TrimSpace(spider.GetText(info)), spider.TrimSpace(spider.GetText(price)))

	}

	u_Time := time.Since(b_Time)
	fmt.Println("用时:", u_Time)

}
