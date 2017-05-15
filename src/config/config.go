package config

func Get(set string) string {
	return Config[set]
}

var Config = map[string]string{
	"ROOT_PATH":     "/home/www/go/webgo/", //安装目录(绝对路径)
	"MYSQL_HOST":    "127.0.0.1",
	"MYSQL_NAME":    "root",
	"MYSQL_PASS":    "",
	"MYSQL_DB":      "godb",
	"MYSQL_CHARSET": "utf8",
}
