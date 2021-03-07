package admin

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/validation"
	"yixiang.co/yshop/controllers"
	"yixiang.co/yshop/jwt"
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
	c.Data["json"] = controllers.SuccessData(vo.ResultList{Content: menus,TotalElements: 0})
	c.ServeJSON()
}

// @Title 菜单添加
// @Description 菜单添加
// @Success 200 {object} controllers.Result
// @router / [post]
func (c *MenuController) Post()  {
	var model models.Menu
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	b, _ := valid.Valid(&model)
	if !b {
		for _, err := range valid.Errors {
			c.Data["json"] = controllers.ErrMsg(err.Message)
		}
	}
	_, e := models.AddMenu(&model)
	if e != nil {
		c.Data["json"] = controllers.ErrMsg(e.Error())
	}
	c.Data["json"] = controllers.SuccessData("操作成功")
	c.ServeJSON()
}

// @Title 菜单修改
// @Description 菜单修改
// @Success 200 {object} controllers.Result
// @router / [put]
func (c *MenuController) Put()  {
	var model models.Menu
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	b, _ := valid.Valid(&model)
	if !b {
		for _, err := range valid.Errors {
			c.Data["json"] = controllers.ErrMsg(err.Message)
		}
	}
	e := models.UpdateByMenu(&model)
	if e != nil {
		c.Data["json"] = controllers.ErrMsg(e.Error())
	}
	c.Data["json"] = controllers.SuccessData("操作成功")
	c.ServeJSON()
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
		c.Data["json"] = controllers.ErrMsg(e.Error())
	}
	c.Data["json"] = controllers.SuccessData("操作成功")
	c.ServeJSON()
}

// @Title 菜单构建
// @Description 菜单构建
// @Success 200 {object} controllers.Result
// @router /build [get]
func (c *MenuController) Build() {
	uid, _:= jwt.GetAdminUserId(c.Ctx.Input)
	menus := models.BuildMenus(uid)
	c.Data["json"] = controllers.SuccessData(menus)
	c.ServeJSON()
}

// @Title 菜单树形
// @Description 菜单树形
// @Success 200 {object} controllers.Result
// @router /tree [get]
func (c *MenuController) GetTree() {
	menus := models.GetAllMenus("")
	c.Data["json"] = controllers.SuccessData(menus)
	c.ServeJSON()
}


