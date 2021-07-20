package admin

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"yixiang.co/yshop/controllers"
	"yixiang.co/yshop/models"
	"yixiang.co/yshop/models/vo"
)

// 角色 API
type LogController struct {
	controllers.BaseController
}

func (c *LogController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Delete", c.Delete)
}


// @Title 日志列表
// @Description 日志列表
// @Success 200 {object} controllers.Result
// @router / [get]
func (c *LogController) GetAll() {
	total,list := models.GetAllLog(c.GetParams())
	c.Ok(vo.ResultList{Content: list,TotalElements: total})
}



// @Title 日志删除
// @Description 日志删除
// @Success 200 {object} controllers.Result
// @router / [delete]
func (c *LogController) Delete() {
	var ids []int64
	json.Unmarshal(c.Ctx.Input.RequestBody, &ids)
	logs.Info(ids)
	e := models.DelBylog(ids)
	if e != nil {
		c.Fail(e.Error(),5005)
	}
	c.Ok("操作成功")
}

