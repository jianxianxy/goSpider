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
	url := "http://www.baidu.com/s?wd="
	_, href := spider.Exec(url)
	u_Time := time.Since(b_Time)
	fmt.Println("用时:", u_Time)
	fmt.Println("内容:", href)
}
