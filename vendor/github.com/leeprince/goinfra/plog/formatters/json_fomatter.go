package formatters

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/5 下午9:36
 * @Desc:
 */

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

var defaultTimestampFormat = time.RFC3339

type fieldKey string

// FieldMap allows customization of the key names for default fields.
type FieldMap map[fieldKey]string

func (f FieldMap) resolve(key fieldKey) string {
	if k, ok := f[key]; ok {
		return k
	}

	return string(key)
}

var _ logrus.Formatter = (*JSONFormatter)(nil)

// JSONFormatter formats logs into parsable json
type JSONFormatter struct {
	// TimestampFormat sets the format used for marshaling timestamps.
	TimestampFormat string

	// DisableTimestamp allows disabling automatic timestamps in output
	DisableTimestamp bool

	// DataKey allows users to put all the log entry parameters into a nested dictionary at a given key.
	DataKey string

	// FieldMap allows users to customize the names of keys for default fields.
	// As an example:
	// formatter := &JSONFormatter{
	//   	FieldMap: FieldMap{
	// 		 FieldKeyTime:  "@timestamp",
	// 		 FieldKeyLevel: "@level",
	// 		 FieldKeyMsg:   "@message",
	// 		 FieldKeyFunc:  "@caller",
	//    },
	// }
	FieldMap FieldMap

	// CallerPrettyfier can be set by the user to modify the content
	// of the function and file keys in the json data when ReportCaller is
	// activated. If any of the returned value is the empty string the
	// corresponding key will be removed from json fields.
	CallerPrettyfier func(*runtime.Frame) (function string, file string)

	// PrettyPrint will indent all json logs
	PrettyPrint bool

	// JSON is an object implemented json.Marshal() and json.Unmarshal()
	JSON jsoniter.API
}

// Format renders a single log entry
// 参考：github.com/sirupsen/logrus@v1.8.1/json_formatter.go@Format
func (f *JSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+4)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	if f.DataKey != "" {
		newData := make(logrus.Fields, 4)
		newData[f.DataKey] = data
		data = newData
	}

	prefixFieldClashes(data, f.FieldMap, entry.HasCaller())

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	// 无法从entry中取出私有成员error
	// if entry.err != "" {
	// 	data[f.FieldMap.resolve(FieldKeyLogrusError)] = entry.err
	// }
	if !f.DisableTimestamp {
		data[f.FieldMap.resolve(logrus.FieldKeyTime)] = entry.Time.Format(timestampFormat)
	}
	data[f.FieldMap.resolve(logrus.FieldKeyMsg)] = entry.Message
	data[f.FieldMap.resolve(logrus.FieldKeyLevel)] = entry.Level.String()
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		if f.CallerPrettyfier != nil {
			funcVal, fileVal = f.CallerPrettyfier(entry.Caller)
		}
		if funcVal != "" {
			data[f.FieldMap.resolve(logrus.FieldKeyFunc)] = funcVal
		}
		if fileVal != "" {
			data[f.FieldMap.resolve(logrus.FieldKeyFile)] = fileVal
		}
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 自定义 Format 部分
	// 根据JSON成员变量是否为空，来决定使用encoding/json还是jsoniter
	if f.JSON != nil {
		encoder := f.JSON.NewEncoder(b)
		if f.PrettyPrint {
			encoder.SetIndent("", "  ")
		}
		if err := encoder.Encode(data); err != nil {
			return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
		}
	} else {
		encoder := json.NewEncoder(b)
		if f.PrettyPrint {
			encoder.SetIndent("", "  ")
		}
		if err := encoder.Encode(data); err != nil {
			return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
		}
	}

	return b.Bytes(), nil
}

// This is to not silently overwrite `time`, `msg`, `func` and `level` fields when
// dumping it. If this code wasn't there doing:
//
//  logrus.WithField("level", 1).Info("hello")
//
// Would just silently drop the user provided level. Instead with this code
// it'll logged as:
//
//  {"level": "info", "fields.level": 1, "msg": "hello", "time": "..."}
//
// It's not exported because it's still using Data in an opinionated way. It's to
// avoid code duplication between the two default formatters.
func prefixFieldClashes(data logrus.Fields, fieldMap FieldMap, reportCaller bool) {
	timeKey := fieldMap.resolve(logrus.FieldKeyTime)
	if t, ok := data[timeKey]; ok {
		data["fields."+timeKey] = t
		delete(data, timeKey)
	}

	msgKey := fieldMap.resolve(logrus.FieldKeyMsg)
	if m, ok := data[msgKey]; ok {
		data["fields."+msgKey] = m
		delete(data, msgKey)
	}

	levelKey := fieldMap.resolve(logrus.FieldKeyLevel)
	if l, ok := data[levelKey]; ok {
		data["fields."+levelKey] = l
		delete(data, levelKey)
	}

	logrusErrKey := fieldMap.resolve(logrus.FieldKeyLogrusError)
	if l, ok := data[logrusErrKey]; ok {
		data["fields."+logrusErrKey] = l
		delete(data, logrusErrKey)
	}

	// If reportCaller is not set, 'func' will not conflict.
	if reportCaller {
		funcKey := fieldMap.resolve(logrus.FieldKeyFunc)
		if l, ok := data[funcKey]; ok {
			data["fields."+funcKey] = l
		}
		fileKey := fieldMap.resolve(logrus.FieldKeyFile)
		if l, ok := data[fileKey]; ok {
			data["fields."+fileKey] = l
		}
	}
}
