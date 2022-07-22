package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"operate-mongodb/common"
)

func InitConf() {
	config := common.YamlConfigFile
	fmt.Printf("您正在使用config的默认值,config的路径为%v\n", config)
	v := viper.New()
	//设置配置文件路径
	v.SetConfigFile(config)
	//从磁盘中加载配置文件,并且还加载 key/value存储
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error conf file: %s \n", err))
	}
	// 监控配置文件变化并热加载程序
	v.WatchConfig()
	//配置文件发生变更之后会调用的回调函数
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("conf file changed:", e.Name)
		if err = v.Unmarshal(&common.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&common.CONFIG); err != nil {
		fmt.Println(err)
	}
	//赋给全局变量
	common.VP = v
}
