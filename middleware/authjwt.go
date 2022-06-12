package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

//拦截解析token，解析成功就放行，用与需要登录的功能服务
func AuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Query("token")
		if len(auth) == 0 {
			c.Abort()
			c.JSON(http.StatusUnauthorized, Response{
				StatusCode: -1,
				StatusMsg:  "Unauthorized",
			})
		}
		tokenClaim, err := GetParseToken(auth)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, Response{
				StatusCode: -1,
				StatusMsg:  "Token error",
			})
		} else {
			//log.Println("Token OK")
			//fmt.Printf("tokenClaim.BaseClaims.User_Id: %v\n", tokenClaim.BaseClaims.User_Id)
			log.Println("id", tokenClaim.User_Id)
			//log.Println("id", tokenClaim.User_Name)
			c.Set("userId", tokenClaim.User_Id)
			//c.Set("userName", tokenClaim.User_Name)
		}
		c.Next()
	}
}

//拦截解析token，用于放行不需要token的服务，如/feed/可以在没有登录的情况下访问
func AuthWithOutLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Query("token")
		if len(auth) != 0 {
			tokenClaim, err := GetParseToken(auth)
			if err != nil {
				c.Abort()
				c.JSON(http.StatusUnauthorized, Response{
					StatusCode: -1,
					StatusMsg:  "Token error",
				})
			} else {
				//log.Println("Token OK")
				c.Set("userId", tokenClaim.User_Id)
				//c.Set("userName", tokenClaim.User_Name)
			}
		} else {
			c.Set("userId", 0)
		}
		c.Next()
	}
}
