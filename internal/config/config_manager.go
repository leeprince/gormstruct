package config

import "fmt"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/12/12 下午11:05
 * @Desc:
 */

// 获取连接的类型
func GetConfigDBType() int {
	return config.DBInfo.Type
}

// 获取连接的数据库
func GetConfigDBDatabase() string {
	db := config.DBInfo.Database
	if flagInfo.database != "" {
		db = flagInfo.database
	}
	
	return db
}

// 获取mysql 连接字符串
func GetConfigDBOfMysqlConStr() string {
	
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&interpolateParams=True",
		config.DBInfo.Username,
		config.DBInfo.Password,
		config.DBInfo.Host,
		config.DBInfo.Port,
		GetConfigDBDatabase(),
	)
}

func GetPrimaryIdType() string {
	return config.PrimaryIdType
}

func GetIsGormModelTag() bool {
	// 已强制为false
	// return config.IsGormModelTag
	return false
}

func GetIsNullToPoint() bool {
	return config.IsNullToPoint
}

func GetTable() string {
	return flagInfo.table
}

func GetPackageName() string {
	packageName := config.PackageName
	if flagInfo.packageName != "" {
		packageName = flagInfo.packageName
	}
	return packageName
}

func GetStructName() string {
	return flagInfo.structName
}

func GetOuputDir() string {
	return config.OuputDir
}

func GetSelfTypeDefine() map[string]string {
	return config.SelfTypeDefine
}

func GenDoc() bool {
	return config.GenDoc
}

func GenDeleteFlagFieldList() []string {
	return config.DeleteFlagFieldList
}
