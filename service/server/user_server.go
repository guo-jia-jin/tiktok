package server

import (
	"Tiktok/dao"
)

type User_Server interface {
	Register(userRegInfo *dao.User) (user *dao.User, err error)
	GetUserPageInfo(user_id uint64) (userPageInfo *UserPageInfo, err error)
	Login(userLogInfo *dao.User) (user *dao.User, err error)
	CreateUserInfo(user_id uint64) (err error)
}
type UserPageInfo struct {
	dao.UserInfo
	Name           string `json:"name"`
	Follow_Count   int64  `json:"follow_count"`
	Follower_Count int64  `json:"follower_count"`
	Is_Follow      bool   `json:"is_follow"`
	TotalFavorited int64  `json:"total_favorited,omitempty"`
	FavoriteCount  int64  `json:"favorite_count,omitempty"`
	WorkCount      int64  `json:"work_count,omitempty"`
}
