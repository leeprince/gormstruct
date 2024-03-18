package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-03-19 02:01:26
 * @Desc:   users 表的 DAO 层
 */

type UsersDAO struct {
	*_BaseDAO
}

// UsersDAO 初始化
func NewUsersDAO(ctx context.Context, db *gorm.DB) *UsersDAO {
	if db == nil {
		panic(fmt.Errorf("UsersDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &UsersDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&Users{}),
			db:               db,
			model:            Users{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          UsersAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *UsersDAO) GetTableName() string {
	users := &Users{}
	return users.TableName()
}

// UpdateOrCreate 存在则更新，否则插入，会忽略零值字段
func (obj *UsersDAO) UpdateOrCreate(users *Users) (rowsAffected int64, err error) {
	if users.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(users, obj.WithPrimaryKey(users.PrimaryKeyValue()))
	}
	return obj.Create(users)
}

// Save gorm 原生的 Save 会保存所有的字段，即使字段是零值。仅会判断主键是否存在，存在则更新，不存在则创建
func (obj *UsersDAO) Save(users *Users) (rowsAffected int64) {
	return obj.db.Save(users).RowsAffected
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *UsersDAO) Create(users interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(users)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithPrimaryKey 设置真正的主键 字段作为 option 条件
func (obj *UsersDAO) WithPrimaryKey(primaryKeyValue interface{}) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[(&Users{}).PrimaryKey()] = primaryKeyValue })
}

// WithID 设置 id(主键) 字段作为 option 条件
func (obj *UsersDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.ID] = id })
}

// WithIDs 设置 id(主键) 字段的切片作为 option 条件
func (obj *UsersDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.ID] = ids })
}

// WithName 设置 name(名称) 字段作为 option 条件
func (obj *UsersDAO) WithName(name string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.Name] = name })
}

// WithNames 设置 name(名称) 字段的切片作为 option 条件
func (obj *UsersDAO) WithNames(names []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.Name] = names })
}

// WithAge 设置 age(年龄) 字段作为 option 条件
func (obj *UsersDAO) WithAge(age int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.Age] = age })
}

// WithAges 设置 age(年龄) 字段的切片作为 option 条件
func (obj *UsersDAO) WithAges(ages []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.Age] = ages })
}

// WithCardNo 设置 card_no(身份证) 字段作为 option 条件
func (obj *UsersDAO) WithCardNo(cardNo string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.CardNo] = cardNo })
}

// WithCardNos 设置 card_no(身份证) 字段的切片作为 option 条件
func (obj *UsersDAO) WithCardNos(cardNos []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.CardNo] = cardNos })
}

// WithHeadImg 设置 head_img(头像) 字段作为 option 条件
func (obj *UsersDAO) WithHeadImg(headImg string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.HeadImg] = headImg })
}

// WithHeadImgs 设置 head_img(头像) 字段的切片作为 option 条件
func (obj *UsersDAO) WithHeadImgs(headImgs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.HeadImg] = headImgs })
}

// WithSchool 设置 school() 字段作为 option 条件
func (obj *UsersDAO) WithSchool(school string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.School] = school })
}

// WithSchools 设置 school() 字段的切片作为 option 条件
func (obj *UsersDAO) WithSchools(schools []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.School] = schools })
}

// WithNullNoDefault 设置 null_no_default() 字段作为 option 条件
func (obj *UsersDAO) WithNullNoDefault(nullNoDefault *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.NullNoDefault] = nullNoDefault })
}

// WithNullNoDefaults 设置 null_no_default() 字段的切片作为 option 条件
func (obj *UsersDAO) WithNullNoDefaults(nullNoDefaults []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.NullNoDefault] = nullNoDefaults })
}

// WithNullDefaultNil 设置 null_default_nil() 字段作为 option 条件
func (obj *UsersDAO) WithNullDefaultNil(nullDefaultNil *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.NullDefaultNil] = nullDefaultNil })
}

// WithNullDefaultNils 设置 null_default_nil() 字段的切片作为 option 条件
func (obj *UsersDAO) WithNullDefaultNils(nullDefaultNils []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.NullDefaultNil] = nullDefaultNils })
}

// WithNullDefaultNoNil 设置 null_default_no_nil() 字段作为 option 条件
func (obj *UsersDAO) WithNullDefaultNoNil(nullDefaultNoNil *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.NullDefaultNoNil] = nullDefaultNoNil })
}

// WithNullDefaultNoNils 设置 null_default_no_nil() 字段的切片作为 option 条件
func (obj *UsersDAO) WithNullDefaultNoNils(nullDefaultNoNils []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.NullDefaultNoNil] = nullDefaultNoNils })
}

// WithNoNullNoDefault 设置 no_null_no_default() 字段作为 option 条件
func (obj *UsersDAO) WithNoNullNoDefault(noNullNoDefault string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.NoNullNoDefault] = noNullNoDefault })
}

// WithNoNullNoDefaults 设置 no_null_no_default() 字段的切片作为 option 条件
func (obj *UsersDAO) WithNoNullNoDefaults(noNullNoDefaults []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.NoNullNoDefault] = noNullNoDefaults })
}

// WithNoNullDefaultNoNil 设置 no_null_default_no_nil() 字段作为 option 条件
func (obj *UsersDAO) WithNoNullDefaultNoNil(noNullDefaultNoNil string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.NoNullDefaultNoNil] = noNullDefaultNoNil })
}

// WithNoNullDefaultNoNils 设置 no_null_default_no_nil() 字段的切片作为 option 条件
func (obj *UsersDAO) WithNoNullDefaultNoNils(noNullDefaultNoNils []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.NoNullDefaultNoNil] = noNullDefaultNoNils })
}

// WithNoNullDefaultNil 设置 no_null_default_nil() 字段作为 option 条件
func (obj *UsersDAO) WithNoNullDefaultNil(noNullDefaultNil string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.NoNullDefaultNil] = noNullDefaultNil })
}

// WithNoNullDefaultNils 设置 no_null_default_nil() 字段的切片作为 option 条件
func (obj *UsersDAO) WithNoNullDefaultNils(noNullDefaultNils []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.NoNullDefaultNil] = noNullDefaultNils })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *UsersDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *UsersDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *UsersDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *UsersDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *UsersDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *UsersDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *UsersDAO) GetByOption(opts ...Option) (result *Users, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *UsersDAO) GetListByOption(opts ...Option) (results []*Users, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *UsersDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *UsersDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
//
//		参数说明：
//	     users: 要更新的数据
//	     opts: 更新的条件
func (obj *UsersDAO) UpdateByOption(users *Users, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(users)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键) 字段值，获取单条记录
func (obj *UsersDAO) GetFromID(id int64) (result *Users, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetListFromID 通过多个 id(主键) 字段值，获取多条记录
func (obj *UsersDAO) GetListFromID(ids []int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromName 通过单个 name(名称) 字段值，获取单条记录
func (obj *UsersDAO) GetFromName(name string) (result *Users, err error) {
	result, err = obj.GetByOption(obj.WithName(name))
	return
}

// GetListFromName 通过多个 name(名称) 字段值，获取多条记录
func (obj *UsersDAO) GetListFromName(names []string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithNames(names))
	return
}

// GetFromAge 通过单个 age(年龄) 字段值，获取多条记录
func (obj *UsersDAO) GetFromAge(age int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithAge(age))
	return
}

// GetListFromAge 通过多个 age(年龄) 字段值，获取多条记录
func (obj *UsersDAO) GetListFromAge(ages []int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithAges(ages))
	return
}

// GetFromCardNo 通过单个 card_no(身份证) 字段值，获取单条记录
func (obj *UsersDAO) GetFromCardNo(cardNo string) (result *Users, err error) {
	result, err = obj.GetByOption(obj.WithCardNo(cardNo))
	return
}

// GetListFromCardNo 通过多个 card_no(身份证) 字段值，获取多条记录
func (obj *UsersDAO) GetListFromCardNo(cardNos []string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithCardNos(cardNos))
	return
}

// GetFromHeadImg 通过单个 head_img(头像) 字段值，获取多条记录
func (obj *UsersDAO) GetFromHeadImg(headImg string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithHeadImg(headImg))
	return
}

// GetListFromHeadImg 通过多个 head_img(头像) 字段值，获取多条记录
func (obj *UsersDAO) GetListFromHeadImg(headImgs []string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithHeadImgs(headImgs))
	return
}

// GetFromSchool 通过单个 school() 字段值，获取多条记录
func (obj *UsersDAO) GetFromSchool(school string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithSchool(school))
	return
}

// GetListFromSchool 通过多个 school() 字段值，获取多条记录
func (obj *UsersDAO) GetListFromSchool(schools []string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithSchools(schools))
	return
}

// GetFromNullNoDefault 通过单个 null_no_default() 字段值，获取多条记录
func (obj *UsersDAO) GetFromNullNoDefault(nullNoDefault *string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithNullNoDefault(nullNoDefault))
	return
}

// GetListFromNullNoDefault 通过多个 null_no_default() 字段值，获取多条记录
func (obj *UsersDAO) GetListFromNullNoDefault(nullNoDefaults []*string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithNullNoDefaults(nullNoDefaults))
	return
}

// GetFromNullDefaultNil 通过单个 null_default_nil() 字段值，获取多条记录
func (obj *UsersDAO) GetFromNullDefaultNil(nullDefaultNil *string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithNullDefaultNil(nullDefaultNil))
	return
}

// GetListFromNullDefaultNil 通过多个 null_default_nil() 字段值，获取多条记录
func (obj *UsersDAO) GetListFromNullDefaultNil(nullDefaultNils []*string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithNullDefaultNils(nullDefaultNils))
	return
}

// GetFromNullDefaultNoNil 通过单个 null_default_no_nil() 字段值，获取多条记录
func (obj *UsersDAO) GetFromNullDefaultNoNil(nullDefaultNoNil *string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithNullDefaultNoNil(nullDefaultNoNil))
	return
}

// GetListFromNullDefaultNoNil 通过多个 null_default_no_nil() 字段值，获取多条记录
func (obj *UsersDAO) GetListFromNullDefaultNoNil(nullDefaultNoNils []*string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithNullDefaultNoNils(nullDefaultNoNils))
	return
}

// GetFromNoNullNoDefault 通过单个 no_null_no_default() 字段值，获取多条记录
func (obj *UsersDAO) GetFromNoNullNoDefault(noNullNoDefault string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithNoNullNoDefault(noNullNoDefault))
	return
}

// GetListFromNoNullNoDefault 通过多个 no_null_no_default() 字段值，获取多条记录
func (obj *UsersDAO) GetListFromNoNullNoDefault(noNullNoDefaults []string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithNoNullNoDefaults(noNullNoDefaults))
	return
}

// GetFromNoNullDefaultNoNil 通过单个 no_null_default_no_nil() 字段值，获取多条记录
func (obj *UsersDAO) GetFromNoNullDefaultNoNil(noNullDefaultNoNil string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithNoNullDefaultNoNil(noNullDefaultNoNil))
	return
}

// GetListFromNoNullDefaultNoNil 通过多个 no_null_default_no_nil() 字段值，获取多条记录
func (obj *UsersDAO) GetListFromNoNullDefaultNoNil(noNullDefaultNoNils []string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithNoNullDefaultNoNils(noNullDefaultNoNils))
	return
}

// GetFromNoNullDefaultNil 通过单个 no_null_default_nil() 字段值，获取多条记录
func (obj *UsersDAO) GetFromNoNullDefaultNil(noNullDefaultNil string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithNoNullDefaultNil(noNullDefaultNil))
	return
}

// GetListFromNoNullDefaultNil 通过多个 no_null_default_nil() 字段值，获取多条记录
func (obj *UsersDAO) GetListFromNoNullDefaultNil(noNullDefaultNils []string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithNoNullDefaultNils(noNullDefaultNils))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *UsersDAO) GetFromCreatedAt(createdAt int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetListFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *UsersDAO) GetListFromCreatedAt(createdAts []int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *UsersDAO) GetFromUpdatedAt(updatedAt int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetListFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *UsersDAO) GetListFromUpdatedAt(updatedAts []int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *UsersDAO) GetFromDeletedAt(deletedAt int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetListFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *UsersDAO) GetListFromDeletedAt(deletedAts []int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *UsersDAO) FetchByPrimaryKey(id int64) (result *Users, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// FetchUniqueByCardNo 通过 card_no 字段值，获取单条记录
func (obj *UsersDAO) FetchUniqueByCardNo(cardNo string) (result *Users, err error) {
	result, err = obj.GetByOption(obj.WithCardNo(cardNo))
	return
}

// FetchUniqueIndexByUnqNameCard 通过 name, card_no 字段值，获取单条记录
func (obj *UsersDAO) FetchUniqueIndexByUnqNameCard(name string, cardNo string) (result *Users, err error) {
	result, err = obj.GetByOption(
		obj.WithName(name),
		obj.WithCardNo(cardNo))
	return
}

// FetchIndexByAge 通过 age 字段值，获取多条记录
func (obj *UsersDAO) FetchIndexByAge(age int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithAge(age))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
