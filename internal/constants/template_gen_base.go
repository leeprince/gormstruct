package constants

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/12/10 上午12:07
 * @Desc:
 */

const (
    TEMP_GENBASE_V1 = `
package {{.PackageName}}
import (
    "context"
    "fmt"
    "time"

    "gorm.io/gorm"
)

// prepare for other
type _BaseMgr struct {
    *gorm.DB
    ctx       context.Context
    cancel    context.CancelFunc
    timeout   time.Duration
    isRelated bool
}

// SetTimeOut set timeout
func (obj *_BaseMgr) SetTimeOut(timeout time.Duration) {
    obj.ctx, obj.cancel = context.WithTimeout(context.Background(), timeout)
    obj.timeout = timeout
}

// SetCtx set context
func (obj *_BaseMgr) SetCtx(c context.Context) {
    if c != nil {
        obj.ctx = c
    }
}

// GetCtx get context
func (obj *_BaseMgr) GetCtx() context.Context {
    return obj.ctx
}

// Cancel cancel context
func (obj *_BaseMgr) Cancel(c context.Context) {
    obj.cancel()
}

// GetDB get gorm.DB info
func (obj *_BaseMgr) GetDB() *gorm.DB {
    return obj.DB
}

// UpdateDB update gorm.DB info
func (obj *_BaseMgr) UpdateDB(db *gorm.DB) {
    obj.DB = db
}

// New new gorm.新gorm,重置条件
func (obj *_BaseMgr) New() {
    obj.DB = obj.NewDB()
}

// NewDB new gorm.新gorm
func (obj *_BaseMgr) NewDB() *gorm.DB {
    return obj.DB.Session(&gorm.Session{NewDB: true, Context: obj.ctx})
}

type options struct {
    query map[string]interface{}
}

// Option overrides behavior of Connect.
type Option interface {
    apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
    f(o)
}

// 自定义sql查询
type Condition struct {
    list []*conditionInfo
}

// And a condition by and .and 一个条件
func (c *Condition) And(column string, cases string, value interface{}) {
    c.list = append(c.list, &conditionInfo{
        andor:  "and",
        column: column, // 列名
        case_:  cases,  // 条件(and,or,in,>=,<=)
        value:  value,
    })
}

// Or a condition by or .or 一个条件
func (c *Condition) Or(column string, cases string, value interface{}) {
    c.list = append(c.list, &conditionInfo{
        andor:  "or",
        column: column, // 列名
        case_:  cases,  // 条件(and,or,in,>=,<=)
        value:  value,
    })
}

func (c *Condition) Get() (where string, out []interface{}) {
    firstAnd := -1
    for i := 0; i < len(c.list); i++ { // 查找第一个and
        if c.list[i].andor == "and" {
            where = fmt.Sprintf("{{GetVV }} %v ?", c.list[i].column, c.list[i].case_)
            out = append(out, c.list[i].value)
            firstAnd = i
            break
        }
    }

    if firstAnd < 0 && len(c.list) > 0 { // 补刀
        where = fmt.Sprintf("{{GetVV }} %v ?", c.list[0].column, c.list[0].case_)
        out = append(out, c.list[0].value)
        firstAnd = 0
    }

    for i := 0; i < len(c.list); i++ { // 添加剩余的
        if firstAnd != i {
            where += fmt.Sprintf(" %v {{GetVV }} %v ?", c.list[i].andor, c.list[i].column, c.list[i].case_)
            out = append(out, c.list[i].value)
        }
    }

    return
}

type conditionInfo struct {
    andor  string
    column string // 列名
    case_  string // 条件(in,>=,<=)
    value  interface{}
}
    `
    
    
    TEMP_GENBASE_V2 = `
package {{.PackageName}}
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
	page struct {
		offset int
		limit int
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
    `
)


