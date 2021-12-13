package constants

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/14 下午6:33
 * @Desc:
 */

const (
    FILE_BASE_NAME = "gen.base.go"
)

const (
    TEMP_GENTABLENAME = `
// TableName get sql table name.获取数据库表名
func (m *{{.StructName}}) TableName() string {
	return "{{.TableName}}"
}
`
)
