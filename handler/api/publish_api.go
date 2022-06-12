package api

import (
	"Tiktok/handler/response"
	"log"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PublishApi struct {
}

// @Router /douyin/publish/action/
func (p *PublishApi) PublishAction(c *gin.Context) {
	file, err := c.FormFile("data")
	if path.Ext(file.Filename) != ".mp4" {
		response.ErrCodeWithMess("上传文件类型错误", c)
	} else if err != nil {
		log.Println("publish_api->PublishAction-> c.FormFile err:", err.Error())
		response.ErrCodeWithMess("视频上传失败", c)
	} else {
		user_id := uint64(c.GetInt64("userId"))
		user_name := c.GetString("userName")
		title := c.PostForm("title")
		log.Println("user_id", user_id)
		log.Println("user_name", user_name)
		err = videoService.Publish(file, user_id, user_name, title)
		if err != nil {
			log.Println("publish_api->PublishAction-> Publish err:", err.Error())
			response.ErrCodeWithMess(err.Error(), c)
		} else {
			response.OKCodeWithMess("视频发布成功", c)
		}
	}
	// c.JSON(200, gin.H{
	// 	"massage": "/douyin/publish/action/",
	// })
}

// @Router /douyin/publish/list/
func (p *PublishApi) PublishList(c *gin.Context) {
	user_id, _ := strconv.ParseUint(c.Query("user_id"), 10, 64)
	cur_id := uint64(c.GetInt64("userId"))
	video_list, err := videoService.GetVideoListByAuthorID(user_id, cur_id)
	if err != nil {
		response.ErrCodeWithMess(err.Error(), c)
	} else {
		response.OkWithDataforFavor(*video_list, c)
	}
}
