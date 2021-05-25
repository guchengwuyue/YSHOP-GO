package admin

import (
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"time"
	"yixiang.co/yshop/common/jwt"
	"yixiang.co/yshop/common/untils"
	"yixiang.co/yshop/controllers"
	"yixiang.co/yshop/dto"
	"yixiang.co/yshop/models"
	"yixiang.co/yshop/vo"
)

// 登录api
type LoginController struct {
	controllers.BaseController
}

type CaptchaResult struct {
	Id          string `json:"id"`
	Base64Blob  string `json:"base_64_blob"`
	VerifyValue string `json:"code"`
}

// 设置自带的store
var store = base64Captcha.DefaultMemStore


func (c *LoginController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Captcha",c.Captcha)
	c.Mapping("Info",c.Info)
	c.Mapping("Logout",c.Logout)
}

// @Title 登录
// @Description 登录
// @Success 200 {object} controllers.Result
// @router /login [post]
func (c *LoginController) Login() {
	var authUser *dto.AuthUser

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &authUser)

	logs.Info(authUser)

	if err == nil {
		currentUser, e := models.GetUserByUsername(authUser.Username)
		//校验验证码
		if !store.Verify(authUser.Id, authUser.Code, true) {
			c.Fail("验证码不对",5001)
		}
		if e != nil {
			c.Fail("用户不存在",5002)
		}
		logs.Info("=======currentUser======")
		logs.Info(currentUser)
		if !untils.ComparePwd(currentUser.Password,[]byte(authUser.Password)) {
			c.Fail("密码错误",5003)
		}else{
			token,_ := jwt.GenerateToken(currentUser,time.Hour*24*100)
			var loginVO = new(vo.LoginVo)
			loginVO.Token = token
			loginVO.User = currentUser
			c.Ok(loginVO)
		}
	}else {
		c.Fail(err.Error(),5004)
	}

}

// @Title 获取用户信息
// @Description 获取用户信息
// @Success 200 {object} controllers.Result
// @router /info [get]
func (c *LoginController) Info() {
	c.Data["json"] = controllers.SuccessData(jwt.GetAdminDetailUser(c.Ctx.Input))
	c.ServeJSON()
}

// @Title 退出登录
// @Description 退出登录
// @Success 200 {object} controllers.Result
// @router /logout [delete]
func (c *LoginController) Logout() {
	err := jwt.RemoveUser(c.Ctx.Input)
	if err != nil {
		c.Fail("退出失败",5005)
	}else{
		c.Ok("退出成功")
	}
}

// @Title 获取验证码
// @Description 获取验证码
// @router /captcha [get]
func (c *LoginController) Captcha(){
	GenerateCaptcha(c.Ctx)
	c.ServeJSON()
}



// 生成图形化验证码  ctx *context.Context
func GenerateCaptcha(ctx *context.Context) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverMath

	// 配置验证码信息
	captchaConfig := base64Captcha.DriverMath{
		Height:          38,
		Width:           110,
		NoiseCount:      0,
		ShowLineOptions: 0,
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	// 自定义配置，如果不需要自定义配置，则上面的结构体和下面这行代码不用写
	driverString = captchaConfig
	driver = driverString.ConvertFonts()

	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		logs.Error(err.Error())
	}
	captchaResult := CaptchaResult{
		Id:         id,
		Base64Blob: b64s,
	}

	ctx.Output.JSON(controllers.SuccessData(captchaResult),true,true)
}