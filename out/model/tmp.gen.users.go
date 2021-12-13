package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _UsersMgr struct {
	*_BaseMgr
}

// 初始化 UsersMgr
func UsersMgr(db *gorm.DB) *_UsersMgr {
	if db == nil {
		panic(fmt.Errorf("UsersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UsersMgr{_BaseMgr: &_BaseMgr{DB: db.Model(Users{}), ctx: ctx, cancel: cancel, timeout: -1}}
}

// 获取表名称
func (obj *_UsersMgr) GetTableName() string {
	users := Users{}
	return users.TableName()
}

// 重置 gorm 会话
func (obj *_UsersMgr) Reset() *_UsersMgr {
	obj.New()
	return UsersMgr(obj.DB)
}

// 获取单条记录
func (obj *_UsersMgr) Get() (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Find(&result).Error
	return
}

// 获取多条记录
func (obj *_UsersMgr) Gets() (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Find(&results).Error
	return
}

// -------------------- 替换 gorm 的方法 -------------
func (obj *_UsersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Count(count)
}

// ----------------------------------------------------------------

// -------------------- 表中的字段作为 option 条件

// 设置 id(主键) 字段作为 option 条件
func (obj *_UsersMgr) WithID(id int) Option {
	return queryOptionFunc(func(o *options) { o.query["id"] = id })
}

// 设置 id(主键) 字段的切片作为 option 条件
func (obj *_UsersMgr) WithBatchID(ids []int) Option {
	return queryOptionFunc(func(o *options) { o.query["id"] = ids })
}

// 设置 name(名称) 字段作为 option 条件
func (obj *_UsersMgr) WithName(name string) Option {
	return queryOptionFunc(func(o *options) { o.query["name"] = name })
}

// 设置 name(名称) 字段的切片作为 option 条件
func (obj *_UsersMgr) WithBatchName(names []string) Option {
	return queryOptionFunc(func(o *options) { o.query["name"] = names })
}

// 设置 age(年龄) 字段作为 option 条件
func (obj *_UsersMgr) WithAge(age int) Option {
	return queryOptionFunc(func(o *options) { o.query["age"] = age })
}

// 设置 age(年龄) 字段的切片作为 option 条件
func (obj *_UsersMgr) WithBatchAge(ages []int) Option {
	return queryOptionFunc(func(o *options) { o.query["age"] = ages })
}

// 设置 card_no(身份证) 字段作为 option 条件
func (obj *_UsersMgr) WithCardNo(cardNo string) Option {
	return queryOptionFunc(func(o *options) { o.query["card_no"] = cardNo })
}

// 设置 card_no(身份证) 字段的切片作为 option 条件
func (obj *_UsersMgr) WithBatchCardNo(cardNos []string) Option {
	return queryOptionFunc(func(o *options) { o.query["card_no"] = cardNos })
}

// 设置 head_img(头像) 字段作为 option 条件
func (obj *_UsersMgr) WithHeadImg(headImg string) Option {
	return queryOptionFunc(func(o *options) { o.query["head_img"] = headImg })
}

// 设置 head_img(头像) 字段的切片作为 option 条件
func (obj *_UsersMgr) WithBatchHeadImg(headImgs []string) Option {
	return queryOptionFunc(func(o *options) { o.query["head_img"] = headImgs })
}

// 设置 created_at(创建时间) 字段作为 option 条件
func (obj *_UsersMgr) WithCreatedAt(createdAt *int) Option {
	return queryOptionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *_UsersMgr) WithBatchCreatedAt(createdAts []*int) Option {
	return queryOptionFunc(func(o *options) { o.query["created_at"] = createdAts })
}

// 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *_UsersMgr) WithUpdatedAt(updatedAt *int) Option {
	return queryOptionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *_UsersMgr) WithBatchUpdatedAt(updatedAts []*int) Option {
	return queryOptionFunc(func(o *options) { o.query["updated_at"] = updatedAts })
}

// 设置 deleted_at(删除 时间) 字段作为 option 条件
func (obj *_UsersMgr) WithDeletedAt(deletedAt *int) Option {
	return queryOptionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// 设置 deleted_at(删除 时间) 字段的切片作为 option 条件
func (obj *_UsersMgr) WithBatchDeletedAt(deletedAts []*int) Option {
	return queryOptionFunc(func(o *options) { o.query["deleted_at"] = deletedAts })
}

// 设置 offset、limit 作为 option 条件支持分页
func (obj *_UsersMgr) WithPage(offset int, limit int) Option {
	return pageOptionFunc(func(o *options) {
		o.page.offset = offset
		o.page.limit = limit
	})
}

// 函数选项模式获取单条记录
func (obj *_UsersMgr) GetByOption(opts ...Option) (result Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Where(options.query).Find(&result).Error
	return
}

// 函数选项模式获取多条记录：支持分页
func (obj *_UsersMgr) GetByOptions(opts ...Option) (results []*Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
		page: struct {
			offset int
			limit  int
		}{offset: 0, limit: 0},
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Where(options.query).Scopes(obj.paginate(&options)).Find(&results).Error
	return
}

// -------------------- 通过存在索引的单个字段作为查询条件

// 通过 id(主键) 获取内容
func (obj *_UsersMgr) GetFromID(id int) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`id` = ?", id).Find(&result).Error
	return
}

// 通过 id(主键) 获取多条记录
func (obj *_UsersMgr) GetBatchFromID(ids []int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`id` IN (?)", ids).Find(&results).Error
	return
}

// 通过 name(名称) 获取内容
func (obj *_UsersMgr) GetFromName(name string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`name` = ?", name).Find(&results).Error
	return
}

// 通过 name(名称) 获取多条记录
func (obj *_UsersMgr) GetBatchFromName(names []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`name` IN (?)", names).Find(&results).Error
	return
}

// 通过 age(年龄) 获取内容
func (obj *_UsersMgr) GetFromAge(age int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`age` = ?", age).Find(&results).Error
	return
}

// 通过 age(年龄) 获取多条记录
func (obj *_UsersMgr) GetBatchFromAge(ages []int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`age` IN (?)", ages).Find(&results).Error
	return
}

// 通过 card_no(身份证) 获取内容
func (obj *_UsersMgr) GetFromCardNo(cardNo string) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`card_no` = ?", cardNo).Find(&result).Error
	return
}

// 通过 card_no(身份证) 获取多条记录
func (obj *_UsersMgr) GetBatchFromCardNo(cardNos []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`card_no` IN (?)", cardNos).Find(&results).Error
	return
}

// 通过 head_img(头像) 获取内容
func (obj *_UsersMgr) GetFromHeadImg(headImg string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`head_img` = ?", headImg).Find(&results).Error
	return
}

// 通过 head_img(头像) 获取多条记录
func (obj *_UsersMgr) GetBatchFromHeadImg(headImgs []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`head_img` IN (?)", headImgs).Find(&results).Error
	return
}

// 通过 created_at(创建时间) 获取内容
func (obj *_UsersMgr) GetFromCreatedAt(createdAt *int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`created_at` = ?", createdAt).Find(&results).Error
	return
}

// 通过 created_at(创建时间) 获取多条记录
func (obj *_UsersMgr) GetBatchFromCreatedAt(createdAts []*int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`created_at` IN (?)", createdAts).Find(&results).Error
	return
}

// 通过 updated_at(更新时间) 获取内容
func (obj *_UsersMgr) GetFromUpdatedAt(updatedAt *int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`updated_at` = ?", updatedAt).Find(&results).Error
	return
}

// 通过 updated_at(更新时间) 获取多条记录
func (obj *_UsersMgr) GetBatchFromUpdatedAt(updatedAts []*int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error
	return
}

// 通过 deleted_at(删除 时间) 获取内容
func (obj *_UsersMgr) GetFromDeletedAt(deletedAt *int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`deleted_at` = ?", deletedAt).Find(&results).Error
	return
}

// 通过 deleted_at(删除 时间) 获取多条记录
func (obj *_UsersMgr) GetBatchFromDeletedAt(deletedAts []*int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`deleted_at` IN (?)", deletedAts).Find(&results).Error
	return
}

// -------------------- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件

// 获取单条记录
func (obj *_UsersMgr) FetchByPrimaryKey(id int) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`id` = ?", id).Find(&result).Error
	return
}

// 获取单条记录
func (obj *_UsersMgr) FetchUniqueByCardNo(cardNo string) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`card_no` = ?", cardNo).Find(&result).Error
	return
}

// 获取单条记录
func (obj *_UsersMgr) FetchUniqueIndexByUnqNameCard(name string, cardNo string) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`name` = ? AND `card_no` = ?", name, cardNo).Find(&result).Error
	return
}

// 获取多条记录
func (obj *_UsersMgr) FetchIndexByAge(age int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where("`age` = ?", age).Find(&results).Error
	return
}
