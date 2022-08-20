package sentry

import (
    "github.com/getsentry/sentry-go"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/17 上午2:09
 * @Desc:
 */

func Init(dsn string) error {
    if dsn == "" {
        return nil
    }
    err := sentry.Init(sentry.ClientOptions{
        Dsn:              dsn,
        Debug:            false,
        AttachStacktrace: false,
        SampleRate:       0,
        TracesSampleRate: 0,
        TracesSampler:    nil,
        IgnoreErrors:     nil,
        BeforeSend:       nil,
        BeforeBreadcrumb: nil,
        Integrations:     nil,
        DebugWriter:      nil,
        Transport:        nil,
        ServerName:       "",
        Release:          "",
        Dist:             "",
        Environment:      "",
        MaxBreadcrumbs:   0,
        HTTPClient:       nil,
        HTTPTransport:    nil,
        HTTPProxy:        "",
        HTTPSProxy:       "",
        CaCerts:          nil,
    })
    return err
}

func CaptureMessage(msg string, flushTime time.Duration) *sentry.EventID {
    defer func() {
        if flushTime == 0 {
            flushTime = SentryDefaultFlushTime
        }
        sentry.Flush(flushTime)
    }()
    
    return sentry.CaptureMessage(msg)
}

func CaptureException(exception error, flushTime time.Duration) *sentry.EventID {
    defer func() {
        if flushTime == 0 {
            flushTime = SentryDefaultFlushTime
        }
        sentry.Flush(flushTime)
    }()
    
    return sentry.CaptureException(exception)
}