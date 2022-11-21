package middleware

import (
	"acc/lib/logger"
	"acc/lib/response"
	"github.com/gin-gonic/gin"
)

// 捕获全局错误
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				appG := response.Gin{C: c}
				logger.Error(err)
				appG.ErrorMsg("系统错误")
				c.Abort()
			}
		}()
		c.Next()
	}
}
