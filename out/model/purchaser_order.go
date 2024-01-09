package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:45:41
 * @Desc:   purchaser_order 表
 */

// 采购订单信息表
type PurchaserOrder struct {
	ID               int64   `gorm:"column:id;primaryKey;type:int(11);autoIncrement" json:"id"`                                // 主键id
	OrgID            string  `gorm:"column:org_id;type:varchar(50);not null;DEFAULT ''" json:"org_id"`                         // 一级供应商（平台企业）租户id
	OpenOrgID        string  `gorm:"column:open_org_id;type:varchar(50);not null;DEFAULT ''" json:"open_org_id"`               // 一级供应商（平台企业）开放租户id
	PurchaserOrderID string  `gorm:"column:purchaser_order_id;type:varchar(50);not null;DEFAULT ''" json:"purchaser_order_id"` // 采购单编号
	OrderID          string  `gorm:"column:order_id;type:varchar(50);not null;DEFAULT ''" json:"order_id"`                     // 订单编号
	IsDel            int     `gorm:"column:is_del;type:int(1);not null;DEFAULT '0'" json:"is_del"`                             // 是否删除：1-删除 0-否
	OrderTime        int64   `gorm:"column:order_time;type:int(11);not null;DEFAULT '0'" json:"order_time"`                    // 订单创建时间
	State            int     `gorm:"column:state;type:int(1);not null;DEFAULT '0'" json:"state"`                               // 订单状态：0待支付 1已支付 2 已取消 3.支付失败 4 支付中
	URL              string  `gorm:"column:url;type:varchar(255);not null;DEFAULT ''" json:"url"`                              // 采购单文件下载URL
	PurchaserID      string  `gorm:"column:purchaser_id;type:varchar(50);not null;DEFAULT ''" json:"purchaser_id"`             // 采购商id
	SupplierID       string  `gorm:"column:supplier_id;type:varchar(50);not null;DEFAULT ''" json:"supplier_id"`               // 供应商id
	TotalAmount      int64   `gorm:"column:total_amount;type:int(11);not null;DEFAULT '0'" json:"total_amount"`                // 订单总金额（单位：分）
	BuyAmount        int64   `gorm:"column:buy_amount;type:int(11);not null;DEFAULT '0'" json:"buy_amount"`                    // 采购金额（单位：分）
	ServiceAmount    int64   `gorm:"column:service_amount;type:int(11);not null;DEFAULT '0'" json:"service_amount"`            // 服务费（单位：分）
	PayType          string  `gorm:"column:pay_type;type:varchar(50);not null;DEFAULT ''" json:"pay_type"`                     // 支付通道
	PayAccount       string  `gorm:"column:pay_account;type:varchar(50);not null;DEFAULT ''" json:"pay_account"`               // 支付账户
	SendType         string  `gorm:"column:send_type;type:varchar(50);not null;DEFAULT ''" json:"send_type"`                   // 发货方式
	ReceiveAreaCode  string  `gorm:"column:receive_area_code;type:varchar(10);not null;DEFAULT ''" json:"receive_area_code"`   // 收货所在地省市区
	ReceiveAddress   string  `gorm:"column:receive_address;type:varchar(50);not null;DEFAULT ''" json:"receive_address"`       // 收货所在地详细地址
	SendAreaCode     string  `gorm:"column:send_area_code;type:varchar(10);not null;DEFAULT ''" json:"send_area_code"`         // 发货所在地省市区
	SendAddress      string  `gorm:"column:send_address;type:varchar(50);not null;DEFAULT ''" json:"send_address"`             // 发货所在地详细地址
	ReceiveURL       *string `gorm:"column:receive_url;type:varchar(255);is null;DEFAULT ''" json:"receive_url"`               // 收货凭据/图片
	WeighURL         *string `gorm:"column:weigh_url;type:varchar(255);is null;DEFAULT ''" json:"weigh_url"`                   // 过磅单图片
	TransName        *string `gorm:"column:trans_name;type:varchar(50);is null;DEFAULT ''" json:"trans_name"`                  // 承运人姓名
	TransCardNo      *string `gorm:"column:trans_card_no;type:varchar(50);is null;DEFAULT ''" json:"trans_card_no"`            // 承运人身份证号
	TransPhone       *string `gorm:"column:trans_phone;type:varchar(11);is null;DEFAULT ''" json:"trans_phone"`                // 承运人手机号码
	TransCarNo       *string `gorm:"column:trans_car_no;type:varchar(20);is null;DEFAULT ''" json:"trans_car_no"`              // 承运人车牌号
	TransURL         *string `gorm:"column:trans_url;type:varchar(255);is null;DEFAULT ''" json:"trans_url"`                   // 承运人图片
	Extend           *string `gorm:"column:extend;type:varchar(255);is null;DEFAULT ''" json:"extend"`                         // 补全字段JSON
}

// 获取结构体对应的表名方法
func (m *PurchaserOrder) TableName() string {
	return "purchaser_order"
}

// 实例化结构体对象
func NewPurchaserOrder() *PurchaserOrder {
	return &PurchaserOrder{}
}

// 获取主键的对应字段
func (m *PurchaserOrder) PrimaryKey() string {
	return PurchaserOrderColumns.ID
}

// 获取主键值
func (m *PurchaserOrder) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var PurchaserOrderColumns = struct {
	ID               string
	OrgID            string
	OpenOrgID        string
	PurchaserOrderID string
	OrderID          string
	IsDel            string
	OrderTime        string
	State            string
	URL              string
	PurchaserID      string
	SupplierID       string
	TotalAmount      string
	BuyAmount        string
	ServiceAmount    string
	PayType          string
	PayAccount       string
	SendType         string
	ReceiveAreaCode  string
	ReceiveAddress   string
	SendAreaCode     string
	SendAddress      string
	ReceiveURL       string
	WeighURL         string
	TransName        string
	TransCardNo      string
	TransPhone       string
	TransCarNo       string
	TransURL         string
	Extend           string
}{
	ID:               "id",
	OrgID:            "org_id",
	OpenOrgID:        "open_org_id",
	PurchaserOrderID: "purchaser_order_id",
	OrderID:          "order_id",
	IsDel:            "is_del",
	OrderTime:        "order_time",
	State:            "state",
	URL:              "url",
	PurchaserID:      "purchaser_id",
	SupplierID:       "supplier_id",
	TotalAmount:      "total_amount",
	BuyAmount:        "buy_amount",
	ServiceAmount:    "service_amount",
	PayType:          "pay_type",
	PayAccount:       "pay_account",
	SendType:         "send_type",
	ReceiveAreaCode:  "receive_area_code",
	ReceiveAddress:   "receive_address",
	SendAreaCode:     "send_area_code",
	SendAddress:      "send_address",
	ReceiveURL:       "receive_url",
	WeighURL:         "weigh_url",
	TransName:        "trans_name",
	TransCardNo:      "trans_card_no",
	TransPhone:       "trans_phone",
	TransCarNo:       "trans_car_no",
	TransURL:         "trans_url",
	Extend:           "extend",
}

// 包含所有表字段的切片
var PurchaserOrderAllColumns = []string{
	PurchaserOrderColumns.ID,               // 主键id
	PurchaserOrderColumns.OrgID,            // 一级供应商（平台企业）租户id
	PurchaserOrderColumns.OpenOrgID,        // 一级供应商（平台企业）开放租户id
	PurchaserOrderColumns.PurchaserOrderID, // 采购单编号
	PurchaserOrderColumns.OrderID,          // 订单编号
	PurchaserOrderColumns.IsDel,            // 是否删除：1-删除 0-否
	PurchaserOrderColumns.OrderTime,        // 订单创建时间
	PurchaserOrderColumns.State,            // 订单状态：0待支付 1已支付 2 已取消 3.支付失败 4 支付中
	PurchaserOrderColumns.URL,              // 采购单文件下载URL
	PurchaserOrderColumns.PurchaserID,      // 采购商id
	PurchaserOrderColumns.SupplierID,       // 供应商id
	PurchaserOrderColumns.TotalAmount,      // 订单总金额（单位：分）
	PurchaserOrderColumns.BuyAmount,        // 采购金额（单位：分）
	PurchaserOrderColumns.ServiceAmount,    // 服务费（单位：分）
	PurchaserOrderColumns.PayType,          // 支付通道
	PurchaserOrderColumns.PayAccount,       // 支付账户
	PurchaserOrderColumns.SendType,         // 发货方式
	PurchaserOrderColumns.ReceiveAreaCode,  // 收货所在地省市区
	PurchaserOrderColumns.ReceiveAddress,   // 收货所在地详细地址
	PurchaserOrderColumns.SendAreaCode,     // 发货所在地省市区
	PurchaserOrderColumns.SendAddress,      // 发货所在地详细地址
	PurchaserOrderColumns.ReceiveURL,       // 收货凭据/图片
	PurchaserOrderColumns.WeighURL,         // 过磅单图片
	PurchaserOrderColumns.TransName,        // 承运人姓名
	PurchaserOrderColumns.TransCardNo,      // 承运人身份证号
	PurchaserOrderColumns.TransPhone,       // 承运人手机号码
	PurchaserOrderColumns.TransCarNo,       // 承运人车牌号
	PurchaserOrderColumns.TransURL,         // 承运人图片
	PurchaserOrderColumns.Extend,           // 补全字段JSON

}

// 设置值：主键id
func (m *PurchaserOrder) SetID(v int64) {
	m.ID = v
}

// 设置值：一级供应商（平台企业）租户id
func (m *PurchaserOrder) SetOrgID(v string) {
	m.OrgID = v
}

// 设置值：一级供应商（平台企业）开放租户id
func (m *PurchaserOrder) SetOpenOrgID(v string) {
	m.OpenOrgID = v
}

// 设置值：采购单编号
func (m *PurchaserOrder) SetPurchaserOrderID(v string) {
	m.PurchaserOrderID = v
}

// 设置值：订单编号
func (m *PurchaserOrder) SetOrderID(v string) {
	m.OrderID = v
}

// 设置值：是否删除：1-删除 0-否
func (m *PurchaserOrder) SetIsDel(v int) {
	m.IsDel = v
}

// 设置值：订单创建时间
func (m *PurchaserOrder) SetOrderTime(v int64) {
	m.OrderTime = v
}

// 设置值：订单状态：0待支付 1已支付 2 已取消 3.支付失败 4 支付中
func (m *PurchaserOrder) SetState(v int) {
	m.State = v
}

// 设置值：采购单文件下载URL
func (m *PurchaserOrder) SetURL(v string) {
	m.URL = v
}

// 设置值：采购商id
func (m *PurchaserOrder) SetPurchaserID(v string) {
	m.PurchaserID = v
}

// 设置值：供应商id
func (m *PurchaserOrder) SetSupplierID(v string) {
	m.SupplierID = v
}

// 设置值：订单总金额（单位：分）
func (m *PurchaserOrder) SetTotalAmount(v int64) {
	m.TotalAmount = v
}

// 设置值：采购金额（单位：分）
func (m *PurchaserOrder) SetBuyAmount(v int64) {
	m.BuyAmount = v
}

// 设置值：服务费（单位：分）
func (m *PurchaserOrder) SetServiceAmount(v int64) {
	m.ServiceAmount = v
}

// 设置值：支付通道
func (m *PurchaserOrder) SetPayType(v string) {
	m.PayType = v
}

// 设置值：支付账户
func (m *PurchaserOrder) SetPayAccount(v string) {
	m.PayAccount = v
}

// 设置值：发货方式
func (m *PurchaserOrder) SetSendType(v string) {
	m.SendType = v
}

// 设置值：收货所在地省市区
func (m *PurchaserOrder) SetReceiveAreaCode(v string) {
	m.ReceiveAreaCode = v
}

// 设置值：收货所在地详细地址
func (m *PurchaserOrder) SetReceiveAddress(v string) {
	m.ReceiveAddress = v
}

// 设置值：发货所在地省市区
func (m *PurchaserOrder) SetSendAreaCode(v string) {
	m.SendAreaCode = v
}

// 设置值：发货所在地详细地址
func (m *PurchaserOrder) SetSendAddress(v string) {
	m.SendAddress = v
}

// 设置值：收货凭据/图片
func (m *PurchaserOrder) SetReceiveURL(v *string) {
	m.ReceiveURL = v
}

// 设置值：过磅单图片
func (m *PurchaserOrder) SetWeighURL(v *string) {
	m.WeighURL = v
}

// 设置值：承运人姓名
func (m *PurchaserOrder) SetTransName(v *string) {
	m.TransName = v
}

// 设置值：承运人身份证号
func (m *PurchaserOrder) SetTransCardNo(v *string) {
	m.TransCardNo = v
}

// 设置值：承运人手机号码
func (m *PurchaserOrder) SetTransPhone(v *string) {
	m.TransPhone = v
}

// 设置值：承运人车牌号
func (m *PurchaserOrder) SetTransCarNo(v *string) {
	m.TransCarNo = v
}

// 设置值：承运人图片
func (m *PurchaserOrder) SetTransURL(v *string) {
	m.TransURL = v
}

// 设置值：补全字段JSON
func (m *PurchaserOrder) SetExtend(v *string) {
	m.Extend = v
}

// 获取值：主键id
func (m *PurchaserOrder) GetID() int64 {
	return m.ID
}

// 获取值：一级供应商（平台企业）租户id
func (m *PurchaserOrder) GetOrgID() string {
	return m.OrgID
}

// 获取值：一级供应商（平台企业）开放租户id
func (m *PurchaserOrder) GetOpenOrgID() string {
	return m.OpenOrgID
}

// 获取值：采购单编号
func (m *PurchaserOrder) GetPurchaserOrderID() string {
	return m.PurchaserOrderID
}

// 获取值：订单编号
func (m *PurchaserOrder) GetOrderID() string {
	return m.OrderID
}

// 获取值：是否删除：1-删除 0-否
func (m *PurchaserOrder) GetIsDel() int {
	return m.IsDel
}

// 获取值：订单创建时间
func (m *PurchaserOrder) GetOrderTime() int64 {
	return m.OrderTime
}

// 获取值：订单状态：0待支付 1已支付 2 已取消 3.支付失败 4 支付中
func (m *PurchaserOrder) GetState() int {
	return m.State
}

// 获取值：采购单文件下载URL
func (m *PurchaserOrder) GetURL() string {
	return m.URL
}

// 获取值：采购商id
func (m *PurchaserOrder) GetPurchaserID() string {
	return m.PurchaserID
}

// 获取值：供应商id
func (m *PurchaserOrder) GetSupplierID() string {
	return m.SupplierID
}

// 获取值：订单总金额（单位：分）
func (m *PurchaserOrder) GetTotalAmount() int64 {
	return m.TotalAmount
}

// 获取值：采购金额（单位：分）
func (m *PurchaserOrder) GetBuyAmount() int64 {
	return m.BuyAmount
}

// 获取值：服务费（单位：分）
func (m *PurchaserOrder) GetServiceAmount() int64 {
	return m.ServiceAmount
}

// 获取值：支付通道
func (m *PurchaserOrder) GetPayType() string {
	return m.PayType
}

// 获取值：支付账户
func (m *PurchaserOrder) GetPayAccount() string {
	return m.PayAccount
}

// 获取值：发货方式
func (m *PurchaserOrder) GetSendType() string {
	return m.SendType
}

// 获取值：收货所在地省市区
func (m *PurchaserOrder) GetReceiveAreaCode() string {
	return m.ReceiveAreaCode
}

// 获取值：收货所在地详细地址
func (m *PurchaserOrder) GetReceiveAddress() string {
	return m.ReceiveAddress
}

// 获取值：发货所在地省市区
func (m *PurchaserOrder) GetSendAreaCode() string {
	return m.SendAreaCode
}

// 获取值：发货所在地详细地址
func (m *PurchaserOrder) GetSendAddress() string {
	return m.SendAddress
}

// 获取值：收货凭据/图片
func (m *PurchaserOrder) GetReceiveURL() *string {
	return m.ReceiveURL
}

// 获取值：过磅单图片
func (m *PurchaserOrder) GetWeighURL() *string {
	return m.WeighURL
}

// 获取值：承运人姓名
func (m *PurchaserOrder) GetTransName() *string {
	return m.TransName
}

// 获取值：承运人身份证号
func (m *PurchaserOrder) GetTransCardNo() *string {
	return m.TransCardNo
}

// 获取值：承运人手机号码
func (m *PurchaserOrder) GetTransPhone() *string {
	return m.TransPhone
}

// 获取值：承运人车牌号
func (m *PurchaserOrder) GetTransCarNo() *string {
	return m.TransCarNo
}

// 获取值：承运人图片
func (m *PurchaserOrder) GetTransURL() *string {
	return m.TransURL
}

// 获取值：补全字段JSON
func (m *PurchaserOrder) GetExtend() *string {
	return m.Extend
}
