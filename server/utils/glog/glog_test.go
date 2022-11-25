package glog

import (
	"go.uber.org/zap"
	"testing"
)

func TestName(t *testing.T) {
	Error("hello")
	zap.L().Debug("hello world")
}
