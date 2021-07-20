package admin

import (
	"encoding/json"
	"yixiang.co/yshop/controllers"
	"yixiang.co/yshop/models"
	"yixiang.co/yshop/models/vo"
)

// 字典api
type DictController struct {
	controllers.BaseController
}

func (c *DictController) URLMapping() {
	c.Mapping("Post", c.Post)
	//c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	//c.Mapping("Put", c.Put)
	//c.Mapping("Delete", c.Delete)
}

// @Title 获取字典列表
// @Description 获取字典列表
// @Success 200 {object} controllers.Result
// @router / [get]
func (c *DictController) GetAll() {
	total,list := models.GetAllDict(c.GetParams())
	c.Ok(vo.ResultList{Content: list,TotalElements: total})
}

// @Title 添加字典
// @Description 添加字典
// @Success 200 {object} controllers.Result
// @router / [post]
func (c *DictController) Post()  {
	var dictModel models.SysDict
	json.Unmarshal(c.Ctx.Input.RequestBody, &dictModel)
	c.Valid(&dictModel)
	_, e := models.AddDict(&dictModel)
	if e != nil {
		c.Fail(e.Error(),5002)
	}
	c.Ok("操作成功")
}

// @Title 修改字典
// @Description 修改字典
// @Success 200 {object} controllers.Result
// @router / [put]
func (c *DictController) Put()  {
	var dictModel models.SysDict
	json.Unmarshal(c.Ctx.Input.RequestBody, &dictModel)
	c.Valid(&dictModel)
	e := models.UpdateByDict(&dictModel)
	if e != nil {
		c.Fail(e.Error(),5004)
	}
	c.Ok("操作成功")
}

// @Title 删除字典
// @Description 删除字典
// @Success 200 {object} controllers.Result
// @router /:id [delete]
func (c *DictController) Delete() {
	id, _ := c.GetInt64(":id",1)
	e := models.DelByDict(id)
	if e != nil {
		c.Fail(e.Error(),5005)
	}
	c.Ok("操作成功")
}
