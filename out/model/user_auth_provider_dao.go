package model

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-07-26 15:20:11
 * @Desc:   user_auth_provider 表的 DAO 层
 */

type UserAuthProviderDAO struct {
	*_BaseDAO
}

// NewUserAuthProviderDAO 初始化
func NewUserAuthProviderDAO(ctx context.Context, db *gorm.DB) *UserAuthProviderDAO {
	if db == nil {
		panic(fmt.Errorf("UserAuthProviderDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &UserAuthProviderDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&UserAuthProvider{}),
			db:               db,
			model:            UserAuthProvider{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          UserAuthProviderAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *UserAuthProviderDAO) GetTableName() string {
	userAuthProvider := &UserAuthProvider{}
	return userAuthProvider.TableName()
}

// UpdateOrCreate 存在则更新，否则插入，会忽略零值字段
func (obj *UserAuthProviderDAO) UpdateOrCreate(userAuthProvider *UserAuthProvider) (rowsAffected int64, err error) {
	if userAuthProvider.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(userAuthProvider, obj.WithPrimaryKey(userAuthProvider.PrimaryKeyValue()))
	}
	return obj.Create(userAuthProvider)
}

// Save gorm 原生的 Save 会保存所有的字段，即使字段是零值。仅会判断主键是否存在，存在则更新，不存在则创建
func (obj *UserAuthProviderDAO) Save(userAuthProvider *UserAuthProvider) (rowsAffected int64, err error) {
	tx := obj.db.Save(userAuthProvider)
	return tx.RowsAffected, tx.Error
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *UserAuthProviderDAO) Create(userAuthProvider interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(userAuthProvider)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithPrimaryKey 设置真正的主键 字段作为 option 条件
func (obj *UserAuthProviderDAO) WithPrimaryKey(primaryKeyValue interface{}) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[(&UserAuthProvider{}).PrimaryKey()] = primaryKeyValue })
}

// WithID 设置 id(主键ID) 字段作为 option 条件
func (obj *UserAuthProviderDAO) WithID(id int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.ID] = id })
}

// WithIDs 设置 id(主键ID) 字段的切片作为 option 条件
func (obj *UserAuthProviderDAO) WithIDs(ids []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.ID] = ids })
}

// WithUserID 设置 user_id(用户ID,user_info的主键) 字段作为 option 条件
func (obj *UserAuthProviderDAO) WithUserID(userID int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.UserID] = userID })
}

// WithUserIDs 设置 user_id(用户ID,user_info的主键) 字段的切片作为 option 条件
func (obj *UserAuthProviderDAO) WithUserIDs(userIDs []int64) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.UserID] = userIDs })
}

// WithProvider 设置 provider(第三方登录提供商。wechatopen：微信开放平台；) 字段作为 option 条件
func (obj *UserAuthProviderDAO) WithProvider(provider string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.Provider] = provider })
}

// WithProviders 设置 provider(第三方登录提供商。wechatopen：微信开放平台；) 字段的切片作为 option 条件
func (obj *UserAuthProviderDAO) WithProviders(providers []string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.Provider] = providers })
}

// WithProviderUserID 设置 provider_user_id(第三方登录提供商用户ID) 字段作为 option 条件
func (obj *UserAuthProviderDAO) WithProviderUserID(providerUserID string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.ProviderUserID] = providerUserID })
}

// WithProviderUserIDs 设置 provider_user_id(第三方登录提供商用户ID) 字段的切片作为 option 条件
func (obj *UserAuthProviderDAO) WithProviderUserIDs(providerUserIDs []string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.ProviderUserID] = providerUserIDs })
}

// WithProviderUserUnionid 设置 provider_user_unionid(第三方登录提供商在开放平台下统一的用户ID) 字段作为 option 条件
func (obj *UserAuthProviderDAO) WithProviderUserUnionid(providerUserUnionid string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.ProviderUserUnionid] = providerUserUnionid })
}

// WithProviderUserUnionids 设置 provider_user_unionid(第三方登录提供商在开放平台下统一的用户ID) 字段的切片作为 option 条件
func (obj *UserAuthProviderDAO) WithProviderUserUnionids(providerUserUnionids []string) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.ProviderUserUnionid] = providerUserUnionids })
}

// WithProviderData 设置 provider_data(第三方登录提供商用户数据) 字段作为 option 条件
func (obj *UserAuthProviderDAO) WithProviderData(providerData datatypes.JSON) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.ProviderData] = providerData })
}

// WithProviderDatas 设置 provider_data(第三方登录提供商用户数据) 字段的切片作为 option 条件
func (obj *UserAuthProviderDAO) WithProviderDatas(providerDatas []datatypes.JSON) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.ProviderData] = providerDatas })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *UserAuthProviderDAO) WithCreatedAt(createdAt time.Time) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *UserAuthProviderDAO) WithCreatedAts(createdAts []time.Time) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *UserAuthProviderDAO) WithUpdatedAt(updatedAt time.Time) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *UserAuthProviderDAO) WithUpdatedAts(updatedAts []time.Time) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *UserAuthProviderDAO) WithDeletedAt(deletedAt *time.Time) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *UserAuthProviderDAO) WithDeletedAts(deletedAts []*time.Time) Option {
	return queryMapOptionFunc(func(o *options) { o.queryMap[UserAuthProviderColumns.DeletedAt] = deletedAts })
}

// WithDeletedAtIsNull 设置 DeletedAt(删除标记) 字段为NULL作为 option 条件
func (obj *UserAuthProviderDAO) WithDeletedAtIsNull() Option {
	return queryArgListOptionFunc(func(o *options) {
		o.queryArgList = append(o.queryArgList, queryArg{query: fmt.Sprintf("%s IS NULL", UserAuthProviderColumns.DeletedAt), arg: nil})
	})
}

// GetByOption 函数选项模式获取单条记录
func (obj *UserAuthProviderDAO) GetByOption(opts ...Option) (result *UserAuthProvider, err error) {
	opts = append(opts, obj.WithDeletedAtIsNull())
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *UserAuthProviderDAO) GetListByOption(opts ...Option) (results []*UserAuthProvider, err error) {
	opts = append(opts, obj.WithDeletedAtIsNull())
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *UserAuthProviderDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAtIsNull())
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *UserAuthProviderDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
//
//		参数说明：
//	     userAuthProvider: 要更新的数据
//	     opts: 更新的条件
func (obj *UserAuthProviderDAO) UpdateByOption(userAuthProvider *UserAuthProvider, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(userAuthProvider)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键ID) 字段值，获取单条记录
func (obj *UserAuthProviderDAO) GetFromID(id int64) (result *UserAuthProvider, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetListFromID 通过多个 id(主键ID) 字段值，获取多条记录
func (obj *UserAuthProviderDAO) GetListFromID(ids []int64) (results []*UserAuthProvider, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromUserID 通过单个 user_id(用户ID,user_info的主键) 字段值，获取多条记录
func (obj *UserAuthProviderDAO) GetFromUserID(userID int64) (results []*UserAuthProvider, err error) {
	results, err = obj.GetListByOption(obj.WithUserID(userID))
	return
}

// GetListFromUserID 通过多个 user_id(用户ID,user_info的主键) 字段值，获取多条记录
func (obj *UserAuthProviderDAO) GetListFromUserID(userIDs []int64) (results []*UserAuthProvider, err error) {
	results, err = obj.GetListByOption(obj.WithUserIDs(userIDs))
	return
}

// GetFromProvider 通过单个 provider(第三方登录提供商。wechatopen：微信开放平台；) 字段值，获取单条记录
func (obj *UserAuthProviderDAO) GetFromProvider(provider string) (result *UserAuthProvider, err error) {
	result, err = obj.GetByOption(obj.WithProvider(provider))
	return
}

// GetListFromProvider 通过多个 provider(第三方登录提供商。wechatopen：微信开放平台；) 字段值，获取多条记录
func (obj *UserAuthProviderDAO) GetListFromProvider(providers []string) (results []*UserAuthProvider, err error) {
	results, err = obj.GetListByOption(obj.WithProviders(providers))
	return
}

// GetFromProviderUserID 通过单个 provider_user_id(第三方登录提供商用户ID) 字段值，获取单条记录
func (obj *UserAuthProviderDAO) GetFromProviderUserID(providerUserID string) (result *UserAuthProvider, err error) {
	result, err = obj.GetByOption(obj.WithProviderUserID(providerUserID))
	return
}

// GetListFromProviderUserID 通过多个 provider_user_id(第三方登录提供商用户ID) 字段值，获取多条记录
func (obj *UserAuthProviderDAO) GetListFromProviderUserID(providerUserIDs []string) (results []*UserAuthProvider, err error) {
	results, err = obj.GetListByOption(obj.WithProviderUserIDs(providerUserIDs))
	return
}

// GetFromProviderUserUnionid 通过单个 provider_user_unionid(第三方登录提供商在开放平台下统一的用户ID) 字段值，获取多条记录
func (obj *UserAuthProviderDAO) GetFromProviderUserUnionid(providerUserUnionid string) (results []*UserAuthProvider, err error) {
	results, err = obj.GetListByOption(obj.WithProviderUserUnionid(providerUserUnionid))
	return
}

// GetListFromProviderUserUnionid 通过多个 provider_user_unionid(第三方登录提供商在开放平台下统一的用户ID) 字段值，获取多条记录
func (obj *UserAuthProviderDAO) GetListFromProviderUserUnionid(providerUserUnionids []string) (results []*UserAuthProvider, err error) {
	results, err = obj.GetListByOption(obj.WithProviderUserUnionids(providerUserUnionids))
	return
}

// GetFromProviderData 通过单个 provider_data(第三方登录提供商用户数据) 字段值，获取多条记录
func (obj *UserAuthProviderDAO) GetFromProviderData(providerData datatypes.JSON) (results []*UserAuthProvider, err error) {
	results, err = obj.GetListByOption(obj.WithProviderData(providerData))
	return
}

// GetListFromProviderData 通过多个 provider_data(第三方登录提供商用户数据) 字段值，获取多条记录
func (obj *UserAuthProviderDAO) GetListFromProviderData(providerDatas []datatypes.JSON) (results []*UserAuthProvider, err error) {
	results, err = obj.GetListByOption(obj.WithProviderDatas(providerDatas))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *UserAuthProviderDAO) GetFromCreatedAt(createdAt time.Time) (results []*UserAuthProvider, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetListFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *UserAuthProviderDAO) GetListFromCreatedAt(createdAts []time.Time) (results []*UserAuthProvider, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *UserAuthProviderDAO) GetFromUpdatedAt(updatedAt time.Time) (results []*UserAuthProvider, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetListFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *UserAuthProviderDAO) GetListFromUpdatedAt(updatedAts []time.Time) (results []*UserAuthProvider, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *UserAuthProviderDAO) GetFromDeletedAt(deletedAt *time.Time) (results []*UserAuthProvider, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetListFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *UserAuthProviderDAO) GetListFromDeletedAt(deletedAts []*time.Time) (results []*UserAuthProvider, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *UserAuthProviderDAO) FetchByPrimaryKey(id int64) (result *UserAuthProvider, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// FetchUniqueIndexByUniqPPuid 通过 provider, provider_user_id 字段值，获取单条记录
func (obj *UserAuthProviderDAO) FetchUniqueIndexByUniqPPuid(provider string, providerUserID string) (result *UserAuthProvider, err error) {
	result, err = obj.GetByOption(
		obj.WithProvider(provider),
		obj.WithProviderUserID(providerUserID))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
