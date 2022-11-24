package jjwt

import (
	"acc/server/global"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type ClaimsInfo struct {
	ID       int    `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	Username string `gorm:"column:username;type:varchar(255);primaryKey" json:"username"`
	Password string `gorm:"column:password;type:varchar(255);not null" json:"password"`
}

// 根据用户登录信息生成token，
func CreateTokenByInfo(auth ClaimsInfo) string {
	j := NewJWT()
	claims := JwtClaims{
		Uid:      auth.ID,
		Username: auth.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			Issuer:    "jjwt",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Debug(err.Error())
	}
	return token
}

func ParseTokenByGin(c *gin.Context) (*JwtClaims, error) {
	tokenHeader := c.Request.Header.Get("Authorization")

	//token是空
	if tokenHeader == "" {
		tokenHeader, _ = c.Cookie("token")
		global.GVA_LOG.Debug("get token by cookie :" + tokenHeader)
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
