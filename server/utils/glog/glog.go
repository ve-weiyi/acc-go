package glog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"sync"
	"time"
)

var _logger *Logger

const (
	prefix = "[GLOG]"
)

type LogConfig struct {
	Level      string `json:"level"`       // Level 最低日志等级，DEBUG<INFO<WARN<ERROR<FATAL 例如：info-->收集info等级以上的日志
	FileName   string `json:"file_name"`   // FileName 日志文件位置
	Format     string `json:"format"`      // 输出
	MaxSize    int    `json:"max_size"`    // MaxSize 进行切割之前，日志文件的最大大小(MB为单位)，默认为100MB
	MaxAge     int    `json:"max_age"`     // MaxAge 是根据文件名中编码的时间戳保留旧日志文件的最大天数。
	MaxBackups int    `json:"max_backups"` // MaxBackups 是要保留的旧日志文件的最大数量。默认是保留所有旧的日志文件（尽管 MaxAge 可能仍会导致它们被删除。）
}

// 默认调用
func init() {
	InitLogger(GetDefaultConfig())
}

func InitLogger(cfg LogConfig) {
	_logger = NewLogger(cfg, 2)
	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(NewLogger(cfg, 0).log)
}

func GetDefaultConfig() LogConfig {
	return LogConfig{
		Level:      "debug",
		FileName:   fmt.Sprintf("runtime/log/%v.log", time.Now().Format("2006-01-02")),
		Format:     "json-",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
	}
}
func GetDefaultLogger() *Logger {
	if _logger == nil {
		InitLogger(GetDefaultConfig())
	}
	return _logger
}

// 负责设置 encoding 的日志格式
func getEncoder(format string) zapcore.Encoder {
	// 获取一个指定的的EncoderConfig，进行自定义
	//encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig := zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		NameKey:    "log",
		TimeKey:    "time",
		CallerKey:  "caller",
		//FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if format == "json" {
		//json格式
		encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		encodeConfig.FunctionKey = "func"
		return zapcore.NewJSONEncoder(encodeConfig)
	}
	//控制台格式
	return zapcore.NewConsoleEncoder(encodeConfig)
}

// 获取日志写入位置, 同时输出控制台和文件
func getLogWriter(lumberJackLogger *lumberjack.Logger) zapcore.WriteSyncer {
	syncFile := zapcore.AddSync(lumberJackLogger) // 打印到文件
	syncConsole := zapcore.AddSync(os.Stderr)     // 打印到控制台
	//同时输出控制台和文件
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}

// CustomTimeEncoder 自定义日志输出时间格式,方便加前缀
func CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(prefix + t.Format("2006/01/02-15:04:05.000"))
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
// Logger.Debug->skip1  glog.Debug->skip2
func NewLogger(cfg LogConfig, skip int) *Logger {

	//初始化内部类
	out := new(Logger)

	out.path = cfg.FileName
	out.level = zapcore.DebugLevel
	out.rotateMu = &sync.Mutex{}
	out.rolling = true
	out.lastRotate = time.Now()

	// 日志切割配置
	lumberJackLogger := getLumberJackLogger(cfg.FileName)
	// 获取日志编码格式
	encoder := getEncoder("json")
	colorEncoder := getEncoder("console")

	//保活信息:debug，info ,交由运维同学监控
	savelog := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.InfoLevel
	})
	//错误日志:Warn，Error，Fatal，Panic ,开发成员关注
	errorlog := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	//// 使用了core的NewTee
	core := zapcore.NewTee(
		// 保活日志
		zapcore.NewCore(encoder, zapcore.AddSync(lumberJackLogger), savelog),
		// 错误日志:输入到文件中，使用json格式，无颜色
		zapcore.NewCore(encoder, zapcore.AddSync(lumberJackLogger), errorlog),
		// 控制台日志:使用彩色的console输出格式
		zapcore.NewCore(colorEncoder, zapcore.AddSync(os.Stderr), out.level),
	)

	// 创建一个将日志写入 WriteSyncer 的核心。
	// Logger.Debug->skip1  glog.Debug->skip2
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(skip))

	out.rlog = lumberJackLogger
	out.log = logger
	out.sugar = logger.Sugar()

	return out
}

func Sync() {
	_logger.log.Sync()
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

func SyslogFormat() {
	logFile, err := os.OpenFile("./log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
