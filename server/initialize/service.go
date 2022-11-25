package initialize

import (
	"acc/server/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func RunService() {

	router := InitRouter()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatalln(server.ListenAndServe())

}

// InitRouter router 入口文件
func InitRouter() *gin.Engine {
	//先设置setMode(),再创建 gin.New() 才能正确运行配置
	gin.SetMode(gin.DebugMode)

	//engine := gin.Default() //默认使用了2个中间件Logger(), Recovery()
	engine := gin.New()
	// 中间件
	engine.Use(middleware.GinLogger())
	engine.Use(gin.Logger())
	// 捕获全局错误 上线开启
	engine.Use(middleware.GinRecovery(true))
	// 跨域设置
	engine.Use(middleware.Cors())

	// 静态文件访问
	engine.StaticFS("/static", http.Dir("./static"))
	//处理路由
	engine.GET("/", func(ctx *gin.Context) {
		//以字符串格式返回
		ctx.JSON(200, "hello world!")
	})
	engine.NoRoute(func(ctx *gin.Context) {
		// 实现内部重定向
		//以字符串格式返回
		ctx.JSON(200, "url error!")
	})

	// 加载模块路由
	_ = engine.Group("/api/v1")
	return engine
}
