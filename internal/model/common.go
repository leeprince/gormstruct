package model

import (
	"fmt"
	infrautils "github.com/leeprince/goinfra/utils"
	"github.com/leeprince/gopublic/mybigcamel"
	"github.com/leeprince/gopublic/tools"
	"github.com/leeprince/gormstruct/internal/config"
	"github.com/leeprince/gormstruct/internal/constants"
	"github.com/leeprince/gormstruct/internal/utils"
	"regexp"
	"strings"
	"text/template"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/22 下午10:08
 * @Desc:
 */

// getTypeName Type acquisition filtering.类型获取过滤
func getTypeName(sourceColumnName, sourceTypeName string, isNull bool) string {
	// 优先匹配自定义类型
	selfDefineTypeMqlDicMap := config.GetSelfTypeDefine()
	if v, ok := selfDefineTypeMqlDicMap[sourceTypeName]; ok {
		return fixNullToPoint(sourceColumnName, v, isNull)
	} else { // 如果精确匹配不到比如tinyint(4) 就找tinyint
		newSourceTypeName := strings.Split(sourceTypeName, "(")[0]
		if v1, ok1 := selfDefineTypeMqlDicMap[newSourceTypeName]; ok1 {
			return fixNullToPoint(newSourceTypeName, v1, isNull)
		}
	}

	// Precise matching first.先精确匹配
	if v, ok := constants.TypeMysqlDicMp[sourceTypeName]; ok {
		return fixNullToPoint(sourceColumnName, v, isNull)
	}

	// Fuzzy Regular Matching.模糊正则匹配
	for _, l := range constants.TypeMysqlMatchList {
		if ok, _ := regexp.MatchString(l.Key, sourceTypeName); ok {
			return fixNullToPoint(sourceColumnName, l.Value, isNull)
		}
	}

	panic(fmt.Sprintf("type (%v) not match in any way.maybe need to add on (https://github.com/xxjwxc/gormt/blob/master/data/view/cnf/def.go)", sourceTypeName))
}

func fixNullToPoint(sourceColumnName, destTypeName string, isNull bool) string {
	if isNull && config.GetIsNullToPoint() {
		if strings.HasPrefix(destTypeName, constants.NullTypeUint) ||
			strings.HasPrefix(destTypeName, constants.NullTypeInt) ||
			strings.HasPrefix(destTypeName, constants.NullTypeFloat) ||
			strings.HasPrefix(destTypeName, constants.NullTypeString) {
			return pointType(destTypeName)
		}
		// 删除字段使用：*time.Time，而不是 time.Time, 避免初始化变量或者数据json转化后使用time.Time的默认值：`0001-01-01 00:00:00 +0000` 插入数据，导致查询删除字段是否为空错误
		if infrautils.InString(sourceColumnName, config.GenDeleteFlagFieldList()) &&
			strings.Contains(destTypeName, constants.ColumnTypeTimeContains) {
			return pointType(destTypeName)
		}
	}
	return destTypeName
}

func pointType(typeName string) string {
	return "*" + typeName
}

func getGormModelElement() []EmInfo {
	var result []EmInfo
	result = append(result, EmInfo{
		IsMulti:       false,
		Notes:         "主键",
		Type:          config.GetPrimaryIdType(), // Type.类型标记
		ColName:       "id",
		ColNameEx:     "id",
		ColStructName: "ID",
	})

	result = append(result, EmInfo{
		IsMulti:       true,
		Notes:         "创建时间",
		Type:          "time.Time", // Type.类型标记
		ColName:       "created_at",
		ColNameEx:     "created_at",
		ColStructName: "CreatedAt",
	})

	result = append(result, EmInfo{
		IsMulti:       true,
		Notes:         "更新时间",
		Type:          "time.Time", // Type.类型标记
		ColName:       "updated_at",
		ColNameEx:     "updated_at",
		ColStructName: "UpdatedAt",
	})

	result = append(result, EmInfo{
		IsMulti:       true,
		Notes:         "删除时间",
		Type:          "time.Time", // Type.类型标记
		ColName:       "deleted_at",
		ColNameEx:     "deleted_at",
		ColStructName: "DeletedAt",
	})
	return result
}

// 构建方法
func buildFList(list *[]FList, tableStructName string, key ColumnsKey, keyName string, tp string, columnName string) {
	for i := 0; i < len(*list); i++ {
		if (*list)[i].KeyName == keyName {
			(*list)[i].KeyNameEl = append((*list)[i].KeyNameEl, KeyInfo{
				Type:          tp,
				ColName:       columnName,
				ColStructName: utils.GetCamelName(columnName),
			})
			return
		}
	}
	// 还未添加过,则添加
	keyName = strings.ReplaceAll(keyName, " ", "_")
	*list = append(*list, FList{
		TableStructName: tableStructName,
		Key:             key,
		KeyName:         keyName,
		KeyNameEl: []KeyInfo{KeyInfo{
			Type:          tp,
			ColName:       columnName,
			ColStructName: utils.GetCamelName(columnName),
		}},
	})
}

// 避免注释中包含换行符
func fixNotes(str string) string {
	return strings.Replace(str, "\n", "\n//", -1)
}

// 过滤 golang 关键字：转换名称
func FilterKeywords(src string) string {
	if tools.IsKeywords(src) {
		return "_" + src
	}
	return src
}

// GenFListIndex 生成list
// status(1:获取函数名,2:获取参数列表,3:获取sql case,4:值列表,5:字段列表字符串,6:参数列表的对应obj的多个 WithXxx 方法)
func GenFListIndex(info FList, status int) string {
	switch status {
	case 1: // 1:获取函数名
		return widthFunctionName(info)
	case 2: // 2:获取参数列表字符串
		var strs []string
		for _, v := range info.KeyNameEl {
			strs = append(strs, fmt.Sprintf("%v %v ", CapLowercase(v.ColStructName), v.Type))
		}
		return strings.Join(strs, ", ")
	case 3: // 3:获取表结构体赋值的定义格式
		var sturctStr string
		for _, v := range info.KeyNameEl {
			sturctStr = fmt.Sprintf("%s %s: %s,", sturctStr, v.ColStructName, CapLowercase(v.ColStructName))
		}
		return fmt.Sprintf("%s{ %s }", info.TableStructName, sturctStr)
	case 4: // 4:值列表字符串
		var strs []string
		for _, v := range info.KeyNameEl {
			strs = append(strs, CapLowercase(v.ColStructName))
		}
		return strings.Join(strs, ", ")
	case 5: // 5: 字段列表字符串
		var strs []string
		for _, v := range info.KeyNameEl {
			strs = append(strs, CapLowercase(v.ColName))
		}
		return strings.Join(strs, ", ")
	case 6: // 6: 参数列表的对应obj的多个 WithXxx 方法
		var withFuncs []string
		for _, v := range info.KeyNameEl {
			withFuncs = append(withFuncs, fmt.Sprintf("obj.With%s(%s)", v.ColStructName, CapLowercase(v.ColStructName)))
		}
		if len(withFuncs) == 1 {
			return strings.Join(withFuncs, ", ")
		}

		return fmt.Sprintf("\n %s", strings.Join(withFuncs, ", \n"))
	}

	return ""
}

func widthFunctionName(info FList) string {
	switch info.Key {
	// case ColumnsKeyDefault:
	case ColumnsKeyPrimary: // primary key.主键
		return "FetchByPrimaryKey"
	case ColumnsKeyUnique: // unique key.唯一索引
		return "FetchUniqueBy" + utils.GetCamelName(info.KeyName)
	case ColumnsKeyIndex: // index key.复合索引
		return "FetchIndexBy" + utils.GetCamelName(info.KeyName)
	case ColumnsKeyUniqueIndex: // unique index key.唯一复合索引
		return "FetchUniqueIndexBy" + utils.GetCamelName(info.KeyName)
	}

	return ""
}

// CapLowercase 小写.且兼容 golint 驼峰命名规则
func CapLowercase(name string) string { // IDAPIID == > idAPIID
	list := strings.Split(mybigcamel.UnMarshal(name), "_")
	if len(list) == 0 {
		return ""
	}

	re := list[0] + name[len(list[0]):]

	return FilterKeywords(re)
}

// 定义表名的结构体
func GetTableStructName(name string) string {
	return name + "{}"
}

// 已表名定义的结构体
func GetCurrentDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 生成格式需要的引用格式
func GetQuoteFormatValue() string {
	return "`%v`"
}

// 基础方法的模版需要的解析的方法
func GenBaseTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"GetVV":              GetQuoteFormatValue,
		"GetCurrentDateTime": GetCurrentDateTime,
	}
}

// 逻辑方法的模版需要的解析的方法
func GenLogicTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"GenFListIndex":      GenFListIndex,
		"CapLowercase":       CapLowercase,
		"GetTableStructName": GetTableStructName,
		"GetCurrentDateTime": GetCurrentDateTime,
	}
}
