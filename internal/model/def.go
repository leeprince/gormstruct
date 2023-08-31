package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/2 上午12:56
 * @Desc:
 */

// DBInfo database default info
type DBInfo struct {
	DbName      string    // database name
	PackageName string    // package name
	Table       TableInfo // 表
}

// TabInfo database table default attribute
type TableInfo struct {
	BaseInfo
	ColumnsElement []ColumnsElementInfo // 字段属性信息
}

// BaseInfo base common attribute. 基础属性
type BaseInfo struct {
	Name    string // 名称
	Comment string // 注释
}

// ColumnsInfo Columns list .表列信息
type ColumnsElementInfo struct {
	BaseInfo
	IsNull  bool      // 是否为空
	Type    string    // 字段的类型
	Default string    // 默认值
	Extra   string    // 主键时，是否为自动递增。自增：Extra==auto_increment
	Keys    []KeyList // 包含该字段的索引列表
}

// 主键唯一索引、普通索引、唯一索引、唯一复合索引、唯一非复合索引
type KeyList struct {
	Key     ColumnsKey // 索引类型
	Multi   bool       // 是否多个(复合组建).主键也可能是联合主键(一个表有多个主键，即联合主键)
	KeyName string     // 索引名称
	KeyFunc string     // 索引方法
}

// ColumnsKey Columns type elem. 类型枚举
type ColumnsKey int

const (
	// ColumnsKeyDefault default
	ColumnsKeyDefault = iota
	// ColumnsKeyPrimary primary key.主键
	ColumnsKeyPrimary
	// ColumnsKeyUnique unique key.唯一索引
	ColumnsKeyUnique
	// ColumnsKeyUniqueIndex unique index key.唯一联合索引
	ColumnsKeyUniqueIndex
	// ColumnsKeyIndex index key.普通索引,包含普通联合索引
	ColumnsKeyIndex
)

// 生成的文件列表
type GenOutInfo struct {
	FileName string // output file name .输出文件名
	
	// 输出的方式：1. 动态生成行并输出；2. 使用模版匹配输出
	FileCtx string // output file context. 输出文件内容.
}

type funDef struct {
	StructName            string
	TableName             string
	Em                    []EmInfo // 字段列表
	Primary               []FList  // 唯一索引（主键、唯一索引、唯一联合索引）方法列表
	Index                 []FList  // 普通索引的方法列表
	IsHaveDeleteFlag      bool     // 行记录是否存在删除的字段
	DeleteFlagIsTime      bool     // 行记录是否存在删除的字段的字段类型是否为time.Time的数据类型（包含字段类型：date,datatime,timestamp)）
	DeleteFlagStructField string   // 存在删除的字段的结构名称
}

// EmInfo element of func info
type EmInfo struct {
	IsMulti       bool   // 一个字段值是否存在多条记录
	Notes         string // 注释
	Type          string // 类型
	ColName       string // 列名
	ColNameEx     string // `列名`
	ColStructName string // 列结构体
}

// FList index of list
type FList struct {
	TableStructName string     // 表名的结构体
	Key             ColumnsKey // 索引的类型
	KeyName         string     // 索引名（键名）。一个键名可能包含多个字段组成复合索引
	KeyNameEl       []KeyInfo  // 索引名（键名）包含的多个字段信息列表
}

// 索引字段信息
type KeyInfo struct {
	Type          string // 类型
	ColName       string // 列名
	ColStructName string // 列结构体
}
