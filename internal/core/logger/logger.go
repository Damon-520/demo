package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var _ log.Logger = (*Logger)(nil)
var _ logrus.Formatter = (*logFormatter)(nil)

type Option func(logger *Logger)

type Config struct {
	Path         string
	Level        string
	Rotationtime time.Duration
	Maxage       time.Duration
	OpenStat     bool
}

type Logger struct {
	logger *logrus.Logger
}

func NewLogger(config Config, ops ...Option) *Logger {
	l := &Logger{}
	l.logger = logrus.New()
	l.logger.SetFormatter(&logFormatter{})
	for _, o := range ops {
		o(l)
	}

	if level, err := logrus.ParseLevel(config.Level); err != nil {
		panic(err.Error())
	} else {
		l.logger.SetLevel(level)
	}

	// 设置日志滚动更新
	writer, err := rotatelogs.New(
		config.Path,
		rotatelogs.WithRotationTime(config.Rotationtime),
		rotatelogs.WithMaxAge(config.Maxage),
	)
	if err != nil {
		panic(err.Error())
	}

	// 开启统计
	if config.OpenStat {
		l.logStat()
	}

	// 设置kratos底层 msg 字段名
	log.DefaultMessageKey = "x_msg"
	return l.setOutput(writer)
}

func (lf *Logger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		lf.errorf("params error")
		return nil
	}

	buf := make(map[string]interface{})
	for i := 0; i < len(keyvals); i += 2 {
		if logKey := keyvals[i].(string); logKey != "" {
			buf[logKey] = keyvals[i+1]
		}
	}

	switch level {
	case log.LevelFatal:
		lf.fatalM(buf)
	case log.LevelWarn, log.LevelError:
		lf.errorM(buf)
	case log.LevelDebug:
		lf.debugM(buf)
	default:
		lf.infoM(buf)
	}

	return nil
}

type logFormatter struct {
}

func (lf *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+6)
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

	data["x_level"] = "log." + entry.Level.String()
	data["x_timestamp"] = time.Now().Unix()
	data["x_date"] = time.Now().In(time.Local).Format("2006-01-02 15:04:05")

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	encoder := json.NewEncoder(b)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v\n", err)
	}

	return b.Bytes(), nil
}

func (l *Logger) setOutput(writer io.Writer) *Logger {
	l.logger.SetOutput(writer)
	return l
}

func (l *Logger) errorf(format string, args ...interface{}) {
	l.errorM(map[string]interface{}{
		"x_msg": fmt.Sprintf(format, args...),
	})
}

func (l *Logger) debugM(msgs map[string]interface{}) {
	l.logger.WithFields(msgs).Debug()
}

func (l *Logger) infoM(msgs map[string]interface{}) {
	l.logger.WithFields(msgs).Info()
}

func (l *Logger) errorM(msgs map[string]interface{}) {
	msgs["x_extra"] = map[string]interface{}{"stack": callFrames(15)}
	l.logger.WithFields(msgs).Error()
}

func (l *Logger) fatalM(msgs map[string]interface{}) {
	msgs["x_extra"] = map[string]interface{}{"stack": callFrames(15)}
	l.logger.WithFields(msgs).Fatal()
}

var packageName string
var once sync.Once

func callFrames(maxDept int) []string {
	var stacks []string
	pcs := make([]uintptr, maxDept)
	depth := runtime.Callers(3, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	// cache this package's fully-qualified name
	once.Do(func() {
		packageName = getPackageName(runtime.FuncForPC(pcs[0]).Name())
	})

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		// If the caller isn't part of this package, we're done
		if pkg != packageName {
			stacks = append(stacks, fmt.Sprintf("%s:%d", f.File, f.Line))
		}
	}

	return stacks
}

func getPackageName(absPath string) string {
	for {
		lastPeriod := strings.LastIndex(absPath, ".")
		lastSlash := strings.LastIndex(absPath, "/")
		if lastPeriod > lastSlash {
			absPath = absPath[:lastPeriod]
		} else {
			break
		}
	}

	return absPath
}

// log 服务统计功能
func (lf *Logger) logStat() {

}
