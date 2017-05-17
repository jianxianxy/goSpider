package config

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
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
	reg := regexp.MustCompile(`<a[^>]*?href\s?=\s?['"]([^'"]*?)['"][^>]*?>(.*?)<\/a>`)
	match := reg.FindAllStringSubmatch(string(body), -1)
	mapCatch := make(map[string]string)
	for _, v := range match {
		if len(v) == 3 {
			mapCatch[v[2]] = v[1]
		}
	}
	return mapCatch
}

//HTMl标签转小写
func TagToLower(con string) string {
	reg, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	html := reg.ReplaceAllStringFunc(con, strings.ToLower)
	return html
}

//删除html中的style样式
func StripStyle(con string) string {
	reg, _ := regexp.Compile(`<style[^>]+?</style[^>]?>`)
	str := reg.ReplaceAllString(string(con), "")
	return str
}

//删除html中的script脚本
func StripScript(con string) string {
	reg, _ := regexp.Compile(`<script[^>]+?</script[^>]?>`)
	str := reg.ReplaceAllString(string(con), "")
	return str
}

//获取htnl中的title
func GetTitle(con string) string {
	reg, _ := regexp.Compile(`<title[^>]*?>(.*?)<\/title[^>]?>`)
	match := reg.FindStringSubmatch(con)
	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
}

//获取html的body
func GetBody(con string) string {
	reg, _ := regexp.Compile(`<body[^>]*?>(.*?)<\/body[^>]?>`)
	match := reg.FindStringSubmatch(con)
	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
}

//获取html仅保留div和h1标签的结构
func GetDivH1(con string) string {
	reg, _ := regexp.Compile(`<[^>]*?>`)
	html := reg.ReplaceAllStringFunc(con, func(str string) string {
		reg, _ := regexp.Compile(`<div|<h1|</div|</h1`)
		match := reg.MatchString(str)
		if match {
			return str
		} else {
			return ""
		}
	})
	return html
}

//获取html的h1
func GetH1(con string) string {
	reg, _ := regexp.Compile(`<h1[^>]*?>(.*?)<\/h1[^>]?>`)
	match := reg.FindStringSubmatch(con)
	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
}
