package admin

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"strconv"
	"yixiang.co/yshop/controllers"
	"yixiang.co/yshop/dto"
	"yixiang.co/yshop/models"
	"yixiang.co/yshop/vo"
)

// 角色 API
type RoleController struct {
	controllers.BaseController
}

func (c *RoleController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}
// @Title 获取单个角色
// @Description 获取单个角色
// @Param    id        path     int    true        "角色ID"
// @Success 200 {object} models.Role
// @router /:id [get]
func (c *RoleController) GetOne() {
	id := c.Ctx.Input.Param(":id")
	id64, _ := strconv.ParseInt(id,10,64)
	role := models.GetOneRole(id64)
	c.Data["json"] = controllers.SuccessData(role)
	c.ServeJSON()
}

// @Title 角色列表
// @Description 角色列表
// @Success 200 {object} controllers.Result
// @router / [get]
func (c *RoleController) GetAll() {
	total,list := models.GetAllRole(c.GetParams())
	c.Data["json"] = controllers.SuccessData(vo.ResultList{Content: list,TotalElements: total})
	c.ServeJSON()
}

// @Title 角色添加
// @Description 角色添加
// @Success 200 {object} controllers.Result
// @router / [post]
func (c *RoleController) Post()  {
	var model models.SysRole
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	b, _ := valid.Valid(&model)
	if !b {
		for _, err := range valid.Errors {
			c.Data["json"] = controllers.ErrMsg(err.Message)
		}
	}
	_, e := models.AddRole(&model)
	if e != nil {
		c.Data["json"] = controllers.ErrMsg(e.Error())
	}
	c.Data["json"] = controllers.SuccessData("操作成功")
	c.ServeJSON()
}

// @router / [put]
func (c *RoleController) Put()  {
	var model models.SysRole
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	b, _ := valid.Valid(&model)
	if !b {
		for _, err := range valid.Errors {
			c.Data["json"] = controllers.ErrMsg(err.Message)
		}
	}
	e := models.UpdateByRole(&model)
	if e != nil {
		c.Data["json"] = controllers.ErrMsg(e.Error())
	}
	c.Data["json"] = controllers.SuccessData("操作成功")
	c.ServeJSON()
}

// @Title 角色删除
// @Description 角色删除
// @Success 200 {object} controllers.Result
// @router / [delete]
func (c *RoleController) Delete() {
	var ids []int64
	json.Unmarshal(c.Ctx.Input.RequestBody, &ids)
	logs.Info(ids)
	e := models.DelByRole(ids)
	if e != nil {
		c.Data["json"] = controllers.ErrMsg(e.Error())
	}
	c.Data["json"] = controllers.SuccessData("操作成功")
	c.ServeJSON()
}

// @Title 角色菜单更新
// @Description 角色菜单更新
// @Success 200 {object} controllers.Result
// @router /menu [put]
func (c *RoleController) Menu()  {
	var model dto.RoleMenu
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	logs.Info("=======menu======")
	logs.Info(model)
	b, _ := valid.Valid(&model)
	if !b {
		for _, err := range valid.Errors {
			c.Data["json"] = controllers.ErrMsg(err.Message)
		}
	}
	models.BatchRoleMenuAdd(model)
	c.Data["json"] = controllers.SuccessData("操作成功")
	c.ServeJSON()
}