package config

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

//获取URL内容
func GetHtmlByUrl(url string) (str string, err error) {
	response, err := http.Get(url)
	if err != nil {
		return "http.Get Error", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}

//从字符串中提取超链接
func GetUrlFromString(body string) map[string]string {
	reg := regexp.MustCompile(`<a[^>]*?href=['"]([^'"]*?)['"][^>]*?>(.*?)<\/a>`)
	match := reg.FindAllStringSubmatch(string(body), -1)
	mapCatch := make(map[string]string)
	for _, v := range match {
		if len(v) == 3 {
			mapCatch[v[2]] = v[1]
		}
	}
	return mapCatch
}
