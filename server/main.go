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
	Index()
	log.Println("init invoke")
}

// 4.只有main包下的main才能运行
func main() {
	log.Println("main invoke")
	initialize.ReadConfig("server/config")
	initialize.Zap()
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
}

func Index() {
	//SetFlags为标准记录器设置输出标志。默认标志位是Ldate、Ltime等。
	//log.Ldate：格式是:2009年1月23日
	//log.Ltime：格式是:01:23:23
	log.SetFlags(log.Ldate | log.Ltime)
	//当发生错误或者查看信息的时候，需要查看日志，
	//默认的日志是不显示行号的，
	//可以通过log.SetFlags函数设置显示行号
	//log.LstdFlags：标准默认的日志信息
	//log.Llongfile：显示日志的文件（绝对路径）和对应行号
	//log.Lshortfile：显示日志的文件（不含路径）和对应行号
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("输出日志")
}
