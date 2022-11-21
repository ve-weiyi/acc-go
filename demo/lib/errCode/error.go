package errCode

import "fmt"

type ApiError struct {
	code    int
	message interface{}
	//	errors.New("用户名不存在")
}

func (api *ApiError) Code() int {
	return api.code
}

// 实现 `error` 接口
func (api *ApiError) Error() string {
	return fmt.Sprintf("%v", api.message)
}

func OnError(err error) *ApiError {
	return &ApiError{CodeError, err.Error()}
}

func NewErrorMsg(msg string) *ApiError {
	return &ApiError{CodeError, msg}
}

func NewError(code int, msg interface{}) *ApiError {
	return &ApiError{code, msg}
}

// 定义错误
var (
	ErrorTokenExpired     error = NewError(301, "Token已过期,请重新登录")
	ErrorTokenNotValidYet error = NewError(302, "Token无效,请重新登录")
	ErrorTokenMalformed   error = NewError(304, "Token不正确,请重新登录")
	ErrorTokenInvalid     error = NewError(303, "这不是一个token,请重新登录")
	ErrorTokenIsEmpty     error = NewError(303, "Token为空，请重新登录")

	ErrorSystem error = NewError(501, "系统错误")
)

// 错误编码
const (
	// 错误码统一格式 A-BB-CC
	// A: 1-2-3-成功
	// A: 4-客户端异常 请求方式错误、请求参数错误、没有权限、没有登录、上传文件过大、重复提交抢锁失败、并发过大触发限流等
	// A: 5-服务端异常 依赖的第三方业务系统异常：比如调用第三方系统超时；第三方系统抛出异常；第三方业务系统限流等
	// A: 6-基础中间件异常 如 MySql、Redis、Mongodb、MQ 等基础中间件出现连接超时、连接池满、访问失败等
	// A: 7-数据问题 数据不一致、记录不存在、主键冲突、字段不能为空等等
	// BB: 模块名称
	// CC: 具体错误编号 自增即可

	// 通用失败 0
	CodeError = 40000
)
