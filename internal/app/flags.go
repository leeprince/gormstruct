package app

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/4 上午12:33
 * @Desc:
 */

type flags struct {
    table string
    packageName string
    structName string
}

var flagInfo flags

func InitFlags(table, packageName, structName string)  {
    flagInfo.table = table
    flagInfo.packageName = packageName
    flagInfo.structName = structName
}

func GetTable() string {
    return flagInfo.table
}

func GetPackageName() string {
    return flagInfo.packageName
}

func GetStructName() string {
    return flagInfo.structName
}

