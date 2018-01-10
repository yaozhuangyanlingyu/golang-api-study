package api

import (
	"github.com/stworker/controllers"
	"github.com/stworker/modules/apimodule/activitymodule"
)

// 活动相关控制器
type ActivityController struct {
	controllers.CommonController
}

/**
 * 获取推广ID列表接口
 * @param page int 	// 分页页数
 * @param page_size	// 分页大小
 * @return json
 */
func (this *ActivityController) AdItemList() {
	// 初始化参数
	page, err := this.GetInt("page", 1)
	if err != nil {
		this.AjaxMsg("page参数不正确", controllers.MSG_ERR)
	}
	pageSize, err := this.GetInt("page_size", 10)
	if err != nil {
		this.AjaxMsg("page_size参数不正确", controllers.MSG_ERR)
	}

	// 查询数据
	m := activitymodule.ActivityModule{}
	data, count, err := m.GetAdItemList(page, pageSize)
	if err != nil {
		this.AjaxMsg(err.Error(), controllers.MSG_ERR)
	}

	// 返回数据
	pages := make(map[string]interface{})
	pages["count"] = count
	pages["page"] = page
	pages["pageSize"] = pageSize
	this.AjaxList("success", controllers.MSG_OK, data, pages)
}
