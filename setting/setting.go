package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() error {
	viper.SetConfigFile("./config/config.yaml")
	errRead := viper.ReadInConfig()
	if errRead != nil {
		panic("Viper Cannot Find Config")
	}
	viper.WatchConfig() //监听config文件配置如果发生变化调用
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("")
	})
	//flag.Parse()
	//conf = config.Init(*configFile)
	//return con
	return errRead
}
