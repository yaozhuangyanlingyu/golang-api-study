package main

import (
	_ "github.com/stworker/routers"

	"github.com/stworker/setting"

	"github.com/astaxie/beego"
)

func main() {
	// 初始化配置
	err := setting.Init()
	if err != nil {
		panic(err)
		return
	}

	// 运行框架
	beego.Run()
}
