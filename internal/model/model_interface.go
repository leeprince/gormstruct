package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/2 上午12:24
 * @Desc:
 */

// IModel Implement the interface to acquire database information and initialize it.实现接口获取数据库信息获取并初始化
type IModel interface {
	GenModel() DBInfo
	GetDbName() string
	GetPkgName() string    // 生成 gorm struct 所在的包名
}
