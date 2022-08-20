package plog

import "github.com/sirupsen/logrus"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/13 上午12:02
 * @Desc:
 */

// Level type
type Level uint32

// A constant exposing all logging levels
var AllLevels = []Level{
    PanicLevel,
    FatalLevel,
    ErrorLevel,
    WarnLevel,
    InfoLevel,
    DebugLevel,
    TraceLevel,
}

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
    // PanicLevel level, highest level of severity. Logs and then calls panic with the
    // message passed to Debug, Info, ...
    PanicLevel Level = iota
    // FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
    // logging level is set to Panic.
    FatalLevel
    // ErrorLevel level. Logs. Used for errors that should definitely be noted.
    // Commonly used for hooks to send errors to an error tracking service.
    ErrorLevel
    // WarnLevel level. Non-critical entries that deserve eyes.
    WarnLevel
    // InfoLevel level. General operational entries about what's going on inside the
    // application.
    InfoLevel
    // DebugLevel level. Usually only enabled when debugging. Very verbose logging.
    DebugLevel
    // TraceLevel level. Designates finer-grained informational events than the Debug.
    TraceLevel
)

func PLevel(l Level) logrus.Level {
    switch l {
    case PanicLevel:
        return logrus.PanicLevel
    case FatalLevel:
        return logrus.FatalLevel
    case ErrorLevel:
        return logrus.ErrorLevel
    case WarnLevel:
        return logrus.WarnLevel
    case InfoLevel:
        return logrus.InfoLevel
    case DebugLevel:
        return logrus.DebugLevel
    case TraceLevel:
        return logrus.TraceLevel
    default:
        return logrus.TraceLevel
    }
}

func PLevels(levels []Level) []logrus.Level {
    var logrusLevels []logrus.Level
    for _, l := range levels {
        switch l {
        case PanicLevel:
            logrusLevels = append(logrusLevels, logrus.PanicLevel)
        case FatalLevel:
            logrusLevels = append(logrusLevels, logrus.FatalLevel)
        case ErrorLevel:
            logrusLevels = append(logrusLevels, logrus.ErrorLevel)
        case WarnLevel:
            logrusLevels = append(logrusLevels, logrus.WarnLevel)
        case InfoLevel:
            logrusLevels = append(logrusLevels, logrus.InfoLevel)
        case DebugLevel:
            logrusLevels = append(logrusLevels, logrus.DebugLevel)
        case TraceLevel:
            logrusLevels = append(logrusLevels, logrus.TraceLevel)
        default:
            logrusLevels = append(logrusLevels, logrus.TraceLevel)
        }
    }
    return logrusLevels
}
