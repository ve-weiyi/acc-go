package api

import (
	"acc/app/model/bean"
	"acc/app/model/dao"
	"acc/app/model/dto"
	"acc/app/service/userService"
	"acc/config"
	"acc/lib/jjwt"
	"acc/lib/orm"
	"acc/lib/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserLogin
// @Summary 用户登录
// @Param	auth	body		dto.AuthReq   	true 	"登录参数"
// @Success 200		{object}  	response.Response{data=bean.LoginResp}
// @Router  /api/v1/login [post]
func UserLogin(c *gin.Context) {
	apiG := &response.Gin{C: c}
	var auth dto.AuthReq

	err := c.ShouldBindJSON(&auth)
	if err != nil {
		apiG.OnError(err)
		return
	}
	login, apierr := userService.UserLogin(auth)

	if apierr != nil {
		apiG.OnError(apierr)
		//apiG.Error(apierr.Code(), apierr.Error(), nil)
		return
	}

	apiG.Success(login)

}

// UserGetInfo
// @Summary 用户信息
// @Description 用户信息
// @Param  	Authorization	header  	string 	 	true	"登录后的token"
// @Success 200 			{object} 	response.Response{data=bean.UserInfoResp}
// @Router  /api/v1/user/info [get]
func UserGetInfo(c *gin.Context) {
	apiG := &response.Gin{C: c}
	uid := c.GetInt("uid")

	data, err := userService.UserGetInfoById(uid)
	if err != nil {
		apiG.OnError(err, nil)
		return
	}

	var userinfo = bean.UserInfoResp{
		Id:       uid,
		Username: c.GetString("username"),
		Roles:    []string{"admin,editor"},
		Details:  data,
	}

	apiG.Success(userinfo)
}

// UserList
// @Summary 用户列表
// @Description 用户列表
// @Param 	Authorization	header  	string 	 	false	"登录后的token"
// @Param  	page           	path       	string      true  	"页码"
// @Success 200 			{object} 	response.Response{data=response.PageResult{datas=[]entity.UserInfo}}
// @Router /api/v1/user/list/{page} [get]
func UserList(c *gin.Context) {
	apiG := &response.Gin{C: c}
	q := dao.Use(orm.DB()).UserInfo

	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		apiG.ErrorMsg(err.Error())
		return
	}

	size := config.AppConfig.PageSize
	users, count, err := q.Order(q.ID).FindByPage(page*size, size)
	if err != nil {
		apiG.ErrorMsg(err.Error())
		return
	}

	apiG.Success(response.PageResult{
		Datas: users,
		Size:  size,
		Page:  page,
		Total: count,
	})
}

// UserTokenParse
// @Summary 用户token信息解析
// @Description
// @Param 	Authorization	header  	string 	 	true	"登录后的token"
// @Success 200 			{object} 	response.Response
// @Router /api/v1/user/token/parse [get]
func UserTokenParse(c *gin.Context) {
	apiG := &response.Gin{C: c}
	claims, err := jjwt.ParseToken(c)
	if err != nil {
		apiG.OnError(err)
	}

	apiG.Success(claims)
}
