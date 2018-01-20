package base

import (
	"errors"
	"strings"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/stworker/util"
)

// model公共类
type BaseModel struct {
	masterOrm map[string]*xorm.Engine // 主库orm链接
	slaveOrm  map[string]*xorm.Engine // 从库orm链接
}

/**
 * 获取主库实例
 * @param groupName string  	// 分组名称
 * @return engine *xorm.Engin, err error // xorm对象， 错误信息
 */
func (this *BaseModel) GetMaster(groupName string) (engine *xorm.Engine, err error) {
	return this.GetOrmEngine(groupName, true)
}

/**
 * 获取从库实例
 * @param groupName string  	// 分组名称
 * @return engine *xorm.Engin, err error // xorm对象， 错误信息
 */
func (this *BaseModel) GetSlave(groupName string) (engine *xorm.Engine, err error) {
	return this.GetOrmEngine(groupName, false)
}

/**
 * 获取数据库实例
 * @param groupName string  	// 分组名称
 * @return engine *xorm.Engin, err error // xorm对象， 错误信息
 */
func (this *BaseModel) GetOrmEngine(groupName string, isMaster bool) (engine *xorm.Engine, err error) {
	// 链接实例
	if isMaster {
		engine = this.masterOrm[groupName]
	} else {
		engine = this.slaveOrm[groupName]
	}

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
func (this *BaseModel) InitOrm(groups string) (err error) {
	groupsArr := strings.Split(groups, ",")
	this.masterOrm = make(map[string]*xorm.Engine)
	this.slaveOrm = make(map[string]*xorm.Engine)

	for _, groupName := range groupsArr {
		// 链接主库
		mDsn := beego.AppConfig.String(groupName + ".master.dsn")
		if len(mDsn) == 0 {
			return errors.New(groupName + "配置不正确")
		}
		mEngine, err := xorm.NewEngine("mysql", mDsn)
		if err != nil {
			return err
		}
		this.masterOrm[groupName] = mEngine

		// 链接Slave
		sDsns := strings.Split(beego.AppConfig.String(groupName+".slave.dsn"), ",")
		if len(sDsns) == 0 {
			return errors.New(groupName + "配置不正确")
		}
		sdsn := sDsns[util.GetRandNum(len(sDsns))]
		sEngine, err := xorm.NewEngine("mysql", sdsn)
		if err != nil {
			return err
		}
		this.slaveOrm[groupName] = sEngine
	}

	return nil
}

// 计算分页位置
func (this *BaseModel) GetLimit(page, pageSize int) int {
	return (page - 1) * pageSize
}
