// Copyright © 2015-2018 Anker Innovations Technology Limited All Rights Reserved.
package glog

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// go使用zap + lumberjack重构项目的日志系统 https://blog.csdn.net/weixin_52000204/article/details/126651319
type Logger struct {
	rlog  *lumberjack.Logger
	log   *zap.Logger        //并重性能与易用性，支持结构化和 printf 风格的日志记录。
	sugar *zap.SugaredLogger // 非常强调性能，不提供 printf 风格的 api（减少了 interface{} 与 反射的性能损耗）

	path       string
	level      zapcore.Level
	rotateMu   *sync.Mutex
	rolling    bool
	lastRotate time.Time
}

func (tlog *Logger) checkRotate() {
	if !tlog.rolling {
		return
	}

	n := time.Now()
	if tlog.differentDay(n) {
		tlog.rotateMu.Lock()
		defer tlog.rotateMu.Unlock()

		// 获得锁之后再次检查是否是不同日期
		// 避免上一次调用已经切割日志,
		if tlog.differentDay(n) {
			tlog.rlog.Rotate()
			tlog.lastRotate = n
		}
	}
}

func (tlog *Logger) differentDay(t time.Time) bool {
	y, m, d := tlog.lastRotate.Year(), tlog.lastRotate.Month(), tlog.lastRotate.Day()
	return y != t.Year() || m != t.Month() || d != t.Day()
}

func (tlog *Logger) EnableDailyFile() {
	tlog.rolling = true
}

func (tlog *Logger) Error(v ...interface{}) {
	tlog.checkRotate()
	tlog.sugar.Error(v...)
}

func (tlog *Logger) Warn(v ...interface{}) {
	tlog.checkRotate()
	tlog.sugar.Warn(v...)
}

func (tlog *Logger) Info(v ...interface{}) {
	tlog.checkRotate()
	tlog.sugar.Info(v...)
}

func (tlog *Logger) Debug(v ...interface{}) {
	tlog.checkRotate()
	tlog.sugar.Debug(v...)
}

func (tlog *Logger) Errorw(format string, v ...interface{}) {
	tlog.checkRotate()
	tlog.sugar.Errorw(format, v...)
}

func (tlog *Logger) Warnw(format string, v ...interface{}) {
	tlog.checkRotate()
	tlog.sugar.Warnw(format, v...)
}

func (tlog *Logger) Infow(format string, v ...interface{}) {
	tlog.checkRotate()
	tlog.sugar.Infow(format, v...)
}

func (tlog *Logger) Debugw(format string, v ...interface{}) {
	tlog.checkRotate()
	tlog.sugar.Debugw(format, v...)
}

func (tlog *Logger) Print(v ...interface{}) {
	tlog.checkRotate()
	tlog.sugar.Info(v...)
}

func (tlog *Logger) Printf(format string, v ...interface{}) {
	tlog.checkRotate()
	tlog.sugar.Info(fmt.Sprintf(format, v...))
}

func (tlog *Logger) GetUnderlyingLogger() *zap.Logger {
	return tlog.log
}
