package dao

import (
	"Tiktok/global"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type Favorite struct {
	ID            uint64
	Video_ID      uint64
	User_ID       uint64
	Favorite_Type uint16
	Created_At    time.Time
	Updated_At    time.Time
}

func (Favorite) TableName() string {
	return "favorites"
}

func (fa *Favorite) InsertFavorite(favorite *Favorite) (err error) {
	err = global.GVA_DB.Model(Favorite{}).Create(&favorite).Error
	if err != nil {
		log.Println("Favorite_Dao->InsertFavorite: created failed")
		log.Println("err:", err.Error())
		return err
	}
	return err
}

func (fa *Favorite) FavoriteCountByVideoID(video_id uint64) (favoriteCount int64, err error) {
	err = global.GVA_DB.Model(Favorite{}).
		Where(map[string]interface{}{"video_id": video_id, "favorite_type": 1}).
		Count(&favoriteCount).Error
	if err != nil {
		return favoriteCount, err
	}
	return favoriteCount, err
}

func (fa *Favorite) UpdateFavoriteAction(video_id uint64, user_id uint64, favoriteType uint16) (err error) {
	err = global.GVA_DB.Model(Favorite{}).
		Where(map[string]interface{}{"video_id": video_id, "user_id": user_id}).
		Updates(map[string]interface{}{"favorite_type": favoriteType, "updated_at": time.Now()}).Error
	if err != nil {
		return err
	}
	return err
}

func (fa *Favorite) IsExist(video_id uint64, user_id uint64) (isExist bool) {
	favorite := Favorite{}
	if !errors.Is(global.GVA_DB.Where("video_id = ? AND user_id = ?",
		video_id, user_id).First(&favorite).Error, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}

func (fa *Favorite) IsFavorite(user_id uint64, video_id uint64) (isFavorite bool, err error) {
	// favorite := Favorite{}
	// err = global.GVA_DB.Where(map[string]interface{}{"video_id": video_id, "user_id": user_id, "favorite_type": 1}).First(&favorite).Error
	// if err == nil {
	// 	return true, err
	// } else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return false, nil
	// } else {
	// 	return false, err
	// }
	var favorite Favorite
	err = global.GVA_DB.Select("favorite_type").Where("video_id = ? AND user_id = ?", video_id, user_id).
		Find(&favorite).Error
	isfavoritetype := favorite.Favorite_Type
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return isFavorite, nil
	} else if err != nil {
		return isFavorite, err
	} else {
		if isfavoritetype == 1 {
			isFavorite = true
			return isFavorite, err
		} else {
			return isFavorite, err
		}
	}
}

//获取该用户点赞的视频id列表
func (fa *Favorite) GetUserFavoriteListID(user_id uint64) (idList *[]uint64, err error) {
	err = global.GVA_DB.Model(Favorite{}).Where("user_id = ? AND favorite_type = ?", user_id, 1).Pluck("video_id", &idList).Error
	if err != nil {
		log.Println("Favorite_Dao->GetUserFavoriteListID err:", err.Error())
		return idList, err
	}
	return idList, err
}

func (fa *Favorite) CountUserFavorite(user_id uint64) (count int64, err error) {
	err = global.GVA_DB.Model(Favorite{}).Where("user_id = ? AND favorite_type = ?", user_id, 1).Count(&count).Error
	if err != nil {
		log.Println("Favorite_Dao->CountUserFavorite err:", err.Error())
		return count, err
	}
	return count, err
}
