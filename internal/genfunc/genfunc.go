package genfunc

import (
    "github.com/leeprince/gormstruct/internal/config"
    "github.com/leeprince/gormstruct/internal/constants"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/14 下午6:34
 * @Desc:
 */

// 生成数据库表名方法
func GetGenTableNameTemp() string {
    return constants.TEMP_GENTABLENAME
}

// 生成操作数据表的基本方法
func GetGenBaseTemp() string {
    switch config.GetGenBaseFuncVersion() {
    case "V1":
        return constants.TEMP_GENBASE_V1
    case "V2":
        return constants.TEMP_GENBASE_V2
    }
    return constants.TEMP_GENBASE_V2
}

// 生成操作数据的方法
func GetGenLogic() string {
    switch config.GetGenLogicFuncVersion() {
    case "V1":
        return constants.TEMP_GENLOGIC_V1
    case "V2":
        return constants.TEMP_GENLOGIC_V2
    }
    return constants.TEMP_GENLOGIC_V2
}
