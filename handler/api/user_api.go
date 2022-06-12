package api

import (
	"Tiktok/handler/request"
	"Tiktok/handler/response"
	"log"

	"strconv"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

// @Router /douyin/user/register/
func (u *UserApi) Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	userRegInfo := request.UserRequest{
		UserName: username,
		Password: password,
	}
	user, err := userService.Register(&userRegInfo)
	if err != nil && err.Error() == "User name already exists" {
		response.ErrCodeWithMess(err.Error(), c)
	} else if err != nil {
		log.Println("err:", err.Error())
	} else {
		response.OkWithDataforUser("注册成功", *user, c)
	}
}

// @Router /douyin/user/
func (u *UserApi) UserInfo(c *gin.Context) {
	user_id := c.Query("user_id")
	//fmt.Printf("user_id: %v\n", user_id)
	//userId, _ := c.Get("userId")
	id, _ := strconv.ParseUint(user_id, 10, 64)
	userPageInfo, err := userService.GetUserPageInfo(id)
	if err != nil {
		log.Println("err:", err.Error())
		response.ErrCodeWithMess(err.Error(), c)
	} else {
		response.OkWithDataforUserPage("", *userPageInfo, c)
	}
}

// @Router /douyin/user/login/
func (u *UserApi) UserLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	userLoginInfo := request.UserRequest{
		UserName: username,
		Password: password,
	}
	user, err := userService.Login(&userLoginInfo)
	if err != nil {
		log.Println("err:", err.Error())
		response.ErrCodeWithMess(err.Error(), c)
	} else {
		response.OkWithDataforUser("登录成功", *user, c)
	}
}
