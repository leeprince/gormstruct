package app

import (
    "fmt"
    "github.com/leeprince/gormstruct/internal/config"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/1 下午11:10
 * @Desc:
 */

type dbApp struct {
    info *config.DBInfo
}

var db dbApp

func InitDB() {
    db.info = config.GetDBInfo()
}

// 获取mysql 连接字符串
func GetConfigDBOfMysqlConStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&interpolateParams=True",
		db.info.Username,
		db.info.Password,
		db.info.Host,
		db.info.Port,
		db.info.Database,
	)
}

// 获取连接的数据库
func GetConfigDBDatabase() string {
	return db.info.Database
}

// 获取连接的类型
func GetConfigDBType() int {
	return db.info.Type
}