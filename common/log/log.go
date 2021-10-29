package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"runtime"
)

var _logger *zap.Logger

const (
	FormatJSON = "json"
	FormatText = "text"
)

type logConfig struct {
	LogPath    string // 日志路径, 空将输出控制台
	LogLevel   string // 日志等级
	Compress   bool   // 压缩日志
	MaxSize    int    // log size (M)
	MaxAge     int    // 日志保存时间 (day)
	MaxBackups int    // 日志保存文件数
	Format     string // 日志类型 text or json
}

func getzapLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func newLogWriter(logpath string, maxsize int, compress bool) io.Writer {
	if logpath == "" || logpath == "-" {
		return os.Stdout
	}
	return &lumberjack.Logger{
		Filename: logpath,
		MaxSize:  maxsize,
		Compress: compress,
	}
}

func newZapEncoder() zapcore.EncoderConfig {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	return encoderConfig
}

func newLoggerCore(log *logConfig) zapcore.Core {

	hook := newLogWriter(log.LogPath, log.MaxSize, log.Compress)

	encoderConfig := newZapEncoder()

	atomLevel := zap.NewAtomicLevelAt(getzapLevel(log.LogLevel))

	var encoder zapcore.Encoder
	if log.Format == FormatJSON {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(hook)),
		atomLevel,
	)
	return core
}

func newLoggerOptions() []zap.Option {
	caller := zap.AddCaller()
	callerskip := zap.AddCallerSkip(1)
	development := zap.Development()
	options := []zap.Option{
		caller,
		callerskip,
		development,
	}
	return options
}

func defaultOption() *logConfig {
	return &logConfig{
		LogPath:    "",
		MaxSize:    10, // 单位MB
		Compress:   true,
		MaxAge:     7,
		MaxBackups: 7,
		LogLevel:   "debug",
		Format:     FormatText,
	}
}

type Option func(*logConfig)

// Path 设置日志路径
// 如果为空 则打印
func Path(logpath string) Option {
	return func(config *logConfig) {
		config.LogPath = logpath
	}
}

// Compress 压缩日志
func Compress(compress bool) Option {
	return func(config *logConfig) {
		config.Compress = compress
	}
}

// Level 设置日志等级
func Level(level string) Option {
	return func(config *logConfig) {
		config.LogLevel = level
	}
}

// MaxSize 设置日志
func MaxSize(size int) Option {
	return func(config *logConfig) {
		config.MaxSize = size
	}
}

// MaxAge 日志储存日
func MaxAge(age int) Option {
	return func(config *logConfig) {
		config.MaxAge = age
	}
}

func MaxBackups(backup int) Option {
	return func(config *logConfig) {
		config.MaxBackups = backup
	}
}

// Format 日志文件格式
func Format(format string) Option {
	return func(config *logConfig) {
		if format == FormatJSON {
			config.Format = FormatJSON
		} else {
			config.Format = FormatText
		}
	}
}

// InitLog conf
func InitLog(opts ...Option) error {
	logcfg := defaultOption()
	for _, opt := range opts {
		opt(logcfg)
	}
	core := newLoggerCore(logcfg)

	zapopts := newLoggerOptions()
	_logger = zap.New(core, zapopts...)
	return nil
}

func Debug(msg string, fields ...zap.Field) {
	_logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	_logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	_logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	_logger.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	_logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	_logger.Fatal(msg, fields...)
}

// Recover 异常捕获
// 务必使用 使其处于defer 函数内
func Recover() {
	if value := recover(); value != nil {
		msg := ""
		for i := 1; ; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			msg = fmt.Sprintf("%s %s:%d(0x%x)", msg, file, line, pc)
		}
		_logger.Error(fmt.Sprintf("%v", value), zap.Error(fmt.Errorf("%s", msg)))
	}
}
