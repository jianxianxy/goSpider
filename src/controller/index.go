package controller

import (
	"config"
	"html/template"
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
	body, err := config.GetHtmlByUrl(url)
	if err == nil {
		locals["body"] = config.GetUrlFromString(body)
	}

	locals["info"] = []string{"Query:", val}
	view.Execute(rp, locals)
}
