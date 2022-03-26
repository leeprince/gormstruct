package genstruct

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/9 下午11:44
 * @Desc:
 */

// GenPackage package of IPackage.包体
type GenPackage struct {
    PackageName string    // 包名字
    Imports     []string  // 引入的包
    Struct      GenStruct // 表的结构体
    FuncStrList []string  // 获取表数据的方法
}

// GenStruct struct of IStruct .结构体
type GenStruct struct {
    Name      string       // 名字（结构体名）
    TableName string       // 表名
    Comment   string       // 表注释
    Elments   []GenElement // 元素组合
}

// GenElement element of sturct.元素类
type GenElement struct {
    Name         string // 元素名（结构体名）
    ColumnName   string // 字段名
    Type         string // 字段类型标记(输出结构体对应的字段类型)
    Comment      string // 字段注释
    TagString    string // 标记
}
