package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022-01-29 23:00:37
 * @Desc:   users 表的 dao 层
 */

type UsersDao struct {
	*_BaseDao
}

// 初始化 UsersDao
func NewUsersDao(db *gorm.DB) *UsersDao {
	if db == nil {
		panic(fmt.Errorf("UsersDao need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &UsersDao{_BaseDao: &_BaseDao{DB: db, ctx: ctx, cancel: cancel, timeout: -1}}
}

// 获取表名称
func (obj *UsersDao) GetTableName() string {
	users := &Users{}
	return users.TableName()
}

// 获取单条记录
func (obj *UsersDao) Get() (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Find(&result).Error
	return
}

// 获取多条记录
func (obj *UsersDao) Gets() (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Find(&results).Error
	return
}

// --- 替换 gorm 的方法 ---
// 统计
func (obj *UsersDao) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Count(count)
}

// 插入
func (obj *UsersDao) Create(users *Users) (rowsAffected int64, err error) {
	tx := obj.DB.WithContext(obj.ctx).Create(users)
	rowsAffected = tx.RowsAffected
	err = tx.Error
	return
}

// --- 替换 gorm 的方法 -END ---

// --- 表中的字段作为 option 条件 ---

// 设置 id(主键) 字段作为 option 条件
func (obj *UsersDao) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.ID] = id })
}

// 设置 id(主键) 字段的切片作为 option 条件
func (obj *UsersDao) WithBatchID(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.ID] = ids })
}

// 设置 name(名称) 字段作为 option 条件
func (obj *UsersDao) WithName(name string) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.Name] = name })
}

// 设置 name(名称) 字段的切片作为 option 条件
func (obj *UsersDao) WithBatchName(names []string) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.Name] = names })
}

// 设置 age(年龄) 字段作为 option 条件
func (obj *UsersDao) WithAge(age int64) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.Age] = age })
}

// 设置 age(年龄) 字段的切片作为 option 条件
func (obj *UsersDao) WithBatchAge(ages []int64) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.Age] = ages })
}

// 设置 card_no(身份证) 字段作为 option 条件
func (obj *UsersDao) WithCardNo(cardNo string) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.CardNo] = cardNo })
}

// 设置 card_no(身份证) 字段的切片作为 option 条件
func (obj *UsersDao) WithBatchCardNo(cardNos []string) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.CardNo] = cardNos })
}

// 设置 head_img(头像) 字段作为 option 条件
func (obj *UsersDao) WithHeadImg(headImg string) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.HeadImg] = headImg })
}

// 设置 head_img(头像) 字段的切片作为 option 条件
func (obj *UsersDao) WithBatchHeadImg(headImgs []string) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.HeadImg] = headImgs })
}

// 设置 created_at(创建时间) 字段作为 option 条件
func (obj *UsersDao) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.CreatedAt] = createdAt })
}

// 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *UsersDao) WithBatchCreatedAt(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.CreatedAt] = createdAts })
}

// 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *UsersDao) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.UpdatedAt] = updatedAt })
}

// 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *UsersDao) WithBatchUpdatedAt(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.UpdatedAt] = updatedAts })
}

// 设置 deleted_at(删除 时间) 字段作为 option 条件
func (obj *UsersDao) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.DeletedAt] = deletedAt })
}

// 设置 deleted_at(删除 时间) 字段的切片作为 option 条件
func (obj *UsersDao) WithBatchDeletedAt(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.query[UsersColumns.DeletedAt] = deletedAts })
}

// 函数选项模式获取单条记录
func (obj *UsersDao) GetByOption(opts ...Option) (result Users, err error) {
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// 函数选项模式获取多条记录：支持分页
func (obj *UsersDao) GetByOptions(opts ...Option) (results []*Users, err error) {
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *UsersDao) GetCountByOptions(opts ...Option) (count int64) {
	obj.prepare(opts...).Model(&Users{}).Count(&count)
	return
}

// 更新数据：非指针的结构体字段更新为零值更新时需配合 WithSelect 更新
// 参数说明：
//    users: 要更新的数据
//    opts: 更新的条件
func (obj *UsersDao) UpdateByOption(users Users, opts ...Option) (rowsAffected int64, err error) {
	tx := obj.prepare(opts...).Updates(&users)
	rowsAffected = tx.RowsAffected
	err = tx.Error
	return
}

// --- 表中的字段作为 option 条件 -END ---

// --- 通过存在索引的单个字段作为查询条件 ---

// 通过 id(主键) 获取内容
func (obj *UsersDao) GetFromID(id int64) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(Users{ID: id}).Find(&result).Error
	return
}

// 通过 id(主键) 获取多条记录
func (obj *UsersDao) GetBatchFromID(ids []int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(fmt.Sprintf("%s IN (?)", UsersColumns.ID), ids).Find(&results).Error
	return
}

// 通过 name(名称) 获取内容
func (obj *UsersDao) GetFromName(name string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(Users{Name: name}).Find(&results).Error
	return
}

// 通过 name(名称) 获取多条记录
func (obj *UsersDao) GetBatchFromName(names []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(fmt.Sprintf("%s IN (?)", UsersColumns.Name), names).Find(&results).Error
	return
}

// 通过 age(年龄) 获取内容
func (obj *UsersDao) GetFromAge(age int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(Users{Age: age}).Find(&results).Error
	return
}

// 通过 age(年龄) 获取多条记录
func (obj *UsersDao) GetBatchFromAge(ages []int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(fmt.Sprintf("%s IN (?)", UsersColumns.Age), ages).Find(&results).Error
	return
}

// 通过 card_no(身份证) 获取内容
func (obj *UsersDao) GetFromCardNo(cardNo string) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(Users{CardNo: cardNo}).Find(&result).Error
	return
}

// 通过 card_no(身份证) 获取多条记录
func (obj *UsersDao) GetBatchFromCardNo(cardNos []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(fmt.Sprintf("%s IN (?)", UsersColumns.CardNo), cardNos).Find(&results).Error
	return
}

// 通过 head_img(头像) 获取内容
func (obj *UsersDao) GetFromHeadImg(headImg string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(Users{HeadImg: headImg}).Find(&results).Error
	return
}

// 通过 head_img(头像) 获取多条记录
func (obj *UsersDao) GetBatchFromHeadImg(headImgs []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(fmt.Sprintf("%s IN (?)", UsersColumns.HeadImg), headImgs).Find(&results).Error
	return
}

// 通过 created_at(创建时间) 获取内容
func (obj *UsersDao) GetFromCreatedAt(createdAt int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(Users{CreatedAt: createdAt}).Find(&results).Error
	return
}

// 通过 created_at(创建时间) 获取多条记录
func (obj *UsersDao) GetBatchFromCreatedAt(createdAts []int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(fmt.Sprintf("%s IN (?)", UsersColumns.CreatedAt), createdAts).Find(&results).Error
	return
}

// 通过 updated_at(更新时间) 获取内容
func (obj *UsersDao) GetFromUpdatedAt(updatedAt int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(Users{UpdatedAt: updatedAt}).Find(&results).Error
	return
}

// 通过 updated_at(更新时间) 获取多条记录
func (obj *UsersDao) GetBatchFromUpdatedAt(updatedAts []int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(fmt.Sprintf("%s IN (?)", UsersColumns.UpdatedAt), updatedAts).Find(&results).Error
	return
}

// 通过 deleted_at(删除 时间) 获取内容
func (obj *UsersDao) GetFromDeletedAt(deletedAt int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(Users{DeletedAt: deletedAt}).Find(&results).Error
	return
}

// 通过 deleted_at(删除 时间) 获取多条记录
func (obj *UsersDao) GetBatchFromDeletedAt(deletedAts []int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(fmt.Sprintf("%s IN (?)", UsersColumns.DeletedAt), deletedAts).Find(&results).Error
	return
}

// --- 通过存在索引的单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// 获取单条记录
func (obj *UsersDao) FetchByPrimaryKey(id int64) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(Users{ID: id}).Find(&result).Error
	return
}

// 获取单条记录
func (obj *UsersDao) FetchUniqueByCardNo(cardNo string) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(Users{CardNo: cardNo}).Find(&result).Error
	return
}

// 获取单条记录
func (obj *UsersDao) FetchUniqueIndexByUnqNameCard(name string, cardNo string) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(Users{Name: name, CardNo: cardNo}).Find(&result).Error
	return
}

// 获取多条记录
func (obj *UsersDao) FetchIndexByAge(age int64) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Where(Users{Age: age}).Find(&results).Error
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
