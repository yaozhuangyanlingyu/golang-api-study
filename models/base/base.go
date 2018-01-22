package base

import (
	"errors"
	"os"
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
 * 初始化db连接
 * @param group string  // 分组名称
 * @return void
 */
func (this *BaseModel) InitDbConn(groupName string) (err error) {
	this.masterOrm = make(map[string]*xorm.Engine)
	this.slaveOrm = make(map[string]*xorm.Engine)

	// 链接主库
	mDsn := beego.AppConfig.String(groupName + ".master.dsn")
	if len(mDsn) == 0 {
		return errors.New(groupName + "配置不正确")
	}
	mEngine, err := this.InitOrm(mDsn)
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
	sEngine, err := this.InitOrm(sdsn)
	if err != nil {
		return err
	}
	this.slaveOrm[groupName] = sEngine

	return nil
}

/**
 * 初始化orm
 * @param dsn string  // 数据库DSN字符串
 * @return orm
 */
func (this *BaseModel) InitOrm(dsn string) (orm *xorm.Engine, err error) {
	// 创建orm链接
	orm, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		return orm, err
	}

	// 输出日志
	orm.ShowSQL(false)
	os.MkdirAll("./logs", os.ModePerm)
	f, err := os.OpenFile("logs/sql.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0664)
	if err != nil {
		return orm, err
	}
	orm.SetLogger(xorm.NewSimpleLogger(f))
	orm.Logger().SetLevel(core.LOG_INFO)

	return orm, nil
}

// 计算分页位置
func (this *BaseModel) GetLimit(page, pageSize int) int {
	return (page - 1) * pageSize
}
