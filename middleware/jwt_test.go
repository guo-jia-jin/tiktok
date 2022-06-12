package middleware

import (
	"Tiktok/dao"
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	user := dao.User{
		ID:   11,
		Name: "guojiajin",
	}
	str := GetJWTByUserInfo(user)
	fmt.Printf("str: %v\n", str)
	b, err := NewJWT().ParseToken(str)
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
	}
	fmt.Printf("b: %v\n", b)
}
