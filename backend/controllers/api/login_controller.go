package api

import (
	"BBS/backend/services"
	"BBS/pkg/jwt"
	"github.com/kataras/iris/v12"
	"log"
)

type LoginController struct {
	ctx iris.Context
}

// @Summary 用户登录接口
// @Description 用户通过用户名密码进行登录
// @Tags　登录相关接口
// @Accept multipart/form-data
// @Produce application/json
// @Param object query model.User false "查询参数"
// @Security ApiKeyAut
//@Success 200 {object} ResponseData
// @Router /login [post]
func (L *LoginController) Post(ctx iris.Context) {
	//获取post的表单值
	userName := ctx.FormValue("Name")
	if len(userName) == 0 {
		log.Println("用户名为空")
		ctx.JSON("用户名不能为空")

	}
	userPwd := ctx.FormValue("Password")
	if len(userPwd) == 0 {
		log.Println("密码为空")
		ctx.JSON("用户密码不能为空")
	}

	//第一次登录校验用户名和密码是否正确
	isOK := services.UserService.IsPwdSuccess(userName, userPwd)
	if !isOK {
		ctx.JSON("登录失败")
		return
	}
	//从用户名获取完整的用用户信息
	User, err := services.UserService.GetUserByName(userName)
	if err != nil || User == nil {
		ctx.JSON("不存在该用户")
	}
	//生成token
	token, err := jwt.GenToken(User.UserID, User.Username)
	if err != nil {
		ctx.JSON("token创建失败")
	}
	//返回tokenString

	ctx.JSON(token)
}
