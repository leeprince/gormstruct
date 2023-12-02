package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-12-02 23:58:01
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

// Save 存在则更新，否则插入
func (obj *UsersDAO) Save(users *Users) (rowsAffected int64, err error) {
	if users.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(users, obj.WithID(users.ID))
	}
	return obj.Create(users)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *UsersDAO) Create(users interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(users)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键) 字段作为 option 条件
func (obj *UsersDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.ID] = id })
}

// WithIDs 设置 id(主键) 字段的切片作为 option 条件
func (obj *UsersDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.ID] = ids })
}

// WithName 设置 name(名称) 字段作为 option 条件
func (obj *UsersDAO) WithName(name *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UsersColumns.Name] = name })
}

// WithNames 设置 name(名称) 字段的切片作为 option 条件
func (obj *UsersDAO) WithNames(names []*string) Option {
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

// GetsFromID 通过多个 id(主键) 字段值，获取多条记录
func (obj *UsersDAO) GetsFromID(ids []int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromName 通过单个 name(名称) 字段值，获取单条记录
func (obj *UsersDAO) GetFromName(name *string) (result *Users, err error) {
	result, err = obj.GetByOption(obj.WithName(name))
	return
}

// GetsFromName 通过多个 name(名称) 字段值，获取多条记录
func (obj *UsersDAO) GetsFromName(names []*string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithNames(names))
	return
}

// GetFromAge 通过单个 age(年龄) 字段值，获取多条记录
func (obj *UsersDAO) GetFromAge(age int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithAge(age))
	return
}

// GetsFromAge 通过多个 age(年龄) 字段值，获取多条记录
func (obj *UsersDAO) GetsFromAge(ages []int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithAges(ages))
	return
}

// GetFromCardNo 通过单个 card_no(身份证) 字段值，获取单条记录
func (obj *UsersDAO) GetFromCardNo(cardNo string) (result *Users, err error) {
	result, err = obj.GetByOption(obj.WithCardNo(cardNo))
	return
}

// GetsFromCardNo 通过多个 card_no(身份证) 字段值，获取多条记录
func (obj *UsersDAO) GetsFromCardNo(cardNos []string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithCardNos(cardNos))
	return
}

// GetFromHeadImg 通过单个 head_img(头像) 字段值，获取多条记录
func (obj *UsersDAO) GetFromHeadImg(headImg string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithHeadImg(headImg))
	return
}

// GetsFromHeadImg 通过多个 head_img(头像) 字段值，获取多条记录
func (obj *UsersDAO) GetsFromHeadImg(headImgs []string) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithHeadImgs(headImgs))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *UsersDAO) GetFromCreatedAt(createdAt int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *UsersDAO) GetsFromCreatedAt(createdAts []int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *UsersDAO) GetFromUpdatedAt(updatedAt int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *UsersDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *UsersDAO) GetFromDeletedAt(deletedAt int64) (results []*Users, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *UsersDAO) GetsFromDeletedAt(deletedAts []int64) (results []*Users, err error) {
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
func (obj *UsersDAO) FetchUniqueIndexByUnqNameCard(name *string, cardNo string) (result *Users, err error) {
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
