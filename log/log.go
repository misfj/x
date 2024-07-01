package log

import "coredx/config"

var (
	DefaultLogger *Logger
)

func Init(conf *config.Logging) error {
	DefaultLogger = NewLogger()

	return DefaultLogger.Init(conf)
}

func Debug(args ...interface{}) {
	if DefaultLogger == nil {
		return
	}
	DefaultLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	if DefaultLogger == nil {
		return
	}
	DefaultLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	if DefaultLogger == nil {
		return
	}
	DefaultLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	if DefaultLogger == nil {
		return
	}
	DefaultLogger.Infof(template, args...)
}

func Error(args ...interface{}) {
	if DefaultLogger == nil {
		return
	}
	DefaultLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	if DefaultLogger == nil {
		return
	}
	DefaultLogger.Errorf(template, args...)
}
