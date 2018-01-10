package routers

import (
	"github.com/stworker/controllers/api"

	"github.com/astaxie/beego"
)

func init() {
	// api相关
	apiControllerRouter()
}

// api相关路由
func apiControllerRouter() {
	// 活动相关控制器
	beego.Router("api/activity/aditemlist", &api.ActivityController{}, "get:AdItemList")
}
