package constants

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/12/10 上午12:07
 * @Desc:
 */

const (
    TempGenLogicV4 = `
{{$allParams := .}}
{{$list := $allParams.Em}}
{{$lowerStructName := CapLowercase $allParams.StructName}}

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   {{GetCurrentDateTime}}
 * @Desc:   {{$allParams.TableName}} 表的 DAO 层
 */

type {{$allParams.StructName}}DAO struct {
    *_BaseDAO
}

// 初始化 {{$allParams.StructName}}DAO
func New{{$allParams.StructName}}DAO(ctx context.Context, db *gorm.DB) *{{$allParams.StructName}}DAO {
    if db == nil {
        panic(fmt.Errorf("{{$allParams.StructName}}DAO need init by db"))
    }
    ctx, cancel := context.WithCancel(ctx)
    return &{{$allParams.StructName}}DAO{
        _BaseDAO: &_BaseDAO{
            DB: db.Model(&{{GetTableStructName $allParams.StructName}}),
            db: db,
            model:{{GetTableStructName $allParams.StructName}},
            ctx:ctx,
            cancel:cancel,
            timeout:-1,
            columns:{{$allParams.StructName}}AllColumns,
            isDefaultColumns:true,
        },
    }
}

// 获取表名称
func (obj *{{$allParams.StructName}}DAO) GetTableName() string {
    {{$lowerStructName}} := &{{GetTableStructName $allParams.StructName}}
    return {{$lowerStructName}}.TableName()
}

// 存在则更新，否则插入
func (obj *{{$allParams.StructName}}DAO) Save({{$lowerStructName}} *{{$allParams.StructName}}) (rowsAffected int64, err error) {
	if {{$lowerStructName}}.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption({{$lowerStructName}}, obj.WithID({{$lowerStructName}}.ID))
	}
    return obj.Create({{$lowerStructName}})
}

// 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *{{$allParams.StructName}}DAO) Create({{$lowerStructName}} interface{}) (rowsAffected int64, err error) {
    tx := obj.withContext().Create({{$lowerStructName}})
    return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---
{{range $oem := $allParams.Em}}
// 设置 {{$oem.ColName}}({{$oem.Notes}}) 字段作为 option 条件
func (obj *{{$allParams.StructName}}DAO) With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) Option {
    return queryOptionFunc(func(o *options) { o.query[{{$allParams.StructName}}Columns.{{$oem.ColStructName}}] = {{CapLowercase $oem.ColStructName}} })
}

// 设置 {{$oem.ColName}}({{$oem.Notes}}) 字段的切片作为 option 条件
func (obj *{{$allParams.StructName}}DAO) With{{$oem.ColStructName}}s({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) Option {
    return queryOptionFunc(func(o *options) { o.query[{{$allParams.StructName}}Columns.{{$oem.ColStructName}}] = {{CapLowercase $oem.ColStructName}}s })
}
{{end}}

// 函数选项模式获取单条记录
func (obj *{{$allParams.StructName}}DAO) GetByOption(opts ...Option) (result *{{$allParams.StructName}}, err error) {
    {{if $allParams.IsHaveDeleteFlag}}opts = append(opts, obj.With{{$allParams.DeleteFlagStructField}}(0)){{end}}
    err = obj.prepare(opts...).Find(&result).Error
    return
}

// 函数选项模式获取多条记录：支持分页
func (obj *{{$allParams.StructName}}DAO) GetByOptions(opts ...Option) (results []*{{$allParams.StructName}}, err error) {
    err = obj.prepare(opts...).Find(&results).Error
    return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *{{$allParams.StructName}}DAO) GetCountByOptions(opts ...Option) (count int64) {
    obj.setIsDefaultColumns(false)
    obj.prepare(opts...).Count(&count)
    return
}

// 更新数据
//  参数说明：
//      {{$lowerStructName}}: 要更新的数据
//      opts: 更新的条件
func (obj *{{$allParams.StructName}}DAO) UpdateByOption({{$lowerStructName}} *{{$allParams.StructName}}, opts ...Option) (rowsAffected int64, err error) {
    obj.setIsDefaultColumns(false)
    tx := obj.prepare(opts...).Updates({{$lowerStructName}})
    return tx.RowsAffected, tx.Error
}
// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---
{{range $oem := $allParams.Em}}
// 通过单个 {{$oem.ColName}}({{$oem.Notes}}) 字段值，{{if $oem.IsMulti}}获取多条记录
func (obj *{{$allParams.StructName}}DAO) GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (results []*{{$allParams.StructName}}, err error) {
    results, err = obj.GetByOptions(obj.With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}))
    return
}
{{else}}获取单条记录
func (obj *{{$allParams.StructName}}DAO) GetFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) (result *{{$allParams.StructName}}, err error) {
    result, err = obj.GetByOption(obj.With{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}))
    return
}
{{end}}
// 通过多个 {{$oem.ColName}}({{$oem.Notes}}) 字段值，获取多条记录
func (obj *{{$allParams.StructName}}DAO) GetsFrom{{$oem.ColStructName}}({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) (results []*{{$allParams.StructName}}, err error) {
    results, err = obj.GetByOptions(obj.With{{$oem.ColStructName}}s({{CapLowercase $oem.ColStructName}}s))
    return
}
{{end}}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---
{{range $ofm := $allParams.Primary}}
// 通过 {{GenFListIndex $ofm 5}} 字段值，获取单条记录
func (obj *{{$allParams.StructName}}DAO) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (result *{{$allParams.StructName}}, err error) {
    result, err = obj.GetByOption({{GenFListIndex $ofm 6}})
    return
}
{{end}}
{{range $ofm := $allParams.Index}}
// 通过 {{GenFListIndex $ofm 5}} 字段值，获取多条记录
func (obj *{{$allParams.StructName}}DAO) {{GenFListIndex $ofm 1}}({{GenFListIndex $ofm 2}}) (results []*{{$allParams.StructName}}, err error) {
    results, err = obj.GetByOptions({{GenFListIndex $ofm 6}})
    return
}
{{end}}
// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
`
)
