// 1.包名：包名一般与目录相同,相同目录只能有一个包名
package main

//2.导包：可以使用 import "p1" 和 import ("p1","p2" )
import (
	"acc/config"
	"acc/lib/logger"
	"acc/lib/orm"
	"acc/lib/redis"
	"acc/routers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// 3.init先于main运行
func init() {
	fmt.Println("init invoke")

	config.Setup()

	if config.AppConfig.MysqlState {
		orm.Setup()
	}

	if config.AppConfig.RedisState {
		redis.Setup()
	}

	if config.AppConfig.LoggerState {
		logger.Setup()
	}

}

// 4.只有main包下的main才能运行
func main() {
	fmt.Println("main invoke")
	// 退出关闭 mysql 连接
	if config.AppConfig.MysqlState {
		defer orm.CloseDB()
	}

	port := fmt.Sprintf(":%s", config.AppConfig.Port)
	router := routers.InitRouter()
	readTimeout := config.AppConfig.ReadTimeout
	writeTimeout := config.AppConfig.WriteTimeout
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		// 服务连接
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("监听: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("关闭服务器...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("服务器关闭:", err)
	}
	log.Println("服务器已关闭")
}
