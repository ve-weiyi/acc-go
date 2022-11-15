package jjwt

import (
	"acc/app/model/entity"
	"acc/config"
	"acc/lib/errCode"
	"acc/lib/logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 根据用户登录信息生成token，
func CreateToken(auth entity.UserAuth) string {
	j := NewJWT()
	claims := JwtClaims{
		Uid:      auth.ID,
		Username: auth.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			Issuer:    config.AppConfig.Name,
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		log.Print(err.Error())
	}
	return token
}

func ParseToken(c *gin.Context) (*JwtClaims, error) {
	tokenHeader := c.Request.Header.Get("Authorization")

	if tokenHeader == "" {
		tokenHeader, _ = c.Cookie("token")
		logger.Debug("get token by cookie :" + tokenHeader)
	}

	j := NewJWT()
	// 解析token
	claims, err := j.ParserToken(tokenHeader)
	if err != nil {
		return nil, errCode.ErrorTokenMalformed
	}
	return claims, nil
}
