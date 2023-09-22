package model

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023-08-25 14:26:41
 * @Desc:   task_log 表
 */

// 任务表
type TaskLog struct {
	ID             int64  `gorm:"column:id;primaryKey;type:int(11);" json:"id"`                                          // 主键ID
	TaskType       int8   `gorm:"column:task_type;type:tinyint(1);not null;DEFAULT '1'" json:"task_type"`                // 任务类型；1:统计月度交易额; 2:统计月末金额
	TaskStatus     int8   `gorm:"column:task_status;type:tinyint(1);not null;DEFAULT '1'" json:"task_status"`            // 任务状态：1:成功；2:失败
	TaskFailReason string `gorm:"column:task_fail_reason;type:varchar(255);not null;DEFAULT ''" json:"task_fail_reason"` // 任务失败原因
	TaskLogid      string `gorm:"column:task_logid;type:varchar(64);not null;DEFAULT ''" json:"task_logid"`              // 任务执行的logid
	TaskInfo       string `gorm:"column:task_info;type:text;not null;" json:"task_info"`                                 // 任务信息
	DeletedAt      int64  `gorm:"column:deleted_at;type:int(11);not null;DEFAULT '0'" json:"deleted_at"`                 // 删除时间
	UpdatedAt      int64  `gorm:"column:updated_at;type:int(11);not null;DEFAULT '0'" json:"updated_at"`                 // 更新时间
	CreatedAt      int64  `gorm:"column:created_at;type:int(11);not null;DEFAULT '0'" json:"created_at"`                 // 创建时间
}

// 获取结构体对应的表名方法
func (m *TaskLog) TableName() string {
	return "task_log"
}

// 实例化结构体对象
func NewTaskLog() *TaskLog {
	return &TaskLog{}
}

// 获取主键的对应字段
func (m *TaskLog) PrimaryKey() string {
	return TaskLogColumns.ID
}

// 获取主键值
func (m *TaskLog) PrimaryKeyValue() int64 {
	return m.ID
}

// 表字段的映射
var TaskLogColumns = struct {
	ID             string
	TaskType       string
	TaskStatus     string
	TaskFailReason string
	TaskLogid      string
	TaskInfo       string
	DeletedAt      string
	UpdatedAt      string
	CreatedAt      string
}{
	ID:             "id",
	TaskType:       "task_type",
	TaskStatus:     "task_status",
	TaskFailReason: "task_fail_reason",
	TaskLogid:      "task_logid",
	TaskInfo:       "task_info",
	DeletedAt:      "deleted_at",
	UpdatedAt:      "updated_at",
	CreatedAt:      "created_at",
}

// 包含所有表字段的切片
var TaskLogAllColumns = []string{
	TaskLogColumns.ID,             // 主键ID
	TaskLogColumns.TaskType,       // 任务类型；1:统计月度交易额; 2:统计月末金额
	TaskLogColumns.TaskStatus,     // 任务状态：1:成功；2:失败
	TaskLogColumns.TaskFailReason, // 任务失败原因
	TaskLogColumns.TaskLogid,      // 任务执行的logid
	TaskLogColumns.TaskInfo,       // 任务信息
	TaskLogColumns.DeletedAt,      // 删除时间
	TaskLogColumns.UpdatedAt,      // 更新时间
	TaskLogColumns.CreatedAt,      // 创建时间

}

// 设置值：主键ID
func (m *TaskLog) SetID(v int64) {
	m.ID = v
}

// 设置值：任务类型；1:统计月度交易额; 2:统计月末金额
func (m *TaskLog) SetTaskType(v int8) {
	m.TaskType = v
}

// 设置值：任务状态：1:成功；2:失败
func (m *TaskLog) SetTaskStatus(v int8) {
	m.TaskStatus = v
}

// 设置值：任务失败原因
func (m *TaskLog) SetTaskFailReason(v string) {
	m.TaskFailReason = v
}

// 设置值：任务执行的logid
func (m *TaskLog) SetTaskLogid(v string) {
	m.TaskLogid = v
}

// 设置值：任务信息
func (m *TaskLog) SetTaskInfo(v string) {
	m.TaskInfo = v
}

// 设置值：删除时间
func (m *TaskLog) SetDeletedAt(v int64) {
	m.DeletedAt = v
}

// 设置值：更新时间
func (m *TaskLog) SetUpdatedAt(v int64) {
	m.UpdatedAt = v
}

// 设置值：创建时间
func (m *TaskLog) SetCreatedAt(v int64) {
	m.CreatedAt = v
}

// 获取值：主键ID
func (m *TaskLog) GetID() int64 {
	return m.ID
}

// 获取值：任务类型；1:统计月度交易额; 2:统计月末金额
func (m *TaskLog) GetTaskType() int8 {
	return m.TaskType
}

// 获取值：任务状态：1:成功；2:失败
func (m *TaskLog) GetTaskStatus() int8 {
	return m.TaskStatus
}

// 获取值：任务失败原因
func (m *TaskLog) GetTaskFailReason() string {
	return m.TaskFailReason
}

// 获取值：任务执行的logid
func (m *TaskLog) GetTaskLogid() string {
	return m.TaskLogid
}

// 获取值：任务信息
func (m *TaskLog) GetTaskInfo() string {
	return m.TaskInfo
}

// 获取值：删除时间
func (m *TaskLog) GetDeletedAt() int64 {
	return m.DeletedAt
}

// 获取值：更新时间
func (m *TaskLog) GetUpdatedAt() int64 {
	return m.UpdatedAt
}

// 获取值：创建时间
func (m *TaskLog) GetCreatedAt() int64 {
	return m.CreatedAt
}
