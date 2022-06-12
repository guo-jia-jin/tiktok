package initialize

import (
	"Tiktok/routers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//初始化总路由
func Routers() *gin.Engine {
	Router := gin.Default()
	SysRouter := routers.RouterGroupAPP
	//访问本地静态文件路由
	Router.StaticFS("/static", http.Dir("./static"))
	//公共分组：如feed视频流这样不需要登录的
	PublicGroup := Router.Group("douyin")
	{
		SysRouter.InitFeedRouter(PublicGroup)
	}
	//私有分组：登录的才能操作的，这里将注册也放置在私有分组中
	PrivateGroup := Router.Group("douyin")
	{
		SysRouter.IintUserRouter(PrivateGroup)
		SysRouter.InitPublishRouter(PrivateGroup)
		SysRouter.InitFavoriteRouter(PrivateGroup)
		SysRouter.InitCommentRouter(PrivateGroup)
		SysRouter.InitRelationRouter(PrivateGroup)
	}
	fmt.Println("总路由初始化成功....")
	return Router
}
