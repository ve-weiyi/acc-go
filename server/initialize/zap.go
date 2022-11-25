package initialize

import (
	"acc/server/global"
	"acc/server/utils/glog"
	"fmt"
	"time"
)

// Glog 获取 zap.Logger
func Glog() {
	cfg := glog.LogConfig{
		Level:      global.GVA_CONFIG.Zap.Level,
		FileName:   fmt.Sprintf("%v/%v.log", global.GVA_CONFIG.Zap.Director, time.Now().Format("2006-01-02")),
		Format:     global.GVA_CONFIG.Zap.Format,
		MaxSize:    1,
		MaxAge:     30,
		MaxBackups: 1,
	}

	global.GVA_LOG = glog.NewLogger(cfg, 1)
}
