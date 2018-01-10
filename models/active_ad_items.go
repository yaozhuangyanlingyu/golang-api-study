package models

import (
	"github.com/stworker/models/base"
)

/**
 * Created by yaoxf on 2018/01/09.
 * @description active_ad_items表
 */
type ActiveAdItemsModel struct {
	Id         int64
	AdId       string `xorm:"char(32) NOT NULL DEFAULT ''"`
	AdPrice    int64  `xorm:"bigint(20) NOT NULL DEFAULT '0'"`
	Status     int8   `xorm:"tinyint(4) NOT NULL DEFAULT '0'"`
	CreateTime int64  `xorm:"int(11) NOT NULL DEFAULT '0'"`
	ModifyTime int64  `xorm:"int(11) NOT NULL DEFAULT '0'"`
}

/**
 * Created by yaoxf on 2018/01/09.
 * @param page int 		// 分页页数
 * @param pageSize int 	// 分页数据条数
 * @return []ActiveAdItemsModel (数据数组), error(错误信息)
 * @description 获取推广ID列表
 */
func (this *ActiveAdItemsModel) GetAdItemList(page, pageSize int) (list []ActiveAdItemsModel, count int64, err error) {
	// 获取数据库链接
	x, _ := base.DbShop.GetSlave()

	// 查询数据
	list = make([]ActiveAdItemsModel, 0)
	start := base.DbShop.GetLimit(page, pageSize)
	if err = x.Table("dbshop.cc_active_ad_items").Where("id > ?", 0).Desc("id").Limit(pageSize, start).Find(&list); err != nil {
		return list, count, err
	}

	// 统计条数
	count, err = x.Table("dbshop.cc_active_ad_items").Where("id > ?", 0).Count()
	if err != nil {
		return list, 0, err
	}

	return list, count, nil
}
