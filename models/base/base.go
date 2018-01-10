package base

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

// model公共类
type BaseModel struct {
	masterOrm map[string]interface{} // 主库orm链接
	slaveOrm  map[string]interface{} // 从库orm链接
}

// 获取数据库实例
func (this *BaseModel) GetSlave() (engine *xorm.Engine, err error) {
	// 链接实例
	engine, err = xorm.NewEngine("mysql", "yaoxf:123456@(192.168.134.128:3306)/dbshop?charset=utf8")

	// 显示SQL
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)

	return engine, err
}

/**
 * 初始化orm
 * @param group string  // 分组名称
 * @return void
 */
func (this *BaseModel) InitOrm(group string) (err error) {
	// 初始化orm
	this.masterOrm = make(map[string]interface{})

	// 链接主库
	masterDns := beego.AppConfig.String("dbshop.master.dsn")
	mOrm, err := xorm.NewEngine("mysql", masterDns)
	if err != nil {
		return err
	}
	this.masterOrm[group] = mOrm

	// 链接Slave
	test := strings.Split(beego.AppConfig.String("dbshop.slave.dsn"), ",")
	fmt.Println(len(test))
	/*
		this.slaveOrm = make(map[string]interface{})
		masterDns := beego.AppConfig.String("dbshop.master.dsn")
		mOrm, err := xorm.NewEngine("mysql", masterDns)
		if err != nil {
			return err
		}*/

	return nil
}

// 计算分页位置
func (this *BaseModel) GetLimit(page, pageSize int) int {
	return (page - 1) * pageSize
}
