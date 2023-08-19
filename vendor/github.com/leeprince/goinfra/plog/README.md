# 日志


## 使用方式


### 默认
默认直接调用

### 自定义配置
**支持自定义的方法：logger_option.go**

- 更快的json序列化方式 `SetFormatterJsonIter`
- 保存日志记录到文件 `SetOutputFile`
- 添加钩子实现 logger 追踪并记录调用日志者的标记, 支持其他包基于该包封装后最终调用者的函数信息（文件、方法名、行号） `addHookReportCaller`
    > 原始的 SetReportCaller，需特殊使用方可记录实际调用日志者的标记。具体说明：logger.go@SetReportCaller
- 日志切割（日志旋转） `SetOutputRotateFile`
- 记录同一个请求的 logID 标识 `WithFiledLogID`
- 添加 `sentry` 告警
- 支持外部基于该包封装成自己包时，追踪并记录调用日志者的标记。Plog 结构体的 `customerTempLoggerPackage` 字段

## 当前库其他包基于该包实现的其他日志的接口
- [x] jaeger: Error、Infof
    > jaegerLogger 基于 github.com/leeprince/goinfra/plog 实现 github.com/uber/jaeger-client-go@logger.go 的 Logger 接口

## 优化
- runtime.Callers 指定底层应跳过的调用栈 minimumCallerDepth
    > minimumCallerDepth = knownLogrusFrames=4 替换为 knownPlogFrames=7
    >   使用 plog.WithXXX 与 plog.Debug/Print/Info/Warn/Error/Fatal/Panic 等记录日志方法会导致底层调用栈的层数不同
