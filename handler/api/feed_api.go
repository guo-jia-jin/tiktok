package api

import (
	"Tiktok/handler/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedApi struct {
}

// @Router /douyin/feed/
func (b *FeedApi) GetFeed(c *gin.Context) {
	userId := c.GetInt64("userId")
	user_id := uint64(userId)
	//fmt.Printf("user_id: %v\n", user_id)
	inputTime := c.Query("latest_time")
	var lastTime time.Time
	if inputTime != "0" && len(inputTime) != 0 {
		me, _ := strconv.ParseInt(inputTime, 10, 64)
		lastTime = time.UnixMilli(me)
	} else {
		lastTime = time.Now()
	}
	videoList, nextTime, err := videoService.Feed(lastTime, user_id)
	if len(*videoList) == 0 && err == nil {
		response.OKWithOutDataforFeed(c)
	} else if err != nil {
		response.ErrCodeWithMess("视频获取失败", c)
	} else {
		response.OkWithDataforFeed(videoList, nextTime, c)
	}
}
