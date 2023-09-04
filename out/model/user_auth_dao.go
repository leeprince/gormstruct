package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-09-05 02:17:30
 * @Desc:   user_auth 表的 DAO 层
 */

type UserAuthDAO struct {
	*_BaseDAO
}

// UserAuthDAO 初始化
func NewUserAuthDAO(ctx context.Context, db *gorm.DB) *UserAuthDAO {
	if db == nil {
		panic(fmt.Errorf("UserAuthDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &UserAuthDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&UserAuth{}),
			db:               db,
			model:            UserAuth{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          UserAuthAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *UserAuthDAO) GetTableName() string {
	userAuth := &UserAuth{}
	return userAuth.TableName()
}

// Save 存在则更新，否则插入
func (obj *UserAuthDAO) Save(userAuth *UserAuth) (rowsAffected int64, err error) {
	if userAuth.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(userAuth, obj.WithID(userAuth.ID))
	}
	return obj.Create(userAuth)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *UserAuthDAO) Create(userAuth interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(userAuth)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键) 字段作为 option 条件
func (obj *UserAuthDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserAuthColumns.ID] = id })
}

// WithIDs 设置 id(主键) 字段的切片作为 option 条件
func (obj *UserAuthDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserAuthColumns.ID] = ids })
}

// WithUserID 设置 user_id(用户ID) 字段作为 option 条件
func (obj *UserAuthDAO) WithUserID(userID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserAuthColumns.UserID] = userID })
}

// WithUserIDs 设置 user_id(用户ID) 字段的切片作为 option 条件
func (obj *UserAuthDAO) WithUserIDs(userIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserAuthColumns.UserID] = userIDs })
}

// WithExpireTime 设置 expire_time(授权过期时间) 字段作为 option 条件
func (obj *UserAuthDAO) WithExpireTime(expireTime time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserAuthColumns.ExpireTime] = expireTime })
}

// WithExpireTimes 设置 expire_time(授权过期时间) 字段的切片作为 option 条件
func (obj *UserAuthDAO) WithExpireTimes(expireTimes []time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserAuthColumns.ExpireTime] = expireTimes })
}

// WithAccessTime 设置 access_time(访问时间) 字段作为 option 条件
func (obj *UserAuthDAO) WithAccessTime(accessTime int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserAuthColumns.AccessTime] = accessTime })
}

// WithAccessTimes 设置 access_time(访问时间) 字段的切片作为 option 条件
func (obj *UserAuthDAO) WithAccessTimes(accessTimes []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserAuthColumns.AccessTime] = accessTimes })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *UserAuthDAO) WithUpdatedAt(updatedAt time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserAuthColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *UserAuthDAO) WithUpdatedAts(updatedAts []time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserAuthColumns.UpdatedAt] = updatedAts })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *UserAuthDAO) WithCreatedAt(createdAt time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserAuthColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *UserAuthDAO) WithCreatedAts(createdAts []time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserAuthColumns.CreatedAt] = createdAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *UserAuthDAO) WithDeletedAt(deletedAt *time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserAuthColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *UserAuthDAO) WithDeletedAts(deletedAts []*time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[UserAuthColumns.DeletedAt] = deletedAts })
}

// WithDeletedAtIsNull 设置 DeletedAt(删除标记) 字段为NULL作为 option 条件
func (obj *UserAuthDAO) WithDeletedAtIsNull() Option {
	return queryArgOptionFunc(func(o *options) { o.queryArg.query = fmt.Sprintf("%s IS NULL", UserAuthColumns.DeletedAt) })
}

// GetByOption 函数选项模式获取单条记录
func (obj *UserAuthDAO) GetByOption(opts ...Option) (result *UserAuth, err error) {
	opts = append(opts, obj.WithDeletedAtIsNull())
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *UserAuthDAO) GetListByOption(opts ...Option) (results []*UserAuth, err error) {
	opts = append(opts, obj.WithDeletedAtIsNull())
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *UserAuthDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *UserAuthDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
//
//		参数说明：
//	     userAuth: 要更新的数据
//	     opts: 更新的条件
func (obj *UserAuthDAO) UpdateByOption(userAuth *UserAuth, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(userAuth)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键) 字段值，获取单条记录
func (obj *UserAuthDAO) GetFromID(id int64) (result *UserAuth, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(主键) 字段值，获取多条记录
func (obj *UserAuthDAO) GetsFromID(ids []int64) (results []*UserAuth, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromUserID 通过单个 user_id(用户ID) 字段值，获取多条记录
func (obj *UserAuthDAO) GetFromUserID(userID int64) (results []*UserAuth, err error) {
	results, err = obj.GetListByOption(obj.WithUserID(userID))
	return
}

// GetsFromUserID 通过多个 user_id(用户ID) 字段值，获取多条记录
func (obj *UserAuthDAO) GetsFromUserID(userIDs []int64) (results []*UserAuth, err error) {
	results, err = obj.GetListByOption(obj.WithUserIDs(userIDs))
	return
}

// GetFromExpireTime 通过单个 expire_time(授权过期时间) 字段值，获取多条记录
func (obj *UserAuthDAO) GetFromExpireTime(expireTime time.Time) (results []*UserAuth, err error) {
	results, err = obj.GetListByOption(obj.WithExpireTime(expireTime))
	return
}

// GetsFromExpireTime 通过多个 expire_time(授权过期时间) 字段值，获取多条记录
func (obj *UserAuthDAO) GetsFromExpireTime(expireTimes []time.Time) (results []*UserAuth, err error) {
	results, err = obj.GetListByOption(obj.WithExpireTimes(expireTimes))
	return
}

// GetFromAccessTime 通过单个 access_time(访问时间) 字段值，获取多条记录
func (obj *UserAuthDAO) GetFromAccessTime(accessTime int64) (results []*UserAuth, err error) {
	results, err = obj.GetListByOption(obj.WithAccessTime(accessTime))
	return
}

// GetsFromAccessTime 通过多个 access_time(访问时间) 字段值，获取多条记录
func (obj *UserAuthDAO) GetsFromAccessTime(accessTimes []int64) (results []*UserAuth, err error) {
	results, err = obj.GetListByOption(obj.WithAccessTimes(accessTimes))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *UserAuthDAO) GetFromUpdatedAt(updatedAt time.Time) (results []*UserAuth, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *UserAuthDAO) GetsFromUpdatedAt(updatedAts []time.Time) (results []*UserAuth, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *UserAuthDAO) GetFromCreatedAt(createdAt time.Time) (results []*UserAuth, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *UserAuthDAO) GetsFromCreatedAt(createdAts []time.Time) (results []*UserAuth, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *UserAuthDAO) GetFromDeletedAt(deletedAt *time.Time) (results []*UserAuth, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *UserAuthDAO) GetsFromDeletedAt(deletedAts []*time.Time) (results []*UserAuth, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *UserAuthDAO) FetchByPrimaryKey(id int64) (result *UserAuth, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
