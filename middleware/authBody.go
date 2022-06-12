package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthBody 鉴权中间件 用于post
func AuthBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.PostFormValue("token")
		fmt.Printf("%v \n", auth)

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
				StatusMsg:  "Token Error",
			})
		} else {
			println("token OK")
		}
		c.Set("userId", tokenClaim.User_Id)
		c.Set("userName", tokenClaim.User_Name)
		c.Next()
	}
}
