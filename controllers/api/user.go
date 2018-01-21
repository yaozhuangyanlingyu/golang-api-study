package api

import (
	"fmt"

	"strconv"

	"github.com/stworker/controllers"
	"github.com/stworker/modules/apimodule/usermodule"
)

// 用户相关控制器
type UserController struct {
	controllers.CommonController
}

/**
 * 根据用户ID获取用户信息
 * @param id int	// 用户ID
 * @return json
 */
func (this *UserController) GetUserById() {
	// 初始化参数
	tmpid := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(tmpid)
	if err != nil {
		this.AjaxMsg(err.Error(), controllers.MSG_ERR)
	}

	// 查询数据
	m := usermodule.UserModule{}
	data, err := m.GetUserInfoById(id)
	if err != nil {
		this.AjaxMsg(err.Error(), controllers.MSG_ERR)
	}
	fmt.Println(data)

}
