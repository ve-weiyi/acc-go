package jjwt

import (
	"acc/config"
	"acc/lib/errCode"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

type JWT struct {
	JwtKey []byte
}

type JwtClaims struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{[]byte(config.AppConfig.JwtSecret)}
}

// CreateToken 生成token
func (j *JWT) CreateToken(claims JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.JwtKey)
}

// ParserToken 解析token
func (j *JWT) ParserToken(tokenString string) (*JwtClaims, error) {
	tokenString, err := checkTokenValid(tokenString)
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errCode.ErrorTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errCode.ErrorTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errCode.ErrorTokenNotValidYet
			} else {
				return nil, errCode.ErrorTokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errCode.ErrorTokenInvalid
	}

	return nil, errCode.ErrorTokenInvalid
}

func TokenPrefix() string {
	//return config.AppConfig.JwtTokenHeader+" "
	return ""
}
func checkTokenValid(tokenHeader string) (token string, err error) {
	//token是空
	if tokenHeader == "" {
		return tokenHeader, errCode.ErrorTokenIsEmpty
	}

	//验证token是否 Bearer 开头的
	ok := strings.HasPrefix(tokenHeader, TokenPrefix())
	if !ok {
		return tokenHeader, errCode.ErrorTokenInvalid
	}

	token = strings.TrimPrefix(tokenHeader, TokenPrefix())
	return token, nil
}
