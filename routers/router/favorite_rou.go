package router

import (
	"Tiktok/handler"
	"Tiktok/middleware"

	"github.com/gin-gonic/gin"
)

type FavoriteRouter struct {
}

func (f *FavoriteRouter) InitFavoriteRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	favoriteRouter := Router.Group("favorite")
	favoriteApi := handler.ApiGroupApp.ApiGroup.FavoriteApi
	{
		favoriteRouter.POST("action/", middleware.AuthLogin(), favoriteApi.FavoriteAction)
		favoriteRouter.GET("list/", middleware.AuthWithOutLogin(), favoriteApi.FavoriteList)
	}
	return favoriteRouter
}
