package main

import (
	"blog/core"
	"blog/flag"
	"blog/global"
	"blog/routers"
)

func main() {
	//读取配置文件
	core.InitConf()
	global.Log = core.InitLogger()
	//global.Log.Warnln("111")
	//global.Log.Error("111")
	//global.Log.Infof("111")
	global.DB = core.InitGorm()
	//fmt.Println(global.DB)

	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("运行在%s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
