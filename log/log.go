package log

import "github.com/astaxie/beego/logs"

const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

type Logger struct {
	*logs.BeeLogger
}

func NewLogger(channelLen int64, adapterName string, config string, logLevel int) *Logger {
	logger := logs.NewLogger(channelLen)
	logger.SetLogger(adapterName, config)
	logger.SetLevel(logLevel)
	logger.EnableFuncCallDepth(true)
	logger.SetLogFuncCallDepth(3)
	return &Logger{logger}
}

func (logger *Logger) Printf(format string, v ...interface{}) {
	logger.Trace(format, v...)
}

var l *Logger

func InitLogger(channelLen int64, adapterName string, config string, logLevel int) {
	l = NewLogger(channelLen, adapterName, config, logLevel)
}

func Criticalf(format string, v ...interface{}) {
	l.Critical(format, v...)
}

func Errorf(format string, v ...interface{}) {
	l.Error(format, v...)
}

func Warnf(format string, v ...interface{}) {
	l.Warn(format, v...)
}

func Infof(format string, v ...interface{}) {
	l.Info(format, v...)
}

func Tracef(format string, v ...interface{}) {
	l.Trace(format, v...)
}

func Debugf(format string, v ...interface{}) {
	l.Debug(format, v...)
}
