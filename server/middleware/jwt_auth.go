package middleware

import (
	"acc/server/utils/jjwt"
	"github.com/gin-gonic/gin"
)

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		claims, err := jjwt.ParseTokenByGin(c)
		if err != nil {
			c.Abort()
			return
		}
		//global.GVA_LOG.Debug("", claims)
		c.Set("uid", claims.Uid)
		c.Set("username", claims.Username)
		c.Next()

	}
}
