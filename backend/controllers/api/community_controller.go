package api

import (
	"BBS/backend/services"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/web"
)

type CommunityController struct {
	ctx iris.Context
}

// @Summary 获取社区列表接口
// @Description 获取社区列表
// @Tags　社区相关接口
// @Accept application/json
// @Produce application/json
// @Param object query model.Community false "查询参数"
// @Security ApiKeyAut
//@Success 200 {object} ResponseData
// @Router /community [get]
func (c *CommunityController) Get(ctx iris.Context) {

	data, err := services.CommunityService.GetCommunityList()
	if len(data) == 0 {
		ctx.JSON("社区列表获取失败")
	}
	if err != nil {
		web.JsonErrorMsg("获取列表失败")
	}
	ctx.JSON(data)
}

//func (c *CommunityController) GetBy(id string, ctx iris.Context) {
//	ID, err := ctx.Params().GetInt64(id)
//	if err != nil {
//		ctx.JSON("参数获取失败")
//	}
//	data, errData := services.CommunityService.GetCommunityDetailById(ID)
//	if errData != nil {
//		web.JsonErrorMsg("无法获取社区详细信息")
//	}
//	ctx.JSON(data)
//
//}
