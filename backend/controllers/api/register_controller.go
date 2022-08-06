package api

import (
	"BBS/backend/services"
	"BBS/model"
	"github.com/kataras/iris/v12"
)

type RegisterController struct {
	ctx iris.Context
}

// @Summary 用户注册接口
// @Description 用户通过填写用户名密码邮箱进行注册
// @Tags　注册相关接口
// @Accept multipart/form-data
// @Produce application/json
// @Param object query model.User false "查询参数"
// @Security ApiKeyAut
//@Success 200 {object} ResponseData
// @Router /register [post]
func (r *RegisterController) Post(ctx iris.Context) {
	userName := ctx.PostValue("Name")
	userPwd := ctx.PostValue("Password")
	userEmail := ctx.PostValue("Email")

	// 判断err是不是 validator.ValidationErrors类型的errors
	user := &model.User{
		Username: userName,
		Password: userPwd,
		Email:    userEmail,
	}
	//validate := validator.New()
	//err := validate.Struct(&user)
	//if err != nil {
	//	for _, err := range err.(validator.ValidationErrors) {
	//		log.Println(err)
	//		ctx.JSON("用户格式不满足")
	//		return
	//	}
	//}
	isOK := services.UserService.Register(user.Username, user.Password, user.Email)

	if !isOK {
		ctx.JSON("注册失败")
		return
	}
	ctx.JSON("注册成功")
}
