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
 * @Desc:
 */

type Plog struct {
	*logrus.Logger
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

func GetLogger() Plog {
	return logger
}

// --- 设置 logger 参数
func SetFormatter(formatter logrus.Formatter) {
	logger.SetFormatter(formatter)
}

func SetLevel(level logrus.Level) {
	logger.SetLevel(level)
}

func AddHook(hook logrus.Hook) {
	logger.AddHook(hook)
}

func SetOutput(output io.Writer) {
	logger.SetOutput(output)
}

// 是否记录日志调用者的标记(位置信息). 默认false
// 原始的 logger.SetReportCaller() 方法的实现说明：github.com/sirupsen/logrus@v1.8.1/entry.go@getCaller 方法中实现 retrieves the name of the first non-logrus calling function（译：检索第一个非logrus调用函数的名称）
// 		1. 外部直接调用 Debug, Print, Info, Warn, Error, Fatal or Panic 等记录日志方法，记录的起始标记实际是 logger 的方法的位置（检索第一个非logrus调用函数的名称）
// 		2. 开启标记之后，外部调用正确方式：外部调用的地方配合使用 WithXXX().记录日志方法(...)
// 目标：兼容原始的 logger.SetReportCaller() 方法，并且检索第一个非 plog 包调用函数的名称。外部可以直接使用本包中的 Debug, Print, Info, Warn, Error, Fatal or Panic 等记录日志方法，即可记录日志调用者的标记
// 请求参数说明： reportCaller == true 自动添加钩子 `AddHookReportCaller` 实现目标
// 使用注意：
// 	1. 返回的 Plog 结构体当作 gorm.io/gorm 的 writer 时，gorm 会自动打印调用者的标记到 msg 中，无需打开此设置。即使打开记录的也是 gorm.io/gorm 的 logger 调用 github.com/sirupsen/logrus 的 Printf 的位置
func SetReportCaller(reportCaller bool, levels ...logrus.Level) {
	// 开启 ReportCaller
 	logger.SetReportCaller(reportCaller)
	
 	if !reportCaller {
 		return
	}
	addHookReportCaller(levels...)
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

// --- 记录日志方法：按等级记录日志。底层每次都获取一个新的 entry
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warning(args ...interface{}) {
	logger.Warning(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func PanicIfError(err error) {
	if err == nil {
		return
	}
	logger.Panic(err)
}
// --- 按等级记录日志 - end
