package mlog

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/mattn/go-colorable"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// Fields 日志数据字段
type Fields = logrus.Fields

// Logger 日志句柄
type Logger struct {
	entry *logrus.Entry
}

// 日志钩子
type logHook struct{}

// 是否可以使用isPanic函数挂掉程序
var isPanic = true

// 是否调试模式
var isDebug = false

// 系统文件夹名称
var syslogname = "systemlog"

// 用户日志文件夹名称
var playerlogname = "playerlog"

func init() {
	flag.BoolVar(&isDebug, "debug", false, "开启调试模式")
}

// PanicToError 把panic错误转换成error错误
func PanicToError() {
	isPanic = false
}

// 获得当前调用椎栈
func getStack() string {
	const size = 64 << 10
	buf := make([]byte, size)
	buf = buf[:runtime.Stack(buf, false)]
	return string(buf)
}

func (h logHook) Fire(e *logrus.Entry) error {
	for i := 1; i < 10; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			continue
		}
		if strings.Contains(file, "battlePlatform/wjrgit.qianz.com/common/vendor/_mlog/mlog.go") {
			_, file, line, ok = runtime.Caller(i + 1)
			if !ok {
				break
			}

			tmpfile := file
			fileslice := strings.Split(file, "/")
			if len(fileslice) >= 4 {
				flen := len(fileslice)
				tmpfile = path.Join(fileslice[flen-4], fileslice[flen-3], fileslice[flen-2], fileslice[flen-1])
			}

			e.Data["File"] = tmpfile
			e.Data["Line"] = line
			break
		}
	}

	return nil
}

func (h logHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// NewPlayerLogger 创建玩家日志句柄
func NewPlayerLogger(playerid int64, fields Fields) *Logger {
	log := logrus.New()
	if isDebug {
		log.SetLevel(logrus.DebugLevel)
	}

	log.Formatter = &logrus.TextFormatter{
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		QuoteEmptyFields: true,
	}
	log.Out = colorable.NewColorableStdout()
	log.AddHook(logHook{})

	// 玩家
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	playerpath := path.Join(dir, playerlogname)
	result, err := pathExists(playerpath)
	if err != nil {
		log.Fatal(err)
	}
	if !result {
		err = os.MkdirAll(playerpath, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 7天
	maxAge := time.Hour * 24 * 7
	baseLogPath := path.Join(dir, "playerlog", fmt.Sprintf("%d", playerid))
	now := time.Now().Unix()
	nowtime := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", nowtime+" 23:59:59", time.Local)
	// 第二天凌晨
	tomts := t.Unix() + 1
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d.log",
		rotatelogs.WithLinkName(""),   // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge), // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(tomts-now)),
	)
	if err != nil {
		log.Fatal(err)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		QuoteEmptyFields: true,
	})
	log.AddHook(lfHook)

	return &Logger{
		entry: log.WithFields(fields),
	}
}

// NewSysLogger 创建日志句柄
func NewSysLogger(fields Fields) *Logger {
	log := logrus.New()
	if isDebug {
		log.SetLevel(logrus.DebugLevel)
	}

	log.Formatter = &logrus.TextFormatter{
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		QuoteEmptyFields: true,
	}
	log.Out = colorable.NewColorableStdout()
	log.AddHook(logHook{})

	return &Logger{
		entry: log.WithFields(fields),
	}
}

// NewLogger TODO
func NewLogger(fields Fields) *Logger {
	return NewSysLogger(fields)
}

// Info Info等级日志
func (l Logger) Info(args ...interface{}) {
	l.entry.Info(args...)
}

// Infof Info等级日志
func (l Logger) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

// Debug Debug等级日志
func (l Logger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

// Debugf Debug等级日志
func (l Logger) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

// Warn Warn等级日志
func (l Logger) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

// Warnf Warn等级日志
func (l Logger) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

// Error Error等级日志
func (l Logger) Error(args ...interface{}) {
	l.entry.Error(args...)
}

// Errorf Error等级日志
func (l Logger) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

// Fatalf error can't recover
func (l Logger) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

// Panic 灾难级别日志，只有加载配置时可以使用
func (l Logger) Panic(args ...interface{}) {
	if isPanic {
		l.entry.Panic(args...)
	} else {
		entry := l.entry.WithFields(logrus.Fields{"Panic": true})
		entry.Error(args...)
		l.entry.Error(getStack())
	}
}

// Panicf 灾难级别日志，只有加载配置时可以使用
func (l Logger) Panicf(format string, args ...interface{}) {
	if isPanic {
		l.entry.Panicf(format, args...)
	} else {
		entry := l.entry.WithFields(logrus.Fields{"Panic": true})
		entry.Errorf(format, args...)
		l.entry.Error(getStack())
	}
}

// WithFields 添加新的数据字段
func (l Logger) WithFields(fields Fields) *Logger {
	return &Logger{
		entry: l.entry.WithFields(fields),
	}
}

// WithField Add a single field to the Entry.
func (l Logger) WithField(key string, value interface{}) *Logger {
	return l.WithFields(Fields{key: value})
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetDebugTag debug mode
func GetDebugTag() bool {
	return isDebug
}
