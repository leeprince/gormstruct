package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-08-30 01:28:13
 * @Desc:   licence 表的 DAO 层
 */

type LicenceDAO struct {
	*_BaseDAO
}

// 初始化 LicenceDAO
func NewLicenceDAO(ctx context.Context, db *gorm.DB) *LicenceDAO {
	if db == nil {
		panic(fmt.Errorf("LicenceDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &LicenceDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&Licence{}),
			db:               db,
			model:            Licence{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          LicenceAllColumns,
			isDefaultColumns: true,
		},
	}
}

// 获取表名称
func (obj *LicenceDAO) GetTableName() string {
	licence := &Licence{}
	return licence.TableName()
}

// 存在则更新，否则插入
func (obj *LicenceDAO) Save(licence *Licence) (rowsAffected int64, err error) {
	if licence.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(licence, obj.WithID(licence.ID))
	}
	return obj.Create(licence)
}

// 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *LicenceDAO) Create(licence interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(licence)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// 设置 id(主键ID) 字段作为 option 条件
func (obj *LicenceDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[LicenceColumns.ID] = id })
}

// 设置 id(主键ID) 字段的切片作为 option 条件
func (obj *LicenceDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[LicenceColumns.ID] = ids })
}

// 设置 client_ip(客户端IP) 字段作为 option 条件
func (obj *LicenceDAO) WithClientIP(clientIP *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[LicenceColumns.ClientIP] = clientIP })
}

// 设置 client_ip(客户端IP) 字段的切片作为 option 条件
func (obj *LicenceDAO) WithClientIPs(clientIPs []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[LicenceColumns.ClientIP] = clientIPs })
}

// 设置 machine_name(设备名称) 字段作为 option 条件
func (obj *LicenceDAO) WithMachineName(machineName *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[LicenceColumns.MachineName] = machineName })
}

// 设置 machine_name(设备名称) 字段的切片作为 option 条件
func (obj *LicenceDAO) WithMachineNames(machineNames []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[LicenceColumns.MachineName] = machineNames })
}

// 设置 expire_time(授权设备过期时间) 字段作为 option 条件
func (obj *LicenceDAO) WithExpireTime(expireTime time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[LicenceColumns.ExpireTime] = expireTime })
}

// 设置 expire_time(授权设备过期时间) 字段的切片作为 option 条件
func (obj *LicenceDAO) WithExpireTimes(expireTimes []time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[LicenceColumns.ExpireTime] = expireTimes })
}

// 设置 deleted_at() 字段作为 option 条件
func (obj *LicenceDAO) WithDeletedAt(deletedAt time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[LicenceColumns.DeletedAt] = deletedAt })
}

// 设置 deleted_at() 字段的切片作为 option 条件
func (obj *LicenceDAO) WithDeletedAts(deletedAts []time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[LicenceColumns.DeletedAt] = deletedAts })
}

// 设置 updated_at() 字段作为 option 条件
func (obj *LicenceDAO) WithUpdatedAt(updatedAt time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[LicenceColumns.UpdatedAt] = updatedAt })
}

// 设置 updated_at() 字段的切片作为 option 条件
func (obj *LicenceDAO) WithUpdatedAts(updatedAts []time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[LicenceColumns.UpdatedAt] = updatedAts })
}

// 设置 created_at() 字段作为 option 条件
func (obj *LicenceDAO) WithCreatedAt(createdAt time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[LicenceColumns.CreatedAt] = createdAt })
}

// 设置 created_at() 字段的切片作为 option 条件
func (obj *LicenceDAO) WithCreatedAts(createdAts []time.Time) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[LicenceColumns.CreatedAt] = createdAts })
}

// 函数选项模式获取单条记录
func (obj *LicenceDAO) GetByOption(opts ...Option) (result *Licence, err error) {
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// 函数选项模式获取多条记录：支持分页
func (obj *LicenceDAO) GetByOptions(opts ...Option) (results []*Licence, err error) {
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *LicenceDAO) GetCountByOptions(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	obj.prepare(opts...).Count(&count)
	return
}

// 更新数据
//
//	参数说明：
//	    licence: 要更新的数据
//	    opts: 更新的条件
func (obj *LicenceDAO) UpdateByOption(licence *Licence, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(licence)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// 通过单个 id(主键ID) 字段值，获取单条记录
func (obj *LicenceDAO) GetFromID(id int64) (result *Licence, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过多个 id(主键ID) 字段值，获取多条记录
func (obj *LicenceDAO) GetsFromID(ids []int64) (results []*Licence, err error) {
	results, err = obj.GetByOptions(obj.WithIDs(ids))
	return
}

// 通过单个 client_ip(客户端IP) 字段值，获取多条记录
func (obj *LicenceDAO) GetFromClientIP(clientIP *string) (results []*Licence, err error) {
	results, err = obj.GetByOptions(obj.WithClientIP(clientIP))
	return
}

// 通过多个 client_ip(客户端IP) 字段值，获取多条记录
func (obj *LicenceDAO) GetsFromClientIP(clientIPs []*string) (results []*Licence, err error) {
	results, err = obj.GetByOptions(obj.WithClientIPs(clientIPs))
	return
}

// 通过单个 machine_name(设备名称) 字段值，获取单条记录
func (obj *LicenceDAO) GetFromMachineName(machineName *string) (result *Licence, err error) {
	result, err = obj.GetByOption(obj.WithMachineName(machineName))
	return
}

// 通过多个 machine_name(设备名称) 字段值，获取多条记录
func (obj *LicenceDAO) GetsFromMachineName(machineNames []*string) (results []*Licence, err error) {
	results, err = obj.GetByOptions(obj.WithMachineNames(machineNames))
	return
}

// 通过单个 expire_time(授权设备过期时间) 字段值，获取多条记录
func (obj *LicenceDAO) GetFromExpireTime(expireTime time.Time) (results []*Licence, err error) {
	results, err = obj.GetByOptions(obj.WithExpireTime(expireTime))
	return
}

// 通过多个 expire_time(授权设备过期时间) 字段值，获取多条记录
func (obj *LicenceDAO) GetsFromExpireTime(expireTimes []time.Time) (results []*Licence, err error) {
	results, err = obj.GetByOptions(obj.WithExpireTimes(expireTimes))
	return
}

// 通过单个 deleted_at() 字段值，获取多条记录
func (obj *LicenceDAO) GetFromDeletedAt(deletedAt time.Time) (results []*Licence, err error) {
	results, err = obj.GetByOptions(obj.WithDeletedAt(deletedAt))
	return
}

// 通过多个 deleted_at() 字段值，获取多条记录
func (obj *LicenceDAO) GetsFromDeletedAt(deletedAts []time.Time) (results []*Licence, err error) {
	results, err = obj.GetByOptions(obj.WithDeletedAts(deletedAts))
	return
}

// 通过单个 updated_at() 字段值，获取多条记录
func (obj *LicenceDAO) GetFromUpdatedAt(updatedAt time.Time) (results []*Licence, err error) {
	results, err = obj.GetByOptions(obj.WithUpdatedAt(updatedAt))
	return
}

// 通过多个 updated_at() 字段值，获取多条记录
func (obj *LicenceDAO) GetsFromUpdatedAt(updatedAts []time.Time) (results []*Licence, err error) {
	results, err = obj.GetByOptions(obj.WithUpdatedAts(updatedAts))
	return
}

// 通过单个 created_at() 字段值，获取多条记录
func (obj *LicenceDAO) GetFromCreatedAt(createdAt time.Time) (results []*Licence, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAt(createdAt))
	return
}

// 通过多个 created_at() 字段值，获取多条记录
func (obj *LicenceDAO) GetsFromCreatedAt(createdAts []time.Time) (results []*Licence, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAts(createdAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// 通过 id 字段值，获取单条记录
func (obj *LicenceDAO) FetchByPrimaryKey(id int64) (result *Licence, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过 machine_name 字段值，获取单条记录
func (obj *LicenceDAO) FetchUniqueByUniqLicenceMachineName(machineName *string) (result *Licence, err error) {
	result, err = obj.GetByOption(obj.WithMachineName(machineName))
	return
}

// 通过 deleted_at 字段值，获取多条记录
func (obj *LicenceDAO) FetchIndexByIDxLicenceDeletedAt(deletedAt time.Time) (results []*Licence, err error) {
	results, err = obj.GetByOptions(obj.WithDeletedAt(deletedAt))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
