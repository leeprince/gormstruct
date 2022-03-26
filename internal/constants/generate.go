package constants

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/2 上午12:59
 * @Desc:
 */

const (
    TagDb           = "gorm"
    TagJson         = "json"
    PrimaryKeyExtra = "auto_increment"
)

const (
    GenGormtagPrimary = "%s:\"column:%s;primaryKey;%s\"" // gorm:"primaryKey;column:字段名;是否自动递增"
    GenGormtagColumun = "%s:\"column:%s;type:%s;%s;%s\"" // gorm:"column:字段名;type:字段类型;是否允许为空;default ''"
    GenJson           = "%s:\"%s\""
)

const (
    GormAutoIncrement = "autoIncrement" // gorm 自动自增的标记
    GormModelWord     = "gorm.Model"
)

const (
    FilePackage = "package"
    FileImport  = "import"
    FileStruct  = "struct"
)

var ImportFile = map[string]string{
    "time.Time":      `"time"`,
    "gorm.Model":     `"gorm.io/gorm"`,
    "fmt":            `"fmt"`,
    "datatypes.JSON": `"gorm.io/datatypes"`,
    "datatypes.Date": `"gorm.io/datatypes"`,
}

const (
    NullTypeUint   = "uint"
    NullTypeInt    = "int"
    NullTypeFloat  = "float"
    NullTypeString = "string"
)

// 精确匹配类型
var TypeMysqlDicMp = map[string]string{
    "smallint":            "int16",
    "smallint unsigned":   "uint16",
    "int":                 "int",
    "int unsigned":        "uint",
    "bigint":              "int64",
    "bigint unsigned":     "uint64",
    "mediumint":           "int32",
    "mediumint unsigned":  "uint32",
    "varchar":             "string",
    "char":                "string",
    "date":                "datatypes.Date",
    "datetime":            "time.Time",
    "bit(1)":              "[]uint8",
    "tinyint":             "int8",
    "tinyint unsigned":    "uint8",
    "tinyint(1)":          "bool", // tinyint(1) 默认设置成bool
    "tinyint(1) unsigned": "bool", // tinyint(1) 默认设置成bool
    "json":                "datatypes.JSON",
    "text":                "string",
    "timestamp":           "time.Time",
    "double":              "float64",
    "double unsigned":     "float64",
    "mediumtext":          "string",
    "longtext":            "string",
    "float":               "float32",
    "float unsigned":      "float32",
    "tinytext":            "string",
    "enum":                "string",
    "time":                "time.Time",
    "tinyblob":            "[]byte",
    "blob":                "[]byte",
    "mediumblob":          "[]byte",
    "longblob":            "[]byte",
    "integer":             "int64",
}

// TypeMysqlMatchList Fuzzy Matching Types.模糊匹配类型
var TypeMysqlMatchList = []struct {
    Key   string
    Value string
}{
    {`^(tinyint)[(]\d+[)] unsigned`, "uint8"},
    {`^(smallint)[(]\d+[)] unsigned`, "uint16"},
    {`^(int)[(]\d+[)] unsigned`, "uint32"},
    {`^(bigint)[(]\d+[)] unsigned`, "uint64"},
    {`^(float)[(]\d+,\d+[)] unsigned`, "float64"},
    {`^(double)[(]\d+,\d+[)] unsigned`, "float64"},
    {`^(tinyint)[(]\d+[)]`, "int8"},
    {`^(smallint)[(]\d+[)]`, "int16"},
    {`^(int)[(]\d+[)]`, "int"},
    {`^(bigint)[(]\d+[)]`, "int64"},
    {`^(char)[(]\d+[)]`, "string"},
    {`^(enum)[(](.)+[)]`, "string"},
    {`^(varchar)[(]\d+[)]`, "string"},
    {`^(varbinary)[(]\d+[)]`, "[]byte"},
    {`^(blob)[(]\d+[)]`, "[]byte"},
    {`^(binary)[(]\d+[)]`, "[]byte"},
    {`^(decimal)[(]\d+,\d+[)]`, "float64"},
    {`^(mediumint)[(]\d+[)]`, "int16"},
    {`^(mediumint)[(]\d+[)] unsigned`, "uint16"},
    {`^(double)[(]\d+,\d+[)]`, "float64"},
    {`^(float)[(]\d+,\d+[)]`, "float64"},
    {`^(datetime)[(]\d+[)]`, "time.Time"},
    {`^(bit)[(]\d+[)]`, "[]uint8"},
    {`^(text)[(]\d+[)]`, "string"},
    {`^(integer)[(]\d+[)]`, "int"},
    {`^(timestamp)[(]\d+[)]`, "time.Time"},
    {`^(geometry)[(]\d+[)]`, "[]byte"},
}
