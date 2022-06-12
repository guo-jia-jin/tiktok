package serverimp

import (
	"Tiktok/dao"
	"Tiktok/service/server"
	"log"
	"time"
)

type Relation_ServerImp struct {
}

func (re *Relation_ServerImp) IsFollow(user_id uint64, tag_id uint64) (isFollow bool, err error) {
	var relaDao dao.Relation
	isFollow, err = relaDao.IsFollow(user_id, tag_id)
	if err != nil {
		log.Println("查询失败:", err.Error())
	}
	return isFollow, nil
}

func (re *Relation_ServerImp) GetFollowCount(user_id uint64) (follow_count int64, err error) {
	var relaDao dao.Relation
	follow_count, err = relaDao.CountFollowByID(user_id)
	if err != nil {
		log.Println("relation_severImp->GetFollowCount happened err:", err.Error())
		return follow_count, err
	}
	return follow_count, err
}

func (re *Relation_ServerImp) GetFollowerCount(user_id uint64) (follower_count int64, err error) {
	var relaDao dao.Relation
	follower_count, err = relaDao.CountFollowerByID(user_id)
	if err != nil {
		log.Println("relation_severImp->GetFollowerCount happened err:", err.Error())
		return follower_count, err
	}
	return follower_count, err
}

func (re *Relation_ServerImp) FollowAction(user_id uint64, tag_id uint64) (err error) {
	var relaDao dao.Relation
	isExist := relaDao.IsExist(user_id, tag_id)
	if isExist {
		err = relaDao.UpdateRelationType(user_id, tag_id, 1)
		if err != nil {
			log.Println("relation_severImp->FollowAction happened err:", err.Error())
		}
	} else {
		relation := dao.Relation{
			Follow_ID:   tag_id,
			Follower_ID: user_id,
			Created_At:  time.Now(),
			Updated_At:  time.Now(),
			Type:        1,
		}
		err = relation.CreateRelation(&relation)
		if err != nil {
			log.Println("relation_severImp->FollowAction happened err:", err.Error())
			return err
		}
	}

	return err
}

func (re *Relation_ServerImp) UnFollowAction(user_id uint64, tag_id uint64) (err error) {
	var relaDao dao.Relation
	isExist := relaDao.IsExist(user_id, tag_id)
	if isExist {
		err = relaDao.UpdateRelationType(user_id, tag_id, 2)
		if err != nil {
			log.Println("relation_severImp->UnFollowAction happened err:", err.Error())
			return err
		}
	}
	return err
}

func (re *Relation_ServerImp) GetFollowList(user_id uint64) (followList *[]server.RelationPageUser, err error) {
	var relaDao dao.Relation
	var user_serverImp User_ServerImp
	followIdList, err := relaDao.GetFollowIDList(user_id)
	if err != nil {
		log.Println("relation_severImp->GetFollowList happened err:", err.Error())
		return followList, err
	}
	follows := make([]server.RelationPageUser, 0, len(*followIdList))
	for _, id := range *followIdList {
		userInfo, err := user_serverImp.GetUserPageInfo(id)
		if err != nil {
			log.Println("GetFollowList in for happened err:", err.Error())
			continue
		}
		follow := server.RelationPageUser{
			ID:             userInfo.UserID,
			Name:           userInfo.Name,
			Follow_Count:   userInfo.Follow_Count,
			Follower_Count: userInfo.Follower_Count,
			Is_Follow:      true,
		}
		follows = append(follows, follow)
	}
	followList = &follows
	return followList, err
}

func (re *Relation_ServerImp) GetFollowerList(user_id uint64) (followerList *[]server.RelationPageUser, err error) {
	var relaDao dao.Relation
	var user_serverImp User_ServerImp
	followerIdList, err := relaDao.GetFollowerIDList(user_id)
	if err != nil {
		log.Println("relation_severImp->GetFollowerList happened err:", err.Error())
		return followerList, err
	}
	follows := make([]server.RelationPageUser, 0, len(*followerIdList))
	for _, id := range *followerIdList {
		userInfo, err := user_serverImp.GetUserPageInfo(id)
		if err != nil {
			log.Println("GetFollowerList in for happened err1:", err.Error())
			continue
		}
		isFollow, err := re.IsFollow(user_id, userInfo.UserID)
		if err != nil {
			log.Println("GetFollowerList in for happened err2:", err.Error())
		}
		follow := server.RelationPageUser{
			ID:             userInfo.UserID,
			Name:           userInfo.Name,
			Follow_Count:   userInfo.Follow_Count,
			Follower_Count: userInfo.Follower_Count,
			Is_Follow:      isFollow,
		}
		follows = append(follows, follow)
	}
	followerList = &follows
	return followerList, err
}
