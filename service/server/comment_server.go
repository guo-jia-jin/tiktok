package server

import "Tiktok/dao"

type Comment_Server interface {
	//根据视频id获取评论数量
	CountByVideoID(id uint64) (int64, error)
	//发表评论
	SendComment(comment *dao.Comment) (CommentInfo, error)
	//删除评论
	DelCommentByID(comment_id uint64) error
	//获取视频评论列表
	GetCommentList(video_id uint64, user_id uint64) ([]CommentInfo, error)
}

type CommentInfo struct {
	ID         uint64       `json:"id,omitempty"`
	UserInfo   UserPageInfo `json:"user,omitempty"`
	Content    string       `json:"content,omitempty"`
	CreateDate string       `json:"create_date,omitempty"`
}
