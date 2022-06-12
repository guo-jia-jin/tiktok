package routers

import "Tiktok/routers/router"

//路由组
//向外提供所有系统路由
type RouterGroup struct {
	router.FeedRouter
	router.UserRouter
	router.PublishRouter
	router.FavoriteRouter
	router.CommentRouter
	router.RelationRouter
}

var RouterGroupAPP = new(RouterGroup)
