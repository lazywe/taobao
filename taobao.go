package taobao

import (
	"taobao/lib"
)

// 初始化配置 InitConfig
func InitConfig(appKey string, secretKey string, debug bool) {
	lib.SetConfig(appKey, secretKey, debug)
}

// // 淘宝客商品查询
// func NewTbkItemGet() services.TbkItemGetInterface {
// 	return services.NewTbkItemGet()
// }

// // 淘宝客商品查询
// func NewTbkItemRecommendGet() services.TbkItemRecommendGetInterface {
// 	return services.NewTbkItemRecommendGet()
// }
