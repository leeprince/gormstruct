package constants

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/12/10 上午12:07
 * @Desc:
 */

const (
    TempGenLogicV1 = `{{$obj := .}}{{$list := $obj.Em}}
type {{$obj.StructName}}Dao struct {
	*_BaseMgr
}

// {{$obj.StructName}}Dao open func
func New{{$obj.StructName}}Dao (db *gorm.DB) *{{$obj.StructName}}Dao {
	if db == nil {
		panic(fmt.Errorf("{{$obj.StructName}}Dao need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &{{$obj.StructName}}Dao{_BaseMgr: &_BaseMgr{DB: db.Dao({{GetTableStructName $obj.StructName}}), ctx:ctx, cancel:cancel, timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *{{$obj.StructName}}Dao) GetTableName() string {
    {{$obj.TableName}} := {{GetTableStructName $obj.StructName}}
	return {{$obj.TableName}}.TableName()
}

// Reset 重置gorm会话
func (obj *{{$obj.StructName}}Dao) Reset() *{{$obj.StructName}}Dao {
	obj.New()
	return obj
}

// Get 获取
func (obj *{{$obj.StructName}}Dao) Get() (result {{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Dao({{$obj.StructName}}{}).Find(&result).Error
	return
}

// Gets 获取批量结果
func (obj *{{$obj.StructName}}Dao) Gets() (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Dao({{$obj.StructName}}{}).Find(&results).Error
	return
}

// ----------------- 替换 gorm 的方法 -------------
func (obj *{{$obj.StructName}}Dao) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Dao({{$obj.StructName}}{}).Count(count)
}

// ----------------------------------------------------------------

// -------------------- 表中的字段作为 option 条件
{{range $oem := $obj.Em}}
// With{{$oem.ColStructName}} {{$oem.ColName}}获取 {{$oem.Notes}}
func (obj *{{$obj.StructName}}Dao) With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) Option {
	return queryOptionFunc(func(o *options) { o.query["{{$oem.ColName}}"] = {{CapLowercase $oem.ColStructName}} })
}
{{end}}

// GetByOption 功能选项模式获取
func (obj *{{$obj.StructName}}Dao) GetByOption(opts ...Option) (result {{$obj.StructName}}, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Dao({{$obj.StructName}}{}).Where(options.query).Find(&result).Error
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *{{$obj.StructName}}Dao) GetByOptions(opts ...Option) (results []*{{$obj.StructName}}, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Dao({{$obj.StructName}}{}).Where(options.query).Find(&results).Error
	return
}
// -------------------- 通过存在索引的单个字段作为查询条件

{{range $oem := $obj.Em}}
// GetFrom{{$oem.ColStructName}} 通过{{$oem.ColName}}获取内容 {{$oem.Notes}} {{if $oem.IsMulti}}
func (obj *{{$obj.StructName}}Dao) GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Dao({{$obj.StructName}}{}).Where("{{$oem.ColNameEx}} = ?", {{CapLowercase $oem.ColStructName}}).Find(&results).Error
	return
}
{{else}}
func (obj *{{$obj.StructName}}Dao)  GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (result {{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Dao({{$obj.StructName}}{}).Where("{{$oem.ColNameEx}} = ?", {{CapLowercase $oem.ColStructName}}).Find(&result).Error
	return
}
{{end}}
// GetBatchFrom{{$oem.ColStructName}} 批量查找 {{$oem.Notes}}
func (obj *{{$obj.StructName}}Dao) GetBatchFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
	err = obj.DB.WithContext(obj.ctx).Dao({{$obj.StructName}}{}).Where("{{$oem.ColNameEx}} IN ?", {{CapLowercase $oem.ColStructName}}s).Find(&results).Error
	return
}
{{end}}

//  -------------------- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件
{{range $ofm := $obj.Primary}}
// {{GenFListIndex $ofm 1}} primary or index 获取唯一内容
func (obj *{{$obj.StructName}}Dao) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (result {{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Dao({{$obj.StructName}}{}).Where("{{GenFListIndex $ofm 3}}", {{GenFListIndex $ofm 4}}).Find(&result).Error
    return
}
{{end}}
{{range $ofm := $obj.Index}}
// {{GenFListIndex $ofm 1}}  获取多个内容
func (obj *{{$obj.StructName}}Dao) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (results []*{{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Dao({{$obj.StructName}}{}).Where("{{GenFListIndex $ofm 3}}", {{GenFListIndex $ofm 4}}).Find(&results).Error
    return
}
 {{end}}

`
    
    TempGenLogicV2 = `
/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   {{GetCurrentDateTime}}
 * @Desc:
 */

{{$obj := .}}{{$list := $obj.Em}}
type {{$obj.StructName}}Dao struct {
    *_BaseMgr
}

// 初始化 {{$obj.StructName}}Dao
func New{{$obj.StructName}}Dao(db *gorm.DB) *{{$obj.StructName}}Dao {
    if db == nil {
        panic(fmt.Errorf("{{$obj.StructName}}Dao need init by db"))
    }
    ctx, cancel := context.WithCancel(context.Background())
    return &{{$obj.StructName}}Dao{_BaseMgr: &_BaseMgr{DB: db.Dao({{GetTableStructName $obj.StructName}}), ctx:ctx, cancel:cancel, timeout:-1}}
}

// 获取表名称
func (obj *{{$obj.StructName}}Dao) GetTableName() string {
    {{$obj.TableName}} := {{GetTableStructName $obj.StructName}}
    return {{$obj.TableName}}.TableName()
}

// 重置 gorm 会话
func (obj *{{$obj.StructName}}Dao) Reset() *{{$obj.StructName}}Dao {
    obj.New()
    return New{{$obj.StructName}}Dao(obj.DB)
}

// 获取单条记录
func (obj *{{$obj.StructName}}Dao) Get() (result {{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Find(&result).Error
    return
}

// 获取多条记录
func (obj *{{$obj.StructName}}Dao) Gets() (results []*{{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Find(&results).Error
    return
}

// -------------------- 替换 gorm 的方法 -------------
func (obj *{{$obj.StructName}}Dao) Count(count *int64) (tx *gorm.DB) {
    return obj.DB.WithContext(obj.ctx).Count(count)
}

// ----------------------------------------------------------------

// -------------------- 表中的字段作为 option 条件
{{range $oem := $obj.Em}}
// 设置 {{$oem.ColName}}({{$oem.Notes}}) 字段作为 option 条件
func (obj *{{$obj.StructName}}Dao) With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) Option {
    return queryOptionFunc(func(o *options) { o.query["{{$oem.ColName}}"] = {{CapLowercase $oem.ColStructName}} })
}

// 设置 {{$oem.ColName}}({{$oem.Notes}}) 字段的切片作为 option 条件
func (obj *{{$obj.StructName}}Dao) WithBatch{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) Option {
    return queryOptionFunc(func(o *options) { o.query["{{$oem.ColName}}"] = {{CapLowercase $oem.ColStructName}}s })
}
{{end}}

// 设置 offset、limit 作为 option 条件支持分页
func (obj *{{$obj.StructName}}Dao) WithPage(offset int, limit int) Option {
	return pageOptionFunc(func(o *options) {
		o.page.offset = offset
		o.page.limit = limit
	})
}

// 函数选项模式获取单条记录
func (obj *{{$obj.StructName}}Dao) GetByOption(opts ...Option) (result {{$obj.StructName}}, err error) {
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
func (obj *{{$obj.StructName}}Dao) GetByOptions(opts ...Option) (results []*{{$obj.StructName}}, err error) {
    options := options{
        query: make(map[string]interface{}, len(opts)),
		page: paging{offset: 0, limit: 0},
    }
    for _, o := range opts {
        o.apply(&options)
    }

	err = obj.DB.WithContext(obj.ctx).Where(options.query).Scopes(obj.paginate(&options)).Find(&results).Error
    return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *{{$obj.StructName}}Dao) GetCountByOptions(opts ...Option) (count int64) {
    options := options{
        query: make(map[string]interface{}, len(opts)),
		page: paging{offset: 0, limit: 0},
    }
    for _, o := range opts {
        o.apply(&options)
    }

	obj.DB.WithContext(obj.ctx).Where(options.query).Scopes(obj.paginate(&options)).Count(&count)
    return
}
// -------------------- 通过存在索引的单个字段作为查询条件

{{range $oem := $obj.Em}}
// 通过 {{$oem.ColName}}({{$oem.Notes}}) 获取内容  {{if $oem.IsMulti}}
func (obj *{{$obj.StructName}}Dao) GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Where("{{$oem.ColNameEx}} = ?", {{CapLowercase $oem.ColStructName}}).Find(&results).Error
    return
}
{{else}}
func (obj *{{$obj.StructName}}Dao) GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (result {{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Where("{{$oem.ColNameEx}} = ?", {{CapLowercase $oem.ColStructName}}).Find(&result).Error
    return
}
{{end}}
// 通过 {{$oem.ColName}}({{$oem.Notes}}) 获取多条记录
func (obj *{{$obj.StructName}}Dao) GetBatchFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Where("{{$oem.ColNameEx}} IN ?", {{CapLowercase $oem.ColStructName}}s).Find(&results).Error
    return
}
{{end}}
// -------------------- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件
{{range $ofm := $obj.Primary}}
// 获取单条记录
func (obj *{{$obj.StructName}}Dao) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (result {{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Where("{{GenFListIndex $ofm 3}}", {{GenFListIndex $ofm 4}}).Find(&result).Error
    return
}
{{end}}
{{range $ofm := $obj.Index}}
// 获取多条记录
func (obj *{{$obj.StructName}}Dao) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (results []*{{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Where("{{GenFListIndex $ofm 3}}", {{GenFListIndex $ofm 4}}).Find(&results).Error
    return
}
{{end}}
`
    
    TempGenLogicV3 = `
/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   {{GetCurrentDateTime}}
 * @Desc:
 */

{{$obj := .}}{{$list := $obj.Em}}
type {{$obj.StructName}}Dao struct {
    *_BaseMgr
}

// 初始化 {{$obj.StructName}}Dao
func New{{$obj.StructName}}Dao(db *gorm.DB) *{{$obj.StructName}}Dao {
    if db == nil {
        panic(fmt.Errorf("{{$obj.StructName}}Dao need init by db"))
    }
    ctx, cancel := context.WithCancel(context.Background())
    return &{{$obj.StructName}}Dao{_BaseMgr: &_BaseMgr{DB: db.Dao({{GetTableStructName $obj.StructName}}), ctx:ctx, cancel:cancel, timeout:-1}}
}

// 获取表名称
func (obj *{{$obj.StructName}}Dao) GetTableName() string {
    {{$obj.TableName}} := {{GetTableStructName $obj.StructName}}
    return {{$obj.TableName}}.TableName()
}

// 获取单条记录
func (obj *{{$obj.StructName}}Dao) Get() (result {{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Find(&result).Error
    return
}

// 获取多条记录
func (obj *{{$obj.StructName}}Dao) Gets() (results []*{{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Find(&results).Error
    return
}

// -------------------- 替换 gorm 的方法 -------------
func (obj *{{$obj.StructName}}Dao) Count(count *int64) (tx *gorm.DB) {
    return obj.DB.WithContext(obj.ctx).Count(count)
}

// ----------------------------------------------------------------

// -------------------- 表中的字段作为 option 条件
{{range $oem := $obj.Em}}
// 设置 {{$oem.ColName}}({{$oem.Notes}}) 字段作为 option 条件
func (obj *{{$obj.StructName}}Dao) With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) Option {
    return queryOptionFunc(func(o *options) { o.query[{{$obj.StructName}}Columns.{{$oem.ColStructName}}] = {{CapLowercase $oem.ColStructName}} })
}

// 设置 {{$oem.ColName}}({{$oem.Notes}}) 字段的切片作为 option 条件
func (obj *{{$obj.StructName}}Dao) WithBatch{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) Option {
    return queryOptionFunc(func(o *options) { o.query[{{$obj.StructName}}Columns.{{$oem.ColStructName}}] = {{CapLowercase $oem.ColStructName}}s })
}
{{end}}

// 设置 offset、limit 作为 option 条件支持分页
func (obj *{{$obj.StructName}}Dao) WithPage(offset int, limit int) Option {
	return pageOptionFunc(func(o *options) {
		o.page.offset = offset
		o.page.limit = limit
	})
}

// 函数选项模式获取单条记录
func (obj *{{$obj.StructName}}Dao) GetByOption(opts ...Option) (result {{$obj.StructName}}, err error) {
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
func (obj *{{$obj.StructName}}Dao) GetByOptions(opts ...Option) (results []*{{$obj.StructName}}, err error) {
    options := options{
        query: make(map[string]interface{}, len(opts)),
		page: paging{offset: 0, limit: 0},
    }
    for _, o := range opts {
        o.apply(&options)
    }

	err = obj.DB.WithContext(obj.ctx).Where(options.query).Scopes(obj.paginate(&options)).Find(&results).Error
    return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *{{$obj.StructName}}Dao) GetCountByOptions(opts ...Option) (count int64) {
    options := options{
        query: make(map[string]interface{}, len(opts)),
		page: paging{offset: 0, limit: 0},
    }
    for _, o := range opts {
        o.apply(&options)
    }

	obj.DB.WithContext(obj.ctx).Where(options.query).Scopes(obj.paginate(&options)).Count(&count)
    return
}
// -------------------- 通过存在索引的单个字段作为查询条件

{{range $oem := $obj.Em}}
// 通过 {{$oem.ColName}}({{$oem.Notes}}) 获取内容  {{if $oem.IsMulti}}
func (obj *{{$obj.StructName}}Dao) GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Where({{$obj.StructName}}{ {{ $oem.ColStructName}}: {{CapLowercase $oem.ColStructName}} }).Find(&results).Error
    return
}
{{else}}
func (obj *{{$obj.StructName}}Dao) GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (result {{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Where({{$obj.StructName}}{ {{ $oem.ColStructName}}: {{CapLowercase $oem.ColStructName}} }).Find(&result).Error
    return
}
{{end}}
// 通过 {{$oem.ColName}}({{$oem.Notes}}) 获取多条记录
func (obj *{{$obj.StructName}}Dao) GetBatchFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Where(fmt.Sprintf("%s IN ?", {{$obj.StructName}}Columns.{{$oem.ColStructName}}), {{CapLowercase $oem.ColStructName}}s).Find(&results).Error
    return
}
{{end}}
// -------------------- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件
{{range $ofm := $obj.Primary}}
// 获取单条记录
func (obj *{{$obj.StructName}}Dao) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (result {{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Where({{GenFListIndex $ofm 3}}).Find(&result).Error
    return
}
{{end}}
{{range $ofm := $obj.Index}}
// 获取多条记录
func (obj *{{$obj.StructName}}Dao) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (results []*{{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Where({{GenFListIndex $ofm 3}}).Find(&results).Error
    return
}
{{end}}
`
    TempGenLogicV4 = `
{{$obj := .}}
{{$list := $obj.Em}}
{{$lowerStructName := CapLowercase $obj.StructName}}

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   {{GetCurrentDateTime}}
 * @Desc:   {{$obj.TableName}} 表的 dao 层
 */

type {{$obj.StructName}}Dao struct {
    *_BaseDao
}

// 表字段的映射
var {{$obj.StructName}}Columns = struct {
{{range $genElement := $obj.Em}}{{$genElement.ColStructName}} string
{{end}}
} {
{{range $genElement := $obj.Em}}{{$genElement.ColStructName}}: "{{$genElement.ColName}}",
{{end}}
}

// 初始化 {{$obj.StructName}}Dao
func New{{$obj.StructName}}Dao(db *gorm.DB) *{{$obj.StructName}}Dao {
    if db == nil {
        panic(fmt.Errorf("{{$obj.StructName}}Dao need init by db"))
    }
    ctx, cancel := context.WithCancel(context.Background())
    return &{{$obj.StructName}}Dao{_BaseDao: &_BaseDao{DB: db.Model({{GetTableStructName $obj.StructName}}), ctx:ctx, cancel:cancel, timeout:-1}}
}

// 获取表名称
func (obj *{{$obj.StructName}}Dao) GetTableName() string {
    {{$obj.TableName}} := &{{GetTableStructName $obj.StructName}}
    return {{$obj.TableName}}.TableName()
}

// 获取单条记录
func (obj *{{$obj.StructName}}Dao) Get() (result {{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Find(&result).Error
    return
}

// 获取多条记录
func (obj *{{$obj.StructName}}Dao) Gets() (results []*{{$obj.StructName}}, err error) {
    err = obj.DB.WithContext(obj.ctx).Find(&results).Error
    return
}

// --- 替换 gorm 的方法 ---
// 统计
func (obj *{{$obj.StructName}}Dao) Count(count *int64) (tx *gorm.DB) {
    return obj.DB.WithContext(obj.ctx).Count(count)
}

// 插入
func (obj *{{$obj.StructName}}Dao) Create({{$lowerStructName}} *{{$obj.StructName}}) (rowsAffected int64, err error) {
    tx := obj.DB.WithContext(obj.ctx).Create({{$lowerStructName}})
    rowsAffected = tx.RowsAffected
    err = tx.Error
    return
}
// --- 替换 gorm 的方法 -END ---

// --- 表中的字段作为 option 条件 ---
{{range $oem := $obj.Em}}
// 设置 {{$oem.ColName}}({{$oem.Notes}}) 字段作为 option 条件
func (obj *{{$obj.StructName}}Dao) With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) Option {
    return queryOptionFunc(func(o *options) { o.query[{{$obj.StructName}}Columns.{{$oem.ColStructName}}] = {{CapLowercase $oem.ColStructName}} })
}

// 设置 {{$oem.ColName}}({{$oem.Notes}}) 字段的切片作为 option 条件
func (obj *{{$obj.StructName}}Dao) WithBatch{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) Option {
    return queryOptionFunc(func(o *options) { o.query[{{$obj.StructName}}Columns.{{$oem.ColStructName}}] = {{CapLowercase $oem.ColStructName}}s })
}
{{end}}

// 函数选项模式获取单条记录
func (obj *{{$obj.StructName}}Dao) GetByOption(opts ...Option) (result {{$obj.StructName}}, err error) {
    err = obj.prepare(opts...).Find(&result).Error
    return
}

// 函数选项模式获取多条记录：支持分页
func (obj *{{$obj.StructName}}Dao) GetByOptions(opts ...Option) (results []*{{$obj.StructName}}, err error) {
    err = obj.prepare(opts...).Find(&results).Error
    return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *{{$obj.StructName}}Dao) GetCountByOptions(opts ...Option) (count int64) {
    obj.prepare(opts...).Model(&{{$obj.StructName}}{}).Count(&count)
    return
}

// 更新数据：非指针的结构体字段更新为零值更新时需配合 WithSelect 更新
// 参数说明：
//    {{$lowerStructName}}: 要更新的数据
//    opts: 更新的条件
func (obj *{{$obj.StructName}}Dao) UpdateByOption({{$lowerStructName}} {{$obj.StructName}}, opts ...Option) (rowsAffected int64, err error) {
    tx := obj.prepare(opts...).Updates(&{{$lowerStructName}})
    rowsAffected = tx.RowsAffected
    err = tx.Error
    return
}
// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---
{{range $oem := $obj.Em}}
// 通过单个 {{$oem.ColName}}({{$oem.Notes}}) 字段值，{{if $oem.IsMulti}}获取多条记录
func (obj *{{$obj.StructName}}Dao) GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
    results, err = obj.GetByOptions(obj.With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}))
    return
}
{{else}}获取单条记录
func (obj *{{$obj.StructName}}Dao) GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (result {{$obj.StructName}}, err error) {
    result, err = obj.GetByOption(obj.With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}))
    return
}
{{end}}
// 通过多个 {{$oem.ColName}}({{$oem.Notes}}) 字段值，获取多条记录
func (obj *{{$obj.StructName}}Dao) GetBatchFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
    results, err = obj.GetByOptions(obj.WithBatch{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}s))
    return
}
{{end}}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---
{{range $ofm := $obj.Primary}}
// 通过 {{GenFListIndex $ofm 5}} 字段值，获取单条记录
func (obj *{{$obj.StructName}}Dao) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (result {{$obj.StructName}}, err error) {
    result, err = obj.GetByOption({{GenFListIndex $ofm 6}})
    return
}
{{end}}
{{range $ofm := $obj.Index}}
// 通过 {{GenFListIndex $ofm 5}} 字段值，获取多条记录
func (obj *{{$obj.StructName}}Dao) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (results []*{{$obj.StructName}}, err error) {
    results, err = obj.GetByOptions({{GenFListIndex $ofm 6}})
    return
}
{{end}}
// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
`
)


