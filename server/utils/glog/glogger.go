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

func (mlog *Logger) checkRotate() {
	if !mlog.rolling {
		return
	}

	n := time.Now()
	if mlog.differentDay(n) {
		mlog.rotateMu.Lock()
		defer mlog.rotateMu.Unlock()

		// 获得锁之后再次检查是否是不同日期
		// 避免上一次调用已经切割日志,
		if mlog.differentDay(n) {
			mlog.rlog.Rotate()
			mlog.lastRotate = n
		}
	}
}

// 判断是不是换天了，如果换天了就要重新调用rotate()
func (mlog *Logger) differentDay(t time.Time) bool {
	y, m, d := mlog.lastRotate.Year(), mlog.lastRotate.Month(), mlog.lastRotate.Day()
	return y != t.Year() || m != t.Month() || d != t.Day()
}

func (mlog *Logger) EnableDailyFile() {
	mlog.rolling = true
}

func (mlog *Logger) Error(v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Error(v...)
}

func (mlog *Logger) Warn(v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Warn(v...)
}

func (mlog *Logger) Info(v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Info(v...)
}

func (mlog *Logger) Debug(v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Debug(v...)
}

func (mlog *Logger) Errorw(format string, v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Errorw(format, v...)
}

func (mlog *Logger) Warnw(format string, v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Warnw(format, v...)
}

func (mlog *Logger) Infow(format string, v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Infow(format, v...)
}

func (mlog *Logger) Debugw(format string, v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Debugw(format, v...)
}

func (mlog *Logger) Print(v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Info(v...)
}

func (mlog *Logger) Printf(format string, v ...interface{}) {
	mlog.checkRotate()
	mlog.sugar.Info(fmt.Sprintf(format, v...))
}

func (mlog *Logger) GetUnderlyingLogger() *zap.Logger {
	return mlog.log
}
