package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// 初始化 gorm 实例的其他字段
type _BaseMgr struct {
	*gorm.DB
	ctx       context.Context
	cancel    context.CancelFunc
	timeout   time.Duration
	isRelated bool
}

// 设置超时
func (obj *_BaseMgr) SetTimeOut(timeout time.Duration) {
	obj.ctx, obj.cancel = context.WithTimeout(context.Background(), timeout)
	obj.timeout = timeout
}

// 设置上下文
func (obj *_BaseMgr) SetCtx(c context.Context) {
	if c != nil {
		obj.ctx = c
	}
}

// 获取上下文
func (obj *_BaseMgr) GetCtx() context.Context {
	return obj.ctx
}

// 取消上下文
func (obj *_BaseMgr) Cancel(c context.Context) {
	obj.cancel()
}

// 获取 DB 实例
func (obj *_BaseMgr) GetDB() *gorm.DB {
	return obj.DB
}

// 更新 DB 实例
func (obj *_BaseMgr) UpdateDB(db *gorm.DB) {
	obj.DB = db
}

// 重置 gorm
func (obj *_BaseMgr) New() {
	obj.DB = obj.NewDB()
}

// 重置 gorm 会话
func (obj *_BaseMgr) NewDB() *gorm.DB {
	return obj.DB.Session(&gorm.Session{NewDB: true, Context: obj.ctx})
}

// 分页器
func (obj *_BaseMgr) paginate(opt *options) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if opt.page.limit > 1 {
			return db.Offset(opt.page.offset).Limit(opt.page.limit)
		}
		return db
	}
}

// 函数选项模式的参数字段
type options struct {
	query map[string]interface{}
	page  struct {
		offset int
		limit  int
	}
}

// 函数选项模式接口
type Option interface {
	apply(*options)
}

// options.query 实现 Option 接口
type queryOptionFunc func(*options)

func (f queryOptionFunc) apply(o *options) {
	f(o)
}

// options.page 实现 Option 接口
type pageOptionFunc func(*options)

func (f pageOptionFunc) apply(o *options) {
	f(o)
}
