package plog

import (
    jsoniter "github.com/json-iterator/go"
    "github.com/leeprince/goinfra/plog/formatters"
    "github.com/leeprince/goinfra/plog/hooks"
    "github.com/leeprince/goinfra/resource/file"
    rotatelogs "github.com/lestrrat-go/file-rotatelogs"
    "github.com/sirupsen/logrus"
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
    
    logger.outFileInfo = FileInfo{
        dirPath:  dirPath,
        filename: filename,
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
        
        // 支持的格式示例：%Y%m%d%H%M%S
        timeFormat := ".%Y-%m-%d"
        // timeFormat := ".%Y-%m-%d %H%M"
        // timeFormat := ".%Y-%m-%d %H%M%S"
        
        if ext == "" {
            filePath = filePath + timeFormat
        } else {
            filePath = strings.TrimSuffix(filePath, ext) + timeFormat + ext
        }
        return filePath, nil
    }
    
    originFilePath := filepath.Join(dirPath, filename)
    if rotationFilePathChangeFunc == nil {
        rotationFilePathChangeFunc = defaultRotationFilePathChangeFunc
    }
    filePath, cErr := rotationFilePathChangeFunc(originFilePath)
    if cErr != nil {
        err = cErr
        return err
    }
    
    // 默认配置
    if rotateOptions == nil || len(rotateOptions) == 0 {
        rotateOptions = []rotatelogs.Option{
            // 当前时区
            rotatelogs.WithClock(rotatelogs.Local),
            
            // 按时间的切割方式。默认按天
            // 与文件名的格式不冲突。这里是切割日志（旋转日志）的时间间隔。
            rotatelogs.WithRotationTime(time.Hour * 24),
            // rotatelogs.WithRotationTime(time.Minute * 1),
            // rotatelogs.WithRotationTime(time.Second * 2),
            
            // 设置指定文件名（支持多级），会自动软链接到当前使用的文件名（切割后当前使用的文件名）。
            //  默认当前路径的文件
            //   application.2022-05-16 172255.log
            //   application.2022-05-16 172735.log
            //   application.log
            //  注意：有些日志采集系统无法读取软链接的文件。如：filebeat 默认无法读取软链接的文件，需单独配置
            rotatelogs.WithLinkName(originFilePath),
            
            // 最大保存的文件数。自动删除过期的历史文件
            rotatelogs.WithRotationCount(7),
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
//  目标：检索第一个非 plog 包的调用函数信息（文件名、方法、行号）
func addHookReportCaller(levels ...logrus.Level) {
    AddHook(hooks.NewReportCallerHook(levels...))
}

// 记录同一个请求的 logID 标识
const LogIDFiled = "logID"
func WithFiledLogID(logID string) *logrus.Entry {
    return WithField(LogIDFiled, logID)
}

// Sentry 实现日志告警功能
func AddHookSentry(dsn string, levels ...Level) {
    AddHook(hooks.NewSentryHook(dsn, PLevels(levels)...))
}

// 支持外部基于该包封装成自己包时，追踪并记录调用日志者的标记
func SetCustomerTempLoggerPackage(loggerPagekage string, levels ...Level) {
	// 不记录日志调用者的标记，则不设置
	if !logger.ReportCaller {
		return
	}
	hooks.NewReportCallerHook(PLevels(levels)...).SetCustomerTempLoggerPackage(loggerPagekage)
}
// --- 设置 logger 参数 - end



