/**
  配置信息
  定义的配置信息使用小写字母，保持对外私有
*/
package config

import (
	"lib/db"
)

var Config = map[string]string{
	"ROOT_PATH": "/home/www/go/goSpider/", //安装目录(绝对路径)
}

//获取配置信息的
func Get(set string) string {
	return Config[set]
}

//数据库链接（单例模式）
type dbStru struct {
	dbcsn db.Mysql
}

var dbCspi *dbStru

func DbSpider() db.Mysql {
	if dbCspi == nil {
		var spiderConn db.Mysql
		spider_conf := make(map[string]string)
		spider_conf["dbhost"] = "tcp(9.9.9.9:3306)"
		spider_conf["dbuser"] = "root"
		spider_conf["dbpass"] = "123456"
		spider_conf["dbname"] = "spider_db"
		spider_conf["charset"] = "utf8"
		spiderConn.GetConn(spider_conf)
		dbCspi = &dbStru{}
		dbCspi.dbcsn = spiderConn
	}
	return dbCspi.dbcsn
}
