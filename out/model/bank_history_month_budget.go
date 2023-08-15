package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-08-15 19:59:34
 * @Desc:   bank_history_month_budget 表
 */

// 银行账号历史月度收支统计表
type BankHistoryMonthBudget struct {
	ID                    uint64 `gorm:"column:id;primaryKey;type:bigint(20) unsigned;autoIncrement" json:"id"`                              // 主键
	OrgID                 int64  `gorm:"column:org_id;type:int(11);not null;DEFAULT '0'" json:"org_id"`                                      // 租户ID
	IncomeExpenditureType int8   `gorm:"column:income_expenditure_type;type:tinyint(1);not null;DEFAULT '1'" json:"income_expenditure_type"` // 收支类型：1(收入)、2(支出)
	BankNumber            string `gorm:"column:bank_number;type:varchar(50);not null;DEFAULT ''" json:"bank_number"`                         // 银行卡号
	BankAccountName       string `gorm:"column:bank_account_name;type:varchar(64);not null;DEFAULT ''" json:"bank_account_name"`             // 银行账户名称
	StatisticsDate        int64  `gorm:"column:statistics_date;type:int(11);not null;DEFAULT '0'" json:"statistics_date"`                    // 统计:年月,202308
	TradingTotalAmount    int64  `gorm:"column:trading_total_amount;type:bigint(20);not null;DEFAULT '0'" json:"trading_total_amount"`       // 交易总金额，单位分
	CreatedAt             int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`                              // 创建时间
}

// 获取结构体对应的表名方法
func (m *BankHistoryMonthBudget) TableName() string {
	return "bank_history_month_budget"
}

// 实例化结构体对象
func NewBankHistoryMonthBudget() *BankHistoryMonthBudget {
	return &BankHistoryMonthBudget{}
}

// 获取主键的对应字段
func (m *BankHistoryMonthBudget) PrimaryKey() string {
	return BankHistoryMonthBudgetColumns.ID
}

// 获取主键值
func (m *BankHistoryMonthBudget) PrimaryKeyValue() uint64 {
	return m.ID
}

// 表字段的映射
var BankHistoryMonthBudgetColumns = struct {
	ID                    string
	OrgID                 string
	IncomeExpenditureType string
	BankNumber            string
	BankAccountName       string
	StatisticsDate        string
	TradingTotalAmount    string
	CreatedAt             string
}{
	ID:                    "id",
	OrgID:                 "org_id",
	IncomeExpenditureType: "income_expenditure_type",
	BankNumber:            "bank_number",
	BankAccountName:       "bank_account_name",
	StatisticsDate:        "statistics_date",
	TradingTotalAmount:    "trading_total_amount",
	CreatedAt:             "created_at",
}

// 包含所有表字段的切片
var BankHistoryMonthBudgetAllColumns = []string{
	BankHistoryMonthBudgetColumns.ID,                    // 主键
	BankHistoryMonthBudgetColumns.OrgID,                 // 租户ID
	BankHistoryMonthBudgetColumns.IncomeExpenditureType, // 收支类型：1(收入)、2(支出)
	BankHistoryMonthBudgetColumns.BankNumber,            // 银行卡号
	BankHistoryMonthBudgetColumns.BankAccountName,       // 银行账户名称
	BankHistoryMonthBudgetColumns.StatisticsDate,        // 统计:年月,202308
	BankHistoryMonthBudgetColumns.TradingTotalAmount,    // 交易总金额，单位分
	BankHistoryMonthBudgetColumns.CreatedAt,             // 创建时间

}

// 设置值：主键
func (m *BankHistoryMonthBudget) SetID(v uint64) {
	m.ID = v
}

// 设置值：租户ID
func (m *BankHistoryMonthBudget) SetOrgID(v int64) {
	m.OrgID = v
}

// 设置值：收支类型：1(收入)、2(支出)
func (m *BankHistoryMonthBudget) SetIncomeExpenditureType(v int8) {
	m.IncomeExpenditureType = v
}

// 设置值：银行卡号
func (m *BankHistoryMonthBudget) SetBankNumber(v string) {
	m.BankNumber = v
}

// 设置值：银行账户名称
func (m *BankHistoryMonthBudget) SetBankAccountName(v string) {
	m.BankAccountName = v
}

// 设置值：统计:年月,202308
func (m *BankHistoryMonthBudget) SetStatisticsDate(v int64) {
	m.StatisticsDate = v
}

// 设置值：交易总金额，单位分
func (m *BankHistoryMonthBudget) SetTradingTotalAmount(v int64) {
	m.TradingTotalAmount = v
}

// 设置值：创建时间
func (m *BankHistoryMonthBudget) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 获取值：主键
func (m *BankHistoryMonthBudget) GetID() uint64 {
	return m.ID
}

// 获取值：租户ID
func (m *BankHistoryMonthBudget) GetOrgID() int64 {
	return m.OrgID
}

// 获取值：收支类型：1(收入)、2(支出)
func (m *BankHistoryMonthBudget) GetIncomeExpenditureType() int8 {
	return m.IncomeExpenditureType
}

// 获取值：银行卡号
func (m *BankHistoryMonthBudget) GetBankNumber() string {
	return m.BankNumber
}

// 获取值：银行账户名称
func (m *BankHistoryMonthBudget) GetBankAccountName() string {
	return m.BankAccountName
}

// 获取值：统计:年月,202308
func (m *BankHistoryMonthBudget) GetStatisticsDate() int64 {
	return m.StatisticsDate
}

// 获取值：交易总金额，单位分
func (m *BankHistoryMonthBudget) GetTradingTotalAmount() int64 {
	return m.TradingTotalAmount
}

// 获取值：创建时间
func (m *BankHistoryMonthBudget) GetCreatedAt() int64 {
	return m.CreatedAt
}
