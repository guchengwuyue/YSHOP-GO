package admin

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"yixiang.co/yshop/controllers"
	"yixiang.co/yshop/models"
	"yixiang.co/yshop/models/vo"
)

// 字典详情api
type DictDetailController struct {
	controllers.BaseController
}

func (c *DictDetailController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title 获取字典详情列表
// @Description 获取字典详情列表
// @Success 200 {object} controllers.Result
// @router / [get]
func (c *DictDetailController) GetAll() {
	dictId, _ := c.GetInt64("dictId")
	dictName := c.GetString("dictName")
	total, list := models.GetAllDictDetail(c.GetParams(), dictId, dictName)
	c.Ok(vo.ResultList{Content: list,TotalElements: total})
}

// @Title 添加字典详情
// @Description 添加字典详情
// @Success 200 {object} controllers.Result
// @router / [post]
func (c *DictDetailController) Post() {
	var model models.SysDictDetail
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	logs.Info(model)
	c.Valid(&model)
	_, e := models.AddDictDetail(&model)
	if e != nil {
		c.Fail(e.Error(),5002)
	}
	c.Ok("操作成功")
}

// @Title 修改字典详情
// @Description 修改字典详情
// @Success 200 {object} controllers.Result
// @router / [put]
func (c *DictDetailController) Put() {
	var model models.SysDictDetail
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	c.Valid(&model)
	e := models.UpdateByDictDetail(&model)
	if e != nil {
		c.Fail(e.Error(),5004)
	}
	c.Ok("操作成功")
}

// @Title 删除字典详情
// @Description 删除字典详情
// @Success 200 {object} controllers.Result
// @router /:id [delete]
func (c *DictDetailController) Delete() {
	id, _ := c.GetInt64(":id", 1)
	e := models.DelByDictDetail(id)
	if e != nil {
		c.Fail(e.Error(),5005)
	}
	c.Ok("操作成功")
}
