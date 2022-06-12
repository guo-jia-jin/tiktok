package router

import (
	"Tiktok/handler"
	"Tiktok/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (u *UserRouter) IintUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := Router.Group("user")
	userRouterApi := handler.ApiGroupApp.ApiGroup.UserApi
	{
		userRouter.GET("/", middleware.AuthWithOutLogin(), userRouterApi.UserInfo)
		userRouter.POST("register/", userRouterApi.Register)
		userRouter.POST("login/", userRouterApi.UserLogin)
	}
	return userRouter
}
