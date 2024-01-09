package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024-01-09 11:45:41
 * @Desc:   purchaser_order 表的 DAO 层
 */

type PurchaserOrderDAO struct {
	*_BaseDAO
}

// PurchaserOrderDAO 初始化
func NewPurchaserOrderDAO(ctx context.Context, db *gorm.DB) *PurchaserOrderDAO {
	if db == nil {
		panic(fmt.Errorf("PurchaserOrderDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &PurchaserOrderDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&PurchaserOrder{}),
			db:               db,
			model:            PurchaserOrder{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          PurchaserOrderAllColumns,
			isDefaultColumns: true,
		},
	}
}

// GetTableName 获取表名称
func (obj *PurchaserOrderDAO) GetTableName() string {
	purchaserOrder := &PurchaserOrder{}
	return purchaserOrder.TableName()
}

// Save 存在则更新，否则插入
func (obj *PurchaserOrderDAO) Save(purchaserOrder *PurchaserOrder) (rowsAffected int64, err error) {
	if purchaserOrder.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(purchaserOrder, obj.WithID(purchaserOrder.ID))
	}
	return obj.Create(purchaserOrder)
}

// Create 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *PurchaserOrderDAO) Create(purchaserOrder interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(purchaserOrder)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// WithID 设置 id(主键id) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.ID] = id })
}

// WithIDs 设置 id(主键id) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.ID] = ids })
}

// WithOrgID 设置 org_id(一级供应商（平台企业）租户id) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithOrgID(orgID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.OrgID] = orgID })
}

// WithOrgIDs 设置 org_id(一级供应商（平台企业）租户id) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithOrgIDs(orgIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.OrgID] = orgIDs })
}

// WithOpenOrgID 设置 open_org_id(一级供应商（平台企业）开放租户id) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithOpenOrgID(openOrgID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.OpenOrgID] = openOrgID })
}

// WithOpenOrgIDs 设置 open_org_id(一级供应商（平台企业）开放租户id) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithOpenOrgIDs(openOrgIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.OpenOrgID] = openOrgIDs })
}

// WithPurchaserOrderID 设置 purchaser_order_id(采购单编号) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithPurchaserOrderID(purchaserOrderID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.PurchaserOrderID] = purchaserOrderID })
}

// WithPurchaserOrderIDs 设置 purchaser_order_id(采购单编号) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithPurchaserOrderIDs(purchaserOrderIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.PurchaserOrderID] = purchaserOrderIDs })
}

// WithOrderID 设置 order_id(订单编号) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithOrderID(orderID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.OrderID] = orderID })
}

// WithOrderIDs 设置 order_id(订单编号) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithOrderIDs(orderIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.OrderID] = orderIDs })
}

// WithIsDel 设置 is_del(是否删除：1-删除 0-否) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithIsDel(isDel int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.IsDel] = isDel })
}

// WithIsDels 设置 is_del(是否删除：1-删除 0-否) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithIsDels(isDels []int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.IsDel] = isDels })
}

// WithOrderTime 设置 order_time(订单创建时间) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithOrderTime(orderTime int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.OrderTime] = orderTime })
}

// WithOrderTimes 设置 order_time(订单创建时间) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithOrderTimes(orderTimes []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.OrderTime] = orderTimes })
}

// WithState 设置 state(订单状态：0待支付 1已支付 2 已取消 3.支付失败 4 支付中) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithState(state int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.State] = state })
}

// WithStates 设置 state(订单状态：0待支付 1已支付 2 已取消 3.支付失败 4 支付中) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithStates(states []int) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.State] = states })
}

// WithURL 设置 url(采购单文件下载URL) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithURL(url string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.URL] = url })
}

// WithURLs 设置 url(采购单文件下载URL) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithURLs(urls []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.URL] = urls })
}

// WithPurchaserID 设置 purchaser_id(采购商id) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithPurchaserID(purchaserID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.PurchaserID] = purchaserID })
}

// WithPurchaserIDs 设置 purchaser_id(采购商id) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithPurchaserIDs(purchaserIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.PurchaserID] = purchaserIDs })
}

// WithSupplierID 设置 supplier_id(供应商id) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithSupplierID(supplierID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.SupplierID] = supplierID })
}

// WithSupplierIDs 设置 supplier_id(供应商id) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithSupplierIDs(supplierIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.SupplierID] = supplierIDs })
}

// WithTotalAmount 设置 total_amount(订单总金额（单位：分）) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithTotalAmount(totalAmount int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.TotalAmount] = totalAmount })
}

// WithTotalAmounts 设置 total_amount(订单总金额（单位：分）) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithTotalAmounts(totalAmounts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.TotalAmount] = totalAmounts })
}

// WithBuyAmount 设置 buy_amount(采购金额（单位：分）) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithBuyAmount(buyAmount int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.BuyAmount] = buyAmount })
}

// WithBuyAmounts 设置 buy_amount(采购金额（单位：分）) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithBuyAmounts(buyAmounts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.BuyAmount] = buyAmounts })
}

// WithServiceAmount 设置 service_amount(服务费（单位：分）) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithServiceAmount(serviceAmount int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.ServiceAmount] = serviceAmount })
}

// WithServiceAmounts 设置 service_amount(服务费（单位：分）) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithServiceAmounts(serviceAmounts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.ServiceAmount] = serviceAmounts })
}

// WithPayType 设置 pay_type(支付通道) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithPayType(payType string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.PayType] = payType })
}

// WithPayTypes 设置 pay_type(支付通道) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithPayTypes(payTypes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.PayType] = payTypes })
}

// WithPayAccount 设置 pay_account(支付账户) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithPayAccount(payAccount string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.PayAccount] = payAccount })
}

// WithPayAccounts 设置 pay_account(支付账户) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithPayAccounts(payAccounts []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.PayAccount] = payAccounts })
}

// WithSendType 设置 send_type(发货方式) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithSendType(sendType string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.SendType] = sendType })
}

// WithSendTypes 设置 send_type(发货方式) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithSendTypes(sendTypes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.SendType] = sendTypes })
}

// WithReceiveAreaCode 设置 receive_area_code(收货所在地省市区) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithReceiveAreaCode(receiveAreaCode string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.ReceiveAreaCode] = receiveAreaCode })
}

// WithReceiveAreaCodes 设置 receive_area_code(收货所在地省市区) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithReceiveAreaCodes(receiveAreaCodes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.ReceiveAreaCode] = receiveAreaCodes })
}

// WithReceiveAddress 设置 receive_address(收货所在地详细地址) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithReceiveAddress(receiveAddress string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.ReceiveAddress] = receiveAddress })
}

// WithReceiveAddresss 设置 receive_address(收货所在地详细地址) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithReceiveAddresss(receiveAddresss []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.ReceiveAddress] = receiveAddresss })
}

// WithSendAreaCode 设置 send_area_code(发货所在地省市区) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithSendAreaCode(sendAreaCode string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.SendAreaCode] = sendAreaCode })
}

// WithSendAreaCodes 设置 send_area_code(发货所在地省市区) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithSendAreaCodes(sendAreaCodes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.SendAreaCode] = sendAreaCodes })
}

// WithSendAddress 设置 send_address(发货所在地详细地址) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithSendAddress(sendAddress string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.SendAddress] = sendAddress })
}

// WithSendAddresss 设置 send_address(发货所在地详细地址) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithSendAddresss(sendAddresss []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.SendAddress] = sendAddresss })
}

// WithReceiveURL 设置 receive_url(收货凭据/图片) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithReceiveURL(receiveURL *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.ReceiveURL] = receiveURL })
}

// WithReceiveURLs 设置 receive_url(收货凭据/图片) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithReceiveURLs(receiveURLs []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.ReceiveURL] = receiveURLs })
}

// WithWeighURL 设置 weigh_url(过磅单图片) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithWeighURL(weighURL *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.WeighURL] = weighURL })
}

// WithWeighURLs 设置 weigh_url(过磅单图片) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithWeighURLs(weighURLs []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.WeighURL] = weighURLs })
}

// WithTransName 设置 trans_name(承运人姓名) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithTransName(transName *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.TransName] = transName })
}

// WithTransNames 设置 trans_name(承运人姓名) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithTransNames(transNames []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.TransName] = transNames })
}

// WithTransCardNo 设置 trans_card_no(承运人身份证号) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithTransCardNo(transCardNo *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.TransCardNo] = transCardNo })
}

// WithTransCardNos 设置 trans_card_no(承运人身份证号) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithTransCardNos(transCardNos []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.TransCardNo] = transCardNos })
}

// WithTransPhone 设置 trans_phone(承运人手机号码) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithTransPhone(transPhone *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.TransPhone] = transPhone })
}

// WithTransPhones 设置 trans_phone(承运人手机号码) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithTransPhones(transPhones []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.TransPhone] = transPhones })
}

// WithTransCarNo 设置 trans_car_no(承运人车牌号) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithTransCarNo(transCarNo *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.TransCarNo] = transCarNo })
}

// WithTransCarNos 设置 trans_car_no(承运人车牌号) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithTransCarNos(transCarNos []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.TransCarNo] = transCarNos })
}

// WithTransURL 设置 trans_url(承运人图片) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithTransURL(transURL *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.TransURL] = transURL })
}

// WithTransURLs 设置 trans_url(承运人图片) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithTransURLs(transURLs []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.TransURL] = transURLs })
}

// WithExtend 设置 extend(补全字段JSON) 字段作为 option 条件
func (obj *PurchaserOrderDAO) WithExtend(extend *string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.Extend] = extend })
}

// WithExtends 设置 extend(补全字段JSON) 字段的切片作为 option 条件
func (obj *PurchaserOrderDAO) WithExtends(extends []*string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[PurchaserOrderColumns.Extend] = extends })
}

// GetByOption 函数选项模式获取单条记录
func (obj *PurchaserOrderDAO) GetByOption(opts ...Option) (result *PurchaserOrder, err error) {

	err = obj.prepare(opts...).Find(&result).Error
	return
}

// GetListByOption 函数选项模式获取多条记录：支持分页
func (obj *PurchaserOrderDAO) GetListByOption(opts ...Option) (results []*PurchaserOrder, err error) {

	err = obj.prepare(opts...).Find(&results).Error
	return
}

// GetCountByOption 函数选项模式获取多条记录：支持统计记录总数
func (obj *PurchaserOrderDAO) GetCountByOption(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)

	obj.prepare(opts...).Count(&count)
	return
}

// GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名;" 标签)
func (obj *PurchaserOrderDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
	err = obj.prepare(opts...).Find(result).Error
	return
}

// UpdateByOption 更新数据
// 	参数说明：
//      purchaserOrder: 要更新的数据
//      opts: 更新的条件
func (obj *PurchaserOrderDAO) UpdateByOption(purchaserOrder *PurchaserOrder, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(purchaserOrder)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// GetFromID 通过单个 id(主键id) 字段值，获取单条记录
func (obj *PurchaserOrderDAO) GetFromID(id int64) (result *PurchaserOrder, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// GetsFromID 通过多个 id(主键id) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromID(ids []int64) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithIDs(ids))
	return
}

// GetFromOrgID 通过单个 org_id(一级供应商（平台企业）租户id) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromOrgID(orgID string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithOrgID(orgID))
	return
}

// GetsFromOrgID 通过多个 org_id(一级供应商（平台企业）租户id) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromOrgID(orgIDs []string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithOrgIDs(orgIDs))
	return
}

// GetFromOpenOrgID 通过单个 open_org_id(一级供应商（平台企业）开放租户id) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromOpenOrgID(openOrgID string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgID(openOrgID))
	return
}

// GetsFromOpenOrgID 通过多个 open_org_id(一级供应商（平台企业）开放租户id) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromOpenOrgID(openOrgIDs []string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithOpenOrgIDs(openOrgIDs))
	return
}

// GetFromPurchaserOrderID 通过单个 purchaser_order_id(采购单编号) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromPurchaserOrderID(purchaserOrderID string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserOrderID(purchaserOrderID))
	return
}

// GetsFromPurchaserOrderID 通过多个 purchaser_order_id(采购单编号) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromPurchaserOrderID(purchaserOrderIDs []string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserOrderIDs(purchaserOrderIDs))
	return
}

// GetFromOrderID 通过单个 order_id(订单编号) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromOrderID(orderID string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithOrderID(orderID))
	return
}

// GetsFromOrderID 通过多个 order_id(订单编号) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromOrderID(orderIDs []string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithOrderIDs(orderIDs))
	return
}

// GetFromIsDel 通过单个 is_del(是否删除：1-删除 0-否) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromIsDel(isDel int) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithIsDel(isDel))
	return
}

// GetsFromIsDel 通过多个 is_del(是否删除：1-删除 0-否) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromIsDel(isDels []int) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithIsDels(isDels))
	return
}

// GetFromOrderTime 通过单个 order_time(订单创建时间) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromOrderTime(orderTime int64) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithOrderTime(orderTime))
	return
}

// GetsFromOrderTime 通过多个 order_time(订单创建时间) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromOrderTime(orderTimes []int64) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithOrderTimes(orderTimes))
	return
}

// GetFromState 通过单个 state(订单状态：0待支付 1已支付 2 已取消 3.支付失败 4 支付中) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromState(state int) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithState(state))
	return
}

// GetsFromState 通过多个 state(订单状态：0待支付 1已支付 2 已取消 3.支付失败 4 支付中) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromState(states []int) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithStates(states))
	return
}

// GetFromURL 通过单个 url(采购单文件下载URL) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromURL(url string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithURL(url))
	return
}

// GetsFromURL 通过多个 url(采购单文件下载URL) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromURL(urls []string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithURLs(urls))
	return
}

// GetFromPurchaserID 通过单个 purchaser_id(采购商id) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromPurchaserID(purchaserID string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserID(purchaserID))
	return
}

// GetsFromPurchaserID 通过多个 purchaser_id(采购商id) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromPurchaserID(purchaserIDs []string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithPurchaserIDs(purchaserIDs))
	return
}

// GetFromSupplierID 通过单个 supplier_id(供应商id) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromSupplierID(supplierID string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithSupplierID(supplierID))
	return
}

// GetsFromSupplierID 通过多个 supplier_id(供应商id) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromSupplierID(supplierIDs []string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithSupplierIDs(supplierIDs))
	return
}

// GetFromTotalAmount 通过单个 total_amount(订单总金额（单位：分）) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromTotalAmount(totalAmount int64) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithTotalAmount(totalAmount))
	return
}

// GetsFromTotalAmount 通过多个 total_amount(订单总金额（单位：分）) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromTotalAmount(totalAmounts []int64) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithTotalAmounts(totalAmounts))
	return
}

// GetFromBuyAmount 通过单个 buy_amount(采购金额（单位：分）) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromBuyAmount(buyAmount int64) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithBuyAmount(buyAmount))
	return
}

// GetsFromBuyAmount 通过多个 buy_amount(采购金额（单位：分）) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromBuyAmount(buyAmounts []int64) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithBuyAmounts(buyAmounts))
	return
}

// GetFromServiceAmount 通过单个 service_amount(服务费（单位：分）) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromServiceAmount(serviceAmount int64) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithServiceAmount(serviceAmount))
	return
}

// GetsFromServiceAmount 通过多个 service_amount(服务费（单位：分）) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromServiceAmount(serviceAmounts []int64) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithServiceAmounts(serviceAmounts))
	return
}

// GetFromPayType 通过单个 pay_type(支付通道) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromPayType(payType string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithPayType(payType))
	return
}

// GetsFromPayType 通过多个 pay_type(支付通道) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromPayType(payTypes []string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithPayTypes(payTypes))
	return
}

// GetFromPayAccount 通过单个 pay_account(支付账户) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromPayAccount(payAccount string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithPayAccount(payAccount))
	return
}

// GetsFromPayAccount 通过多个 pay_account(支付账户) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromPayAccount(payAccounts []string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithPayAccounts(payAccounts))
	return
}

// GetFromSendType 通过单个 send_type(发货方式) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromSendType(sendType string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithSendType(sendType))
	return
}

// GetsFromSendType 通过多个 send_type(发货方式) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromSendType(sendTypes []string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithSendTypes(sendTypes))
	return
}

// GetFromReceiveAreaCode 通过单个 receive_area_code(收货所在地省市区) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromReceiveAreaCode(receiveAreaCode string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithReceiveAreaCode(receiveAreaCode))
	return
}

// GetsFromReceiveAreaCode 通过多个 receive_area_code(收货所在地省市区) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromReceiveAreaCode(receiveAreaCodes []string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithReceiveAreaCodes(receiveAreaCodes))
	return
}

// GetFromReceiveAddress 通过单个 receive_address(收货所在地详细地址) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromReceiveAddress(receiveAddress string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithReceiveAddress(receiveAddress))
	return
}

// GetsFromReceiveAddress 通过多个 receive_address(收货所在地详细地址) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromReceiveAddress(receiveAddresss []string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithReceiveAddresss(receiveAddresss))
	return
}

// GetFromSendAreaCode 通过单个 send_area_code(发货所在地省市区) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromSendAreaCode(sendAreaCode string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithSendAreaCode(sendAreaCode))
	return
}

// GetsFromSendAreaCode 通过多个 send_area_code(发货所在地省市区) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromSendAreaCode(sendAreaCodes []string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithSendAreaCodes(sendAreaCodes))
	return
}

// GetFromSendAddress 通过单个 send_address(发货所在地详细地址) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromSendAddress(sendAddress string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithSendAddress(sendAddress))
	return
}

// GetsFromSendAddress 通过多个 send_address(发货所在地详细地址) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromSendAddress(sendAddresss []string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithSendAddresss(sendAddresss))
	return
}

// GetFromReceiveURL 通过单个 receive_url(收货凭据/图片) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromReceiveURL(receiveURL *string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithReceiveURL(receiveURL))
	return
}

// GetsFromReceiveURL 通过多个 receive_url(收货凭据/图片) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromReceiveURL(receiveURLs []*string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithReceiveURLs(receiveURLs))
	return
}

// GetFromWeighURL 通过单个 weigh_url(过磅单图片) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromWeighURL(weighURL *string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithWeighURL(weighURL))
	return
}

// GetsFromWeighURL 通过多个 weigh_url(过磅单图片) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromWeighURL(weighURLs []*string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithWeighURLs(weighURLs))
	return
}

// GetFromTransName 通过单个 trans_name(承运人姓名) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromTransName(transName *string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithTransName(transName))
	return
}

// GetsFromTransName 通过多个 trans_name(承运人姓名) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromTransName(transNames []*string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithTransNames(transNames))
	return
}

// GetFromTransCardNo 通过单个 trans_card_no(承运人身份证号) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromTransCardNo(transCardNo *string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithTransCardNo(transCardNo))
	return
}

// GetsFromTransCardNo 通过多个 trans_card_no(承运人身份证号) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromTransCardNo(transCardNos []*string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithTransCardNos(transCardNos))
	return
}

// GetFromTransPhone 通过单个 trans_phone(承运人手机号码) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromTransPhone(transPhone *string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithTransPhone(transPhone))
	return
}

// GetsFromTransPhone 通过多个 trans_phone(承运人手机号码) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromTransPhone(transPhones []*string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithTransPhones(transPhones))
	return
}

// GetFromTransCarNo 通过单个 trans_car_no(承运人车牌号) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromTransCarNo(transCarNo *string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithTransCarNo(transCarNo))
	return
}

// GetsFromTransCarNo 通过多个 trans_car_no(承运人车牌号) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromTransCarNo(transCarNos []*string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithTransCarNos(transCarNos))
	return
}

// GetFromTransURL 通过单个 trans_url(承运人图片) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromTransURL(transURL *string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithTransURL(transURL))
	return
}

// GetsFromTransURL 通过多个 trans_url(承运人图片) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromTransURL(transURLs []*string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithTransURLs(transURLs))
	return
}

// GetFromExtend 通过单个 extend(补全字段JSON) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetFromExtend(extend *string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithExtend(extend))
	return
}

// GetsFromExtend 通过多个 extend(补全字段JSON) 字段值，获取多条记录
func (obj *PurchaserOrderDAO) GetsFromExtend(extends []*string) (results []*PurchaserOrder, err error) {
	results, err = obj.GetListByOption(obj.WithExtends(extends))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// FetchByPrimaryKey 通过 id 字段值，获取单条记录
func (obj *PurchaserOrderDAO) FetchByPrimaryKey(id int64) (result *PurchaserOrder, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
