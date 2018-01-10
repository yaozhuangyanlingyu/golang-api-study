package activitymodule

import (
	"github.com/stworker/models"
)

/**
 * Created by yaoxf on 2018/01/09.
 * @description 活动相关模块数据
 */
type ActivityModule struct {
}

/**
 * Created by yaoxf on 2018/01/09.
 * @param page int 		// 分页页数
 * @param pageSize int 	// 分页数据条数
 * @return []ActiveAdItemsModel (数据数组), error(错误信息)
 * @description 获取推广ID列表
 */
func (this *ActivityModule) GetAdItemList(page, pageSize int) (list []models.ActiveAdItemsModel, count int64, err error) {
	m := models.ActiveAdItemsModel{}
	return m.GetAdItemList(page, pageSize)
}
