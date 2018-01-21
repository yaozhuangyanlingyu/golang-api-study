package usermodule

import (
	"github.com/stworker/models"
)

/**
 * Created by yaoxf on 2018/01/20.
 * @description 用户相关模块数据
 */
type UserModule struct {
}

/**
 * Created by yaoxf on 2018/01/20.
 * @param page int 		// 分页页数
 * @param pageSize int 	// 分页数据条数
 * @return UserModel, error(错误信息)
 * @description 根据用户ID获取用户信息
 */
func (this *UserModule) GetUserInfoById(id int) (userInfo models.UserModel, err error) {
	m := models.UserModel{}
	return m.GetUserInfoById(id)
}
