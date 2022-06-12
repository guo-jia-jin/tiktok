package response

import (
	"Tiktok/service/server"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	NextTime   int64                  `json:"next_time,omitempty"`
	Video_List []server.VideoPageInfo `json:"video_list,omitempty"`
}

func OkWithDataforFeed(video_list *[]server.VideoPageInfo, next_time time.Time, c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:   Response{SUCCESS, ""},
		NextTime:   next_time.UnixMilli(),
		Video_List: *video_list,
	})
}

func OKWithOutDataforFeed(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		StatusCode: SUCCESS,
		StatusMsg:  "",
	})
}
