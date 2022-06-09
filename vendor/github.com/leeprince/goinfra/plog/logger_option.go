package plog

import (
    jsoniter "github.com/json-iterator/go"
    rotatelogs "github.com/lestrrat-go/file-rotatelogs"
    "github.com/sirupsen/logrus"
    "github.com/leeprince/goinfra/plog/formatters"
    "github.com/leeprince/goinfra/plog/hooks"
    "github.com/leeprince/goinfra/resource/file"
    "io"
    "os"
    "path/filepath"
    "strings"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/5 上午8:40
 * @Desc:
 */

// --- 设置 logger 参数
// jsoniter 更快的json序列化方式
// jsoniterAPI =
// 	jsoniter.ConfigDefault |
// 	jsoniter.ConfigCompatibleWithStandardLibrary() (推荐：100% 兼容json标准库:"encoding/json") |
// 	jsoniter.ConfigFastest
func SetFormatterJsonInter(jsoniterAPI jsoniter.API) {
    formatter := &formatters.JSONFormatter{
        TimestampFormat:  "2006-01-02 15:04:05.000",
        DisableTimestamp: false,
        DataKey:          "data",
        FieldMap: formatters.FieldMap{
            logrus.FieldKeyTime: "logTime",
        },
        CallerPrettyfier: nil,
        PrettyPrint:      false,
        JSON:             jsoniterAPI,
    }
    SetFormatter(formatter)
}

// dirPath 支持多层级目录结构
// isBothStdout：是否支持同时通过 os.Stdout 记录日志
func SetOutputFile(dirPath, filename string, isBothStdout bool) error {
    var writer io.Writer
    var err error
    
    filePath := filepath.Join(dirPath, filename)
    if _, ok := file.CheckFileExist(filePath); ok {
        writer, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    } else {
        // 创建目录
        err = os.MkdirAll(dirPath, os.ModePerm)
        if ok = os.IsNotExist(err); ok {
            return err
        }
        writer, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
    }
    if err != nil {
        return err
    }
    
    if logger.Out == nil {
        SetOutput(writer)
    }
    
    if isBothStdout {
        // 支持 logger.Out 和 *io.file(可用于io) 的记录日志方式
        // MultiWriter同时接受了3种数据类型，分别是io.Writer、*os.File、io.WriteCloser
        SetOutput(io.MultiWriter(logger.Out, writer))
    } else {
        SetOutput(writer)
    }
    
    return nil
}

// 日志切分
// 依赖库：github.com/lestrrat-go/file-rotatelogs
// dirPath 支持多层级目录结构
// isBothStdout：是否支持同时通过 os.Stdout 记录日志
func SetOutputRotateFile(dirPath, filename string, isBothStdout bool, rotationFilePathChangeFunc func(string) (string, error), rotateOptions ...rotatelogs.Option) error {
    var writer io.Writer
    var err error
    
    defaultRotationFilePathChangeFunc := func(filePath string) (string, error) {
        // filePath 修改: /path/to/your.log => /path/to/your.2020-09-27.log
        ext := filepath.Ext(filePath)
        timeFormat := ".%Y-%m-%d"
        if ext == "" {
            filePath = filePath + timeFormat
        } else {
            filePath = strings.TrimSuffix(filePath, ext) + timeFormat + ext
        }
        return filePath, nil
    }
    
    filePath := filepath.Join(dirPath, filename)
    if rotationFilePathChangeFunc == nil {
        rotationFilePathChangeFunc = defaultRotationFilePathChangeFunc
    }
    filePath, err = rotationFilePathChangeFunc(filePath)
    if err != nil {
        return err
    }
    
    // 默认配置：Local时区，按天分割
    if rotateOptions == nil || len(rotateOptions) == 0 {
        rotateOptions = []rotatelogs.Option{
            rotatelogs.WithClock(rotatelogs.Local),
            // 与文件名的格式不冲突。这里是切割日志（旋转日志）的时间间隔。
            rotatelogs.WithRotationTime(time.Hour * 24),
        }
    }
    
    writer, err = rotatelogs.New(
        filePath,
        rotateOptions...,
    )
    if err != nil {
        return err
    }
    
    if isBothStdout {
        // 支持 logger.Out 和 *io.file(可用于io) 的记录日志方式
        // MultiWriter同时接受了3种数据类型，分别是io.Writer、*os.File、io.WriteCloser
        SetOutput(io.MultiWriter(logger.Out, writer))
    } else {
        SetOutput(writer)
    }
    
    return nil
}

// AddHookReportCaller
//  目标：检索第一个非 plog 包调用函数的名称
func addHookReportCaller(levels ...logrus.Level) {
    AddHook(hooks.NewReportCallerHook(levels...))
}

// 记录同一个请求的 logID 标识
const LogIDFiled = "logID"
func WithFiledLogID(logID string) *logrus.Entry {
    return WithField(LogIDFiled, logID)
}

// Sentry 实现日志告警功能
func AddHookSentry(dsn string, levels ...logrus.Level) {
    AddHook(hooks.NewSentryHook(dsn, levels...))
}

// --- 设置 logger 参数 - end
