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
    return constants.TempGenTableName
}

// 生成表字段映射结构体
func GetGenTableFieldTemp() string {
    return constants.TempGenTableField
}

// 生成操作数据表的基本方法
func GetGenBaseTemp() string {
    switch config.GetGenBaseFuncVersion() {
    case "V1":
        return constants.TempGenBaseV1
    case "V2":
        return constants.TempGenBaseV2
    case "V3":
        return constants.TempGenBaseV3
    case "V4":
        return constants.TempGenBaseV4
    }
    return constants.TempGenBaseV4
}

// 生成操作数据的方法
func GetGenLogic() string {
    switch config.GetGenLogicFuncVersion() {
    case "V1":
        return constants.TempGenLogicV1
    case "V2":
        return constants.TempGenLogicV2
    case "V3":
        return constants.TempGenLogicV3
    case "V4":
        return constants.TempGenLogicV4
    }
    return constants.TempGenLogicV4
}

