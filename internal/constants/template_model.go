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

{{$primaryKey := ""}}
{{$primaryKeyType := ""}}
{{range $genElement := $obj.Elments}}
{{if $genElement.IsPrimaryKey}}
{{$primaryKey = $genElement.Name}}
{{$primaryKeyType = $genElement.Type}}
{{end}}
{{end}}
// 获取主键的对应字段
func (m *{{$obj.Name}}) PrimaryKey() string {
	return {{$obj.Name}}Columns.{{$primaryKey}}
}

// 获取主键值
func (m *{{$obj.Name}}) PrimaryKeyValue() {{$primaryKeyType}} {
	return m.{{$primaryKey}}
}

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
// 设置值：{{$genElement.Comment}}
func (m *{{$obj.Name}}) Set{{$genElement.Name}}(v {{$genElement.Type}}) {
	m.{{$genElement.Name}} = v
}
{{end}}

{{range $genElement := $obj.Elments}}
// 获取值：{{$genElement.Comment}}
func (m *{{$obj.Name}}) Get{{$genElement.Name}}() {{$genElement.Type}} {
	return m.{{$genElement.Name}}
}
{{end}}
`
)
