package controller

import (
	"xanthing/internal/utils"

	"github.com/gin-gonic/gin"
)

type WxminiCtl struct{}

func (ctl *WxminiCtl) WxLogin(c *gin.Context) {
	utils.Success(c, "hello world")
	params, err := utils.NewRequestParams(c)
	if err != nil {
		utils.BadRequest(c, "参数解析失败: "+err.Error())
		return
	}
	utils.Success(c, params.GetAll())
	return
}
