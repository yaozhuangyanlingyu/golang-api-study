package base

// 数据库实例变量
var DbShop = &ShopModel{BaseModel{}, "dbshop"}

// ShopModel
type ShopModel struct {
	BaseModel
	GroupName string
}
