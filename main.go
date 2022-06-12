package main

import (
	"Tiktok/global"
	"Tiktok/initialize"
)

func main() {
	global.GVA_VP = initialize.Viper()               //初始化全局变量
	global.GVA_DB = initialize.GormMysql()           //gorm连接mysql数据库
	global.GVA_STATICURL = initialize.GetStaticUrl() //静态资源url前缀，无奈之举
	Router := initialize.Routers()                   //初始化总路由
	Router.Run(global.GVA_CONFIG.RunPort.Content)
}
