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
    knownLogrusFrames  int = 4
)

type ReportCallerHook struct {
    levels []logrus.Level
}

func NewReportCallerHook(levels ...logrus.Level) *ReportCallerHook {
    if levels == nil {
        levels = logrus.AllLevels
    }
    return &ReportCallerHook{
        levels: levels,
    }
}

func (h *ReportCallerHook) Levels() []logrus.Level {
    return h.levels
}

// 参考：github.com/sirupsen/logrus@v1.8.1/entry.go@getCaller 方法
// 本方法可以优化为直接利用 frames := runtime.CallersFrames(pcs[:depth]) 去固定 plog 包应跳过的追踪栈
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
        
        minimumCallerDepth = knownLogrusFrames
    })
    
    // Restrict the lookback frames to avoid runaway lookups
    pcs := make([]uintptr, maximumCallerDepth)
    depth := runtime.Callers(minimumCallerDepth, pcs)
    frames := runtime.CallersFrames(pcs[:depth])
    
    isMyLoggerPackage := false // 是否匹配到定义的包名的方法 - prince@comm 2022/3/15 下午9:27
    for f, again := frames.Next(); again; f, again = frames.Next() {
        pkg := h.getPackageName(f.Function)
        
        // 匹配到定义的包名的方法之后不等于定于包名的即为实际调用方的位置 - prince@comm 2022/3/15 下午9:37
        if isMyLoggerPackage && pkg != myLoggerPackage {
			entry.Caller = &f
			break
        }
        
        // If the caller isn't part of this package, we're done
        if pkg == myLoggerPackage {
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


