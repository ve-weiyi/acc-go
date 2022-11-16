package v1

import (
	"acc/lib/response"
	"github.com/gin-gonic/gin"
)

// Report
// @Summary 用户token信息解析
// @Description
// @Param 	Authorization	header  	string 	 	false	"登录后的token"
// @Success 200 			{object} 	response.Response
// @Router /api/v1/report [get]
func Report(c *gin.Context) {
	apiG := &response.Gin{C: c}
	apiG.Success("操作成功")
}
