package controller

import (
	"config"
	"fmt"
	"html/template"
	"lib/function"
	"lib/seg"
	"net/http"
)

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
	val := rq.FormValue("keyword")
	url := rq.FormValue("url")
	view, err := template.ParseFiles(config.Get("ROOT_PATH") + "static/view.html")
	if err != nil {
		http.Error(rp, err.Error(), http.StatusInternalServerError)
		return
	}

	locals := make(map[string]interface{})
	if len(url) < 1 {
		url = "http://www.baidu.com/s?wd=" + val
	}
	//获取HTML
	html, err := function.GetHtmlByUrl(url)
	if err == nil {
		//HTML标签转小写并去除样式、脚本
		html = function.TagToLower(html)
		//提取页面的超链接
		locals["href"] = function.GetUrlFromString(html)
		//获取title
		locals["title"] = function.GetTitle(html)
		//获取Body
		body := function.GetBody(function.StripNote(function.StripStyle(function.StripScript(html))))
		locals["body"] = body
		//保留div和h1的结构
		boh1 := function.GetDivH1(body)
		//获取h1内容(文章标题)
		locals["conTitle"] = function.GetH1(boh1)

	}
	//分词
	plan := seg.SegString(val)
	fmt.Println(plan)

	view.Execute(rp, locals)
}
