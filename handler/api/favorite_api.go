package api

import (
	"Tiktok/handler/response"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteApi struct {
}

// @Router /douyin/favorite/action/
func (f *FavoriteApi) FavoriteAction(c *gin.Context) {
	user_id := uint64(c.GetInt64("userId"))
	video_id, _ := strconv.ParseUint(c.Query("video_id"), 10, 64)
	action_type, _ := strconv.ParseUint(c.Query("action_type"), 10, 16)
	err := favoriteService.FavoriteAction(user_id, video_id, uint16(action_type))
	if err != nil {
		log.Println("err:", err.Error())
		response.ErrCodeWithMess("点赞失败", c)
	} else {
		response.OKCodeWithMess("点赞成功", c)
	}
}

// @Router /douyin/favorite/list/
func (f *FavoriteApi) FavoriteList(c *gin.Context) {
	user_id, _ := strconv.ParseUint(c.Query("user_id"), 10, 64)
	cur_id := uint64(c.GetInt64("userId"))
	videoList, err := favoriteService.GetUserFavoriteList(user_id, cur_id)
	if err != nil {
		log.Println("err:", err.Error())
		response.ErrCodeWithMess("获取失败", c)
	} else {
		response.OkWithDataforFavor(*videoList, c)
	}
}
