package jjwt

import (
	"acc/server/global"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SigningKey []byte
}

type JwtClaims struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var (
	TokenExpired     = errors.New("token 已过期")
	TokenNotValidYet = errors.New("token 无效")
	TokenMalformed   = errors.New("token 格式错误")
	TokenInvalid     = errors.New("token 不可用")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}

// CreateToken 生成token
func (j *JWT) CreateToken(claims JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParserToken 解析token
func (j *JWT) ParserToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}

	return nil, TokenInvalid
}
