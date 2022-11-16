package routers

import (
	"acc/app/api"
	v1 "acc/app/api/v1"
	"acc/app/middleware"
	"acc/config"
	"acc/docs"
	"acc/lib/file"
	"acc/lib/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.AppConfig.RunMode)

	//gin.DisableConsoleColor()

	gin.DefaultWriter = io.MultiWriter(logger.F, os.Stdout)
	//engine := gin.Default() //默认使用了2个中间件Logger(), Recovery()
	engine := gin.New()
	//engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// 中间件
	// 捕获全局错误 上线开启
	//engine.Use(middleware.Recover())
	// 跨域设置
	engine.Use(middleware.Cors())

	// 静态文件访问
	engine.StaticFS(config.ImageConfig.StaticUrl, http.Dir(file.ImageStaticDir()))
	//处理路由
	engine.GET("/", func(ctx *gin.Context) {
		//以字符串格式返回
		ctx.JSON(200, "hello world!")
	})

	swaggerInit(engine)
	// 加载模块路由
	group := engine.Group("/api/v1")
	AddSystemRoute(group)
	AddAuthRoute(group)
	return engine
}

func AddSystemRoute(r *gin.RouterGroup) {
	r.GET("/", func(ctx *gin.Context) {
		//以字符串格式返回
		ctx.JSON(200, "hello world!")
	})
	r.POST("/login", api.UserLogin)
	r.POST("/report", v1.Report)
	auth := r.Group("/")
	{
		auth.POST("/tag/add", api.TagAdd)
		auth.DELETE("/tag/delete", api.TagDelete)
		auth.PUT("/tag/update", api.TagUpdate)
		auth.GET("/tag/list/:page", api.TagList)

	}
}
func AddAuthRoute(r *gin.RouterGroup) {
	auth := r.Group("/")
	auth.Use(middleware.JwtToken())
	{
		auth.GET("/user/info", api.UserGetInfo)
		auth.GET("/user/list/:page", api.UserList)
		auth.GET("/user/token/parse", api.UserTokenParse)

		auth.DELETE("/todo/delete", api.TodoDelete)
		auth.PUT("/todo/update", api.TodoUpdate)
		auth.POST("/todo/add", api.TodoAdd)
		auth.GET("/todo/list/:page", api.TodoList)
	}

}

func swaggerInit(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	var host = config.AppConfig.Host + ":" + config.AppConfig.Port
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	logger.Debug(fmt.Sprintf("http://%s/swagger/doc/index.html", host))
}
