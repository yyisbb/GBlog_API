package main

import (
	"gblog_api/cmd"
	"gblog_api/config"
)

func main() {
	//注册配置中心并启动
	cmd.Run(config.RegisterConfigCenter)
}
