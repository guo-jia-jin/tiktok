package router

import (
	"Tiktok/handler"
	"Tiktok/middleware"

	"github.com/gin-gonic/gin"
)

type FeedRouter struct {
}

// XXX/douyin/feed
func (b *FeedRouter) InitFeedRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	feedRouter := Router.Group("feed")
	feedRouterApi := handler.ApiGroupApp.ApiGroup.FeedApi
	{
		//相关分组请求
		feedRouter.GET("/", middleware.AuthWithOutLogin(), feedRouterApi.GetFeed)
	}
	return feedRouter
}
