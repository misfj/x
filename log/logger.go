package log

import (
	"coredx/config"
	"os"
	"strings"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

const timeLayout = "2006-01-02 15:04:05.999999"

type Config struct {
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	DebugMode  bool
}

type Logger struct {
	config *Config
	file   *lumberjack.Logger
	sugar  *zap.SugaredLogger
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) fileWriter(conf *config.Logging) *lumberjack.Logger {
	return &lumberjack.Logger{
		// 日志名称
		Filename: conf.File,
		// 日志大小限制，单位MB
		MaxSize: int(conf.MaxSize),
		// 历史日志文件保留天数
		MaxAge: int(conf.MaxAge),
		// 最大保留历史日志数量
		MaxBackups: int(conf.MaxBackup),
		// 本地时区
		LocalTime: true,
		// 历史日志文件压缩标识
		Compress: false,
	}
}

func (l *Logger) Init(conf *config.Logging) error {
	atomicLevel := zap.NewAtomicLevel()

	if strings.EqualFold("debug", conf.Level) {
		atomicLevel.SetLevel(zapcore.DebugLevel)
	} else {
		atomicLevel.SetLevel(zapcore.InfoLevel)
	}

	zapconf := zap.NewDevelopmentConfig()
	zapconf.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	zapconf.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	zapconf.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(timeLayout))
	}

	enc := zapcore.NewConsoleEncoder(zapconf.EncoderConfig)

	var zapCore zapcore.Core
	var zaplogger *zap.Logger

	if strings.EqualFold("debug", conf.Level) {
		zapCore = zapcore.NewCore(&encoderWrapper{Encoder: enc}, zapcore.AddSync(os.Stdout), atomicLevel)
		zaplogger = zap.New(zapcore.NewTee(zapCore))
	} else {
		zapCore = zapcore.NewCore(&encoderWrapper{Encoder: enc}, zapcore.AddSync(l.fileWriter(conf)), atomicLevel)
		zaplogger = zap.New(zapcore.NewTee(zapCore))
	}

	zaplogger = zaplogger.WithOptions(zap.AddCaller(), zap.AddCallerSkip(2))
	l.sugar = zaplogger.Sugar()

	return nil
}

func (l *Logger) Flush() error {
	err := l.sugar.Sync()
	if l.file != nil {
		err = l.file.Close()
	}
	return err
}

func (l *Logger) Debug(args ...interface{}) {
	l.sugar.Debug(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.sugar.Debugf(template, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.sugar.Info(args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.sugar.Infof(template, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.sugar.Error(args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.sugar.Errorf(template, args...)
}

type encoderWrapper struct {
	zapcore.Encoder
}

func (ew *encoderWrapper) EncodeEntry(entry zapcore.Entry, field []zapcore.Field) (*buffer.Buffer, error) {
	buf, err := ew.Encoder.EncodeEntry(entry, field)
	if err != nil {
		return nil, err
	}

	bs := buf.String()
	bs = strings.ReplaceAll(bs, string('\t'), " ")
	buf.Reset()
	buf.Write([]byte(bs))
	return buf, nil
}
