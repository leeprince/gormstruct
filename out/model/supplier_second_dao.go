package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:47:03
 * @Desc:   supplier_second 表的 DAO 层
 */

type SupplierSecondDAO struct {
	*_BaseDAO
}

// SupplierSecondDAO 初始化
func NewSupplierSecondDAO(ctx context.Context, db *gorm.DB) *SupplierSecondDAO {
	if db == nil {
		panic(fmt.Errorf("SupplierSecondDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &SupplierSecondDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&SupplierSecond{}),
			db:               db,
			model:            SupplierSecond{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          SupplierSecondAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *SupplierSecondDAO) GetTableName() string {
	supplierSecond := &SupplierSecond{}
	return supplierSecond.TableName()
}

// Save 存在则更新，否则插入
func (obj *SupplierSecondDAO) Save(supplierSecond *SupplierSecond) (rowsAffected int64, err error) {
	if supplierSecond.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(supplierSecond, obj.WithID(supplierSecond.ID))
	}
	return obj.Create(supplierSecond)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *SupplierSecondDAO) Create(supplierSecond interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(supplierSecond)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键id) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.ID] = id })
}

// WithIDs 设置 id(主键id) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.ID] = ids })
}

// WithOpenOrgIDs 设置 open_org_ids(一级供应商（平台企业）租户id集合) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithOpenOrgIDs(openOrgIDs string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.OpenOrgIDs] = openOrgIDs })
}

// WithOpenOrgIDss 设置 open_org_ids(一级供应商（平台企业）租户id集合) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithOpenOrgIDss(openOrgIDss []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.OpenOrgIDs] = openOrgIDss })
}

// WithSupplierID 设置 supplier_id(供应商在平台侧的唯一标识 主键-按照它来更新或插入) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithSupplierID(supplierID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.SupplierID] = supplierID })
}

// WithSupplierIDs 设置 supplier_id(供应商在平台侧的唯一标识 主键-按照它来更新或插入) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithSupplierIDs(supplierIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.SupplierID] = supplierIDs })
}

// WithRegisterDate 设置 register_date(供应商注册时间) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithRegisterDate(registerDate int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.RegisterDate] = registerDate })
}

// WithRegisterDates 设置 register_date(供应商注册时间) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithRegisterDates(registerDates []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.RegisterDate] = registerDates })
}

// WithState 设置 state(供应商状态：0-异常 1-正常 2-退出) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithState(state int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.State] = state })
}

// WithStates 设置 state(供应商状态：0-异常 1-正常 2-退出) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithStates(states []int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.State] = states })
}

// WithStateDesc 设置 state_desc(供应商状态说明) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithStateDesc(stateDesc string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.StateDesc] = stateDesc })
}

// WithStateDescs 设置 state_desc(供应商状态说明) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithStateDescs(stateDescs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.StateDesc] = stateDescs })
}

// WithName 设置 name(供应商姓名) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithName(name string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.Name] = name })
}

// WithNames 设置 name(供应商姓名) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithNames(names []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.Name] = names })
}

// WithCardType 设置 card_type(供应商身份证件种类) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithCardType(cardType string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.CardType] = cardType })
}

// WithCardTypes 设置 card_type(供应商身份证件种类) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithCardTypes(cardTypes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.CardType] = cardTypes })
}

// WithCardNo 设置 card_no(身份证号) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithCardNo(cardNo string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.CardNo] = cardNo })
}

// WithCardNos 设置 card_no(身份证号) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithCardNos(cardNos []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.CardNo] = cardNos })
}

// WithPhone 设置 phone(手机号码) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithPhone(phone string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.Phone] = phone })
}

// WithPhones 设置 phone(手机号码) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithPhones(phones []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.Phone] = phones })
}

// WithIDentityType 设置 identity_type(供应商业务身份标识) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithIDentityType(identityType string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.IDentityType] = identityType })
}

// WithIDentityTypes 设置 identity_type(供应商业务身份标识) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithIDentityTypes(identityTypes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.IDentityType] = identityTypes })
}

// WithCardFaceURL 设置 card_face_url(身份证人像面) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithCardFaceURL(cardFaceURL string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.CardFaceURL] = cardFaceURL })
}

// WithCardFaceURLs 设置 card_face_url(身份证人像面) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithCardFaceURLs(cardFaceURLs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.CardFaceURL] = cardFaceURLs })
}

// WithCardEmblemURL 设置 card_emblem_url(身份证国徽面) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithCardEmblemURL(cardEmblemURL string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.CardEmblemURL] = cardEmblemURL })
}

// WithCardEmblemURLs 设置 card_emblem_url(身份证国徽面) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithCardEmblemURLs(cardEmblemURLs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.CardEmblemURL] = cardEmblemURLs })
}

// WithIsIDentity 设置 is_identity(是否完成实名认证：0-否 1-是) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithIsIDentity(isIDentity int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.IsIDentity] = isIDentity })
}

// WithIsIDentitys 设置 is_identity(是否完成实名认证：0-否 1-是) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithIsIDentitys(isIDentitys []int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.IsIDentity] = isIDentitys })
}

// WithIDentityDate 设置 identity_date(实名认证时间) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithIDentityDate(identityDate int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.IDentityDate] = identityDate })
}

// WithIDentityDates 设置 identity_date(实名认证时间) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithIDentityDates(identityDates []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.IDentityDate] = identityDates })
}

// WithIDentityState 设置 identity_state(实名认证状态：1 一致 2 不一致 3 无记录 -1 异常状态) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithIDentityState(identityState int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.IDentityState] = identityState })
}

// WithIDentityStates 设置 identity_state(实名认证状态：1 一致 2 不一致 3 无记录 -1 异常状态) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithIDentityStates(identityStates []int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.IDentityState] = identityStates })
}

// WithIDentityDesc 设置 identity_desc(实名认证结果说明) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithIDentityDesc(identityDesc string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.IDentityDesc] = identityDesc })
}

// WithIDentityDescs 设置 identity_desc(实名认证结果说明) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithIDentityDescs(identityDescs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.IDentityDesc] = identityDescs })
}

// WithIDentityUpdateTime 设置 identity_update_time(实名认证更新时间，时间戳) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithIDentityUpdateTime(identityUpdateTime int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.IDentityUpdateTime] = identityUpdateTime })
}

// WithIDentityUpdateTimes 设置 identity_update_time(实名认证更新时间，时间戳) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithIDentityUpdateTimes(identityUpdateTimes []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.IDentityUpdateTime] = identityUpdateTimes })
}

// WithCleanProtocalFile 设置 clean_protocal_file(是否清理协议文件：1：是 非1：否) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithCleanProtocalFile(cleanProtocalFile int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.CleanProtocalFile] = cleanProtocalFile })
}

// WithCleanProtocalFiles 设置 clean_protocal_file(是否清理协议文件：1：是 非1：否) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithCleanProtocalFiles(cleanProtocalFiles []int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.CleanProtocalFile] = cleanProtocalFiles })
}

// WithCleanAccounts 设置 clean_accounts(是否清理账户信息：1：是 非1：否) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithCleanAccounts(cleanAccounts int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.CleanAccounts] = cleanAccounts })
}

// WithCleanAccountss 设置 clean_accounts(是否清理账户信息：1：是 非1：否) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithCleanAccountss(cleanAccountss []int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.CleanAccounts] = cleanAccountss })
}

// WithCreatedAt 设置 created_at(创建时间) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.CreatedAt] = createdAt })
}

// WithCreatedAts 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.CreatedAt] = createdAts })
}

// WithUpdatedAt 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.UpdatedAt] = updatedAt })
}

// WithUpdatedAts 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.UpdatedAt] = updatedAts })
}

// WithDeletedAt 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *SupplierSecondDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.DeletedAt] = deletedAt })
}

// WithDeletedAts 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *SupplierSecondDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[SupplierSecondColumns.DeletedAt] = deletedAts })
}

// GetByOption 函数选项模式获取单条记录
func (obj *SupplierSecondDAO) GetByOption(opts ...Option) (result *SupplierSecond, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *SupplierSecondDAO) GetListByOption(opts ...Option) (results []*SupplierSecond, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *SupplierSecondDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	opts = append(opts, obj.WithDeletedAt(0))
	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *SupplierSecondDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      supplierSecond: 要更新的数据
//      opts: 更新的条件
func (obj *SupplierSecondDAO) UpdateByOption(supplierSecond *SupplierSecond, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(supplierSecond)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键id) 字段值，获取单条记录
func (obj *SupplierSecondDAO) GetFromID(id int64) (result *SupplierSecond, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(主键id) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromID(ids []int64) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromOpenOrgIDs 通过单个 open_org_ids(一级供应商（平台企业）租户id集合) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromOpenOrgIDs(openOrgIDs string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgIDs(openOrgIDs))
	return
}

// GetsFromOpenOrgIDs 通过多个 open_org_ids(一级供应商（平台企业）租户id集合) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromOpenOrgIDs(openOrgIDss []string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgIDss(openOrgIDss))
	return
}

// GetFromSupplierID 通过单个 supplier_id(供应商在平台侧的唯一标识 主键-按照它来更新或插入) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromSupplierID(supplierID string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithSupplierID(supplierID))
	return
}

// GetsFromSupplierID 通过多个 supplier_id(供应商在平台侧的唯一标识 主键-按照它来更新或插入) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromSupplierID(supplierIDs []string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithSupplierIDs(supplierIDs))
	return
}

// GetFromRegisterDate 通过单个 register_date(供应商注册时间) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromRegisterDate(registerDate int64) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithRegisterDate(registerDate))
	return
}

// GetsFromRegisterDate 通过多个 register_date(供应商注册时间) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromRegisterDate(registerDates []int64) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithRegisterDates(registerDates))
	return
}

// GetFromState 通过单个 state(供应商状态：0-异常 1-正常 2-退出) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromState(state int) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithState(state))
	return
}

// GetsFromState 通过多个 state(供应商状态：0-异常 1-正常 2-退出) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromState(states []int) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithStates(states))
	return
}

// GetFromStateDesc 通过单个 state_desc(供应商状态说明) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromStateDesc(stateDesc string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithStateDesc(stateDesc))
	return
}

// GetsFromStateDesc 通过多个 state_desc(供应商状态说明) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromStateDesc(stateDescs []string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithStateDescs(stateDescs))
	return
}

// GetFromName 通过单个 name(供应商姓名) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromName(name string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithName(name))
	return
}

// GetsFromName 通过多个 name(供应商姓名) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromName(names []string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithNames(names))
	return
}

// GetFromCardType 通过单个 card_type(供应商身份证件种类) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromCardType(cardType string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithCardType(cardType))
	return
}

// GetsFromCardType 通过多个 card_type(供应商身份证件种类) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromCardType(cardTypes []string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithCardTypes(cardTypes))
	return
}

// GetFromCardNo 通过单个 card_no(身份证号) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromCardNo(cardNo string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithCardNo(cardNo))
	return
}

// GetsFromCardNo 通过多个 card_no(身份证号) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromCardNo(cardNos []string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithCardNos(cardNos))
	return
}

// GetFromPhone 通过单个 phone(手机号码) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromPhone(phone string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithPhone(phone))
	return
}

// GetsFromPhone 通过多个 phone(手机号码) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromPhone(phones []string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithPhones(phones))
	return
}

// GetFromIDentityType 通过单个 identity_type(供应商业务身份标识) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromIDentityType(identityType string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithIDentityType(identityType))
	return
}

// GetsFromIDentityType 通过多个 identity_type(供应商业务身份标识) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromIDentityType(identityTypes []string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithIDentityTypes(identityTypes))
	return
}

// GetFromCardFaceURL 通过单个 card_face_url(身份证人像面) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromCardFaceURL(cardFaceURL string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithCardFaceURL(cardFaceURL))
	return
}

// GetsFromCardFaceURL 通过多个 card_face_url(身份证人像面) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromCardFaceURL(cardFaceURLs []string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithCardFaceURLs(cardFaceURLs))
	return
}

// GetFromCardEmblemURL 通过单个 card_emblem_url(身份证国徽面) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromCardEmblemURL(cardEmblemURL string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithCardEmblemURL(cardEmblemURL))
	return
}

// GetsFromCardEmblemURL 通过多个 card_emblem_url(身份证国徽面) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromCardEmblemURL(cardEmblemURLs []string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithCardEmblemURLs(cardEmblemURLs))
	return
}

// GetFromIsIDentity 通过单个 is_identity(是否完成实名认证：0-否 1-是) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromIsIDentity(isIDentity int) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithIsIDentity(isIDentity))
	return
}

// GetsFromIsIDentity 通过多个 is_identity(是否完成实名认证：0-否 1-是) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromIsIDentity(isIDentitys []int) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithIsIDentitys(isIDentitys))
	return
}

// GetFromIDentityDate 通过单个 identity_date(实名认证时间) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromIDentityDate(identityDate int64) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithIDentityDate(identityDate))
	return
}

// GetsFromIDentityDate 通过多个 identity_date(实名认证时间) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromIDentityDate(identityDates []int64) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithIDentityDates(identityDates))
	return
}

// GetFromIDentityState 通过单个 identity_state(实名认证状态：1 一致 2 不一致 3 无记录 -1 异常状态) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromIDentityState(identityState int) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithIDentityState(identityState))
	return
}

// GetsFromIDentityState 通过多个 identity_state(实名认证状态：1 一致 2 不一致 3 无记录 -1 异常状态) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromIDentityState(identityStates []int) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithIDentityStates(identityStates))
	return
}

// GetFromIDentityDesc 通过单个 identity_desc(实名认证结果说明) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromIDentityDesc(identityDesc string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithIDentityDesc(identityDesc))
	return
}

// GetsFromIDentityDesc 通过多个 identity_desc(实名认证结果说明) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromIDentityDesc(identityDescs []string) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithIDentityDescs(identityDescs))
	return
}

// GetFromIDentityUpdateTime 通过单个 identity_update_time(实名认证更新时间，时间戳) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromIDentityUpdateTime(identityUpdateTime int64) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithIDentityUpdateTime(identityUpdateTime))
	return
}

// GetsFromIDentityUpdateTime 通过多个 identity_update_time(实名认证更新时间，时间戳) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromIDentityUpdateTime(identityUpdateTimes []int64) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithIDentityUpdateTimes(identityUpdateTimes))
	return
}

// GetFromCleanProtocalFile 通过单个 clean_protocal_file(是否清理协议文件：1：是 非1：否) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromCleanProtocalFile(cleanProtocalFile int) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithCleanProtocalFile(cleanProtocalFile))
	return
}

// GetsFromCleanProtocalFile 通过多个 clean_protocal_file(是否清理协议文件：1：是 非1：否) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromCleanProtocalFile(cleanProtocalFiles []int) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithCleanProtocalFiles(cleanProtocalFiles))
	return
}

// GetFromCleanAccounts 通过单个 clean_accounts(是否清理账户信息：1：是 非1：否) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromCleanAccounts(cleanAccounts int) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithCleanAccounts(cleanAccounts))
	return
}

// GetsFromCleanAccounts 通过多个 clean_accounts(是否清理账户信息：1：是 非1：否) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromCleanAccounts(cleanAccountss []int) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithCleanAccountss(cleanAccountss))
	return
}

// GetFromCreatedAt 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromCreatedAt(createdAt int64) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAt(createdAt))
	return
}

// GetsFromCreatedAt 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromCreatedAt(createdAts []int64) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithCreatedAts(createdAts))
	return
}

// GetFromUpdatedAt 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromUpdatedAt(updatedAt int64) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAt(updatedAt))
	return
}

// GetsFromUpdatedAt 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithUpdatedAts(updatedAts))
	return
}

// GetFromDeletedAt 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetFromDeletedAt(deletedAt int64) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAt(deletedAt))
	return
}

// GetsFromDeletedAt 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *SupplierSecondDAO) GetsFromDeletedAt(deletedAts []int64) (results []*SupplierSecond, err error) {
	results, err = obj.GetListByOption(obj.WithDeletedAts(deletedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *SupplierSecondDAO) FetchByPrimaryKey(id int64) (result *SupplierSecond, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
