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
