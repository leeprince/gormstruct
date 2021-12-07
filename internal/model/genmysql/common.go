package genmysql

import "strings"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/4 下午11:21
 * @Desc:
 */

func checkGormModel(columns *[]genColumns) bool {
    var gormFieldCount int
    var tmpNotGormFields []genColumns
    for _, i2 := range *columns {
        if strings.EqualFold(i2.Field, GORM_MODEL_FIELD_ID)  ||
            strings.EqualFold(i2.Field, GORM_MODEL_FIELD_CREATE) ||
            strings.EqualFold(i2.Field, GORM_MODEL_FIELD_UPDATE) ||
            strings.EqualFold(i2.Field, GORM_MODEL_FIELD_DELETE) {
            gormFieldCount++
        } else {
            tmpNotGormFields = append(tmpNotGormFields, i2)
        }
    }
    
    if gormFieldCount >= 4 {
        columns = &tmpNotGormFields
        return true
    }
    return false
}