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
    
    TEMP_GENBASE = `
package {{.PackageName}}
import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// prepare for other
type _BaseMgr struct {
	*gorm.DB
	ctx       context.Context
	cancel    context.CancelFunc
	timeout   time.Duration
	isRelated bool
}

// SetTimeOut set timeout
func (obj *_BaseMgr) SetTimeOut(timeout time.Duration) {
	obj.ctx, obj.cancel = context.WithTimeout(context.Background(), timeout)
	obj.timeout = timeout
}

// SetCtx set context
func (obj *_BaseMgr) SetCtx(c context.Context) {
	if c != nil {
		obj.ctx = c
	}
}

// GetCtx get context
func (obj *_BaseMgr) GetCtx() context.Context {
	return obj.ctx
}

// Cancel cancel context
func (obj *_BaseMgr) Cancel(c context.Context) {
	obj.cancel()
}

// GetDB get gorm.DB info
func (obj *_BaseMgr) GetDB() *gorm.DB {
	return obj.DB
}

// UpdateDB update gorm.DB info
func (obj *_BaseMgr) UpdateDB(db *gorm.DB) {
	obj.DB = db
}

// New new gorm.新gorm,重置条件
func (obj *_BaseMgr) New() {
	obj.DB = obj.NewDB()
}

// NewDB new gorm.新gorm
func (obj *_BaseMgr) NewDB() *gorm.DB {
	return obj.DB.Session(&gorm.Session{NewDB: true, Context: obj.ctx})
}

type options struct {
	query map[string]interface{}
}

// Option overrides behavior of Connect.
type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

// 自定义sql查询
type Condition struct {
	list []*conditionInfo
}

// And a condition by and .and 一个条件
func (c *Condition) And(column string, cases string, value interface{}) {
	c.list = append(c.list, &conditionInfo{
		andor:  "and",
		column: column, // 列名
		case_:  cases,  // 条件(and,or,in,>=,<=)
		value:  value,
	})
}

// Or a condition by or .or 一个条件
func (c *Condition) Or(column string, cases string, value interface{}) {
	c.list = append(c.list, &conditionInfo{
		andor:  "or",
		column: column, // 列名
		case_:  cases,  // 条件(and,or,in,>=,<=)
		value:  value,
	})
}

func (c *Condition) Get() (where string, out []interface{}) {
	firstAnd := -1
	for i := 0; i < len(c.list); i++ { // 查找第一个and
		if c.list[i].andor == "and" {
			where = fmt.Sprintf("{{GetVV }} %v ?", c.list[i].column, c.list[i].case_)
			out = append(out, c.list[i].value)
			firstAnd = i
			break
		}
	}

	if firstAnd < 0 && len(c.list) > 0 { // 补刀
		where = fmt.Sprintf("{{GetVV }} %v ?", c.list[0].column, c.list[0].case_)
		out = append(out, c.list[0].value)
		firstAnd = 0
	}

	for i := 0; i < len(c.list); i++ { // 添加剩余的
		if firstAnd != i {
			where += fmt.Sprintf(" %v {{GetVV }} %v ?", c.list[i].andor, c.list[i].column, c.list[i].case_)
			out = append(out, c.list[i].value)
		}
	}

	return
}

type conditionInfo struct {
	andor  string
	column string // 列名
	case_  string // 条件(in,>=,<=)
	value  interface{}
}
	`
    
    TEMP_GENLOGIC = `{{$obj := .}}{{$list := $obj.Em}}
type _{{$obj.StructName}}Mgr struct {
	*_BaseMgr
}

// {{$obj.StructName}}Mgr open func
func {{$obj.StructName}}Mgr(db *gorm.DB) *_{{$obj.StructName}}Mgr {
	if db == nil {
		panic(fmt.Errorf("{{$obj.StructName}}Mgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	// return &_{{$obj.StructName}}Mgr{_BaseMgr: &_BaseMgr{DB: db.Table("{{GetTablePrefixName $obj.TableName}}"), ctx:ctx, cancel:cancel, timeout:-1}}
	// return &_{{$obj.StructName}}Mgr{_BaseMgr: &_BaseMgr{DB: db.Table("{{GetTablePrefixName .TableName}}"), ctx:ctx, cancel:cancel, timeout:-1}}
	// return &_{{$obj.StructName}}Mgr{_BaseMgr: &_BaseMgr{DB: db.Model({{GetTableStructName .StructName}}), ctx:ctx, cancel:cancel, timeout:-1}}
	return &_{{$obj.StructName}}Mgr{_BaseMgr: &_BaseMgr{DB: db.Model({{GetTableStructName $obj.StructName}}), ctx:ctx, cancel:cancel, timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_{{$obj.StructName}}Mgr) GetTableName() string {
	return "{{GetTablePrefixName $obj.TableName}}"
}

// Reset 重置gorm会话
func (obj *_{{$obj.StructName}}Mgr) Reset() *_{{$obj.StructName}}Mgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_{{$obj.StructName}}Mgr) Get() (result {{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Find(&result).Error
	return
}

// Gets 获取批量结果
func (obj *_{{$obj.StructName}}Mgr) Gets() (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Find(&results).Error
	return
}

////////////////////////////////// 替换 gorm 的方法 /////////////////////////////////
func (obj *_{{$obj.StructName}}Mgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

////////////////////////// 表中的字段作为 option 条件 ////////////////////////////////////////////
{{range $oem := $obj.Em}}
// With{{$oem.ColStructName}} {{$oem.ColName}}获取 {{$oem.Notes}}
func (obj *_{{$obj.StructName}}Mgr) With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) Option {
	return optionFunc(func(o *options) { o.query["{{$oem.ColName}}"] = {{CapLowercase $oem.ColStructName}} })
}
{{end}}

// GetByOption 功能选项模式获取
func (obj *_{{$obj.StructName}}Mgr) GetByOption(opts ...Option) (result {{$obj.StructName}}, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Where(options.query).Find(&result).Error
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_{{$obj.StructName}}Mgr) GetByOptions(opts ...Option) (results []*{{$obj.StructName}}, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Where(options.query).Find(&results).Error
	return
}
////////////////////////// 通过存在索引的单个字段作为查询条件 ////////////////////////////////////////////

{{range $oem := $obj.Em}}
// GetFrom{{$oem.ColStructName}} 通过{{$oem.ColName}}获取内容 {{$oem.Notes}} {{if $oem.IsMulti}}
func (obj *_{{$obj.StructName}}Mgr) GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Where("{{$oem.ColNameEx}} = ?", {{CapLowercase $oem.ColStructName}}).Find(&results).Error
	return
}
{{else}}
func (obj *_{{$obj.StructName}}Mgr)  GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (result {{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Where("{{$oem.ColNameEx}} = ?", {{CapLowercase $oem.ColStructName}}).Find(&result).Error
	return
}
{{end}}
// GetBatchFrom{{$oem.ColStructName}} 批量查找 {{$oem.Notes}}
func (obj *_{{$obj.StructName}}Mgr) GetBatchFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Where("{{$oem.ColNameEx}} IN (?)", {{CapLowercase $oem.ColStructName}}s).Find(&results).Error
	return
}
 {{end}}
 ////////////////////////// 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ////////////////////////////////////////////
 {{range $ofm := $obj.Primary}}
 // {{GenFListIndex $ofm 1}} primary or index 获取唯一内容
 func (obj *_{{$obj.StructName}}Mgr) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (result {{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Where("{{GenFListIndex $ofm 3}}", {{GenFListIndex $ofm 4}}).Find(&result).Error
	return
}
 {{end}}

 {{range $ofm := $obj.Index}}
 // {{GenFListIndex $ofm 1}}  获取多个内容
 func (obj *_{{$obj.StructName}}Mgr) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Where("{{GenFListIndex $ofm 3}}", {{GenFListIndex $ofm 4}}).Find(&results).Error
	return
}
 {{end}}

`
)
