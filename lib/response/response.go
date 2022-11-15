package response

import (
	"acc/lib/errCode"
	"acc/lib/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}

func (g *Gin) HandleError(data interface{}, err *errCode.ApiError) {
	if err != nil {
		g.Response(err.Code(), err.Error(), nil)
	} else {
		g.Success(data)
	}
}

// NewError 失败返回
func (g *Gin) OnError(err error, data ...interface{}) {
	switch v := reflect.ValueOf(err); v {
	case reflect.ValueOf(errCode.ApiError{}):
		code := v.MethodByName("Code").Call(nil)
		msg := v.MethodByName("Message").Call(nil)
		int, _ := strconv.Atoi(fmt.Sprintln(code))
		g.C.JSON(http.StatusOK, Response{
			Code:    int,
			Message: fmt.Sprintln(msg),
			Data:    utils.CheckData(data),
		})

	default:
		g.C.JSON(http.StatusOK, Response{
			Code:    501,
			Message: err.Error(),
			Data:    utils.CheckData(data),
		})
	}
}

// Response 返回封装 {code,msg,data}
func (g *Gin) Response(code int, message string, data ...interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    utils.CheckData(data),
	})
}

// Success 成功返回
func (g *Gin) Success(data ...interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    utils.CheckData(data),
	})
}

// NewError 自定义错误信息返回
func (g *Gin) ErrorMsg(msg string, data ...interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code:    errCode.CodeError,
		Message: msg,
		Data:    utils.CheckData(data),
	})
	//结束请求
	//g.C.AbortWithStatus(http.StatusOK)
}
