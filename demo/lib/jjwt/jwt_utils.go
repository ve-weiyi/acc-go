package jjwt

import (
	"acc/app/model/entity"
	"acc/config"
	"acc/lib/logger"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
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
		logger.Debug(err.Error())
	}
	return token
}

func ParseToken(c *gin.Context) (*JwtClaims, error) {
	tokenHeader := c.Request.Header.Get("Authorization")

	//token是空
	if tokenHeader == "" {
		tokenHeader, _ = c.Cookie("token")
		logger.Debug("get token by cookie :" + tokenHeader)
		return nil, TokenInvalid
	}

	//验证token是否 Bearer 开头的
	ok := strings.HasPrefix(tokenHeader, TokenPrefix())
	if !ok {
		return nil, TokenInvalid
	}

	token := strings.TrimPrefix(tokenHeader, TokenPrefix())

	j := NewJWT()
	// 解析token
	claims, err := j.ParserToken(token)
	if err != nil {
		return nil, TokenMalformed
	}
	return claims, nil
}

func TokenPrefix() string {
	//return config.AppConfig.JwtTokenHeader+" "
	return ""
}
