package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:46:00
 * @Desc:   purchaser_order_pre 表
 */

// 采购单
type PurchaserOrderPre struct {
	ID                 int64  `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`                                    // 主键id
	OrgID              int64  `gorm:"column:org_id;type:int(11);not null;DEFAULT '0'" json:"org_id"`                                // 一级供应商（平台企业）租户id
	OpenOrgID          string `gorm:"column:open_org_id;type:varchar(32);not null;DEFAULT ''" json:"open_org_id"`                   // 一级供应商（平台企业）开放租户id
	PurchaserID        string `gorm:"column:purchaser_id;type:varchar(50);not null;DEFAULT ''" json:"purchaser_id"`                 // 采购商id
	PurchaserOrderID   string `gorm:"column:purchaser_order_id;type:varchar(50);not null;DEFAULT ''" json:"purchaser_order_id"`     // 采购单编号
	PurchaserOrderName string `gorm:"column:purchaser_order_name;type:varchar(50);not null;DEFAULT ''" json:"purchaser_order_name"` // 采购单名称
	OrderTime          int64  `gorm:"column:order_time;type:int(11);not null;DEFAULT '0'" json:"order_time"`                        // 订单创建时间
	State              int64  `gorm:"column:state;type:int(11);not null;DEFAULT '0'" json:"state"`                                  // 状态。0采购中 1已取消 2待支付 3已支付 4已完成 5支付中
	URL                string `gorm:"column:url;type:varchar(255);not null;DEFAULT ''" json:"url"`                                  // 采购单文件下载URL
	Cycle              string `gorm:"column:cycle;type:varchar(50);not null;DEFAULT ''" json:"cycle"`                               // 采购周期
	ClassName          string `gorm:"column:class_name;type:varchar(50);not null;DEFAULT '0'" json:"class_name"`                    // 商品品类
	Count              int64  `gorm:"column:count;type:int(11);not null;DEFAULT '0'" json:"count"`                                  // 采购数量
	Amount             int64  `gorm:"column:amount;type:int(11);not null;DEFAULT '0'" json:"amount"`                                // 采购金额
	ReceiveAreaCode    string `gorm:"column:receive_area_code;type:varchar(20);not null;DEFAULT ''" json:"receive_area_code"`       // 收货所在地省市区
	ReceiveAddress     string `gorm:"column:receive_address;type:varchar(50);not null;DEFAULT ''" json:"receive_address"`           // 收货所在地详细地址
	Special            string `gorm:"column:special;type:varchar(50);not null;DEFAULT ''" json:"special"`                           // 品质规格
	RequireText        string `gorm:"column:require_text;type:varchar(50);not null;DEFAULT ''" json:"require_text"`                 // 采购要求
	Rate               string `gorm:"column:rate;type:varchar(50);not null;DEFAULT '0'" json:"rate"`                                // 签约服务费率
	FeeAmount          int64  `gorm:"column:fee_amount;type:int(11);not null;DEFAULT '0'" json:"fee_amount"`                        // 服务费【单位：分】
	BillAmount         int64  `gorm:"column:bill_amount;type:int(11);not null;DEFAULT '0'" json:"bill_amount"`                      // 账单金额【单位：分】
	RealAmount         int64  `gorm:"column:real_amount;type:int(11);not null;DEFAULT '0'" json:"real_amount"`                      // 实际支付采购金额【单位分】
	CreatedAt          int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`                        // 创建时间
	UpdatedAt          int64  `gorm:"column:updated_at;type:int(11);not null;DEFAULT '0'" json:"updated_at"`                        // 更新时间
	DeletedAt          int64  `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"`                        // 删除时间
}

// 获取结构体对应的表名方法
func (m *PurchaserOrderPre) TableName() string {
	return "purchaser_order_pre"
}

// 实例化结构体对象
func NewPurchaserOrderPre() *PurchaserOrderPre {
	return &PurchaserOrderPre{}
}

// 获取主键的对应字段
func (m *PurchaserOrderPre) PrimaryKey() string {
	return PurchaserOrderPreColumns.ID
}

// 获取主键值
func (m *PurchaserOrderPre) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var PurchaserOrderPreColumns = struct {
	ID                 string
	OrgID              string
	OpenOrgID          string
	PurchaserID        string
	PurchaserOrderID   string
	PurchaserOrderName string
	OrderTime          string
	State              string
	URL                string
	Cycle              string
	ClassName          string
	Count              string
	Amount             string
	ReceiveAreaCode    string
	ReceiveAddress     string
	Special            string
	RequireText        string
	Rate               string
	FeeAmount          string
	BillAmount         string
	RealAmount         string
	CreatedAt          string
	UpdatedAt          string
	DeletedAt          string
}{
	ID:                 "id",
	OrgID:              "org_id",
	OpenOrgID:          "open_org_id",
	PurchaserID:        "purchaser_id",
	PurchaserOrderID:   "purchaser_order_id",
	PurchaserOrderName: "purchaser_order_name",
	OrderTime:          "order_time",
	State:              "state",
	URL:                "url",
	Cycle:              "cycle",
	ClassName:          "class_name",
	Count:              "count",
	Amount:             "amount",
	ReceiveAreaCode:    "receive_area_code",
	ReceiveAddress:     "receive_address",
	Special:            "special",
	RequireText:        "require_text",
	Rate:               "rate",
	FeeAmount:          "fee_amount",
	BillAmount:         "bill_amount",
	RealAmount:         "real_amount",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	DeletedAt:          "deleted_at",
}

// 包含所有表字段的切片
var PurchaserOrderPreAllColumns = []string{
	PurchaserOrderPreColumns.ID,                 // 主键id
	PurchaserOrderPreColumns.OrgID,              // 一级供应商（平台企业）租户id
	PurchaserOrderPreColumns.OpenOrgID,          // 一级供应商（平台企业）开放租户id
	PurchaserOrderPreColumns.PurchaserID,        // 采购商id
	PurchaserOrderPreColumns.PurchaserOrderID,   // 采购单编号
	PurchaserOrderPreColumns.PurchaserOrderName, // 采购单名称
	PurchaserOrderPreColumns.OrderTime,          // 订单创建时间
	PurchaserOrderPreColumns.State,              // 状态。0采购中 1已取消 2待支付 3已支付 4已完成 5支付中
	PurchaserOrderPreColumns.URL,                // 采购单文件下载URL
	PurchaserOrderPreColumns.Cycle,              // 采购周期
	PurchaserOrderPreColumns.ClassName,          // 商品品类
	PurchaserOrderPreColumns.Count,              // 采购数量
	PurchaserOrderPreColumns.Amount,             // 采购金额
	PurchaserOrderPreColumns.ReceiveAreaCode,    // 收货所在地省市区
	PurchaserOrderPreColumns.ReceiveAddress,     // 收货所在地详细地址
	PurchaserOrderPreColumns.Special,            // 品质规格
	PurchaserOrderPreColumns.RequireText,        // 采购要求
	PurchaserOrderPreColumns.Rate,               // 签约服务费率
	PurchaserOrderPreColumns.FeeAmount,          // 服务费【单位：分】
	PurchaserOrderPreColumns.BillAmount,         // 账单金额【单位：分】
	PurchaserOrderPreColumns.RealAmount,         // 实际支付采购金额【单位分】
	PurchaserOrderPreColumns.CreatedAt,          // 创建时间
	PurchaserOrderPreColumns.UpdatedAt,          // 更新时间
	PurchaserOrderPreColumns.DeletedAt,          // 删除时间

}

// 设置值：主键id
func (m *PurchaserOrderPre) SetID(v int64) {
	m.ID = v
}

// 设置值：一级供应商（平台企业）租户id
func (m *PurchaserOrderPre) SetOrgID(v int64) {
	m.OrgID = v
}

// 设置值：一级供应商（平台企业）开放租户id
func (m *PurchaserOrderPre) SetOpenOrgID(v string) {
	m.OpenOrgID = v
}

// 设置值：采购商id
func (m *PurchaserOrderPre) SetPurchaserID(v string) {
	m.PurchaserID = v
}

// 设置值：采购单编号
func (m *PurchaserOrderPre) SetPurchaserOrderID(v string) {
	m.PurchaserOrderID = v
}

// 设置值：采购单名称
func (m *PurchaserOrderPre) SetPurchaserOrderName(v string) {
	m.PurchaserOrderName = v
}

// 设置值：订单创建时间
func (m *PurchaserOrderPre) SetOrderTime(v int64) {
	m.OrderTime = v
}

// 设置值：状态。0采购中 1已取消 2待支付 3已支付 4已完成 5支付中
func (m *PurchaserOrderPre) SetState(v int64) {
	m.State = v
}

// 设置值：采购单文件下载URL
func (m *PurchaserOrderPre) SetURL(v string) {
	m.URL = v
}

// 设置值：采购周期
func (m *PurchaserOrderPre) SetCycle(v string) {
	m.Cycle = v
}

// 设置值：商品品类
func (m *PurchaserOrderPre) SetClassName(v string) {
	m.ClassName = v
}

// 设置值：采购数量
func (m *PurchaserOrderPre) SetCount(v int64) {
	m.Count = v
}

// 设置值：采购金额
func (m *PurchaserOrderPre) SetAmount(v int64) {
	m.Amount = v
}

// 设置值：收货所在地省市区
func (m *PurchaserOrderPre) SetReceiveAreaCode(v string) {
	m.ReceiveAreaCode = v
}

// 设置值：收货所在地详细地址
func (m *PurchaserOrderPre) SetReceiveAddress(v string) {
	m.ReceiveAddress = v
}

// 设置值：品质规格
func (m *PurchaserOrderPre) SetSpecial(v string) {
	m.Special = v
}

// 设置值：采购要求
func (m *PurchaserOrderPre) SetRequireText(v string) {
	m.RequireText = v
}

// 设置值：签约服务费率
func (m *PurchaserOrderPre) SetRate(v string) {
	m.Rate = v
}

// 设置值：服务费【单位：分】
func (m *PurchaserOrderPre) SetFeeAmount(v int64) {
	m.FeeAmount = v
}

// 设置值：账单金额【单位：分】
func (m *PurchaserOrderPre) SetBillAmount(v int64) {
	m.BillAmount = v
}

// 设置值：实际支付采购金额【单位分】
func (m *PurchaserOrderPre) SetRealAmount(v int64) {
	m.RealAmount = v
}

// 设置值：创建时间
func (m *PurchaserOrderPre) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 设置值：更新时间
func (m *PurchaserOrderPre) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：删除时间
func (m *PurchaserOrderPre) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 获取值：主键id
func (m *PurchaserOrderPre) GetID() int64 {
	return m.ID
}

// 获取值：一级供应商（平台企业）租户id
func (m *PurchaserOrderPre) GetOrgID() int64 {
	return m.OrgID
}

// 获取值：一级供应商（平台企业）开放租户id
func (m *PurchaserOrderPre) GetOpenOrgID() string {
	return m.OpenOrgID
}

// 获取值：采购商id
func (m *PurchaserOrderPre) GetPurchaserID() string {
	return m.PurchaserID
}

// 获取值：采购单编号
func (m *PurchaserOrderPre) GetPurchaserOrderID() string {
	return m.PurchaserOrderID
}

// 获取值：采购单名称
func (m *PurchaserOrderPre) GetPurchaserOrderName() string {
	return m.PurchaserOrderName
}

// 获取值：订单创建时间
func (m *PurchaserOrderPre) GetOrderTime() int64 {
	return m.OrderTime
}

// 获取值：状态。0采购中 1已取消 2待支付 3已支付 4已完成 5支付中
func (m *PurchaserOrderPre) GetState() int64 {
	return m.State
}

// 获取值：采购单文件下载URL
func (m *PurchaserOrderPre) GetURL() string {
	return m.URL
}

// 获取值：采购周期
func (m *PurchaserOrderPre) GetCycle() string {
	return m.Cycle
}

// 获取值：商品品类
func (m *PurchaserOrderPre) GetClassName() string {
	return m.ClassName
}

// 获取值：采购数量
func (m *PurchaserOrderPre) GetCount() int64 {
	return m.Count
}

// 获取值：采购金额
func (m *PurchaserOrderPre) GetAmount() int64 {
	return m.Amount
}

// 获取值：收货所在地省市区
func (m *PurchaserOrderPre) GetReceiveAreaCode() string {
	return m.ReceiveAreaCode
}

// 获取值：收货所在地详细地址
func (m *PurchaserOrderPre) GetReceiveAddress() string {
	return m.ReceiveAddress
}

// 获取值：品质规格
func (m *PurchaserOrderPre) GetSpecial() string {
	return m.Special
}

// 获取值：采购要求
func (m *PurchaserOrderPre) GetRequireText() string {
	return m.RequireText
}

// 获取值：签约服务费率
func (m *PurchaserOrderPre) GetRate() string {
	return m.Rate
}

// 获取值：服务费【单位：分】
func (m *PurchaserOrderPre) GetFeeAmount() int64 {
	return m.FeeAmount
}

// 获取值：账单金额【单位：分】
func (m *PurchaserOrderPre) GetBillAmount() int64 {
	return m.BillAmount
}

// 获取值：实际支付采购金额【单位分】
func (m *PurchaserOrderPre) GetRealAmount() int64 {
	return m.RealAmount
}

// 获取值：创建时间
func (m *PurchaserOrderPre) GetCreatedAt() int64 {
	return m.CreatedAt
}

// 获取值：更新时间
func (m *PurchaserOrderPre) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：删除时间
func (m *PurchaserOrderPre) GetDeletedAt() int64 {
	return m.DeletedAt
}
