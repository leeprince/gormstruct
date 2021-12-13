package config

import "fmt"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/12/12 下午11:05
 * @Desc:
 */
func GetDBInfo() *DBInfo {
    return &config.DBInfo
}
// 获取连接的类型
func GetConfigDBType() int {
	return config.DBInfo.Type
}
// 获取连接的数据库
func GetConfigDBDatabase() string {
	return config.DBInfo.Database
}
// 获取mysql 连接字符串
func GetConfigDBOfMysqlConStr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&interpolateParams=True",
		config.DBInfo.Username,
		config.DBInfo.Password,
		config.DBInfo.Host,
		config.DBInfo.Port,
		config.DBInfo.Database,
	)
}
func GetGenBaseFuncVersion() string {
    return config.GenBaseFuncVersion
}
func GetGenLogicFuncVersion() string {
    return config.GenLogicFuncVersion
}
func GetPrimaryIdType() string {
    return config.PrimaryIdType
}
func GetIsGormModelTag() bool {
    return config.IsGormModelTag
}
func GetIsNullToPoint() bool {
    return config.IsNullToPoint
}