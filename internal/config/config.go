package config

import (
    "sync"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/1 上午12:42
 * @Desc:
 */

// Config custom config struct
type configDef struct {
    // 使用全小写
    DBInfo      DBInfo `yaml:"dbinfo"`
    TablePrefix string `yaml:"tableprefix"` // 表前缀
}

// DBInfo mysql database information. mysql 数据库信息
type DBInfo struct {
    Host     string `validate:"required"` // Host. 地址
    Port     int    // Port 端口号
    Username string // Username 用户名
    Password string // Password 密码
    Database string // Database 数据库名
    Type     int    // 数据库类型: 0:mysql , 1:sqlite , 2:mssql
}

var (
    configPath string
    config     *configDef
    once       sync.Once
)

func GetDBInfo() *DBInfo {
    return &config.DBInfo
}

func GetTablePrefix() string {
    return config.TablePrefix
}
