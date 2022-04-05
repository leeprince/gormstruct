package gendoc

import (
    "fmt"
    "github.com/leeprince/gormstruct/internal/config"
    "sync"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/4 下午2:30
 * @Desc:
 */

var onceColumnLine sync.Once
var onceKeyLine sync.Once
var columnLine []string
var keyLine []string

func AddColumnLine(str string) {
    onceColumnLine.Do(func() {
        if config.GenDoc() {
            columnLine = append(columnLine, fmt.Sprintf("字段名\\t字段类型\\t是否NULL\\t默认值\\t注释"))
        }
    })
    columnLine = append(columnLine, str)
}
func AddKeyLine(str string) {
    onceKeyLine.Do(func() {
        if config.GenDoc() {
            keyLine = append(keyLine, fmt.Sprintf("\n是否唯一索引\\t索引名称\\t索引字段\\t"))
        }
    })
    keyLine = append(keyLine, str)
}

func FmtGenDoc(table string) {
    if !config.GenDoc() {
        return
    }
    
    fmt.Printf("---------------------------生成数据库 %s 表的文档----------------------\n", table)
    for _, i2 := range columnLine {
        fmt.Println(i2)
    }
    
    for _, i3 := range keyLine {
        fmt.Println(i3)
    }
    fmt.Printf("---------------------------生成数据库 %s 表的文档----------------------\n", table)
}
