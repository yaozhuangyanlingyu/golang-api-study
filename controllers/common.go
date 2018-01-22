package controllers

import (
	"fmt"
	"math"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

const (
	MSG_OK  = 0
	MSG_ERR = 1001
)

// 公共控制器
type CommonController struct {
	beego.Controller
}

// 初始化
func (this *CommonController) Prepare() {
	// 记录日志
	this.WriteRequestLog()
}

// 记录请求日志
func (this *CommonController) WriteRequestLog() {
	logs.Trace(fmt.Sprintf("%s %s Header: %s Body: %s", this.Ctx.Request.Method, this.Ctx.Request.URL, this.Ctx.Request.Header, this.Ctx.Request.Body))
}

// AJAX返回提示信息
func (this *CommonController) AjaxMsg(msg interface{}, code int) {
	out := make(map[string]interface{})
	out["msg"] = msg
	out["code"] = code
	out["data"] = make(map[string]interface{})
	out["count"] = 0
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

// Ajax返回列表数据
func (this *CommonController) AjaxList(msg interface{}, code int, data interface{}, pages map[string]interface{}) {
	out := make(map[string]interface{})
	out["msg"] = msg
	out["code"] = code
	out["data"] = data
	tmpCount := float64(pages["count"].(int64))
	tmpPageSize := float64(pages["pageSize"].(int))
	pages["pageTotal"] = math.Ceil(tmpCount / tmpPageSize)
	out["pages"] = pages
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}
