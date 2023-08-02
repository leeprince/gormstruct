package test

import (
	"fmt"
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/2 14:06
 * @Desc:
 */

var mysqlDns = "root:root@tcp(127.0.0.1:3306)/tmp?charset=utf8&parseTime=true&loc=Local&interpolateParams=True"

const IsDebug = true

func InitLooger() {
	// err := plog.SetOutputFile("./logs", "gorm.log", false)
	err := plog.SetOutputFile("./logs", "gorm.log", true)
	if err != nil {
		panic(fmt.Sprintf("plog.SetOutputFile err:%v", err))
	}
	// plog.SetReportCaller(true)

	plog.Debugln("0001", "0002")
	plog.WithFiledLogID(utils.UniqID()).WithField("key01", "v001").Debugln("0001", "0002")
}

var logWriterStdout = log.New(os.Stdout, "\r\n", log.LstdFlags) // io writer（日志输出的目标，前缀和日志包含的内容——译者注）

var (
	DBLogger = logger.New(
		logWriterStdout, // 标准输出
		//plog.GetLogger(), // 指定日志文件输出

		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			Colorful:                  true,        // 彩色打印
			IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
			LogLevel:                  logger.Info, // 日志级别
		},
	)
)

func InitDB() *gorm.DB {
	InitLooger()

	mysqlConnDns := mysqlDns

	// --- gorm.io/gorm
	// 需 import  "gorm.io/driver/mysql","gorm.io/gorm"
	db, err := gorm.Open(mysql.Open(mysqlConnDns), &gorm.Config{
		PrepareStmt:              false,
		Logger:                   DBLogger,
		DisableNestedTransaction: false,
	})
	if err != nil {
		panic("gorm open err:" + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("db.DB() err:" + err.Error())
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if IsDebug {
		return db.Debug()
	}
	// --- gorm.io/gorm -end

	// --- 	github.com/jinzhu/gorm
	// > 注意：使用 github.com/jinzhu/gorm 替换 gorm.io/gorm 需先阅读：base_dao.go 文件的 @Desc 说明
	// 连接需 import "github.com/jinzhu/gorm", _ "github.com/jinzhu/gorm/dialects/mysql"

	/*mysqlConnDns := mysqlDns

	  db, err := gorm.Open("mysql", mysqlConnDns)
	  if err != nil {
	      panic("gorm open err:" + err.Error())
	  }
	  sqlDB := db.DB()

	  // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	  sqlDB.SetMaxIdleConns(10)
	  // SetMaxOpenConns 设置打开数据库连接的最大数量。
	  sqlDB.SetMaxOpenConns(100)
	  // SetConnMaxLifetime 设置了连接可复用的最大时间。
	  sqlDB.SetConnMaxLifetime(time.Hour)
	*/
	// --- 	github.com/jinzhu/gorm -end

	return db

}
