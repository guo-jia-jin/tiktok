package api

import (
	"Tiktok/dao"
	"Tiktok/handler/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CommentApi struct {
}

// @Router /douyin/comment/action/
func (co *CommentApi) CommentAction(c *gin.Context) {
	actionType := c.Query("action_type")
	user_id := uint64(c.GetInt64("userId"))
	if actionType == "1" {
		video_id, _ := strconv.ParseUint(c.Query("video_id"), 10, 64)
		comment_text := c.Query("comment_text")
		comment := dao.Comment{
			User_ID:     user_id,
			Video_ID:    video_id,
			Content:     comment_text,
			Create_Date: time.Now(),
		}
		commInfo, err := commentService.SendComment(&comment)
		if err != nil {
			response.ErrCodeWithMess("评论发表失败", c)
		} else {
			response.OkWithDataforComm("评论发表成功", &commInfo, c)
		}
	} else if actionType == "2" {
		comment_id, _ := strconv.ParseUint(c.Query("comment_id"), 10, 64)
		err := commentService.DelCommentByID(comment_id, user_id)
		if err != nil && err.Error() == "Permission denied" {
			response.ErrCodeWithMess("非评论作者", c)
		} else if err != nil {
			response.ErrCodeWithMess("评论删除失败", c)
		} else {
			response.OKCodeWithMess("评论删除成功", c)
		}
	} else {
		response.ErrCodeWithMess("操作类型错误", c)
	}
}

// @Router /douyin/comment/list/
func (co *CommentApi) CommentList(c *gin.Context) {
	user_id := uint64(c.GetInt64("userId"))
	//log.Println("user_id:", user_id)
	//log.Println("video_id:", videoId)
	video_id, _ := strconv.ParseUint(c.Query("video_id"), 10, 64)
	commentList, err := commentService.GetCommentList(video_id, user_id)
	if err != nil {
		response.ErrCodeWithMess("评论加载失败", c)
	} else {
		response.OkWithDataforCommList("", commentList, c)
	}
}
