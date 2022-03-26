package genmysql

import (
    "encoding/json"
    "fmt"
    "github.com/leeprince/gopublic/mysqldb"
    "github.com/leeprince/gormstruct/internal/config"
    "github.com/leeprince/gormstruct/internal/constants"
    "github.com/leeprince/gormstruct/internal/logger"
    "github.com/leeprince/gormstruct/internal/model"
    "strings"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/2 上午12:52
 * @Desc:
 */

type mysqlModel struct {
}

// MySQLModel mysql model from IModel
var MySQLModel mysqlModel

func NewMySQLModel() model.IModel {
    return &MySQLModel
}

// GenModel get model.DBInfo info.获取数据库相关属性
func (m *mysqlModel) GenModel() model.DBInfo {
    orm := mysqldb.OnInitDBOrm(config.GetConfigDBOfMysqlConStr())
    defer orm.OnDestoryDB()
    
    var dbInfo model.DBInfo
    m.getTableInfo(orm, &dbInfo)
    dbInfo.PackageName = m.GetPkgName()
    dbInfo.DbName = m.GetDbName()
    
    debugInfo, _ := json.Marshal(dbInfo)
    fmt.Printf("获取数据库相关属性:%s", string(debugInfo))
    return dbInfo
}

// 获取数据库名字
func (m *mysqlModel) GetDbName() string {
    return config.GetConfigDBDatabase()
}

// 获取包名
func (m *mysqlModel) GetPkgName() string {
    return config.GetPackageName()
}

// 获取表信息
func (m *mysqlModel) getTableInfo(orm *mysqldb.MySqlDB, info *model.DBInfo) {
    // 获取表注释
    tableComment, err := m.getTableComment(orm, config.GetTable())
    if err != nil {
        logger.Error("获取表注释错误：", err)
    }
    logger.Info(fmt.Sprintf("获取表信息。config：%s; 表注释：%s", config.GetTable(), tableComment))
    
    // 获取字段信息
    info.Table.Name = config.GetTable()
    info.Table.Comment = tableComment
    info.Table.ColumnsElement, err = m.getTableColumns(orm, info.Table.Name)
}

func (m *mysqlModel) getTableComment(orm *mysqldb.MySqlDB, table string) (string, error) {
    var desc string
    sql := fmt.Sprintf("SELECT TABLE_COMMENT FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s'", m.GetDbName(), table)
    err := orm.Raw(sql).Scan(&desc).Error
    if err != nil {
        return "", err
    }
    
    return desc, nil
}

func (m *mysqlModel) getTableColumns(orm *mysqldb.MySqlDB, table string) (columnsElement []model.ColumnsElementInfo, err error) {
    keyCount := make(map[string]int)
    keyMap := make(map[string][]keys)
    
    // 获取索引信息
    var keys []keys
    keysSql := fmt.Sprintf("SHOW KEYS FROM %s", table)
    orm.Raw(keysSql).Scan(&keys)
    for _, key := range keys {
        keyCount[key.KeyName]++
        keyMap[key.ColumnName] = append(keyMap[key.ColumnName], key)
    }
    
    // 获取字段信息
    var columns []genColumns
    colSql := fmt.Sprintf("SHOW FULL COLUMNS FROM %s", table)
    orm.Raw(colSql).Scan(&columns)
    // 判断数据库设计是否符合 gorm.Model， 符合则定义表存在 gorm.Model
    if checkGormModel(&columns) {
        columnsElement = append(columnsElement, model.ColumnsElementInfo{
            Type: constants.GormModelWord,
        })
    }
    // 获取字段名称/类型/注释/默认值/是否允许null/索引信息(主键唯一索引、普通索引、唯一索引、唯一复合索引、唯一非复合索引)
    for _, i2 := range columns {
        var tmpColumns model.ColumnsElementInfo
        // 字段名称
        tmpColumns.Name = i2.Field
        // 字段类型
        tmpColumns.Type = i2.Type
        
        // 字段注释
        tmpColumns.Comment = i2.Desc
        
        // 主键时，是否为自动递增。自增：Extra==auto_increment
        tmpColumns.Extra = i2.Extra
        
        // 默认值
        if i2.Default != nil {
            if *i2.Default == "" {
                tmpColumns.Default = "DEFAULT ''"
            } else {
                tmpColumns.Default = fmt.Sprintf("DEFAULT '%s'", *i2.Default)
            }
        }
        
        // 是否允许null
        tmpColumns.IsNull = strings.EqualFold(i2.Null, FIELD_IS_NULL)
        
        // 索引信息(主键唯一索引、普通索引、唯一索引、唯一复合索引、唯一非复合索引)
        if keys, ok := keyMap[i2.Field]; ok {
            for _, key := range keys {
                if key.NonUnique == 0 { // 唯一索引：主键、唯一索引、唯一复合索引
                    // 主键
                    if strings.EqualFold(key.KeyName, KEY_NAME_PRIMARY) {
                        tmpColumns.Keys = append(tmpColumns.Keys, model.KeyList{
                            Key:     model.ColumnsKeyPrimary,
                            Multi:   (keyCount[key.KeyName] > 1),
                            KeyName: key.KeyName,
                            KeyFunc: key.IndexType,
                        })
                        continue
                    }
                    // 唯一索引
                    if keyCount[key.KeyName] == 1 {
                        tmpColumns.Keys = append(tmpColumns.Keys, model.KeyList{
                            Key:     model.ColumnsKeyUnique,
                            Multi:   (keyCount[key.KeyName] > 1),
                            KeyName: key.KeyName,
                            KeyFunc: key.IndexType,
                        })
                        continue
                    }
                    // 唯一复合索引
                    tmpColumns.Keys = append(tmpColumns.Keys, model.KeyList{
                        Key:     model.ColumnsKeyUniqueIndex,
                        Multi:   (keyCount[key.KeyName] > 1),
                        KeyName: key.KeyName,
                        KeyFunc: key.IndexType,
                    })
                } else { // 非唯一索引。包含：普通索引｜普通复合索引
                    tmpColumns.Keys = append(tmpColumns.Keys, model.KeyList{
                        Key:     model.ColumnsKeyIndex,
                        Multi:   (keyCount[key.KeyName] > 1),
                        KeyName: key.KeyName,
                        KeyFunc: key.IndexType,
                    })
                }
            }
        }
        columnsElement = append(columnsElement, tmpColumns)
    }
    
    return
}
