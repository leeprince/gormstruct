package constants

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/14 下午6:33
 * @Desc:
 */

const (
    TempGenTableName = `
// 获取结构体对应的表名方法
func (m *{{.StructName}}) TableName() string {
	return "{{.TableName}}"
}
`
    TempGenTableField = `
// 表字段的映射
var {{.Name}}Columns = struct {
{{range $genElement := .Elments}}{{$genElement.Name}} string
{{end}}
} {
{{range $genElement := .Elments}}{{$genElement.Name}}: "{{$genElement.ColumnName}}",
{{end}}
}
`
)
