// 1.包名：包名一般与目录相同,相同目录只能有一个包名
package main

//2.导包：可以使用 import "p1" 和 import ("p1","p2" )
import (
	"acc/server/global"
	"acc/server/initialize"
	"log"
)

// 3.init先于main运行
func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("init invoke")
}

// 4.只有main包下的main才能运行
func main() {
	log.Println("main invoke")
	initialize.ReadConfig("config")
	//日志文件
	initialize.Glog()
	//初始化数据库
	initialize.Gorm()

	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	if global.GVA_DB != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	initialize.RunService()
}
