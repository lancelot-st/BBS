package api

import (
	"BBS/backend/repositories"
	"BBS/backend/services"
	"BBS/model"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/web"
	"log"
)

type VoteController struct {
	ctx iris.Context
}

// @Summary 用户投票接口
// @Description 用户为喜爱或者讨厌的帖子进行投票有赞成票、反对票和放弃
// @Tags　投票相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer token"
// @Security ApiKeyAut
//@Success 200 {object} ResponseData
// @Router /vote [post]
func (v *VoteController) Post(ctx iris.Context) {

	//参数校验
	p := new(model.VoteData)
	err := ctx.ReadJSON(p)
	if err != nil {
		ctx.JSON(p)
		return
	}
	//post_id := ctx.URLParam("post_id")
	//PostId, err := strconv.ParseInt(post_id, 10, 64)
	//if err != nil {
	//	log.Println("参数获取失败")
	//}
	//
	//p.PostID = PostId
	//direction := ctx.URLParam("direction")
	//Direction, errD := strconv.ParseInt(direction, 10, 64)
	//if errD != nil {
	//	log.Println("参数获取失败")
	//}
	//p.Direction = int(Direction)
	valid := validator.New()
	errValid := valid.Struct(p)
	if errValid != nil {
		for _, err := range errValid.(validator.ValidationErrors) {
			log.Println(err)
			web.JsonErrorMsg(err.Error())
		}
	}
	//获取userID
	userID, err := repositories.UserRepository.GetCurrentUserID(ctx)
	if err != nil {
		log.Println("获取用户ID失败")
	}

	err = services.VoteService.VoteForArticle(userID, p)
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON("投票成功")
}
