package api

import (
	"Tiktok/handler/response"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RelationApi struct {
}

// @Router /douyin/relation/action/
func (re *RelationApi) RelationAction(c *gin.Context) {
	actionType := c.Query("action_type")
	userId := c.GetInt64("userId")
	user_id := uint64(userId)
	tag_id, _ := strconv.ParseUint(c.Query("to_user_id"), 10, 64)

	//log.Println("user_id:", user_id)
	//log.Println("tag_id:", tag_id)
	if actionType == "1" {
		err := relationService.FollowAction(user_id, tag_id)
		if err != nil {
			log.Println("in if actionType == 1 -> err:", err.Error())
			response.ErrCodeWithMess(err.Error(), c)
		} else {
			response.OKCodeWithMess("关注失败", c)
		}
	} else if actionType == "2" {
		err := relationService.UnFollowAction(user_id, tag_id)
		if err != nil {
			log.Println("in if actionType == 2 -> err:", err.Error())
			response.ErrCodeWithMess("取消关注失败", c)
		} else {
			response.OKCodeWithMess("取消关注成功", c)
		}
	} else {
		response.ErrCodeWithMess("操作类型错误", c)
	}
}

// @Router /douyin/relation/follow/list/
func (re *RelationApi) FollowList(c *gin.Context) {
	// userId := c.GetInt64("userId")
	// user_id := uint64(userId)
	user_id, _ := strconv.ParseUint(c.Query("user_id"), 10, 64)
	followList, err := relationService.GetFollowList(user_id)
	if err != nil {
		log.Println("FollowList err:", err.Error())
		response.ErrCodeWithMess("获取关注列表失败", c)
	} else {
		response.OkWithDataforRelation("", followList, c)
	}
}

// @Router /douyin/relation/follower/list/
func (re *RelationApi) FollowerList(c *gin.Context) {
	user_id, _ := strconv.ParseUint(c.Query("user_id"), 10, 64)
	followerList, err := relationService.GetFollowerList(user_id)
	if err != nil {
		log.Println("FollowerList err:", err.Error())
		response.ErrCodeWithMess("获取粉丝列表失败", c)
	} else {
		response.OkWithDataforRelation("", followerList, c)
	}
}
