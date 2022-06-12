package router

import (
	"Tiktok/handler"
	"Tiktok/middleware"

	"github.com/gin-gonic/gin"
)

type RelationRouter struct {
}

func (re *RelationRouter) InitRelationRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	relationRouter := Router.Group("relation")
	relationApi := handler.ApiGroupApp.ApiGroup.RelationApi
	{
		relationRouter.POST("action/", middleware.AuthLogin(), relationApi.RelationAction)
		relationRouter.GET("follow/list/", middleware.AuthLogin(), relationApi.FollowList)
		relationRouter.GET("follower/list/", middleware.AuthLogin(), relationApi.FollowerList)
	}
	return relationRouter
}
