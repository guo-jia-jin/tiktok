package router

import (
	"Tiktok/handler"
	"Tiktok/middleware"

	"github.com/gin-gonic/gin"
)

type PublishRouter struct {
}

func (p *PublishRouter) InitPublishRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	publishRouter := Router.Group("publish")
	publishRouterApi := handler.ApiGroupApp.ApiGroup.PublishApi
	{
		publishRouter.POST("action/", middleware.AuthBody(), publishRouterApi.PublishAction)
		publishRouter.GET("list/", middleware.AuthWithOutLogin(), publishRouterApi.PublishList)
	}
	return publishRouter
}
