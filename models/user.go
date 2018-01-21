package models

import (
	"github.com/stworker/models/base"
)

/**
 * Created by yaoxf on 2018/01/20.
 * @description user表
 */
type UserModel struct {
	Id         int64
	LoginName  string `xorm:"varchar(20) NOT NULL DEFAULT ''"`
	RealName   string `xorm:"varchar(20) NOT NULL DEFAULT ''"`
	Password   string `xorm:"char(32) NOT NULL DEFAULT ''"`
	Email      string `xorm:"varchar(50) NOT NULL DEFAULT ''"`
	Status     int    `xorm:"tinyint(1) NOT NULL DEFAULT 0"`
	CreateTime int64  `xorm:int unsigned NOT NULL DEFAULT 0`
	UpdateTime int64  `xorm:int unsigned NOT NULL DEFAULT 0`
}

/**
 * 根据用户信息获取用户信息
 * @param id int	// 用户ID
 * @return json
 */
func (this *UserModel) GetUserInfoById(id int) (info UserModel, err error) {
	// 获取数据库链接
	x, err := base.DbUser.GetSlave(base.DbUser.GroupName)
	if err != nil {
		return info, err
	}

	// 查询数据
	x.Table("cc_user").Where("id = ?", id).Get(&info)
	return info, err
}
