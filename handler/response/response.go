package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

const (
	SUCCESS = 0
	ERROR   = 1
)

func ErrCodeWithMess(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		StatusCode: ERROR,
		StatusMsg:  msg,
	})
}

func OKCodeWithMess(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		StatusCode: SUCCESS,
		StatusMsg:  msg,
	})
}
