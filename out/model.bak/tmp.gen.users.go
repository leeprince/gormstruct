package model_bak

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _UsersMgr struct {
	*_BaseMgr
}

// UsersMgr open func
func UsersMgr(db *gorm.DB) *_UsersMgr {
	if db == nil {
		panic(fmt.Errorf("UsersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	// return &_UsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("users"), ctx:ctx, cancel:cancel, timeout:-1}}
	// return &_UsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("users"), ctx:ctx, cancel:cancel, timeout:-1}}
	// return &_UsersMgr{_BaseMgr: &_BaseMgr{DB: db.Model(Users{}), ctx:ctx, cancel:cancel, timeout:-1}}
	return &_UsersMgr{_BaseMgr: &_BaseMgr{DB: db.Model(Users{}), ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UsersMgr) GetTableName() string {
	return "users"
}

// Reset 重置gorm会话
func (obj *_UsersMgr) Reset() *_UsersMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UsersMgr) Get() (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Find(&result).Error
	return
}

// Gets 获取批量结果
func (obj *_UsersMgr) Gets() (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Find(&results).Error
	return
}

////////////////////////////////// 替换 gorm 的方法 /////////////////////////////////
func (obj *_UsersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Users{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

////////////////////////// 表中的字段作为 option 条件 ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_UsersMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_UsersMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAge age获取 年龄
func (obj *_UsersMgr) WithAge(age int) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// WithCardNo card_no获取 身份证
func (obj *_UsersMgr) WithCardNo(cardNo string) Option {
	return optionFunc(func(o *options) { o.query["card_no"] = cardNo })
}

// WithHeadImg head_img获取 头像
func (obj *_UsersMgr) WithHeadImg(headImg string) Option {
	return optionFunc(func(o *options) { o.query["head_img"] = headImg })
}

// GetByOption 功能选项模式获取
func (obj *_UsersMgr) GetByOption(opts ...Option) (result Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where(options.query).Find(&result).Error
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UsersMgr) GetByOptions(opts ...Option) (results []*Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where(options.query).Find(&results).Error
	return
}

////////////////////////// 通过存在索引的单个字段作为查询条件 ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键
func (obj *_UsersMgr) GetFromID(id int) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`id` = ?", id).Find(&result).Error
	return
}

// GetBatchFromID 批量查找 主键
func (obj *_UsersMgr) GetBatchFromID(ids []int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`id` IN (?)", ids).Find(&results).Error
	return
}

// GetFromName 通过name获取内容 名称
func (obj *_UsersMgr) GetFromName(name string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`name` = ?", name).Find(&results).Error
	return
}

// GetBatchFromName 批量查找 名称
func (obj *_UsersMgr) GetBatchFromName(names []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`name` IN (?)", names).Find(&results).Error
	return
}

// GetFromAge 通过age获取内容 年龄
func (obj *_UsersMgr) GetFromAge(age int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`age` = ?", age).Find(&results).Error
	return
}

// GetBatchFromAge 批量查找 年龄
func (obj *_UsersMgr) GetBatchFromAge(ages []int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`age` IN (?)", ages).Find(&results).Error
	return
}

// GetFromCardNo 通过card_no获取内容 身份证
func (obj *_UsersMgr) GetFromCardNo(cardNo string) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`card_no` = ?", cardNo).Find(&result).Error
	return
}

// GetBatchFromCardNo 批量查找 身份证
func (obj *_UsersMgr) GetBatchFromCardNo(cardNos []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`card_no` IN (?)", cardNos).Find(&results).Error
	return
}

// GetFromHeadImg 通过head_img获取内容 头像
func (obj *_UsersMgr) GetFromHeadImg(headImg string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`head_img` = ?", headImg).Find(&results).Error
	return
}

// GetBatchFromHeadImg 批量查找 头像
func (obj *_UsersMgr) GetBatchFromHeadImg(headImgs []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`head_img` IN (?)", headImgs).Find(&results).Error
	return
}

////////////////////////// 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UsersMgr) FetchByPrimaryKey(id int) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`id` = ?", id).Find(&result).Error
	return
}

// FetchUniqueByCardNo primary or index 获取唯一内容
func (obj *_UsersMgr) FetchUniqueByCardNo(cardNo string) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`card_no` = ?", cardNo).Find(&result).Error
	return
}

// FetchUniqueIndexByUnqNameCard primary or index 获取唯一内容
func (obj *_UsersMgr) FetchUniqueIndexByUnqNameCard(name string, cardNo string) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`name` = ? AND `card_no` = ?", name, cardNo).Find(&result).Error
	return
}

// FetchIndexByAge  获取多个内容
func (obj *_UsersMgr) FetchIndexByAge(age int) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`age` = ?", age).Find(&results).Error
	return
}
