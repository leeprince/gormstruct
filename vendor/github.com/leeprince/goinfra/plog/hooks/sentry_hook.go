package hooks

import (
    "errors"
    "github.com/sirupsen/logrus"
    "github.com/leeprince/goinfra/alert/sentry_util"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/16 上午11:31
 * @Desc:
 */


type SentryHook struct {
    levels []logrus.Level
}

func NewSentryHook(dsn string, levels ...logrus.Level) *SentryHook {
    if levels == nil {
        levels = logrus.AllLevels
    }
    err := sentry_util.Init(dsn)
    if err != nil {
        return nil
    }
    return &SentryHook{
        levels: levels,
    }
}

func (h *SentryHook) Levels() []logrus.Level {
    return h.levels
}

// 本方法可以优化为直接利用 frames := runtime.CallersFrames(pcs[:depth]) 去固定 plog 包应跳过的追踪栈
func (h *SentryHook) Fire(entry *logrus.Entry) error {
    sentry_util.CaptureException(errors.New(entry.Message), time.Millisecond * 200)
    return nil
}