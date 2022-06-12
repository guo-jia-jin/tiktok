package server

import (
	"Tiktok/dao"
	"mime/multipart"
	"time"
)

type Video_Server interface {
	//Feed
	Feed(lastTime time.Time, user_id uint64) (*[]VideoPageInfo, time.Time, error)
	//GetVideoByID
	GetVideoByID(video_id uint64, user_id uint64) (*VideoPageInfo, error)
	//Publish
	Publish(file *multipart.FileHeader, user_id uint64, user_name string, title string) error
	//获取当前用户的视频发布数
	PublishCount(user_id uint64) int64
	//通过用户id获取其发布视频的id列表
	GetVideoIDList(user_id uint64) (*[]uint64, error)
	//通过用户获取其发布的视频
	GetVideoListByAuthorID(user_id uint64, cur_id uint64) (*[]VideoPageInfo, error)
}

type VideoPageInfo struct {
	dao.Video
	Author         UserPageInfo `json:"author"`
	Favorite_Count int64        `json:"favorite_count"`
	Comment_Count  int64        `json:"comment_count"`
	Is_Favorite    bool         `json:"is_favorite"`
}
