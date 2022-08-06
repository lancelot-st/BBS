package api

import (
	"BBS/backend/services"
	"BBS/model"
	"BBS/pkg/jwt"
	"BBS/pkg/snowflake"
	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/web"
	"log"
	"strconv"
	"strings"
	"time"
)

type ArticleController struct {
	ctx iris.Context
}

// @Summary 发布帖子接口
// @Description 创建帖子
// @Tags　帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer token"
// @Param object query model.Article false "查询参数"
// @Security ApiKeyAut
// @Success 200 {object} ResponseData
// @Router /posts [post]
func (a *ArticleController) Post(ctx iris.Context) {
	//获取参数

	postID := snowflake.GenID() //雪花算法随机生成id

	AuthorID, errID := GetCurrentUserID(ctx) //获取当前的用户id即是当前帖子的作者ID
	if errID != nil {
		log.Println("获取信息失败")
		ctx.JSON("用户id获取失败")
		return
	}

	p := &model.Article{
		AuthorID:   AuthorID,
		PostID:     postID,
		CreateTime: time.Now(),
	}
	err := ctx.ReadJSON(p)
	if err != nil {
		ctx.JSON("请求获取失败")
	}

	//创建帖子
	err = services.ArticleService.CreateArticleServ(p) //创建帖子
	if err != nil {
		log.Println("传入数据库失败")
		return
	}
	//返回响应
	ctx.JSON("发帖成功")

}

func (a *ArticleController) GetBy(id int64, ctx iris.Context) {

	//获取帖子的id
	id, errID := a.ctx.Params().GetInt64("id")
	if errID != nil {
		web.JsonErrorMsg("帖子id获取失败")
	}

	//根据id取出帖子的数据
	data, errData := services.ArticleService.GetArticleByIDServ(id)
	if errData != nil {
		web.JsonErrorMsg("获取帖子内容失败")
	}
	//返回响应
	ctx.JSON(data)
}

// @Summary 获取帖子列表接口
// @Description 按帖子发布的时间的从新到旧排列获取帖子列表
// @Tags　帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param object query model.ArticleList false "查询参数"
// @Security ApiKeyAut
// @Success 200 {object} ResponseData
// @Router /posts/list [get]
func (a *ArticleController) GetList(ctx iris.Context) {
	page, size := GetPageInfo(ctx) //获取页码，及每页所需要展示的帖子数

	data, err := services.ArticleService.GetArticleListServ(page, size)
	if err != nil {
		web.JsonErrorMsg("获取失败")
	}
	ctx.JSON(data)
}

// @Summary 获取帖子列表接口
// @Description 按分数的大小从大到小排列获取帖子列表
// @Tags　帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param object query model.ArticleList false "查询参数"
// @Security ApiKeyAut
//@Success 200 {object} ResponseData
// @Router /posts/list2 [get]
func (a *ArticleController) GetList2(ctx iris.Context) {
	page, size := GetPageInfo(ctx)
	p := &model.ArticleList{
		Page:  page,
		Size:  size,
		Order: "Score",
	}
	data, err := services.ArticleService.GetArticleListServ2(p)
	if err != nil {
		web.JsonErrorMsg("获取失败")
	}
	ctx.JSON(data)
}

func GetPageInfo(ctx iris.Context) (page int64, size int64) {
	Page := ctx.URLParam("page")
	limit := ctx.URLParam("size")

	var (
		offset   int64
		limitset int64
		err      error
	)
	offset, err = strconv.ParseInt(Page, 10, 64)
	if err != nil {
		offset = 0
	}
	limitset, err = strconv.ParseInt(limit, 10, 64)
	if err != nil {
		limitset = 2
	}
	return offset, limitset
}

func GetCurrentUserID(ctx iris.Context) (int64, error) {
	token := ctx.Request().Header.Get("Authorization")
	kv := strings.Split(token, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		ctx.JSON("token获取失败")

	}

	tokenString := kv[1]
	jwtToken, err := jwt.ParseToken(tokenString)
	if err != nil {
		ctx.JSON("token解析失败")
	}
	userID := jwtToken.UserID
	return userID, nil
}
