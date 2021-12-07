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
    TableName string       // table_name.表名
    Name      string       // name.名字
    Comment   string       // notes.注释
    Elments   []GenElement // em.元素组合
}

// GenElement element of sturct.元素类
type GenElement struct {
    Name       string // Name.元素名
    ColumnName string // table name.表名
    Type       string // Type.类型标记
    Comment    string // Notes.注释
    TagString  string // tages.标记
}
