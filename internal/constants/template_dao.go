package constants

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/12/10 上午12:07
 * @Desc:
 */

const (
    TempGenLogicV4 = `
{{$obj := .}}
{{$list := $obj.Em}}
{{$lowerStructName := CapLowercase $obj.StructName}}

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   {{GetCurrentDateTime}}
 * @Desc:   {{$obj.TableName}} 表的 DAO 层
 */

type {{$obj.StructName}}DAO struct {
    *_BaseDAO
}

// 初始化 {{$obj.StructName}}DAO
func New{{$obj.StructName}}DAO(ctx context.Context, db *gorm.DB) *{{$obj.StructName}}DAO {
    if db == nil {
        panic(fmt.Errorf("{{$obj.StructName}}DAO need init by db"))
    }
    ctx, cancel := context.WithCancel(ctx)
    return &{{$obj.StructName}}DAO{
        _BaseDAO: &_BaseDAO{
            DB: db.Model(&{{GetTableStructName $obj.StructName}}),
            model:{{GetTableStructName $obj.StructName}},
            ctx:ctx,
            cancel:cancel,
            timeout:-1,
            columns:{{$obj.StructName}}AllColumns,
            isDefaultColumns:true,
        },
    }
}

// 获取表名称
func (obj *{{$obj.StructName}}DAO) GetTableName() string {
    {{$obj.TableName}} := &{{GetTableStructName $obj.StructName}}
    return {{$obj.TableName}}.TableName()
}

// 存在则更新，否则插入
func (obj *{{$obj.StructName}}DAO) Save({{$lowerStructName}} *{{$obj.StructName}}) (rowsAffected int64, err error) {
	if {{$lowerStructName}}.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption({{$lowerStructName}}, obj.WithID({{$lowerStructName}}.ID))
	}
    return obj.Create({{$lowerStructName}})
}

// 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *{{$obj.StructName}}DAO) Create({{$lowerStructName}} interface{}) (rowsAffected int64, err error) {
    tx := obj.WithContext().Create({{$lowerStructName}})
    return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---
{{range $oem := $obj.Em}}
// 设置 {{$oem.ColName}}({{$oem.Notes}}) 字段作为 option 条件
func (obj *{{$obj.StructName}}DAO) With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) Option {
    return queryOptionFunc(func(o *options) { o.query[{{$obj.StructName}}Columns.{{$oem.ColStructName}}] = {{CapLowercase $oem.ColStructName}} })
}

// 设置 {{$oem.ColName}}({{$oem.Notes}}) 字段的切片作为 option 条件
func (obj *{{$obj.StructName}}DAO) With{{$oem.ColStructName}}s({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) Option {
    return queryOptionFunc(func(o *options) { o.query[{{$obj.StructName}}Columns.{{$oem.ColStructName}}] = {{CapLowercase $oem.ColStructName}}s })
}
{{end}}

// 函数选项模式获取单条记录
func (obj *{{$obj.StructName}}DAO) GetByOption(opts ...Option) (result *{{$obj.StructName}}, err error) {
    err = obj.prepare(opts...).Find(&result).Error
    return
}

// 函数选项模式获取多条记录：支持分页
func (obj *{{$obj.StructName}}DAO) GetByOptions(opts ...Option) (results []*{{$obj.StructName}}, err error) {
    err = obj.prepare(opts...).Find(&results).Error
    return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *{{$obj.StructName}}DAO) GetCountByOptions(opts ...Option) (count int64) {
    obj.setIsDefaultColumns(false)
    obj.prepare(opts...).Count(&count)
    return
}

// 更新数据
//  参数说明：
//      users: 要更新的数据
//      opts: 更新的条件
func (obj *{{$obj.StructName}}DAO) UpdateByOption({{$lowerStructName}} *{{$obj.StructName}}, opts ...Option) (rowsAffected int64, err error) {
    obj.setIsDefaultColumns(false)
    tx := obj.prepare(opts...).Updates({{$lowerStructName}})
    return tx.RowsAffected, tx.Error
}
// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---
{{range $oem := $obj.Em}}
// 通过单个 {{$oem.ColName}}({{$oem.Notes}}) 字段值，{{if $oem.IsMulti}}获取多条记录
func (obj *{{$obj.StructName}}DAO) GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
    results, err = obj.GetByOptions(obj.With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}))
    return
}
{{else}}获取单条记录
func (obj *{{$obj.StructName}}DAO) GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (result *{{$obj.StructName}}, err error) {
    result, err = obj.GetByOption(obj.With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}))
    return
}
{{end}}
// 通过多个 {{$oem.ColName}}({{$oem.Notes}}) 字段值，获取多条记录
func (obj *{{$obj.StructName}}DAO) GetsFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) (results []*{{$obj.StructName}}, err error) {
    results, err = obj.GetByOptions(obj.With{{$oem.ColStructName}}s({{CapLowercase $oem.ColStructName}}s))
    return
}
{{end}}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---
{{range $ofm := $obj.Primary}}
// 通过 {{GenFListIndex $ofm 5}} 字段值，获取单条记录
func (obj *{{$obj.StructName}}DAO) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (result *{{$obj.StructName}}, err error) {
    result, err = obj.GetByOption({{GenFListIndex $ofm 6}})
    return
}
{{end}}
{{range $ofm := $obj.Index}}
// 通过 {{GenFListIndex $ofm 5}} 字段值，获取多条记录
func (obj *{{$obj.StructName}}DAO) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (results []*{{$obj.StructName}}, err error) {
    results, err = obj.GetByOptions({{GenFListIndex $ofm 6}})
    return
}
{{end}}
// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
`
)
