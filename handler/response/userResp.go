package response

import (
	"Tiktok/dao"
	"Tiktok/middleware"
	"Tiktok/service/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	Response
	User_ID uint64 `json:"user_id"`
	Token   string `json:"token"`
}

type UserPageInfoResp struct {
	Response
	UserPageInfo server.UserPageInfo `json:"user"`
}

func OkWithDataforUser(msg string, user dao.User, c *gin.Context) {
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{SUCCESS, msg},
		User_ID:  user.ID,
		Token:    middleware.GetJWTByUserInfo(user),
	})
}

func OkWithDataforUserPage(msg string, userPage server.UserPageInfo, c *gin.Context) {
	c.JSON(http.StatusOK, UserPageInfoResp{
		Response:     Response{SUCCESS, msg},
		UserPageInfo: userPage,
	})
}
