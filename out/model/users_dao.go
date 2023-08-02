package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022-08-20 12:14:41
 * @Desc:   users 表的 DAO 层
 */

type UsersDAO struct {
	*_BaseDAO
}

// 初始化 UsersDAO
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

// 获取表名称
func (obj *UsersDAO) GetTableName() string {
	users := &Users{}
	return users.TableName()
}

// 存在则更新，否则插入
func (obj *UsersDAO) Save(users *Users) (rowsAffected int64, err error) {
	if users.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(users, obj.WithID(users.ID))
	}
	return obj.Create(users)
}

// 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *UsersDAO) Create(users interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(users)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// 设置 id(主键) 字段作为 option 条件
func (obj *UsersDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.ID] = id })
}

// 设置 id(主键) 字段的切片作为 option 条件
func (obj *UsersDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.ID] = ids })
}

// 设置 name(名称) 字段作为 option 条件
func (obj *UsersDAO) WithName(name *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.Name] = name })
}

// 设置 name(名称) 字段的切片作为 option 条件
func (obj *UsersDAO) WithNames(names []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.Name] = names })
}

// 设置 age(年龄) 字段作为 option 条件
func (obj *UsersDAO) WithAge(age int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.Age] = age })
}

// 设置 age(年龄) 字段的切片作为 option 条件
func (obj *UsersDAO) WithAges(ages []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.Age] = ages })
}

// 设置 card_no(身份证) 字段作为 option 条件
func (obj *UsersDAO) WithCardNo(cardNo string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.CardNo] = cardNo })
}

// 设置 card_no(身份证) 字段的切片作为 option 条件
func (obj *UsersDAO) WithCardNos(cardNos []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.CardNo] = cardNos })
}

// 设置 head_img(头像) 字段作为 option 条件
func (obj *UsersDAO) WithHeadImg(headImg string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.HeadImg] = headImg })
}

// 设置 head_img(头像) 字段的切片作为 option 条件
func (obj *UsersDAO) WithHeadImgs(headImgs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.HeadImg] = headImgs })
}

// 设置 created_at(创建时间) 字段作为 option 条件
func (obj *UsersDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.CreatedAt] = createdAt })
}

// 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *UsersDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.CreatedAt] = createdAts })
}

// 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *UsersDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.UpdatedAt] = updatedAt })
}

// 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *UsersDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.UpdatedAt] = updatedAts })
}

// 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *UsersDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.DeletedAt] = deletedAt })
}

// 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *UsersDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.DeletedAt] = deletedAts })
}

// 函数选项模式获取单条记录
func (obj *UsersDAO) GetByOption(opts ...Option) (result *Users, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// 函数选项模式获取单条记录
func (obj *UsersDAO) GetByOptionGLTID(gltID int64, opts ...Option) (result *Users, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Where(fmt.Sprintf("%s >= ?", UsersColumns.ID), gltID).Find(&result).Error
	return
}

// 函数选项模式获取多条记录：支持分页
func (obj *UsersDAO) GetByOptions(opts ...Option) (results []*Users, err error) {
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *UsersDAO) GetCountByOptions(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	obj.prepare(opts...).Count(&count)
	return
}

// 更新数据
//  参数说明：
//      users: 要更新的数据
//      opts: 更新的条件
func (obj *UsersDAO) UpdateByOption(users *Users, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(users)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// 通过单个 id(主键) 字段值，获取单条记录
func (obj *UsersDAO) GetFromID(id int64) (result *Users, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过多个 id(主键) 字段值，获取多条记录
func (obj *UsersDAO) GetsFromID(ids []int64) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithIDs(ids))
	return
}

// 通过单个 name(名称) 字段值，获取多条记录
func (obj *UsersDAO) GetFromName(name *string) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithName(name))
	return
}

// 通过多个 name(名称) 字段值，获取多条记录
func (obj *UsersDAO) GetsFromName(names []*string) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithNames(names))
	return
}

// 通过单个 age(年龄) 字段值，获取多条记录
func (obj *UsersDAO) GetFromAge(age int64) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithAge(age))
	return
}

// 通过多个 age(年龄) 字段值，获取多条记录
func (obj *UsersDAO) GetsFromAge(ages []int64) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithAges(ages))
	return
}

// 通过单个 card_no(身份证) 字段值，获取单条记录
func (obj *UsersDAO) GetFromCardNo(cardNo string) (result *Users, err error) {
	result, err = obj.GetByOption(obj.WithCardNo(cardNo))
	return
}

// 通过多个 card_no(身份证) 字段值，获取多条记录
func (obj *UsersDAO) GetsFromCardNo(cardNos []string) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithCardNos(cardNos))
	return
}

// 通过单个 head_img(头像) 字段值，获取多条记录
func (obj *UsersDAO) GetFromHeadImg(headImg string) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithHeadImg(headImg))
	return
}

// 通过多个 head_img(头像) 字段值，获取多条记录
func (obj *UsersDAO) GetsFromHeadImg(headImgs []string) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithHeadImgs(headImgs))
	return
}

// 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *UsersDAO) GetFromCreatedAt(createdAt int64) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAt(createdAt))
	return
}

// 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *UsersDAO) GetsFromCreatedAt(createdAts []int64) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAts(createdAts))
	return
}

// 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *UsersDAO) GetFromUpdatedAt(updatedAt int64) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithUpdatedAt(updatedAt))
	return
}

// 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *UsersDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithUpdatedAts(updatedAts))
	return
}

// 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *UsersDAO) GetFromDeletedAt(deletedAt int64) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithDeletedAt(deletedAt))
	return
}

// 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *UsersDAO) GetsFromDeletedAt(deletedAts []int64) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// 通过 id 字段值，获取单条记录
func (obj *UsersDAO) FetchByPrimaryKey(id int64) (result *Users, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过 card_no 字段值，获取单条记录
func (obj *UsersDAO) FetchUniqueByCardNo(cardNo string) (result *Users, err error) {
	result, err = obj.GetByOption(obj.WithCardNo(cardNo))
	return
}

// 通过 name, card_no 字段值，获取单条记录
func (obj *UsersDAO) FetchUniqueIndexByUnqNameCard(name *string, cardNo string) (result *Users, err error) {
	result, err = obj.GetByOption(
		obj.WithName(name),
		obj.WithCardNo(cardNo))
	return
}

// 通过 age 字段值，获取多条记录
func (obj *UsersDAO) FetchIndexByAge(age int64) (results []*Users, err error) {
	results, err = obj.GetByOptions(obj.WithAge(age))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
