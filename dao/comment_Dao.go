package dao

import (
	"Tiktok/global"
	"errors"
	"log"
	"time"
)

type Comment struct {
	ID          uint64
	Video_ID    uint64
	User_ID     uint64
	Content     string
	Create_Date time.Time
}

func (Comment) TableName() string {
	return "comments"
}

func (c *Comment) CountByVideoID(videoId uint64) (count int64, err error) {
	err = global.GVA_DB.Model(Comment{}).Where(map[string]interface{}{"video_id": videoId}).
		Count(&count).Error
	if err != nil {
		log.Println("Comment_Dao-CountByVideoID:failed")
		log.Println("err:", err.Error()) //打印错误详细信息
		return -1, errors.New("comments count failed")
	}
	return count, err
}

func (c *Comment) CreateComment(comment *Comment) (com *Comment, err error) {
	err = global.GVA_DB.Model(Comment{}).Create(&comment).Error
	if err != nil {
		log.Println("Comment_Dao-CreateComment:failed")
		log.Println("err:", err.Error()) //打印错误详细信息
		return nil, errors.New("comment create failed")
	}
	com = comment
	return com, err
}

func (c *Comment) DeleteCommentByID(commId uint64) (err error) {
	comment := Comment{}
	err = global.GVA_DB.Model(Comment{}).Delete(&comment, commId).Error
	if err != nil {
		log.Println("Comment_Dao-DeleteCommentByID:failed")
		log.Println("err:", err.Error()) //打印错误详细信息
		return errors.New("comment delete failed")
	}
	return err
}

func (c *Comment) GetCommentListByVideoID(videoId uint64) (commentList *[]Comment, err error) {
	err = global.GVA_DB.Model(Comment{}).Where(map[string]interface{}{"video_id": videoId}).
		Find(&commentList).Order("create_date desc").Error
	if err != nil {
		log.Println("Comment_Dao-GetCommentListByVideoID:failed")
		log.Println("err:", err.Error())
		return nil, errors.New("Get CommnetList failed")
	}
	return commentList, err
}
