package base

// 数据库实例变量
var DbUser = &UserModel{BaseModel{}, "dbuser"}

// UserModel
type UserModel struct {
	BaseModel
	GroupName string
}
