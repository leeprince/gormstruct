package plog

import (
    "github.com/sirupsen/logrus"
)

func NewDefaultLogger() {
    l := logrus.New()
    /*l.Formatter = &logrus.TextFormatter{
        TimestampFormat: "2006-01-02 15:04:05.000000",
        FullTimestamp:   true,
    }*/
    l.Formatter = &logrus.JSONFormatter{
        TimestampFormat:   "2006-01-02 15:04:05.000000",
        DisableTimestamp:  false,
        DisableHTMLEscape: true,
        DataKey:           "context", // 允许将用户通过WithXXX设置的所有参数，放入该字段中，并且支持嵌套。不设置则平铺所有参数
        FieldMap: logrus.FieldMap{
            logrus.FieldKeyTime: "time",
        },
        CallerPrettyfier: nil,
        PrettyPrint:      false,
    }
    l.Level = logrus.TraceLevel
    
    SetLogger(l)
}

// --- 记录日志方法：按等级记录日志。底层每次都获取一个新的 entry
func Trace(args ...interface{}) {
	logger.Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	logger.Tracef(format, args...)
}

func Traceln(args ...interface{}) {
	logger.Traceln(args...)
}

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

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func PanicIfError(err error) {
	if err == nil {
		return
	}
	logger.Panic(err)
}
// --- 按等级记录日志 - end
