package api

import (
	"acc/app/model/dto"
	"acc/app/service/todoService"
	"acc/lib/response"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
)

// TodoAdd
// @Summary todo添加
// @Description
// @Param  Authorization	header  	string 	 	false	"登录后的token"
// @Param  data		      	body   		dto.TodoReq  true 	"请求体"
// @Success 200 			{object} 	response.Response
// @Router /api/v1/todo/add [post]
func TodoAdd(ctx *gin.Context) {
	apiG := &response.Gin{C: ctx}

	uid := ctx.GetInt("uid")
	if uid <= 0 {
		apiG.ErrorMsg("无操作权限")
		return
	}

	var data *dto.TodoReq
	err := ctx.ShouldBindJSON(data)
	if err != nil {
		body, _ := io.ReadAll(ctx.Request.Body)
		apiG.ErrorMsg("输入格式不正确", body)
		return
	}

	//用户关联
	data.Todo.UID = uid
	apiG.HandleError(todoService.TodoAdd(data))
}

// TodoDelete
// @Summary todo删除
// @Description
// @Param  Authorization	header  	string 	 	false	"登录后的token"
// @Param  id      		    query     	int  		true 	"id"
// @Success 200 			{object} 	response.Response
// @Router /api/v1/todo/delete [delete]
func TodoDelete(ctx *gin.Context) {
	apiG := &response.Gin{C: ctx}
	id, _ := strconv.Atoi(ctx.Query("id"))

	apiG.HandleError(todoService.TodoDelete(id))
}

// TodoUpdate
// @Summary todo修改
// @Description
// @Param  Authorization	header  	string 	 	false	"登录后的token"
// @Param  data		      	body   		dto.TodoReq  true 	"请求体"
// @Success 200 			{object} 	response.Response
// @Router /api/v1/todo/update [put]
func TodoUpdate(ctx *gin.Context) {
	apiG := &response.Gin{C: ctx}
	var todo *dto.TodoReq
	err := ctx.ShouldBindJSON(todo)
	if err != nil {
		apiG.ErrorMsg("输入格式不正确")
		return
	}

	result, err := todoService.TodoUpdate(todo)
	if err != nil {
		apiG.ErrorMsg("查询错误")
		return
	}
	apiG.Success(result)
}

// TodoList
// @Summary  todo查询
// @Param  Authorization  header  	 string         true  "登录后的token"
// @Param  page           path       string         true  "页码"
// @Param  data  	      query      dto.Condition  true  "查询参数"
// @router /api/v1/todo/list/{page} [get]
func TodoList(ctx *gin.Context) {
	apiG := &response.Gin{C: ctx}
	// 查询参数
	p := ctx.Param("page")
	key := ctx.Query("keywords")

	page, _ := strconv.Atoi(p)

	var condition dto.Condition
	condition.Keywords = key
	condition.Uid = ctx.GetInt("uid")

	query, err := todoService.TodoQuery(page, condition)
	if err != nil {
		apiG.ErrorMsg("操作失败")
		return
	}
	apiG.Success(query)
}
