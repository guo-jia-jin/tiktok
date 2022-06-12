package response

import (
	"Tiktok/service/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FavoriteResponse struct {
	Response
	Video_List []server.VideoPageInfo `json:"video_list,omitempty"`
}

func OkWithDataforFavor(video_list []server.VideoPageInfo, c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:   Response{SUCCESS, ""},
		Video_List: video_list,
	})
}
