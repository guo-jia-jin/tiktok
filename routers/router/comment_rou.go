package router

import (
	"Tiktok/handler"
	"Tiktok/middleware"

	"github.com/gin-gonic/gin"
)

type CommentRouter struct {
}

func (co *CommentRouter) InitCommentRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	commentRouter := Router.Group("comment")
	commentApi := handler.ApiGroupApp.ApiGroup.CommentApi
	{
		commentRouter.POST("action/", middleware.AuthLogin(), commentApi.CommentAction)
		commentRouter.GET("list/", middleware.AuthLogin(), commentApi.CommentList)
	}
	return commentRouter
}
