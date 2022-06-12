package dao

import (
	"Tiktok/global"
	"log"
	"time"
)

//用于对应数据库实体
type Video struct {
	ID           uint64    `json:"id"`
	Title        string    `json:"title"`
	UserID       uint64    `json:"-"`
	Playurl      string    `json:"play_url"`
	Coverurl     string    `json:"cover_url"`
	Published_At time.Time `json:"-"`
}

func (Video) TableName() string {
	return "videos"
}

func (v *Video) InsertVideo(video *Video) (err error) {
	err = global.GVA_DB.Model(Video{}).Create(&video).Error
	if err != nil {
		return err
	}
	return err
}

func (v *Video) GetVideoListByAuthorID(user_id uint64) (videoList *[]Video, err error) {
	err = global.GVA_DB.Model(Video{}).Where("user_id = ?", user_id).Find(&videoList).Error
	if err != nil {
		return videoList, err
	}
	return videoList, err
}

func (v *Video) GetVideoByVideoID(video_id uint64) (video *Video, err error) {
	err = global.GVA_DB.Model(Video{}).Where("id = ?", video_id).First(&video).Error
	if err != nil {
		return video, err
	}
	return video, err
}

func (v *Video) GetVideoListByLastTime(lastTime time.Time, user_id uint64) (videoList *[]Video, err error) {
	err = global.GVA_DB.Model(Video{}).Limit(3).Order("published_at desc").
		Where("published_at < ? AND user_id <> ?", lastTime, user_id).Find(&videoList).Error
	if err != nil {
		log.Println("video_dao -> GetVideoListByLastTime err:", err.Error())
		return videoList, err
	}
	return videoList, err
}

func (v *Video) GetVideoIDList(user_id uint64) (videoId_list *[]uint64, err error) {
	err = global.GVA_DB.Model(Video{}).Where("user_id = ?", user_id).
		Pluck("id", &videoId_list).Error
	if err != nil {
		log.Println("video_dao -> GetVideoIDList err:", err.Error())
		return videoId_list, err
	}
	return videoId_list, err
}

func (v *Video) CountVideoByUserID(user_id uint64) (count int64, err error) {
	err = global.GVA_DB.Model(Video{}).Where("user_id = ?", user_id).Count(&count).Error
	if err != nil {
		log.Println("video_dao -> CountVideoByUserID err:", err.Error())
		return count, err
	}
	return count, err
}
