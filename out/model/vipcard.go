package model

import (
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-05-20 14:36:31
 * @Desc:   vipcard 表
 */

// 任务表
type Vipcard struct {
	ID           int64     `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`                      // id
	VipID        *int64    `gorm:"column:vip_id;type:int(11);is null;DEFAULT:0" json:"vip_id"`                     // vip id
	V2TypeID     int64     `gorm:"column:v2_type_id;type:int(11);not null;DEFAULT:0" json:"v2_type_id"`            // 类型id
	IsPublicSale int8      `gorm:"column:is_public_sale;type:tinyint(1);not null;DEFAULT:1" json:"is_public_sale"` // 0 非公开售卖 1公开售卖
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;is null;" json:"create_time"`                   // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime;is null;" json:"update_time"`                   // 更新时间
}

// 获取结构体对应的表名方法
func (m *Vipcard) TableName() string {
	return "vipcard"
}

// 实例化结构体对象
func NewVipcard() *Vipcard {
	return &Vipcard{}
}

// 获取主键的对应字段
func (m *Vipcard) PrimaryKey() string {
	return VipcardColumns.ID
}

// 获取主键值
func (m *Vipcard) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var VipcardColumns = struct {
	ID           string
	VipID        string
	V2TypeID     string
	IsPublicSale string
	CreateTime   string
	UpdateTime   string
}{
	ID:           "id",
	VipID:        "vip_id",
	V2TypeID:     "v2_type_id",
	IsPublicSale: "is_public_sale",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// 包含所有表字段的切片
var VipcardAllColumns = []string{
	VipcardColumns.ID,           // id
	VipcardColumns.VipID,        // vip id
	VipcardColumns.V2TypeID,     // 类型id
	VipcardColumns.IsPublicSale, // 0 非公开售卖 1公开售卖
	VipcardColumns.CreateTime,   // 创建时间
	VipcardColumns.UpdateTime,   // 更新时间

}

// 设置值：id
func (m *Vipcard) SetID(v int64) {
	m.ID = v
}

// 设置值：vip id
func (m *Vipcard) SetVipID(v *int64) {
	m.VipID = v
}

// 设置值：类型id
func (m *Vipcard) SetV2TypeID(v int64) {
	m.V2TypeID = v
}

// 设置值：0 非公开售卖 1公开售卖
func (m *Vipcard) SetIsPublicSale(v int8) {
	m.IsPublicSale = v
}

// 设置值：创建时间
func (m *Vipcard) SetCreateTime(v time.Time) {
	m.CreateTime = v
}

// 设置值：更新时间
func (m *Vipcard) SetUpdateTime(v time.Time) {
	m.UpdateTime = v
}

// 获取值：id
func (m *Vipcard) GetID() int64 {
	return m.ID
}

// 获取值：vip id
func (m *Vipcard) GetVipID() *int64 {
	return m.VipID
}

// 获取值：类型id
func (m *Vipcard) GetV2TypeID() int64 {
	return m.V2TypeID
}

// 获取值：0 非公开售卖 1公开售卖
func (m *Vipcard) GetIsPublicSale() int8 {
	return m.IsPublicSale
}

// 获取值：创建时间
func (m *Vipcard) GetCreateTime() time.Time {
	return m.CreateTime
}

// 获取值：更新时间
func (m *Vipcard) GetUpdateTime() time.Time {
	return m.UpdateTime
}
