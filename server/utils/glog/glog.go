package glog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync"
	"time"
)

var _logger *Logger

func InitLogger(filepath string, level zapcore.Level, format string) *Logger {
	_logger, _ = NewLogger(filepath, level, format)
	return _logger
}

// 负责设置 encoding 的日志格式
func getEncoder(format string) zapcore.Encoder {
	// 获取一个指定的的EncoderConfig，进行自定义
	//encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "log",
		TimeKey:        "time",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if format == "json" {
		//json格式
		return zapcore.NewJSONEncoder(encodeConfig)
	}
	//控制台格式
	return zapcore.NewConsoleEncoder(encodeConfig)
}

// 负责日志切割配置
func getLumberJackLogger(filepath string) *lumberjack.Logger {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath, // 文件位置
		MaxSize:    1,        // 进行切割之前,日志文件的最大大小(MB为单位)
		MaxAge:     30,       // 保留旧文件的最大天数
		MaxBackups: 5,        // 保留旧文件的最大个数
		Compress:   false,    // 是否压缩/归档旧文件
	}
	return lumberJackLogger
}

// NewLogger 初始化Logger
func NewLogger(filepath string, level zapcore.Level, format string) (*Logger, error) {
	lumberJackLogger := getLumberJackLogger(filepath)
	// 获取日志写入位置, 同时输出控制台和文件
	writeSyncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	// 获取日志编码格式
	encoder := getEncoder(format)

	// 创建一个将日志写入 WriteSyncer 的核心。
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	//初始化内部类
	out := new(Logger)

	out.rlog = lumberJackLogger
	out.log = logger
	out.sugar = logger.Sugar()

	out.path = filepath
	out.level = level
	out.rotateMu = &sync.Mutex{}
	out.rolling = true
	out.lastRotate = time.Now()

	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(logger)

	return out, nil
}

func Sync() {
	_logger.log.Sync()
}

func GetDefault() *Logger {
	return _logger
}

// active

func Error(v ...interface{}) {
	_logger.Error(v...)
}

func Warn(v ...interface{}) {
	_logger.Warn(v...)
}

func Info(v ...interface{}) {
	_logger.Info(v...)
}

func Debug(v ...interface{}) {
	_logger.Debug(v...)
}

func Errorw(format string, v ...interface{}) {
	_logger.Errorw(format, v...)
}

func Warnw(format string, v ...interface{}) {
	_logger.Warnw(format, v...)
}

func Infow(format string, v ...interface{}) {
	_logger.Infow(format, v...)
}

func Debugw(format string, v ...interface{}) {
	_logger.Debugw(format, v...)
}
