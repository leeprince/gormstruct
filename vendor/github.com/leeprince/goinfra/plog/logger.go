package plog

import (
    "context"
	"github.com/sirupsen/logrus"
    "io"
    "sync"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/1 下午10:45
 * @Desc:	Plog 实例
 */

type Plog struct {
	// logrus Logger 实例
    *logrus.Logger
    
    // 导出的文件信息
    outFileInfo               FileInfo
}

// 文件信息
type FileInfo struct {
    dirPath  string
    filename string
}

var (
    logger Plog
    once   sync.Once
)

func init() {
    once.Do(func() {
        NewDefaultLogger()
    })
}

func SetLogger(l *logrus.Logger) {
    logger.Logger = l
}

func GetLogger() *Plog {
    return &logger
}

// --- 设置 logger 参数
func SetFormatter(formatter logrus.Formatter) {
    logger.SetFormatter(formatter)
}

func SetLevel(level Level) {
    logger.SetLevel(PLevel(level))
}

func AddHook(hook logrus.Hook) {
    logger.AddHook(hook)
}

func SetOutput(output io.Writer) {
    logger.SetOutput(output)
}

// 是否记录日志调用者的函数信息(（文件名、方法、行号）). 默认false
// 原始的 logger.SetReportCaller() 方法的实现说明：github.com/sirupsen/logrus@v1.8.1/entry.go@getCaller 方法中实现 retrieves the name of the first non-logrus calling function（译：检索第一个非logrus调用函数的名称）
// 		1. 外部直接调用 Debug, Print, Info, Warn, Error, Fatal or Panic 等记录日志方法，记录的起始标记实际是 logger 的方法的位置（检索第一个非logrus调用函数的名称）
// 		2. 开启标记之后，外部调用正确方式：外部调用的地方配合使用 WithXXX().记录日志方法(...)
// 目标：兼容原始的 logger.SetReportCaller() 方法，并且检索第一个非 plog 包调用函数的名称。外部可以直接使用本包中的 Debug, Print, Info, Warn, Error, Fatal or Panic 等记录日志方法，即可记录日志调用者的标记
// 请求参数说明： reportCaller == true 自动添加钩子 `AddHookReportCaller` 实现目标
// 使用注意：
// 	1. 返回的 Plog 结构体当作 gorm.io/gorm 的 writer 时，gorm 会自动打印调用者的标记到 msg 中，无需打开此设置。即使打开记录的也是 gorm.io/gorm 的 logger 调用 github.com/sirupsen/logrus 的 Printf 的位置
func SetReportCaller(reportCaller bool, levels ...Level) {
    // 开启 ReportCaller
    logger.SetReportCaller(reportCaller)
    
    if !reportCaller {
        return
    }
    addHookReportCaller(PLevels(levels)...)
}

// --- 设置 logger 参数 - end

// --- 设置 entry 参数, 并返回 Entry 对象
func WithField(key string, value interface{}) *logrus.Entry {
    return logger.WithField(key, value)
}

func WithFields(fields map[string]interface{}) *logrus.Entry {
    return logger.WithFields(fields)
}

func WithError(err error) *logrus.Entry {
    return logger.WithError(err)
}

func WithContext(ctx context.Context) *logrus.Entry {
    return logger.WithContext(ctx)
}

// ---  设置 entry 参数, 并返回 Entry 对象 - end

// -- Plog 的方法
func (p *Plog) GetOutFileInfo() (dirPath, filename string) {
    return p.outFileInfo.dirPath, p.outFileInfo.filename
}

// -- Plog 的方法 -end
