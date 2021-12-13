package constants

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/12/10 上午12:07
 * @Desc:
 */

const (
    TEMP_GENLOGIC_V1 = `{{$obj := .}}{{$list := $obj.Em}}
type _{{$obj.StructName}}Mgr struct {
	*_BaseMgr
}

// {{$obj.StructName}}Mgr open func
func {{$obj.StructName}}Mgr(db *gorm.DB) *_{{$obj.StructName}}Mgr {
	if db == nil {
		panic(fmt.Errorf("{{$obj.StructName}}Mgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_{{$obj.StructName}}Mgr{_BaseMgr: &_BaseMgr{DB: db.Model({{GetTableStructName $obj.StructName}}), ctx:ctx, cancel:cancel, timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_{{$obj.StructName}}Mgr) GetTableName() string {
    {{$obj.TableName}} := {{GetTableStructName $obj.StructName}}
	return {{$obj.TableName}}.TableName()
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

// ----------------- 替换 gorm 的方法 -------------
func (obj *_{{$obj.StructName}}Mgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model({{$obj.StructName}}{}).Count(count)
}

// ----------------------------------------------------------------

// -------------------- 表中的字段作为 option 条件
{{range $oem := $obj.Em}}
// With{{$oem.ColStructName}} {{$oem.ColName}}获取 {{$oem.Notes}}
func (obj *_{{$obj.StructName}}Mgr) With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) Option {
	return queryOptionFunc(func(o *options) { o.query["{{$oem.ColName}}"] = {{CapLowercase $oem.ColStructName}} })
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
// -------------------- 通过存在索引的单个字段作为查询条件

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

//  -------------------- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件
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
    
    TEMP_GENLOGIC_V2 = `{{$obj := .}}{{$list := $obj.Em}}
type _{{$obj.StructName}}Mgr struct {
    *_BaseMgr
}

// 初始化 {{$obj.StructName}}Mgr
func {{$obj.StructName}}Mgr(db *gorm.DB) *_{{$obj.StructName}}Mgr {
    if db == nil {
        panic(fmt.Errorf("{{$obj.StructName}}Mgr need init by db"))
    }
    ctx, cancel := context.WithCancel(context.Background())
    return &_{{$obj.StructName}}Mgr{_BaseMgr: &_BaseMgr{DB: db.Model({{GetTableStructName $obj.StructName}}), ctx:ctx, cancel:cancel, timeout:-1}}
}

// 获取表名称
func (obj *_{{$obj.StructName}}Mgr) GetTableName() string {
    {{$obj.TableName}} := {{GetTableStructName $obj.StructName}}
    return {{$obj.TableName}}.TableName()
}

// 重置 gorm 会话
func (obj *_{{$obj.StructName}}Mgr) Reset() *_{{$obj.StructName}}Mgr {
    obj.New()
    return {{$obj.StructName}}Mgr(obj.DB)
}

// 获取单条记录
func (obj *_{{$obj.StructName}}Mgr) Get() (result {{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Find(&result).Error
    return
}

// 获取多条记录
func (obj *_{{$obj.StructName}}Mgr) Gets() (results []*{{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Find(&results).Error
    return
}

// -------------------- 替换 gorm 的方法 -------------
func (obj *_{{$obj.StructName}}Mgr) Count(count *int64) (tx *gorm.DB) {
    return obj.DB.WithContext(obj.ctx).Count(count)
}

// ----------------------------------------------------------------

// -------------------- 表中的字段作为 option 条件
{{range $oem := $obj.Em}}
// 设置 {{$oem.ColName}}({{$oem.Notes}}) 字段作为 option 条件
func (obj *_{{$obj.StructName}}Mgr) With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) Option {
    return queryOptionFunc(func(o *options) { o.query["{{$oem.ColName}}"] = {{CapLowercase $oem.ColStructName}} })
}

// 设置 {{$oem.ColName}}({{$oem.Notes}}) 字段的切片作为 option 条件
func (obj *_{{$obj.StructName}}Mgr) WithBatch{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) Option {
    return queryOptionFunc(func(o *options) { o.query["{{$oem.ColName}}"] = {{CapLowercase $oem.ColStructName}}s })
}
{{end}}

// 设置 offset、limit 作为 option 条件支持分页
func (obj *_UsersMgr) WithPage(offset int, limit int) Option {
	return pageOptionFunc(func(o *options) {
		o.page.offset = offset
		o.page.limit = limit
	})
}

// 函数选项模式获取单条记录
func (obj *_{{$obj.StructName}}Mgr) GetByOption(opts ...Option) (result {{$obj.StructName}}, err error) {
    options := options{
        query: make(map[string]interface{}, len(opts)),
    }
    for _, o := range opts {
        o.apply(&options)
    }

    err = obj.DB.WithContext(obj.ctx).Where(options.query).Find(&result).Error
    return
}

// 函数选项模式获取多条记录：支持分页
func (obj *_{{$obj.StructName}}Mgr) GetByOptions(opts ...Option) (results []*{{$obj.StructName}}, err error) {
    options := options{
        query: make(map[string]interface{}, len(opts)),
		page: struct {
			offset int
			limit  int
		}{offset: 0, limit: 0},
    }
    for _, o := range opts {
        o.apply(&options)
    }

	err = obj.DB.WithContext(obj.ctx).Where(options.query).Scopes(obj.paginate(&options)).Find(&results).Error
    return
}
// -------------------- 通过存在索引的单个字段作为查询条件

{{range $oem := $obj.Em}}
// 通过 {{$oem.ColName}}({{$oem.Notes}}) 获取内容  {{if $oem.IsMulti}}
func (obj *_{{$obj.StructName}}Mgr) GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Where("{{$oem.ColNameEx}} = ?", {{CapLowercase $oem.ColStructName}}).Find(&results).Error
    return
}
{{else}}
func (obj *_{{$obj.StructName}}Mgr) GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (result {{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Where("{{$oem.ColNameEx}} = ?", {{CapLowercase $oem.ColStructName}}).Find(&result).Error
    return
}
{{end}}
// 通过 {{$oem.ColName}}({{$oem.Notes}}) 获取多条记录
func (obj *_{{$obj.StructName}}Mgr) GetBatchFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Where("{{$oem.ColNameEx}} IN (?)", {{CapLowercase $oem.ColStructName}}s).Find(&results).Error
    return
}
{{end}}
// -------------------- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件
{{range $ofm := $obj.Primary}}
// 获取单条记录
func (obj *_{{$obj.StructName}}Mgr) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (result {{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Where("{{GenFListIndex $ofm 3}}", {{GenFListIndex $ofm 4}}).Find(&result).Error
    return
}
{{end}}
{{range $ofm := $obj.Index}}
// 获取多条记录
func (obj *_{{$obj.StructName}}Mgr) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (results []*{{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Where("{{GenFListIndex $ofm 3}}", {{GenFListIndex $ofm 4}}).Find(&results).Error
    return
}
{{end}}
`
)


