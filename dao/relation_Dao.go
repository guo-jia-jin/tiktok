package dao

import (
	"Tiktok/global"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type Relation struct {
	ID          uint64
	Follow_ID   uint64
	Follower_ID uint64
	Created_At  time.Time
	Updated_At  time.Time
	Type        uint16
}

func (Relation) TableName() string {
	return "relations"
}

func (re *Relation) CreateRelation(relation *Relation) (err error) {
	err = global.GVA_DB.Create(&relation).Error
	if err != nil {
		log.Println("Relation_Dao->CreateRelation: created failed")
		log.Println("err:", err.Error())
		return err
	}
	return err
}

func (re *Relation) CountFollowByID(user_id uint64) (follow_count int64, err error) {
	err = global.GVA_DB.Model(Relation{}).
		Where(map[string]interface{}{"follower_id": user_id, "type": 1}).
		Count(&follow_count).Error
	if err != nil {
		log.Println("Relation_Dao->CountFollowByID: count failed")
		log.Println("err:", err.Error())
		return follow_count, err
	}
	return follow_count, err
}

func (re *Relation) CountFollowerByID(user_id uint64) (follower_count int64, err error) {
	err = global.GVA_DB.Model(Relation{}).
		Where(map[string]interface{}{"follow_id": user_id, "type": 1}).
		Count(&follower_count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Relation_Dao->CountFollowerByID: count failed")
		log.Println("err:", err.Error())
		return follower_count, err
	}
	return follower_count, err
}

func (re *Relation) UpdateRelationType(user_id uint64, user_to_id uint64, relationType uint16) (err error) {
	err = global.GVA_DB.Model(Relation{}).
		Where(map[string]interface{}{"follow_id": user_to_id, "follower_id": user_id}).
		Updates(map[string]interface{}{"type": relationType, "updated_at": time.Now()}).Error
	if err != nil {
		log.Println("Relation_Dao->UpdateRelationType: update type failed")
		log.Println("err:", err.Error())
		return err
	}
	return err
}

func (re *Relation) IsExist(user_id uint64, user_to_id uint64) (isExist bool) {
	relation := Relation{}
	if !errors.Is(global.GVA_DB.Where("follower_id = ? AND follow_id = ?",
		user_id, user_to_id).First(&relation).Error, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}

func (re *Relation) IsFollow(user_id uint64, to_user_id uint64) (isFollow bool, err error) {
	var relation Relation
	err = global.GVA_DB.Select("type").Where("follow_id = ? AND follower_id = ?", to_user_id, user_id).
		Find(&relation).Error
	isfollowtype := relation.Type
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return isFollow, nil
	} else if err != nil {
		return isFollow, err
	} else {
		if isfollowtype == 1 {
			isFollow = true
			return isFollow, err
		} else {
			return isFollow, err
		}
	}
	// relation := Relation{}
	// err = global.GVA_DB.
	// 	Where("follow_id = ? AND follower_id = ? And type = ?", to_user_id, user_id, 1).
	// 	First(&relation).Error
	// fmt.Printf("gorm.ErrRecordNotFound: %v\n", gorm.ErrRecordNotFound)
	// if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return isFollow, err
	// } else if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return isFollow, nil
	// }
	// isFollow = true
	// return isFollow, err
}

func (re *Relation) GetFollowIDList(user_id uint64) (followIdList *[]uint64, err error) {
	err = global.GVA_DB.Model(Relation{}).
		Where(map[string]interface{}{"follower_id": user_id, "type": 1}).
		Pluck("follow_id", &followIdList).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return followIdList, nil
	} else if err != nil {
		log.Println("Relation_Dao->GetFollowIDList happened err:", err.Error())
		return followIdList, err
	}
	return followIdList, err
}

func (re *Relation) GetFollowerIDList(user_id uint64) (followerIdList *[]uint64, err error) {
	err = global.GVA_DB.Model(Relation{}).
		Where(map[string]interface{}{"follow_id": user_id, "type": 1}).
		Pluck("follower_id", &followerIdList).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return followerIdList, nil
	} else if err != nil {
		log.Println("Relation_Dao->GetFollowIDList happened err:", err.Error())
		return followerIdList, err
	}
	return followerIdList, err
}
