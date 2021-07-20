package admin

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"yixiang.co/yshop/controllers"
	"yixiang.co/yshop/models"
	"yixiang.co/yshop/models/vo"
)

// 岗位api
type JobController struct {
	controllers.BaseController
}

func (c *JobController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title 岗位列表
// @Description 岗位列表
// @Success 200 {object} controllers.Result
// @router / [get]
func (c *JobController) GetAll() {
	enabled, _ := c.GetInt64("enabled",-1)
	total,list := models.GetAllJob(c.GetParams(),enabled)
	c.Ok(vo.ResultList{Content: list,TotalElements: total})
}

// @Title 岗位添加
// @Description 岗位添加
// @Success 200 {object} controllers.Result
// @router / [post]
func (c *JobController) Post()  {
	var model models.SysJob
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	c.Valid(&model)
	_, e := models.AddJob(&model)
	if e != nil {
		c.Fail(e.Error(),5002)
	}
	c.Ok("操作成功")
}

// @Title 岗位修改
// @Description 岗位修改
// @Success 200 {object} controllers.Result
// @router / [put]
func (c *JobController) Put()  {
	var model models.SysJob
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	c.Valid(&model)
	e := models.UpdateByJob(&model)
	if e != nil {
		c.Fail(e.Error(),5004)
	}
	c.Ok("操作成功")
}

// @Title 岗位删除
// @Description 岗位删除
// @Success 200 {object} controllers.Result
// @router / [delete]
func (c *JobController) Delete() {
	var ids []int64
	json.Unmarshal(c.Ctx.Input.RequestBody, &ids)
	logs.Info(ids)
	e := models.DelByJob(ids)
	if e != nil {
		c.Fail(e.Error(),5005)
	}
	c.Ok("操作成功")
}