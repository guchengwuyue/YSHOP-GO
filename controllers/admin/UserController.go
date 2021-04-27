package admin

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	"yixiang.co/yshop/controllers"
	"yixiang.co/yshop/dto"
	"yixiang.co/yshop/jwt"
	"yixiang.co/yshop/models"
	"yixiang.co/yshop/untils"
	"yixiang.co/yshop/vo"
)

// 用户 API
type UserController struct {
	controllers.BaseController
}

func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title 用户列表
// @Description 用户列表
// @Success 200 {object} controllers.Result
// @router / [get]
func (c *UserController) GetAll() {
	deptId, _ := c.GetInt64("deptId",-1)
	enabled, _ := c.GetInt64("enabled",-1)
	total,list := models.GetAllUser(c.GetParams(),deptId,enabled)
	c.Data["json"] = controllers.SuccessData(vo.ResultList{Content: list,TotalElements: total})
	c.ServeJSON()
}

// @Title 用户添加
// @Description 用户添加
// @Success 200 {object} controllers.Result
// @router / [post]
func (c *UserController) Post()  {
	var model models.SysUser
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	b, _ := valid.Valid(&model)
	if !b {
		for _, err := range valid.Errors {
			c.Data["json"] = controllers.ErrMsg(err.Message)
		}
	}
	_, e := models.AddUser(&model)
	if e != nil {
		c.Data["json"] = controllers.ErrMsg(e.Error())
	}
	c.Data["json"] = controllers.SuccessData("操作成功")
	c.ServeJSON()
}

// @Title 用户编辑
// @Description 用户编辑
// @Success 200 {object} controllers.Result
// @router / [put]
func (c *UserController) Put()  {
	var model models.SysUser
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	logs.Info("======start======")
	logs.Info(model)
	b, _ := valid.Valid(&model)
	if !b {
		for _, err := range valid.Errors {
			c.Data["json"] = controllers.ErrMsg(err.Message)
		}
	}
	e := models.UpdateByUser(&model)
	if e != nil {
		c.Data["json"] = controllers.ErrMsg(e.Error())
	}
	c.Data["json"] = controllers.SuccessData("操作成功")
	c.ServeJSON()
}

// @Title 用户删除
// @Description 用户删除
// @Success 200 {object} controllers.Result
// @router / [delete]
func (c *UserController) Delete() {
	var ids []int64
	json.Unmarshal(c.Ctx.Input.RequestBody, &ids)
	e := models.DelByUser(ids)
	if e != nil {
		c.Data["json"] = controllers.ErrMsg(e.Error())
	}
	c.Data["json"] = controllers.SuccessData("操作成功")
	c.ServeJSON()
}

// @Title 用户上传图像
// @Description 用户上传图像
// @Success 200 {object} controllers.Result
// @router /updateAvatar [post]
func (c *UserController) Avatar()  {
	logs.Info("======file start======")
	f, h, err := c.GetFile("file")
	if err != nil {
		logs.Error(err)
	}
	defer f.Close()
	var path = "static/upload/" + h.Filename
	e := c.SaveToFile("file", path) // 保存位置在 static/upload, 没有文件夹要先创建
	logs.Error(e)
	apiUrl, _ := beego.AppConfig.String("api_url")
	avatarUrl := apiUrl + "/" +path

	//save user
	uid, _ := jwt.GetAdminUserId(c.Ctx.Input)
	user, _ := models.GetUserById(uid)
	if user == nil {
		c.Data["json"] = controllers.ErrMsg("非法操作")
	}else {
		user.Avatar = avatarUrl
		models.UpdateCurrentUser(user)
		c.Data["json"] = controllers.SuccessData("ok")
	}
	c.ServeJSON()
}

// @Title 用户修改密码
// @Description 用户修改密码
// @Success 200 {object} controllers.Result
// @router /updatePass [post]
func (c *UserController) Pass()  {
	var model dto.UserPass
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	b, _ := valid.Valid(&model)
	if !b {
		for _, err := range valid.Errors {
			c.Data["json"] = controllers.ErrMsg(err.Message)
		}
	}
	//save user
	uid, _ := jwt.GetAdminUserId(c.Ctx.Input)
	user, _ := models.GetUserById(uid)
	if user == nil {
		c.Data["json"] = controllers.ErrMsg("非法操作")
	}else {
		if !untils.ComparePwd(user.Password,[]byte(model.OldPass)) {
			c.Data["json"] = controllers.ErrMsg("旧密码错误密码错误")
			c.ServeJSON()
		}
		user.Password = untils.HashAndSalt([]byte(model.NewPass))
		models.UpdateCurrentUser(user)
		c.Data["json"] = controllers.SuccessData("ok")
	}
	c.ServeJSON()
}

// @Title 用户修改个人信息
// @Description 用户修改个人信息
// @Success 200 {object} controllers.Result
// @router /center [put]
func (c *UserController) Center()  {
	var model dto.UserPost
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &model)
	b, _ := valid.Valid(&model)
	if !b {
		for _, err := range valid.Errors {
			c.Data["json"] = controllers.ErrMsg(err.Message)
		}
		c.ServeJSON()
	}
	//save user
	uid, _ := jwt.GetAdminUserId(c.Ctx.Input)
	user, _ := models.GetUserById(uid)
	if user == nil {
		c.Data["json"] = controllers.ErrMsg("非法操作")
	}else {
		user.Phone = model.Phone
		user.Sex = model.Sex
		user.NickName = model.NickName
		models.UpdateCurrentUser(user)
		c.Data["json"] = controllers.SuccessData("ok")
	}
	c.ServeJSON()
}