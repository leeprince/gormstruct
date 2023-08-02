package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-07-26 18:10:28
 * @Desc:   import_invoice_task 表的 DAO 层
 */

type ImportInvoiceTaskDAO struct {
	*_BaseDAO
}

// 初始化 ImportInvoiceTaskDAO
func NewImportInvoiceTaskDAO(ctx context.Context, db *gorm.DB) *ImportInvoiceTaskDAO {
	if db == nil {
		panic(fmt.Errorf("ImportInvoiceTaskDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &ImportInvoiceTaskDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&ImportInvoiceTask{}),
			db:               db,
			model:            ImportInvoiceTask{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          ImportInvoiceTaskAllColumns,
			isDefaultColumns: true,
		},
	}
}

// 获取表名称
func (obj *ImportInvoiceTaskDAO) GetTableName() string {
	importInvoiceTask := &ImportInvoiceTask{}
	return importInvoiceTask.TableName()
}

// 存在则更新，否则插入
func (obj *ImportInvoiceTaskDAO) Save(importInvoiceTask *ImportInvoiceTask) (rowsAffected int64, err error) {
	if importInvoiceTask.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(importInvoiceTask, obj.WithID(importInvoiceTask.ID))
	}
	return obj.Create(importInvoiceTask)
}

// 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *ImportInvoiceTaskDAO) Create(importInvoiceTask interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(importInvoiceTask)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// 设置 id(主键) 字段作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.ID] = id })
}

// 设置 id(主键) 字段的切片作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.ID] = ids })
}

// 设置 org_id(租户ID) 字段作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithOrgID(orgID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.OrgID] = orgID })
}

// 设置 org_id(租户ID) 字段的切片作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithOrgIDs(orgIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.OrgID] = orgIDs })
}

// 设置 client_id(纳税人ID) 字段作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithClientID(clientID int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.ClientID] = clientID })
}

// 设置 client_id(纳税人ID) 字段的切片作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithClientIDs(clientIDs []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.ClientID] = clientIDs })
}

// 设置 enterprise_id(企业ID) 字段作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithEnterpriseID(enterpriseID string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.EnterpriseID] = enterpriseID })
}

// 设置 enterprise_id(企业ID) 字段的切片作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithEnterpriseIDs(enterpriseIDs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.EnterpriseID] = enterpriseIDs })
}

// 设置 tax_code(税号) 字段作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithTaxCode(taxCode string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.TaxCode] = taxCode })
}

// 设置 tax_code(税号) 字段的切片作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithTaxCodes(taxCodes []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.TaxCode] = taxCodes })
}

// 设置 file_md5(源文件md5) 字段作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithFileMd5(fileMd5 string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.FileMd5] = fileMd5 })
}

// 设置 file_md5(源文件md5) 字段的切片作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithFileMd5s(fileMd5s []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.FileMd5] = fileMd5s })
}

// 设置 source_file_name(源文件名称) 字段作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithSourceFileName(sourceFileName string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.SourceFileName] = sourceFileName })
}

// 设置 source_file_name(源文件名称) 字段的切片作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithSourceFileNames(sourceFileNames []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.SourceFileName] = sourceFileNames })
}

// 设置 file_cos_url(源文件cos地址) 字段作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithFileCosURL(fileCosURL string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.FileCosURL] = fileCosURL })
}

// 设置 file_cos_url(源文件cos地址) 字段的切片作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithFileCosURLs(fileCosURLs []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.FileCosURL] = fileCosURLs })
}

// 设置 task_status(任务状态：1(等待中)；2(成功)；3(失败)) 字段作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithTaskStatus(taskStatus int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.TaskStatus] = taskStatus })
}

// 设置 task_status(任务状态：1(等待中)；2(成功)；3(失败)) 字段的切片作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithTaskStatuss(taskStatuss []int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.TaskStatus] = taskStatuss })
}

// 设置 fail_type(失败的原因类型；1：文件错误；2：识别/解析错误，允许重试) 字段作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithFailType(failType int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.FailType] = failType })
}

// 设置 fail_type(失败的原因类型；1：文件错误；2：识别/解析错误，允许重试) 字段的切片作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithFailTypes(failTypes []int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.FailType] = failTypes })
}

// 设置 fail_content(失败原因) 字段作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithFailContent(failContent string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.FailContent] = failContent })
}

// 设置 fail_content(失败原因) 字段的切片作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithFailContents(failContents []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.FailContent] = failContents })
}

// 设置 retry_num(重试次数) 字段作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithRetryNum(retryNum int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.RetryNum] = retryNum })
}

// 设置 retry_num(重试次数) 字段的切片作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithRetryNums(retryNums []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.RetryNum] = retryNums })
}

// 设置 created_at(创建时间) 字段作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.CreatedAt] = createdAt })
}

// 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.CreatedAt] = createdAts })
}

// 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.UpdatedAt] = updatedAt })
}

// 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *ImportInvoiceTaskDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[ImportInvoiceTaskColumns.UpdatedAt] = updatedAts })
}

// 函数选项模式获取单条记录
func (obj *ImportInvoiceTaskDAO) GetByOption(opts ...Option) (result *ImportInvoiceTask, err error) {

	err = obj.prepare(opts...).Find(&result).Error
	return
}

// 函数选项模式获取多条记录：支持分页
func (obj *ImportInvoiceTaskDAO) GetByOptions(opts ...Option) (results []*ImportInvoiceTask, err error) {
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *ImportInvoiceTaskDAO) GetCountByOptions(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	obj.prepare(opts...).Count(&count)
	return
}

// 更新数据
//  参数说明：
//      importInvoiceTask: 要更新的数据
//      opts: 更新的条件
func (obj *ImportInvoiceTaskDAO) UpdateByOption(importInvoiceTask *ImportInvoiceTask, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(importInvoiceTask)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// 通过单个 id(主键) 字段值，获取单条记录
func (obj *ImportInvoiceTaskDAO) GetFromID(id int64) (result *ImportInvoiceTask, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过多个 id(主键) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetsFromID(ids []int64) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithIDs(ids))
	return
}

// 通过单个 org_id(租户ID) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetFromOrgID(orgID int64) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithOrgID(orgID))
	return
}

// 通过多个 org_id(租户ID) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetsFromOrgID(orgIDs []int64) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithOrgIDs(orgIDs))
	return
}

// 通过单个 client_id(纳税人ID) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetFromClientID(clientID int64) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithClientID(clientID))
	return
}

// 通过多个 client_id(纳税人ID) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetsFromClientID(clientIDs []int64) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithClientIDs(clientIDs))
	return
}

// 通过单个 enterprise_id(企业ID) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetFromEnterpriseID(enterpriseID string) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithEnterpriseID(enterpriseID))
	return
}

// 通过多个 enterprise_id(企业ID) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetsFromEnterpriseID(enterpriseIDs []string) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithEnterpriseIDs(enterpriseIDs))
	return
}

// 通过单个 tax_code(税号) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetFromTaxCode(taxCode string) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithTaxCode(taxCode))
	return
}

// 通过多个 tax_code(税号) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetsFromTaxCode(taxCodes []string) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithTaxCodes(taxCodes))
	return
}

// 通过单个 file_md5(源文件md5) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetFromFileMd5(fileMd5 string) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithFileMd5(fileMd5))
	return
}

// 通过多个 file_md5(源文件md5) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetsFromFileMd5(fileMd5s []string) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithFileMd5s(fileMd5s))
	return
}

// 通过单个 source_file_name(源文件名称) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetFromSourceFileName(sourceFileName string) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithSourceFileName(sourceFileName))
	return
}

// 通过多个 source_file_name(源文件名称) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetsFromSourceFileName(sourceFileNames []string) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithSourceFileNames(sourceFileNames))
	return
}

// 通过单个 file_cos_url(源文件cos地址) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetFromFileCosURL(fileCosURL string) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithFileCosURL(fileCosURL))
	return
}

// 通过多个 file_cos_url(源文件cos地址) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetsFromFileCosURL(fileCosURLs []string) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithFileCosURLs(fileCosURLs))
	return
}

// 通过单个 task_status(任务状态：1(等待中)；2(成功)；3(失败)) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetFromTaskStatus(taskStatus int8) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithTaskStatus(taskStatus))
	return
}

// 通过多个 task_status(任务状态：1(等待中)；2(成功)；3(失败)) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetsFromTaskStatus(taskStatuss []int8) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithTaskStatuss(taskStatuss))
	return
}

// 通过单个 fail_type(失败的原因类型；1：文件错误；2：识别/解析错误，允许重试) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetFromFailType(failType int8) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithFailType(failType))
	return
}

// 通过多个 fail_type(失败的原因类型；1：文件错误；2：识别/解析错误，允许重试) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetsFromFailType(failTypes []int8) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithFailTypes(failTypes))
	return
}

// 通过单个 fail_content(失败原因) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetFromFailContent(failContent string) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithFailContent(failContent))
	return
}

// 通过多个 fail_content(失败原因) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetsFromFailContent(failContents []string) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithFailContents(failContents))
	return
}

// 通过单个 retry_num(重试次数) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetFromRetryNum(retryNum int64) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithRetryNum(retryNum))
	return
}

// 通过多个 retry_num(重试次数) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetsFromRetryNum(retryNums []int64) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithRetryNums(retryNums))
	return
}

// 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetFromCreatedAt(createdAt int64) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAt(createdAt))
	return
}

// 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetsFromCreatedAt(createdAts []int64) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAts(createdAts))
	return
}

// 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetFromUpdatedAt(updatedAt int64) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithUpdatedAt(updatedAt))
	return
}

// 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *ImportInvoiceTaskDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*ImportInvoiceTask, err error) {
	results, err = obj.GetByOptions(obj.WithUpdatedAts(updatedAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// 通过 id 字段值，获取单条记录
func (obj *ImportInvoiceTaskDAO) FetchByPrimaryKey(id int64) (result *ImportInvoiceTask, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
