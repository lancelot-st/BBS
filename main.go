package main

import (
	"BBS/model"
	"BBS/model/redis"
	"BBS/pkg/snowflake"
	"BBS/routes"
	"BBS/setting"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

//@title BBS
//@version 1.0
//@description 采用redit论坛的算法进行简化版制作的论坛
//@contact.name XZ
//@contact.url https://lancelot-st.github.io
//@contact.email XZupup4088@163.com
//@license.name Apache 2.0
//@license.url http://www.apache.org/licenses/LICENSE-2.0.html
//@host 127.0.0.1:8084
//@BasePath /api
func main() {
	//初始话配置
	errViper := setting.Init()
	if errViper != nil {
		log.Fatal("viper 读取配置失败")
	}
	//雪花算法id初始化
	snowflake.Init(1)

	//初始话gorm连接
	err := model.NewMySqlConn()
	if err != nil {
		log.Fatal("gorm 连接失败")
	}
	//初始话redis连接
	err = redis.Init()
	if err != nil {
		log.Fatalf("reids 读取配置失败")
	}
	//关闭redis连接
	defer redis.Close()
	//初始化雪花算法
	err = snowflake.Init(1)
	if err != nil {
		log.Fatalf("雪花算法初始化失败")
	}
	//注册路由
	routes.Router()

}
