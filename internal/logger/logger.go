package logger

import (
    "github.com/sirupsen/logrus"
    "os"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/1 下午10:45
 * @Desc:
 */

func DefaultLogger() *logrus.Logger {
    return &logrus.Logger{
        Out: os.Stdout,
        Formatter: &logrus.TextFormatter{
            TimestampFormat: "2006-01-02 15:04:05.000",
            FullTimestamp:   true,
        },
        Level: logrus.InfoLevel,
    }
}