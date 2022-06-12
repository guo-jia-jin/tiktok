package api

import (
	"Tiktok/service"
)

type ApiGroup struct {
	FeedApi
	UserApi
	PublishApi
	FavoriteApi
	CommentApi
	RelationApi
}

var (
	videoService    = service.ServerGroupApp.Servers.Video_ServerImp
	userService     = service.ServerGroupApp.Servers.User_ServerImp
	commentService  = service.ServerGroupApp.Servers.Comment_ServerImp
	relationService = service.ServerGroupApp.Servers.Relation_ServerImp
	favoriteService = service.ServerGroupApp.Servers.Favorite_ServerImp
)
