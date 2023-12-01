package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-12-02 02:13:49
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

// GetByOption 函数选项模式获取单条记录
func (obj *UsersDAO) GetByOption(opts ...Option) (result *Users, err error) {

	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *UsersDAO) GetListByOption(opts ...Option) (results []*Users, err error) {

	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *UsersDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)

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

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
