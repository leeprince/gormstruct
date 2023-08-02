package genfunc

import (
	"github.com/leeprince/gormstruct/internal/constants"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/14 下午6:34
 * @Desc:
 */

// 生成数据库表名方法
func GetGenTableNameTemp() string {
	return constants.TempGenTableName
}

// 生成表字段相关的方法。1.字段映射；2.主键相关
func GetGenTableFieldTemp() string {
	return constants.TempGenTableField
}

// 生成操作数据表的基本方法
func GetGenBaseTemp() string {

	return constants.TempGenBaseV4
}

// 生成操作数据的方法
func GetGenLogic() string {
	return constants.TempGenLogicV4
}
