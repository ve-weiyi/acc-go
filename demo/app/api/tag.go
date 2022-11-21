package api

import (
	"acc/app/model/dao"
	"acc/app/model/entity"
	"acc/config"
	"acc/lib/errCode"
	"acc/lib/logger"
	"acc/lib/orm"
	"acc/lib/response"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
)

// TagAdd
// @Summary 标签添加
// @Description
// @Param  Authorization	header  	string 	 	false	"登录后的token"
// @Param  data		      	body   		entity.Tag  true 	"请求体"
// @Success 200 			{object} 	response.Response
// @Router /api/v1/tag/add [post]
func TagAdd(ctx *gin.Context) {
	apiG := &response.Gin{C: ctx}
	q := dao.Use(orm.DB()).Tag

	var tag entity.Tag
	err := ctx.ShouldBindJSON(&tag)
	body, _ := io.ReadAll(ctx.Request.Body)
	logger.Debug(body)
	if err != nil {
		apiG.ErrorMsg(err.Error())
		return
	}

	err = q.Create(&tag)
	if err != nil {
		apiG.ErrorMsg(err.Error())
		return
	}

	apiG.Success(tag)
}

// TagDelete
// @Summary 标签删除
// @Description
// @Param  Authorization	header  	string 	 	false	"登录后的token"
// @Param  id      		    query   	int  		true 	"id"
// @Success 200 			{object} 	response.Response
// @Router /api/v1/tag/delete [delete]
func TagDelete(ctx *gin.Context) {
	apiG := &response.Gin{C: ctx}
	q := dao.Use(orm.DB()).Tag

	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		apiG.ErrorMsg(err.Error())
		return
	}
	//先查找
	tag, err := q.Where(q.ID.Eq(id)).First()
	if err != nil {
		apiG.ErrorMsg(err.Error())
		return
	}
	//找到后删除
	res, err := q.Delete(tag)
	if err != nil {
		apiG.ErrorMsg(err.Error())
		return
	}

	apiG.Success(res.RowsAffected)
}

// TagUpdate
// @Summary 标签更新
// @Description
// @Param  Authorization	header  	string 	 	false	"登录后的token"
// @Param  tag      	    body   	entity.Tag  true 	"参数"
// @Success 200 			{object} 	response.Response
// @Router /api/v1/tag/update [put]
func TagUpdate(ctx *gin.Context) {
	apiG := &response.Gin{C: ctx}
	q := dao.Use(orm.DB()).Tag

	var tag entity.Tag
	err := ctx.ShouldBindJSON(&tag)
	if err != nil {
		apiG.ErrorMsg(err.Error())
		return
	}
	//先查找
	_, err = q.Where(q.ID.Eq(tag.ID)).First()
	if err != nil {
		apiG.ErrorMsg(err.Error())
		return
	}

	//找到后更新
	_, err = q.Updates(&tag)
	if err != nil {
		apiG.ErrorMsg(err.Error())
		return
	}

	apiG.Success(tag)
}

// TagList
// @Summary 标签查询
// @Description
// @Param	Authorization	header  	string 	 		false	"登录后的token"
// @Param	page            path      	string  		false  	"页码"
// @Param  data  	      	query      	dto.Condition  	true  "查询参数"
// @Success 200 			{object} 	response.Response
// @Router /api/v1/tag/list/{page} [get]
func TagList(ctx *gin.Context) {
	apiG := &response.Gin{C: ctx}
	q := dao.Use(orm.DB()).Tag

	key := ctx.Query("keywords")
	page, err := strconv.Atoi(ctx.Param("page"))
	if err != nil {
		apiG.ErrorMsg(err.Error())
		return
	}

	limit := config.AppConfig.PageSize
	offset := page * limit

	res, total, err := q.Where(q.Name.Like("%"+key+"%")).
		Order(q.ID).FindByPage(offset, limit)

	if err != nil {
		apiG.ErrorMsg(err.Error())
		return
	}

	if total < int64(offset) {
		apiG.HandleError(nil, errCode.NewErrorMsg("超出页码范围"))
		return
	}

	result := &response.PageResult{
		Datas: res,
		Size:  limit,
		Page:  page,
		Total: total,
	}

	apiG.Success(result)
}
