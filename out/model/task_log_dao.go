package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-08-25 14:26:41
 * @Desc:   task_log 表的 DAO 层
 */

type TaskLogDAO struct {
	*_BaseDAO
}

// 初始化 TaskLogDAO
func NewTaskLogDAO(ctx context.Context, db *gorm.DB) *TaskLogDAO {
	if db == nil {
		panic(fmt.Errorf("TaskLogDAO need init by db"))
	}
	ctx, cancel := context.WithCancel(ctx)
	return &TaskLogDAO{
		_BaseDAO: &_BaseDAO{
			DB:               db.Model(&TaskLog{}),
			db:               db,
			model:            TaskLog{},
			ctx:              ctx,
			cancel:           cancel,
			timeout:          -1,
			columns:          TaskLogAllColumns,
			isDefaultColumns: true,
		},
	}
}

// 获取表名称
func (obj *TaskLogDAO) GetTableName() string {
	taskLog := &TaskLog{}
	return taskLog.TableName()
}

// 存在则更新，否则插入
func (obj *TaskLogDAO) Save(taskLog *TaskLog) (rowsAffected int64, err error) {
	if taskLog.PrimaryKeyValue() > 0 {
		return obj.UpdateByOption(taskLog, obj.WithID(taskLog.ID))
	}
	return obj.Create(taskLog)
}

// 创建数据:允许单条/批量创建，批量创建时传入切片
func (obj *TaskLogDAO) Create(taskLog interface{}) (rowsAffected int64, err error) {
	tx := obj.withContext().Create(taskLog)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 ---

// 设置 id(主键ID) 字段作为 option 条件
func (obj *TaskLogDAO) WithID(id int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.ID] = id })
}

// 设置 id(主键ID) 字段的切片作为 option 条件
func (obj *TaskLogDAO) WithIDs(ids []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.ID] = ids })
}

// 设置 task_type(任务类型；1:统计月度交易额; 2:统计月末金额) 字段作为 option 条件
func (obj *TaskLogDAO) WithTaskType(taskType int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.TaskType] = taskType })
}

// 设置 task_type(任务类型；1:统计月度交易额; 2:统计月末金额) 字段的切片作为 option 条件
func (obj *TaskLogDAO) WithTaskTypes(taskTypes []int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.TaskType] = taskTypes })
}

// 设置 task_status(任务状态：1:成功；2:失败) 字段作为 option 条件
func (obj *TaskLogDAO) WithTaskStatus(taskStatus int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.TaskStatus] = taskStatus })
}

// 设置 task_status(任务状态：1:成功；2:失败) 字段的切片作为 option 条件
func (obj *TaskLogDAO) WithTaskStatuss(taskStatuss []int8) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.TaskStatus] = taskStatuss })
}

// 设置 task_fail_reason(任务失败原因) 字段作为 option 条件
func (obj *TaskLogDAO) WithTaskFailReason(taskFailReason string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.TaskFailReason] = taskFailReason })
}

// 设置 task_fail_reason(任务失败原因) 字段的切片作为 option 条件
func (obj *TaskLogDAO) WithTaskFailReasons(taskFailReasons []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.TaskFailReason] = taskFailReasons })
}

// 设置 task_logid(任务执行的logid) 字段作为 option 条件
func (obj *TaskLogDAO) WithTaskLogid(taskLogid string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.TaskLogid] = taskLogid })
}

// 设置 task_logid(任务执行的logid) 字段的切片作为 option 条件
func (obj *TaskLogDAO) WithTaskLogids(taskLogids []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.TaskLogid] = taskLogids })
}

// 设置 task_info(任务信息) 字段作为 option 条件
func (obj *TaskLogDAO) WithTaskInfo(taskInfo string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.TaskInfo] = taskInfo })
}

// 设置 task_info(任务信息) 字段的切片作为 option 条件
func (obj *TaskLogDAO) WithTaskInfos(taskInfos []string) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.TaskInfo] = taskInfos })
}

// 设置 deleted_at(删除时间) 字段作为 option 条件
func (obj *TaskLogDAO) WithDeletedAt(deletedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.DeletedAt] = deletedAt })
}

// 设置 deleted_at(删除时间) 字段的切片作为 option 条件
func (obj *TaskLogDAO) WithDeletedAts(deletedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.DeletedAt] = deletedAts })
}

// 设置 updated_at(更新时间) 字段作为 option 条件
func (obj *TaskLogDAO) WithUpdatedAt(updatedAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.UpdatedAt] = updatedAt })
}

// 设置 updated_at(更新时间) 字段的切片作为 option 条件
func (obj *TaskLogDAO) WithUpdatedAts(updatedAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.UpdatedAt] = updatedAts })
}

// 设置 created_at(创建时间) 字段作为 option 条件
func (obj *TaskLogDAO) WithCreatedAt(createdAt int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.CreatedAt] = createdAt })
}

// 设置 created_at(创建时间) 字段的切片作为 option 条件
func (obj *TaskLogDAO) WithCreatedAts(createdAts []int64) Option {
	return queryOptionFunc(func(o *options) { o.queryMap[TaskLogColumns.CreatedAt] = createdAts })
}

// 函数选项模式获取单条记录
func (obj *TaskLogDAO) GetByOption(opts ...Option) (result *TaskLog, err error) {
	opts = append(opts, obj.WithDeletedAt(0))
	err = obj.prepare(opts...).Find(&result).Error
	return
}

// 函数选项模式获取多条记录：支持分页
func (obj *TaskLogDAO) GetByOptions(opts ...Option) (results []*TaskLog, err error) {
	err = obj.prepare(opts...).Find(&results).Error
	return
}

// 函数选项模式获取多条记录：支持统计记录总数
func (obj *TaskLogDAO) GetCountByOptions(opts ...Option) (count int64) {
	obj.setIsDefaultColumns(false)
	obj.prepare(opts...).Count(&count)
	return
}

// 更新数据
//  参数说明：
//      taskLog: 要更新的数据
//      opts: 更新的条件
func (obj *TaskLogDAO) UpdateByOption(taskLog *TaskLog, opts ...Option) (rowsAffected int64, err error) {
	obj.setIsDefaultColumns(false)
	tx := obj.prepare(opts...).Updates(taskLog)
	return tx.RowsAffected, tx.Error
}

// --- 表中的字段作为 option 条件 -END ---

// --- 单个字段作为查询条件 ---

// 通过单个 id(主键ID) 字段值，获取单条记录
func (obj *TaskLogDAO) GetFromID(id int64) (result *TaskLog, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// 通过多个 id(主键ID) 字段值，获取多条记录
func (obj *TaskLogDAO) GetsFromID(ids []int64) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithIDs(ids))
	return
}

// 通过单个 task_type(任务类型；1:统计月度交易额; 2:统计月末金额) 字段值，获取多条记录
func (obj *TaskLogDAO) GetFromTaskType(taskType int8) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithTaskType(taskType))
	return
}

// 通过多个 task_type(任务类型；1:统计月度交易额; 2:统计月末金额) 字段值，获取多条记录
func (obj *TaskLogDAO) GetsFromTaskType(taskTypes []int8) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithTaskTypes(taskTypes))
	return
}

// 通过单个 task_status(任务状态：1:成功；2:失败) 字段值，获取多条记录
func (obj *TaskLogDAO) GetFromTaskStatus(taskStatus int8) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithTaskStatus(taskStatus))
	return
}

// 通过多个 task_status(任务状态：1:成功；2:失败) 字段值，获取多条记录
func (obj *TaskLogDAO) GetsFromTaskStatus(taskStatuss []int8) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithTaskStatuss(taskStatuss))
	return
}

// 通过单个 task_fail_reason(任务失败原因) 字段值，获取多条记录
func (obj *TaskLogDAO) GetFromTaskFailReason(taskFailReason string) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithTaskFailReason(taskFailReason))
	return
}

// 通过多个 task_fail_reason(任务失败原因) 字段值，获取多条记录
func (obj *TaskLogDAO) GetsFromTaskFailReason(taskFailReasons []string) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithTaskFailReasons(taskFailReasons))
	return
}

// 通过单个 task_logid(任务执行的logid) 字段值，获取多条记录
func (obj *TaskLogDAO) GetFromTaskLogid(taskLogid string) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithTaskLogid(taskLogid))
	return
}

// 通过多个 task_logid(任务执行的logid) 字段值，获取多条记录
func (obj *TaskLogDAO) GetsFromTaskLogid(taskLogids []string) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithTaskLogids(taskLogids))
	return
}

// 通过单个 task_info(任务信息) 字段值，获取多条记录
func (obj *TaskLogDAO) GetFromTaskInfo(taskInfo string) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithTaskInfo(taskInfo))
	return
}

// 通过多个 task_info(任务信息) 字段值，获取多条记录
func (obj *TaskLogDAO) GetsFromTaskInfo(taskInfos []string) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithTaskInfos(taskInfos))
	return
}

// 通过单个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *TaskLogDAO) GetFromDeletedAt(deletedAt int64) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithDeletedAt(deletedAt))
	return
}

// 通过多个 deleted_at(删除时间) 字段值，获取多条记录
func (obj *TaskLogDAO) GetsFromDeletedAt(deletedAts []int64) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithDeletedAts(deletedAts))
	return
}

// 通过单个 updated_at(更新时间) 字段值，获取多条记录
func (obj *TaskLogDAO) GetFromUpdatedAt(updatedAt int64) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithUpdatedAt(updatedAt))
	return
}

// 通过多个 updated_at(更新时间) 字段值，获取多条记录
func (obj *TaskLogDAO) GetsFromUpdatedAt(updatedAts []int64) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithUpdatedAts(updatedAts))
	return
}

// 通过单个 created_at(创建时间) 字段值，获取多条记录
func (obj *TaskLogDAO) GetFromCreatedAt(createdAt int64) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAt(createdAt))
	return
}

// 通过多个 created_at(创建时间) 字段值，获取多条记录
func (obj *TaskLogDAO) GetsFromCreatedAt(createdAts []int64) (results []*TaskLog, err error) {
	results, err = obj.GetByOptions(obj.WithCreatedAts(createdAts))
	return
}

// --- 单个字段作为查询条件 -END ---

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 ---

// 通过 id 字段值，获取单条记录
func (obj *TaskLogDAO) FetchByPrimaryKey(id int64) (result *TaskLog, err error) {
	result, err = obj.GetByOption(obj.WithID(id))
	return
}

// --- 通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件 -END ---
