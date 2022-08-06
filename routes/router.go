package routes

import (
	"BBS/backend/controllers/api"
	"BBS/backend/services"
	_ "BBS/docs" //这个不要忘了
	"BBS/pkg/jwt"
	"github.com/go-playground/validator/v10"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/iris-contrib/swagger/v12"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/mvc"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}

func Router() {
	//iris对象的创建
	app := iris.New()
	//设置logger模式为debug
	app.Logger().SetLevel("Debug")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})
	app.WrapRouter(c.ServeHTTP)
	app.Use(logger.New())
	//全局使用validate
	app.Validator = validator.New()
	//注册模板
	tmpl := iris.HTML("./templates", ".html") // 加载html
	app.HandleDir("static", "./static")       // 加载静态文件
	app.RegisterView(tmpl)

	app.Get("/", func(context iris.Context) {
		context.View("index.html")
	})
	//tmpl := iris.HTML("./frontend/view/html/zh", ".html") // 加载html
	//app.HandleDir("static", "./frontend/view")            // 加载静态文件
	//app.RegisterView(tmpl)
	//app.Get("/", func(context iris.Context) {
	//	context.View("index.html")
	//})

	// swagger配置方法一：其他文档
	//config := swagger.Config{
	//	// 指向swagger init生成文档的路径
	//	URL:          "http://www.xxx.com/swagger/doc.json",
	//	DeepLinking:  true,
	//
	//}
	//app.Get("/swagger/*any", swagger.CustomWrapHandler(&config,swaggerFiles.Handler))
	// swagger配置方法二：默认文档
	//app.Get("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	//采用mvc架构模式注册路由
	mvc.Configure(app.Party("/api"), func(context *mvc.Application) {
		app.Get("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
		context.Party("/register").Handle(new(api.RegisterController))   //注册模块
		context.Party("/login").Handle(new(api.LoginController))         //登录模块
		context.Party("/community").Handle(new(api.CommunityController)) //社区列表显示模块
		context.Party("/posts").Handle(new(api.ArticleController))       //帖子显示模块
		context.Router.Use(jwt.JWTAuthMiddleware().Serve)                //这个鉴权中间件放在需要绑定的路由前面没有经过鉴权就无法访问后面的路由
		app.Use(jwt.JWTAuthMiddleware().Serve)

		context.Party("/vote").Handle(new(api.VoteController)) //投票模块

		app.Get("/api/community/{id:string}", func(context iris.Context) {
			ID, _ := context.Params().GetInt64("id")
			data, _ := services.CommunityService.GetCommunityDetailById(ID)
			context.JSON(data)
		})
	})

	//app.Get("/community/{id:string}", func(context iris.Context) {
	//	ID, _ := context.Params().GetInt64("id")
	//	data, _ := services.CommunityService.GetCommunityDetailById(ID)
	//	context.JSON(data)
	//})

	//程序运行
	addr := "localhost:" + viper.GetString("app.Port")
	app.Run(iris.Addr(addr),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
