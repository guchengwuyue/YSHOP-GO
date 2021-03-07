package admin

import (
	"yixiang.co/yshop/controllers"
)

type ToolsController struct {
	controllers.BaseController
}

// @router / [post]
func (c *ToolsController) Post()  {

	c.Data["json"] = controllers.SuccessData("操作成功")
	c.ServeJSON()
}