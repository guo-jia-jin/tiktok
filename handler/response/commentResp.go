package response

import (
	"Tiktok/service/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentsResponse struct {
	Response
	CommentList []server.CommentInfo `json:"comment_list,omitempty"`
}

type CommentResponse struct {
	Response
	CommentList server.CommentInfo `json:"comment,omitempty"`
}

func OkWithDataforCommList(msg string, commlist *[]server.CommentInfo, c *gin.Context) {
	c.JSON(http.StatusOK, CommentsResponse{
		Response:    Response{SUCCESS, msg},
		CommentList: *commlist,
	})
}

func OkWithDataforComm(msg string, commlist *server.CommentInfo, c *gin.Context) {
	c.JSON(http.StatusOK, CommentResponse{
		Response:    Response{SUCCESS, msg},
		CommentList: *commlist,
	})
}
