package serverimp

import (
	"Tiktok/dao"
	"Tiktok/service/server"
	"log"
	"time"
)

type Favorite_ServerImp struct {
	// //根据用户id获取此用户被点赞数
	// TotalUserBeFavorited(user_id uint64) (int64, error)
}

func (fa *Favorite_ServerImp) Is_Favorite(video_id uint64, user_id uint64) (isFavorite bool, err error) {
	var favorDao dao.Favorite
	isFavorite, err = favorDao.IsFavorite(user_id, video_id)
	if err != nil {
		log.Println("get isFavorite failed err:", err.Error())
		err = nil
	}
	return isFavorite, err
}

func (fa *Favorite_ServerImp) VideoFvaoriteCount(video_id uint64) (count int64, err error) {
	var favorDao dao.Favorite
	count, err = favorDao.FavoriteCountByVideoID(video_id)
	if err != nil {
		log.Println("VideoFvaoriteCount->get fvaoriteCount failed err:", err.Error())
		err = nil
	}
	return count, err
}

func (fa *Favorite_ServerImp) TotalUserFavorite(user_id uint64) (total int64, err error) {
	var video_serverImp Video_ServerImp
	//1.获取这个用户发布的所有视频id列表
	videoIds, err := video_serverImp.GetVideoIDList(user_id)
	if err != nil {
		return total, err
	}
	//2.通过循环统计用户的每一个视频的点赞数，并且累加
	for _, videoId := range *videoIds {
		count, err := fa.VideoFvaoriteCount(videoId)
		if err != nil {
			log.Println("TotalUserFavorite in for err1:", err.Error())
			err = nil
			continue
		}
		total += count
	}
	//3.输出结果
	return total, err
}

func (fa *Favorite_ServerImp) UserFavoriteCount(user_id uint64) (count int64, err error) {
	var favorDao dao.Favorite
	count, err = favorDao.CountUserFavorite(user_id)
	if err != nil {
		log.Println("UserFavoriteCount->get count failed err:", err.Error())
		err = nil
	}
	return count, err
}

func (fa *Favorite_ServerImp) FavoriteAction(user_id uint64, video_id uint64, action_type uint16) (err error) {
	var favorDao dao.Favorite
	isExist := favorDao.IsExist(video_id, user_id)
	//fmt.Printf("isExist: %v\n", isExist)
	if isExist {
		err = favorDao.UpdateFavoriteAction(video_id, user_id, action_type)
		if err != nil {
			log.Println("In FavoriteAction happened err1:", err.Error())
			return err
		}
		return err
	}
	favorite := dao.Favorite{
		User_ID:       user_id,
		Video_ID:      video_id,
		Created_At:    time.Now(),
		Updated_At:    time.Now(),
		Favorite_Type: action_type,
	}
	err = favorDao.InsertFavorite(&favorite)
	if err != nil {
		log.Println("In FavoriteAction happened err2:", err.Error())
		return err
	}
	return err
}

func (fa *Favorite_ServerImp) GetUserFavoriteList(user_id uint64, cur_id uint64) (video_list *[]server.VideoPageInfo, err error) {
	var video_serverImp Video_ServerImp
	//user_id 为当前页面用户的id cur_id登录用户id，其浏览了当前页面用户
	var favorDao dao.Favorite
	video_idList, err := favorDao.GetUserFavoriteListID(user_id) //获取当前页面用户喜欢视频的id列表
	//fmt.Printf("video_idList: %v\n", *video_idList)
	if err != nil {
		log.Println("Favorite_severImp->GetUserFavoriteList err:", err.Error())
		return video_list, err
	}
	var vdieoDao dao.Video
	vdieos := make([]dao.Video, 0, len(*video_idList))
	for _, video_id := range *video_idList {
		video, err := vdieoDao.GetVideoByVideoID(video_id)
		if err != nil {
			log.Println("Favorite_severImp->GetUserFavoriteList in for err1:", err.Error())
			err = nil
			continue
		}
		vdieos = append(vdieos, *video)
	}
	//组装视频页面信息返回体，其中cur_id用于确定当前页面用户喜欢的视频和浏览页面用户的关系
	videoList := make([]server.VideoPageInfo, 0, len(vdieos))
	err = video_serverImp.ProcessVideos(&videoList, &vdieos, cur_id)
	if err != nil {
		return video_list, err
	}
	video_list = &videoList
	return video_list, err
}
