package hooks

import (
    "errors"
    "github.com/sirupsen/logrus"
    "github.com/leeprince/goinfra/alert/sentry"
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
    err := sentry.Init(dsn)
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

func (h *SentryHook) Fire(entry *logrus.Entry) error {
    sentry.CaptureException(errors.New(entry.Message), time.Millisecond * 200)
    return nil
}