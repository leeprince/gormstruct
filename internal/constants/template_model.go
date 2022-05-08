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

// 实例化结构体对象
func New{{.StructName}}() *{{.StructName}} {
	return &{{.StructName}}{}
}
`
    TempGenTableField = `
{{$obj := .}}
// 表字段的映射
var {{$obj.Name}}Columns = struct {
{{range $genElement := $obj.Elments}}{{$genElement.Name}} string
{{end}}
} {
{{range $genElement := $obj.Elments}}{{$genElement.Name}}: "{{$genElement.ColumnName}}",
{{end}}
}

// 包含所有表字段的切片
var {{$obj.Name}}AllColumns = []string{
    {{range $genElement := $obj.Elments}}{{$obj.Name}}Columns. {{$genElement.Name}}, // {{$genElement.Comment}}
    {{end}}
}

{{range $genElement := $obj.Elments}}
// 设置：{{$genElement.Comment}}
func (m *{{$obj.Name}}) Set{{$genElement.Name}}(v {{$genElement.Type}}) {
	m.{{$genElement.Name}} = v
}
{{end}}
`
)
