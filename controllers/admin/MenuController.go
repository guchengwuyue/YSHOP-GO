package admin

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/validation"
	"yixiang.co/yshop/common/jwt"
	"yixiang.co/yshop/controllers"
	"yixiang.co/yshop/models"
	"yixiang.co/yshop/vo"
)

// 菜单api
type MenuController struct {
	controllers.BaseController
}

func (c *MenuController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Build", c.Build)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title 菜单列表
// @Description 菜单列表
// @Success 200 {object} controllers.Result
// @router / [get]
func (c *MenuController) GetAll() {
	name := c.GetString("blurry")
	menus := models.GetAllMenus(name)
	c.Ok(vo.ResultList{Content: menus,TotalElements: 0})
}

// @Title 菜单添加
// @Description 菜单添加
// @Success 200 {object} controllers.Result
// @router / [post]
func (c *MenuController) Post()  {
	var model models.SysMenu
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	b, _ := valid.Valid(&model)
	if !b {
		for _, err := range valid.Errors {
			c.Fail(err.Message,5001)
		}
	}
	_, e := models.AddMenu(&model)
	if e != nil {
		c.Fail(e.Error(),5002)
	}
	c.Ok("操作成功")
}

// @Title 菜单修改
// @Description 菜单修改
// @Success 200 {object} controllers.Result
// @router / [put]
func (c *MenuController) Put()  {
	var model models.SysMenu
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	b, _ := valid.Valid(&model)
	if !b {
		for _, err := range valid.Errors {
			c.Fail(err.Message,5003)
		}
	}
	e := models.UpdateByMenu(&model)
	if e != nil {
		c.Fail(e.Error(),5004)
	}
	c.Ok("操作成功")
}

// @Title 菜单删除
// @Description 菜单删除
// @Success 200 {object} controllers.Result
// @router / [delete]
func (c *MenuController) Delete() {
	var ids []int64
	json.Unmarshal(c.Ctx.Input.RequestBody, &ids)
	e := models.DelByMenu(ids)
	if e != nil {
		c.Fail(e.Error(),5005)
	}
	c.Ok("操作成功")
}

// @Title 菜单构建
// @Description 菜单构建
// @Success 200 {object} controllers.Result
// @router /build [get]
func (c *MenuController) Build() {
	uid, _:= jwt.GetAdminUserId(c.Ctx.Input)
	menus := models.BuildMenus(uid)
	c.Ok(menus)
}

// @Title 菜单树形
// @Description 菜单树形
// @Success 200 {object} controllers.Result
// @router /tree [get]
func (c *MenuController) GetTree() {
	menus := models.GetAllMenus("")
	c.Ok(menus)
}


