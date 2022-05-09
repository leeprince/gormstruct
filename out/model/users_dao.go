package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022-05-10 01:54:51
 * @Desc:   users 表的 dao 层
 */

type UsersDao struct {
	*_BaseDao
}

// 初始化 UsersDao
func NewUsersDao(ctx context.Context, db *gorm.DB) *UsersDao {
	if db == nil {
		panic(fmt.Errorf("UsersDao need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &UsersDao{
		_BaseDao: &_BaseDao{
			DB:      db,
			ctx:     ctx,
			cancel:  cancel,
			timeout: -1,
			columns: UsersAllColumns,
		},
	}
}

// 获取表名称
func (obj *UsersDao) GetTableName() string {
	users := &Users{}
	return users.TableName()
}

// 获取单条记录
func (obj *UsersDao) Get() (result Users, err error) {
	err = obj.WithContext().Find(&result).Error
	return
}

// 获取多条记录
func (obj *UsersDao) Gets() (results []*Users, err error) {
	err = obj.WithContext().Find(&results).Error
	return
}

// --- 替换 gorm 的方法 ---
// 统计
func (obj *UsersDao) Count(count *int64) (tx *gorm.DB) {
	return obj.WithContext().Table(obj.GetTableName()).Count(count)
}

// 插入
func (obj *UsersDao) Create(users *Users) (rowsAffected int64, err error) {
	tx := obj.WithContext().Create(users)
	rowsAffected = tx.RowsAffected
	err = tx.Error
	return
}

// --- 替换 gorm 的方法 -END ---

// --- 表中的字段作为 option 条件 ---

// 设置 id(主键) 字段作为 option 条件
func (obj *UsersDao) WithID(id int32) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.ID] = id })
}

// 设置 id(主键) 字段的切片作为 option 条件
func (obj *UsersDao) WithIDs(ids []int32) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.ID] = ids })
}

// 设置 name(名称) 字段作为 option 条件
func (obj *UsersDao) WithName(name *string) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.Name] = name })
}

// 设置 name(名称) 字段的切片作为 option 条件
func (obj *UsersDao) WithNames(names []*string) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.Name] = names })
}

// 设置 age(年龄) 字段作为 option 条件
func (obj *UsersDao) WithAge(age int32) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.Age] = age })
}

// 设置 age(年龄) 字段的切片作为 option 条件
func (obj *UsersDao) WithAges(ages []int32) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.Age] = ages })
}

// 设置 card_no(身份证) 字段作为 option 条件
func (obj *UsersDao) WithCardNo(cardNo string) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.CardNo] = cardNo })
}

// 设置 card_no(身份证) 字段的切片作为 option 条件
func (obj *UsersDao) WithCardNos(cardNos []string) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.CardNo] = cardNos })
}

// 设置 head_img(头像) 字段作为 option 条件
func (obj *UsersDao) WithHeadImg(headImg string) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.HeadImg] = headImg })
}

// 设置 head_img(头像) 字段的切片作为 option 条件
func (obj *UsersDao) WithHeadImgs(headImgs []string) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.HeadImg] = headImgs })
}

// 设置 created_at(创建时间) 字段作为 option 条件
func (obj *UsersDao) WithCreatedAt(createdAt int32) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.CreatedAt] = createdAt })
}

// 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *UsersDao) WithCreatedAts(createdAts []int32) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.CreatedAt] = createdAts })
}

// 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *UsersDao) WithUpdatedAt(updatedAt int32) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.UpdatedAt] = updatedAt })
}

// 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *UsersDao) WithUpdatedAts(updatedAts []int32) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.UpdatedAt] = updatedAts })
}

// 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *UsersDao) WithDeletedAt(deletedAt int32) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.DeletedAt] = deletedAt })
}

// 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *UsersDao) WithDeletedAts(deletedAts []int32) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.DeletedAt] = deletedAts })
}

// 函数选项模式获取单条记录
func (obj *UsersDao) GetByOption(opts ...Option) (result Users, err error) {
	obj.setIsUpdateSql(false)
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// 函数选项模式获取多条记录：支持分页
func (obj *UsersDao) GetByOptions(opts ...Option) (results []*Users, err error) {
	obj.setIsUpdateSql(false)
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *UsersDao) GetCountByOptions(opts ...Option) (count int64) {
	obj.setIsUpdateSql(false)
	obj.prepare(opts...).Model(&Users{}).Count(&count)
	return
}

// 更新数据
//  参数说明：
//      users: 要更新的数据
//      opts: 更新的条件
func (obj *UsersDao) UpdateByOption(users Users, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsUpdateSql(true)
	tx := obj.prepare(opts...).Updates(&users)
	rowsAffected = tx.RowsAffected
	err = tx.Error
	return
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// 通过单个 id(主键) 字段值，获取单条记录
func (obj *UsersDao) GetFromID(id int32) (result Users, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过多个 id(主键) 字段值，获取多条记录
func (obj *UsersDao) GetsFromID(ids []int32) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithIDs(ids))
	return
}

// 通过单个 name(名称) 字段值，获取多条记录
func (obj *UsersDao) GetFromName(name *string) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithName(name))
	return
}

// 通过多个 name(名称) 字段值，获取多条记录
func (obj *UsersDao) GetsFromName(names []*string) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithNames(names))
	return
}

// 通过单个 age(年龄) 字段值，获取多条记录
func (obj *UsersDao) GetFromAge(age int32) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithAge(age))
	return
}

// 通过多个 age(年龄) 字段值，获取多条记录
func (obj *UsersDao) GetsFromAge(ages []int32) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithAges(ages))
	return
}

// 通过单个 card_no(身份证) 字段值，获取单条记录
func (obj *UsersDao) GetFromCardNo(cardNo string) (result Users, err error) {
	result, err = obj.GetByOption(obj.WithCardNo(cardNo))
	return
}

// 通过多个 card_no(身份证) 字段值，获取多条记录
func (obj *UsersDao) GetsFromCardNo(cardNos []string) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithCardNos(cardNos))
	return
}

// 通过单个 head_img(头像) 字段值，获取多条记录
func (obj *UsersDao) GetFromHeadImg(headImg string) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithHeadImg(headImg))
	return
}

// 通过多个 head_img(头像) 字段值，获取多条记录
func (obj *UsersDao) GetsFromHeadImg(headImgs []string) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithHeadImgs(headImgs))
	return
}

// 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *UsersDao) GetFromCreatedAt(createdAt int32) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAt(createdAt))
	return
}

// 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *UsersDao) GetsFromCreatedAt(createdAts []int32) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAts(createdAts))
	return
}

// 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *UsersDao) GetFromUpdatedAt(updatedAt int32) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithUpdatedAt(updatedAt))
	return
}

// 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *UsersDao) GetsFromUpdatedAt(updatedAts []int32) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithUpdatedAts(updatedAts))
	return
}

// 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *UsersDao) GetFromDeletedAt(deletedAt int32) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithDeletedAt(deletedAt))
	return
}

// 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *UsersDao) GetsFromDeletedAt(deletedAts []int32) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// 通过 id 字段值，获取单条记录
func (obj *UsersDao) FetchByPrimaryKey(id int32) (result Users, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过 card_no 字段值，获取单条记录
func (obj *UsersDao) FetchUniqueByCardNo(cardNo string) (result Users, err error) {
	result, err = obj.GetByOption(obj.WithCardNo(cardNo))
	return
}

// 通过 name, card_no 字段值，获取单条记录
func (obj *UsersDao) FetchUniqueIndexByUnqNameCard(name *string, cardNo string) (result Users, err error) {
	result, err = obj.GetByOption(
		obj.WithName(name),
		obj.WithCardNo(cardNo))
	return
}

// 通过 age 字段值，获取多条记录
func (obj *UsersDao) FetchIndexByAge(age int32) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithAge(age))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
