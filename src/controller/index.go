package controller

import (
	"config"
	"html/template"
	"lib/seg"
	"lib/spider"
	"net/http"
	"time"
)

//页面分析结果
type pageSeg struct {
	Title string
	Url   string
	Seg   map[string]int
}

//首页
func Index(rp http.ResponseWriter, rq *http.Request) {
	rp.Header().Set("Content-Type", "text/html")
	//调用模版
	view, err := template.ParseFiles(config.Get("ROOT_PATH") + "static/index.html")
	if err != nil {
		http.Error(rp, err.Error(), http.StatusInternalServerError)
		return
	}
	locals := make(map[string]interface{})
	locals["info"] = []string{}
	view.Execute(rp, locals)
}

//搜索
func Search(rp http.ResponseWriter, rq *http.Request) {
	rp.Header().Set("Content-Type", "text/html")
	url := rq.FormValue("url")
	keyword := rq.FormValue("keyword")
	//设置模版
	view, err := template.ParseFiles(config.Get("ROOT_PATH") + "static/view.html")
	if err != nil {
		http.Error(rp, err.Error(), http.StatusInternalServerError)
		return
	}

	locals := make([]pageSeg, 0)
	if len(url) < 1 {
		url = "http://www.baidu.com/s?wd=" + keyword
	}
	_, href := spider.Exec(url)
	if len(href) > 0 {
		//设置超时时间
		timeout := make(chan bool, 1)
		go func() {
			time.Sleep(6e9)
			timeout <- true
		}()
		//加载字典(why is here 节省内存,否则内存会溢出的)
		seg.LoadDict()
		//并行计算
		var charr []chan pageSeg
		for key, val := range href {
			chspi := make(chan pageSeg, 1)
			charr = append(charr, chspi)
			go goSpider(key, val, chspi)
		}
		for _, cha := range charr {
			select {
			case chret := <-cha:
				locals = append(locals, chret)
			case <-timeout:
				break
			}
		}
	}

	view.Execute(rp, locals)
}

func goSpider(keyword string, url string, ch chan pageSeg) {
	info, _ := spider.Exec(url)
	if body, ok := info["body"]; ok {
		plan := seg.SegString(keyword)
		match := seg.MatchLevel(plan, body)
		ch <- pageSeg{Title: keyword, Url: url, Seg: match}
	}
}
