package main

import (
	"BBS/model"
	"BBS/model/redis"
	"BBS/pkg/snowflake"
	"BBS/setting"
	"fmt"
	"log"
)

func main() {
	//初始话配置
	errViper := setting.Init()
	if errViper != nil {
		log.Fatal("viper 读取配置失败")
	}
	err := model.NewMySqlConn()
	if err != nil {
		return
	}
	snowflake.Init(1)

	//errUser := repositories.UserRepository.CreateUser("xiaoming", "123456", "88887990@qq.com")
	//if errUser != nil {
	//	fmt.Println("用户创建失败")
	//}
	err = redis.Init()
	if err != nil {
		log.Fatal("redis 启动失败")
	}

	client := redis.GetRedis()
	//pipeline := client.TxPipeline()
	//
	//pipeline.ZAdd(redis.KeyPostTimeZSet, redis2.Z{
	//	Score:  16000004,
	//	Member: 252464758586868,
	//})
	//_, err = pipeline.Exec()
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	time := client.ZScore(redis.KeyPostTimeZSet, "20519968726585344")
	fmt.Print(time)
	//rd := redis.NewClient(&redis.Options{
	//	Addr:     "", // url
	//	Password: "",
	//	DB:       0, // 0号数据库
	//})
	//result, err := rd.Ping().Result()
	//if err != nil {
	//	fmt.Println("ping err :", err)
	//	return
	//}

	//tmpl := iris.HTML("./frontend/view/html/zh", ".html") // 加载html
	//app.HandleDir("static", "./frontend/view")            // 加载静态文件
	//app.RegisterView(tmpl)
	//app.Get("/", func(context iris.Context) {
	//	context.View("index.html")
	//})

}
