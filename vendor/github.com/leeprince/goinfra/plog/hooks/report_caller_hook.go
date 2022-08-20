package hooks

import (
    "github.com/sirupsen/logrus"
    "runtime"
    "strings"
    "sync"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/6 下午11:52
 * @Desc:
 */

var (
    // qualified package name, cached at first use
    myLoggerPackage string
    
    // Positions in the call stack when tracing to report the calling method
    minimumCallerDepth int
    
    // Used for caller information initialisation
    callerInitOnce sync.Once
)

const (
    myLoggerPackageFuncMark = "github.com/leeprince/goinfra/plog."
)

const (
    maximumCallerDepth int = 25
    
    // 已知 plog 包底层应跳过的调用栈
    knownPlogFrames  int = 6 // knownLogrusFrames int = 4
)

type ReportCallerHook struct {
    // 日志等级切片
    levels                    []logrus.Level
    
    // 支持外部基于该包 plog 封装成自己包时，追踪并记录调用日志者的标记
    customerTempLoggerPackage string
}

var reportCallerHook *ReportCallerHook

func NewReportCallerHook(levels ...logrus.Level) *ReportCallerHook {
    if levels == nil {
        levels = logrus.AllLevels
    }
    reportCallerHook = &ReportCallerHook{
        levels: levels,
    }
    return reportCallerHook
}

func (h *ReportCallerHook) Levels() []logrus.Level {
    return h.levels
}

// 支持外部基于该包封装成自己包时，追踪并记录调用日志者的标记
func (h *ReportCallerHook) SetCustomerTempLoggerPackage(loggerPagekage string) {
    reportCallerHook.customerTempLoggerPackage = loggerPagekage
}

// 参考：github.com/sirupsen/logrus@v1.8.1/entry.go@getCaller 方法
// 执行钩子的 Fire 方法前已经通过 github.com/sirupsen/logrus@v1.8.1/entry.go 的 getCaller() 对 entry.Caller 进行赋值，所以本方法实现的核心就是对 entry.Caller 进行重新赋值。
func (h *ReportCallerHook) Fire(entry *logrus.Entry) error {
    // cache this package's fully-qualified name
    callerInitOnce.Do(func() {
        pcs := make([]uintptr, maximumCallerDepth)
        _ = runtime.Callers(0, pcs)
        
        // dynamic get the package name and the minimum caller depth
        for i := 0; i < maximumCallerDepth; i++ {
            funcName := runtime.FuncForPC(pcs[i]).Name()
            if strings.Contains(funcName, myLoggerPackageFuncMark) {
                myLoggerPackage = h.getPackageName(funcName)
                break
            }
        }
        
        minimumCallerDepth = knownPlogFrames
    })
    
    loggerPackageName := h.loggerPackageName()
    if loggerPackageName == "" {
        return nil
    }
    defer func() {
        reportCallerHook.customerTempLoggerPackage = ""
    }()
    
    // ReportCallerHook@Fire 方法中提升追踪并记录调用日志者的函数信息（文件、方法名、行号）
    // > 方案一：利用 `pc, file, line, ok := runtime.Caller(i)` 直接跳过底层调用栈获取真正的调用者的函数信息（文件名、方法、行号）。优点：性能好，缺点：不方便使用，需使用者调试确定i值
    // > 方案二（最终方案）：哈希包名并存储并记录应跳过底层调用栈的i值（文件、方法名、行号）到内存中，下次检查并直接使用结合`方案一`去记录调用者的函数信息
    
    // Restrict the lookback frames to avoid runaway lookups
    pcs := make([]uintptr, maximumCallerDepth)
    depth := runtime.Callers(minimumCallerDepth, pcs)
    frames := runtime.CallersFrames(pcs[:depth])
    
    isMyLoggerPackage := false // 是否匹配到定义的包名的方法
    for f, again := frames.Next(); again; f, again = frames.Next() {
        pkg := h.getPackageName(f.Function)
        
        // 匹配到定义的包名的方法之后不等于指定包名的即为实际调用方的位置
        if isMyLoggerPackage && pkg != loggerPackageName {
            entry.Caller = &f
            break
        }
        
        // If the caller isn't part of this package, we're done
        if pkg == loggerPackageName {
            isMyLoggerPackage = true
        }
    }
    
    return nil
}

// 参考：github.com/sirupsen/logrus@v1.8.1/entry.go@getPackageName
// getPackageName reduces a fully qualified function name to the package name
// There really ought to be to be a better way...
func (h *ReportCallerHook) getPackageName(f string) string {
    for {
        lastPeriod := strings.LastIndex(f, ".")
        lastSlash := strings.LastIndex(f, "/")
        if lastPeriod > lastSlash {
            f = f[:lastPeriod]
        } else {
            break
        }
    }
    
    return f
}

// 要跳过的日志包名
func (h *ReportCallerHook) loggerPackageName() string {
    loggerPackageName := myLoggerPackage
    if reportCallerHook.customerTempLoggerPackage != "" {
        return reportCallerHook.customerTempLoggerPackage
    }
    return loggerPackageName
}
