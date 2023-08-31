package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-08-31 23:21:05
 * @Desc:   user_base 表的 DAO 层
 */

type UserBaseDAO struct {
	*_BaseDAO
}

// UserBaseDAO 初始化
func NewUserBaseDAO(ctx context.Context, db *gorm.DB) *UserBaseDAO {
	if db == nil {
		panic(fmt.Errorf("UserBaseDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &UserBaseDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&UserBase{}),
			db:               db,
			model:            UserBase{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          UserBaseAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *UserBaseDAO) GetTableName() string {
	userBase := &UserBase{}
	return userBase.TableName()
}

// Save 存在则更新，否则插入
func (obj *UserBaseDAO) Save(userBase *UserBase) (rowsAffected int64, err error) {
	if userBase.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(userBase, obj.WithID(userBase.ID))
	}
	return obj.Create(userBase)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *UserBaseDAO) Create(userBase interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(userBase)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id() 字段作为 option 条件
func (obj *UserBaseDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserBaseColumns.ID] = id })
}

// WithIDs 设置 id() 字段的切片作为 option 条件
func (obj *UserBaseDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserBaseColumns.ID] = ids })
}

// WithName 设置 name() 字段作为 option 条件
func (obj *UserBaseDAO) WithName(name string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserBaseColumns.Name] = name })
}

// WithNames 设置 name() 字段的切片作为 option 条件
func (obj *UserBaseDAO) WithNames(names []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserBaseColumns.Name] = names })
}

// WithAge 设置 age() 字段作为 option 条件
func (obj *UserBaseDAO) WithAge(age int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserBaseColumns.Age] = age })
}

// WithAges 设置 age() 字段的切片作为 option 条件
func (obj *UserBaseDAO) WithAges(ages []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserBaseColumns.Age] = ages })
}

// GetByOption 函数选项模式获取单条记录
func (obj *UserBaseDAO) GetByOption(opts ...Option) (result *UserBase, err error) {

	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *UserBaseDAO) GetListByOption(opts ...Option) (results []*UserBase, err error) {

	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *UserBaseDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加添加 gorm:"column:字段的别名;" 标签)
func (obj *UserBaseDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
//
//		参数说明：
//	     userBase: 要更新的数据
//	     opts: 更新的条件
func (obj *UserBaseDAO) UpdateByOption(userBase *UserBase, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(userBase)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id() 字段值，获取单条记录
func (obj *UserBaseDAO) GetFromID(id int64) (result *UserBase, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id() 字段值，获取多条记录
func (obj *UserBaseDAO) GetsFromID(ids []int64) (results []*UserBase, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromName 通过单个 name() 字段值，获取多条记录
func (obj *UserBaseDAO) GetFromName(name string) (results []*UserBase, err error) {
	results, err = obj.GetListByOption(obj.WithName(name))
	return
}

// GetsFromName 通过多个 name() 字段值，获取多条记录
func (obj *UserBaseDAO) GetsFromName(names []string) (results []*UserBase, err error) {
	results, err = obj.GetListByOption(obj.WithNames(names))
	return
}

// GetFromAge 通过单个 age() 字段值，获取多条记录
func (obj *UserBaseDAO) GetFromAge(age int64) (results []*UserBase, err error) {
	results, err = obj.GetListByOption(obj.WithAge(age))
	return
}

// GetsFromAge 通过多个 age() 字段值，获取多条记录
func (obj *UserBaseDAO) GetsFromAge(ages []int64) (results []*UserBase, err error) {
	results, err = obj.GetListByOption(obj.WithAges(ages))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *UserBaseDAO) FetchByPrimaryKey(id int64) (result *UserBase, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// FetchIndexByAge 通过 age 字段值，获取多条记录
func (obj *UserBaseDAO) FetchIndexByAge(age int64) (results []*UserBase, err error) {
	results, err = obj.GetListByOption(obj.WithAge(age))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
