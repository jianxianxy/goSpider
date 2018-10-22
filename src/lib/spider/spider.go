package spider

import (
	"io/ioutil"
	"lib/mahonia"
	"net/http"
	"regexp"
	"strings"
)

//根据url获取页面相关信息
func Exec(url string) (map[string]string, map[string]string) {
	//获取HTML
	info := make(map[string]string)
	href := make(map[string]string)
	html, err := GetHtmlByUrl(url)
	if err == nil {
		html = TagToLower(html)       //HTML标签转小写
		href = GetUrlFromString(html) //提取页面的超链接

		body := GetBody(StripNote(StripStyle(StripScript(html)))) //获取Body并去除注释、样式、脚本
		info["body"] = body
		title := GetH1(body) //获取h1内容(文章标题)
		if len(title) > 0 {
			info["title"] = title
		}
	}
	return info, href
}

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
	reg := regexp.MustCompile(`<a[^>]*?href\s?=\s?['"](http[^'"]*?)['"][^>]*?>([\p{Han}]+.*?)<\/a>`)
	match := reg.FindAllStringSubmatch(string(body), -1)
	mapCatch := make(map[string]string)
	for _, v := range match {
		if len(v) == 3 {
			mapCatch[strings.TrimSpace(v[2])] = strings.TrimSpace(v[1])
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
	reg, _ := regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	str := reg.ReplaceAllString(string(con), "")
	return str
}

//删除html中的script脚本
func StripScript(con string) string {
	reg, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	str := reg.ReplaceAllString(string(con), "")
	return str
}

//删除html中的注释
func StripNote(con string) string {
	reg, _ := regexp.Compile(`<!-[.\s\S]*?->`)
	str := reg.ReplaceAllString(string(con), "")
	return str
}

//获取html中的body
func GetBody(con string) string {
	reg, _ := regexp.Compile(`<body[^>]*?>([.\s\S]*?)<\/body[^>]?>`)
	match := reg.FindStringSubmatch(con)
	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
}

//根据条件匹配
func FindByIC(con, tag, attr string) string {
	html := string([]rune(con))
	var model, mval string
	if attr != "" {
		model = string([]rune(attr[0:1]))
		mval = string([]rune(attr[1:]))
	}
	var regStr string
	if model == "#" {
		regStr = `<` + tag + `[^>]+id[^>]+` + mval + `[^>]*>`
	} else if model == "." {
		regStr = `<` + tag + `[^>]+class[^>]+` + mval + `[^>]*>`
	} else {
		regStr = `<` + tag + `[^>]*>`
	}
	reg, _ := regexp.Compile(regStr)
	loc := reg.FindStringIndex(html)
	if len(loc) < 1 {
		return ""
	}
	loop := len(html)
	var can = 0 //内含同名标签标识，等于0才可以截取
	for i := loc[1]; i < loop; i++ {
		if string(html[i]) == "<" && string(html[i+1]) == "/" {
			beg := i + 2
			end := beg + len(tag)
			if end > loop {
				end = loop
			}
			cur := string(html[beg:end])
			if can == 0 && cur == tag {
				for ie := end; ie < loop; ie++ {
					if string(html[ie]) == ">" {
						return html[loc[0] : ie+1]
					}
				}
			}
			if cur == tag && can > 0 {
				can = can - 1
			}
		} else if string(html[i]) == "<" {
			beg := i + 1
			end := beg + len(tag)
			if end > loop {
				end = loop
			}
			cur := string(html[beg:end])
			if cur == tag {
				can = can + 1
			}
		}
	}
	return ""
}

//根据条件匹配全部
func FindByICA(con, tag, attr string, matRet *[]string) int {
	html := string([]rune(con))
	var model, mval string
	if attr != "" {
		model = string([]rune(attr[0:1]))
		mval = string([]rune(attr[1:]))
	}
	var regStr string
	if model == "#" {
		regStr = `<` + tag + `[^>]+id[^>]+` + mval + `[^>]*>`
	} else if model == "." {
		regStr = `<` + tag + `[^>]+class[^>]+` + mval + `[^>]*>`
	} else {
		regStr = `<` + tag + `[^>]*>`
	}
	reg, _ := regexp.Compile(regStr)
	loc := reg.FindStringIndex(html)
	if len(loc) < 1 {
		return 0
	}
	loop := len(html)
	var can = 0 //内含同名标签标识，等于0才可以截取
	var next int
	for i := loc[1]; i < loop; i++ {
		if string(html[i]) == "<" && string(html[i+1]) == "/" {
			beg := i + 2
			end := beg + len(tag)
			if end > loop {
				end = loop
			}
			cur := string(html[beg:end])
			if can == 0 && cur == tag {
				for ie := end; ie < loop; ie++ {
					if string(html[ie]) == ">" {
						next = ie + 1
						*matRet = append(*matRet, html[loc[0]:next])
						FindByICA(string(html[next:]), tag, attr, matRet)
						return 1
					}
				}
			}
			if cur == tag && can > 0 {
				can = can - 1
			}
		} else if string(html[i]) == "<" {
			beg := i + 1
			end := beg + len(tag)
			if end > loop {
				end = loop
			}
			cur := string(html[beg:end])
			if cur == tag {
				can = can + 1
			}
		}
	}
	return 1
}

//获取标签中的文本
func GetText(html string) string {
	reg, _ := regexp.Compile(`<[^>]*?>`)
	return reg.ReplaceAllString(html, "")
}

//获取html中的title
func GetTitle(con string) string {
	reg, _ := regexp.Compile(`<title[^>]*?>(.*?)<\/title[^>]?>`)
	match := reg.FindStringSubmatch(con)
	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
}

//获取html仅保留div和h1标签的结构[用于分析正文位置]
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

//获取html的图片
func GetImg(con string) []string {
	reg, _ := regexp.Compile(`<img[^>]*src=["']([^"']*)["'][^>]*>`)
	ret := make([]string, 0)
	match := reg.FindAllStringSubmatch(con, -1)
	for _, v := range match {
		ret = append(ret, v[1])
	}
	return ret
}

//转码
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

//去除换行和空格
func TrimSpace(str string) string {
	reg, _ := regexp.Compile(`\s`)
	return reg.ReplaceAllString(str, "")
}
