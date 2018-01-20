package setting

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/stworker/models/base"
)

// 初始化框架
func Init() (err error) {
	// 初始化日志
	err = InitLogger()
	if err != nil {
		return err
	}

	// 初始化数据库链接，多个数据库名称逗号隔开
	err = base.DbShop.InitOrm("dbshop")
	if err != nil {
		return err
	}
	//x, _ := base.DbShop.GetSlave(base.DbShop.GroupName)
	//fmt.Println(x)

	return nil
}

// 初始化日志
func InitLogger() error {
	// 配置信息
	config := make(map[string]interface{})
	config["filename"] = beego.AppConfig.String("log.path")
	config["level"] = GetLogLevel(beego.AppConfig.String("log.level"))
	config["maxlines"] = 100
	config["separate"] = []string{"error"}
	configStr, err := json.Marshal(config)
	if err != nil {
		return err
	}

	// 记录日志
	logs.Async()
	logs.EnableFuncCallDepth(true)
	logs.SetLogger(logs.AdapterMultiFile, string(configStr))
	logs.Debug("Start stworker for beego.")
	return nil
}

// 日志级别
func GetLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}
