package middleware

import (
	"acc/lib/jjwt"
	"acc/lib/response"
	"github.com/gin-gonic/gin"
	"log"
)

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiG := &response.Gin{C: c}

		claims, err := jjwt.ParseToken(c)
		if err != nil {
			apiG.OnError(err)
			c.Abort()
			return
		}

		log.Print(claims)
		c.Set("uid", claims.Uid)
		c.Set("username", claims.Username)
		c.Next()
	}
}
