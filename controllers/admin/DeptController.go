package admin

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"yixiang.co/yshop/controllers"
	"yixiang.co/yshop/models"
	"yixiang.co/yshop/models/vo"
)
// 部门api
type DeptController struct {
	controllers.BaseController
}

func (c *DeptController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title 获取部门列表
// @Description 获取部门列表
// @Success 200 {object} controllers.Result
// @router / [get]
func (c *DeptController) GetAll() {
	name := c.GetString("name")
	enabled, _ := c.GetInt8("enabled",-1)
	list := models.GetAllDepts(name,enabled)
	c.Ok(vo.ResultList{Content: list,TotalElements: 0})
}

// @Title 添加部门
// @Description 添加部门
// @Success 200 {object} controllers.Result
// @router / [post]
func (c *DeptController) Post()  {
	var model models.SysDept
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	c.Valid(&model)
	_, e := models.AddDept(&model)
	if e != nil {
		c.Fail(e.Error(),5002)
	}
	c.Ok("操作成功")
}

// @Title 修改部门
// @Description 修改部门
// @Success 200 {object} controllers.Result
// @router / [put]
func (c *DeptController) Put()  {
	var model models.SysDept
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	c.Valid(&model)
	e := models.UpdateByDept(&model)
	if e != nil {
		c.Fail(e.Error(),5004)
	}
	c.Ok("操作成功")
}

// @Title 删除部门
// @Description 删除部门
// @Success 200 {object} controllers.Result
// @router / [delete]
func (c *DeptController) Delete() {
	var ids []int64
	json.Unmarshal(c.Ctx.Input.RequestBody, &ids)
	logs.Info(ids)
	e := models.DelByDept(ids)
	if e != nil {
		c.Fail(e.Error(),5005)
	}
	c.Ok("操作成功")
}