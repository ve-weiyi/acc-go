package initialize

import (
	"acc/server/global"
	"acc/server/utils"
	"acc/server/utils/glog"
	"fmt"
	"os"
	"path"
	"time"
)

// Glog 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Glog() {
	if ok, _ := utils.PathExists(global.GVA_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.GVA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GVA_CONFIG.Zap.Director, os.ModePerm)
	}

	// 获取日志最低等级，即>=该等级，才会被写入。
	level := global.GVA_CONFIG.Zap.TransportLevel()
	filepath := path.Join(global.GVA_CONFIG.Zap.Director, time.Now().Format("2006-01-01"), level.String()+".log")
	format := global.GVA_CONFIG.Zap.Format
	global.GVA_LOG = glog.InitLogger(filepath, level, format)
}
