package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022-06-11 14:37:17
 * @Desc:   user_base 表的 dao 层
 */

type UserBaseDao struct {
	*_BaseDao
}

// 初始化 UserBaseDao
func NewUserBaseDao(ctx context.Context, db *gorm.DB) *UserBaseDao {
	if db == nil {
		panic(fmt.Errorf("UserBaseDao need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &UserBaseDao{
		_BaseDao: &_BaseDao{
			DB:               db.Model(&UserBase{}),
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          UserBaseAllColumns,
			isDefaultColumns: true,
		},
	}
}

// 获取表名称
func (obj *UserBaseDao) GetTableName() string {
	user_base := &UserBase{}
	return user_base.TableName()
}

// 存在则更新，否则插入：检查模型主键(默认是ID为字段的整型数据类型)存在则更新，否则插入
func (obj *UserBaseDao) Save(userBase *UserBase) (rowsAffected int64, err error) {
	if userBase.ID > 0 {
		return obj.UpdateByOption(userBase, obj.WithID(userBase.ID))
	}
	tx := obj.WithContext().Save(userBase)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// 设置 id() 字段作为 option 条件
func (obj *UserBaseDao) WithID(id int32) Option {
	return queryOptionFunc(func(o *options) { o.query[UserBaseColumns.ID] = id })
}

// 设置 id() 字段的切片作为 option 条件
func (obj *UserBaseDao) WithIDs(ids []int32) Option {
	return queryOptionFunc(func(o *options) { o.query[UserBaseColumns.ID] = ids })
}

// 设置 name() 字段作为 option 条件
func (obj *UserBaseDao) WithName(name string) Option {
	return queryOptionFunc(func(o *options) { o.query[UserBaseColumns.Name] = name })
}

// 设置 name() 字段的切片作为 option 条件
func (obj *UserBaseDao) WithNames(names []string) Option {
	return queryOptionFunc(func(o *options) { o.query[UserBaseColumns.Name] = names })
}

// 设置 age() 字段作为 option 条件
func (obj *UserBaseDao) WithAge(age int32) Option {
	return queryOptionFunc(func(o *options) { o.query[UserBaseColumns.Age] = age })
}

// 设置 age() 字段的切片作为 option 条件
func (obj *UserBaseDao) WithAges(ages []int32) Option {
	return queryOptionFunc(func(o *options) { o.query[UserBaseColumns.Age] = ages })
}

// 函数选项模式获取单条记录
func (obj *UserBaseDao) GetByOption(opts ...Option) (result *UserBase, err error) {
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// 函数选项模式获取多条记录：支持分页
func (obj *UserBaseDao) GetByOptions(opts ...Option) (results []*UserBase, err error) {
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *UserBaseDao) GetCountByOptions(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	obj.prepare(opts...).Count(&count)
	return
}

// 更新数据
//  参数说明：
//      users: 要更新的数据
//      opts: 更新的条件
func (obj *UserBaseDao) UpdateByOption(userBase *UserBase, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(userBase)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// 通过单个 id() 字段值，获取单条记录
func (obj *UserBaseDao) GetFromID(id int32) (result *UserBase, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过多个 id() 字段值，获取多条记录
func (obj *UserBaseDao) GetsFromID(ids []int32) (results []*UserBase, err error) {
	results, err = obj.GetByOptions(obj.WithIDs(ids))
	return
}

// 通过单个 name() 字段值，获取多条记录
func (obj *UserBaseDao) GetFromName(name string) (results []*UserBase, err error) {
	results, err = obj.GetByOptions(obj.WithName(name))
	return
}

// 通过多个 name() 字段值，获取多条记录
func (obj *UserBaseDao) GetsFromName(names []string) (results []*UserBase, err error) {
	results, err = obj.GetByOptions(obj.WithNames(names))
	return
}

// 通过单个 age() 字段值，获取多条记录
func (obj *UserBaseDao) GetFromAge(age int32) (results []*UserBase, err error) {
	results, err = obj.GetByOptions(obj.WithAge(age))
	return
}

// 通过多个 age() 字段值，获取多条记录
func (obj *UserBaseDao) GetsFromAge(ages []int32) (results []*UserBase, err error) {
	results, err = obj.GetByOptions(obj.WithAges(ages))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// 通过 id 字段值，获取单条记录
func (obj *UserBaseDao) FetchByPrimaryKey(id int32) (result *UserBase, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过 age 字段值，获取多条记录
func (obj *UserBaseDao) FetchIndexByAge(age int32) (results []*UserBase, err error) {
	results, err = obj.GetByOptions(obj.WithAge(age))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
