package serverimp

import (
	"Tiktok/dao"
	"Tiktok/middleware"
	"Tiktok/service/server"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

//为feedApi和publishApi提供实现服务
type Video_ServerImp struct {
}

func (vi *Video_ServerImp) Feed(lastTime time.Time, user_id uint64) (videos *[]server.VideoPageInfo, nextTime time.Time, err error) {
	var videoDao dao.Video
	videoList, err := videoDao.GetVideoListByLastTime(lastTime, user_id)
	videosTemp := make([]server.VideoPageInfo, 0, len(*videoList))
	err = vi.ProcessVideos(&videosTemp, videoList, user_id)
	if err != nil {
		log.Println("video_serverImp->Feed err:", err.Error())
		return videos, nextTime, errors.New("获取视频流失败")
	}
	if len(videosTemp) == 0 {
		nextTime = time.Now()
	} else {
		nextTime = videosTemp[len(videosTemp)-1].Published_At
	}
	videos = &videosTemp
	return videos, nextTime, err
}

func (vi *Video_ServerImp) ProcessVideos(videos *[]server.VideoPageInfo, videoList *[]dao.Video, user_id uint64) (err error) {
	for _, video := range *videoList {
		var videoPage server.VideoPageInfo
		err = vi.InsertVeidoPageInfo(&videoPage, &video, user_id)
		if err != nil {
			log.Println("ProcessVideos in for err1:", err.Error())
			err = nil
			continue
		}
		*videos = append(*videos, videoPage)
	}
	return err
}

func (vi *Video_ServerImp) InsertVeidoPageInfo(videoPage *server.VideoPageInfo, video *dao.Video, user_id uint64) (err error) {
	//到时候可以加入协程
	//获取所需的外部服务
	var server_utils Server_Utils
	var comment_serverImp Comment_ServerImp
	var relation_serverImp Relation_ServerImp
	var favorite_serverImp Favorite_ServerImp
	//组装video
	video.Playurl = server_utils.UrlUnParse(video.Playurl)
	video.Coverurl = server_utils.UrlUnParse(video.Coverurl)
	//组装videoPage
	com_count, err := comment_serverImp.CountByVideoID(video.ID)
	if err != nil {
		log.Println("获取count失败:", err.Error())
		err = nil
	}
	videoPage.Favorite_Count, err = favorite_serverImp.VideoFvaoriteCount(video.ID)
	if err != nil {
		log.Println("获取Favorite_Count失败:", err.Error())
		err = nil
	}
	if err != nil {
		return err
	}
	videoPage.Video = *video
	videoPage.Comment_Count = com_count
	//组装Author
	var user_severImp User_ServerImp
	author, err := user_severImp.GetUserPageInfo(uint64(video.UserID))
	videoPage.Author = *author
	work_count, _ := vi.PublishCount(uint64(video.UserID))
	videoPage.Author.WorkCount = work_count
	//组装需要关联用户的信息
	if user_id != 0 && user_id != uint64(video.UserID) {
		isFollow, err := relation_serverImp.IsFollow(user_id, uint64(video.UserID))
		if err != nil {
			log.Println("获取isFollow失败:", err.Error())
			err = nil
		}
		isFavorite, err := favorite_serverImp.Is_Favorite(video.ID, user_id)
		if err != nil {
			log.Println("获取isFavorite失败:", err.Error())
			err = nil
		}
		videoPage.Author.Is_Follow = isFollow
		videoPage.Is_Favorite = isFavorite
	}
	return err
}
func (vi *Video_ServerImp) Publish(file *multipart.FileHeader, user_id uint64, user_name string, title string) (err error) {
	rename := user_name + strconv.Itoa(int(user_id)) + strconv.FormatInt(time.Now().Unix(), 36)
	dir := "./static"
	dstvideo := dir + "/videos/" + rename + ".mp4"
	dstcover := dir + "/covers/" + rename + ".jpg"
	err = vi.saveVideoFile(file, dstvideo)
	if err != nil {
		log.Println("video_serverImp -> Publish -> saveVideoFile err:", err.Error())
		return errors.New("视频发布失败")
	}
	err = middleware.SaveFirstFrame(dstvideo, dstcover)
	if err != nil {
		log.Println("video_serverImp -> Publish-> SaveFirstFrame err:", err.Error())
		return errors.New("视频封面获取失败")
	}
	playUrl := "/static/videos/" + rename + ".mp4"
	coverUrl := "/static/covers/" + rename + ".jpg"
	video := dao.Video{
		Title:        title,
		UserID:       user_id,
		Playurl:      playUrl,
		Coverurl:     coverUrl,
		Published_At: time.Now(),
	}
	err = vi.InsertVideo(&video)
	if err != nil {
		log.Println("video_serverImp -> Publish -> InsertVideo err:", err.Error())
		return errors.New("视频创建失败")
	}
	return err
}

//保存video到本地
func (vi *Video_ServerImp) saveVideoFile(file *multipart.FileHeader, dst string) (err error) {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return nil
}

//保存video相关信息到数据库
func (vi *Video_ServerImp) InsertVideo(video *dao.Video) (err error) {
	var videoDao dao.Video
	err = videoDao.InsertVideo(video)
	if err != nil {
		log.Println("video_serverImp -> InsertVideo err:", err.Error())
		return err
	}
	return err
}

func (vi *Video_ServerImp) GetVideoIDList(user_id uint64) (videoId_list *[]uint64, err error) {
	var videoDao dao.Video
	videoId_list, err = videoDao.GetVideoIDList(user_id)
	if err != nil {
		log.Println("获取用户视频id列表失败:", err.Error())
		return videoId_list, err
	}
	return videoId_list, err
}

func (vi *Video_ServerImp) PublishCount(user_id uint64) (count int64, err error) {
	var videoDao dao.Video
	count, err = videoDao.CountVideoByUserID(user_id)
	if err != nil {
		log.Println("获取视频数量失败:", err.Error())
		err = nil
	}
	return count, err
}

//用于publish/list服务
func (vi *Video_ServerImp) GetVideoListByAuthorID(user_id uint64, cur_id uint64) (videos *[]server.VideoPageInfo, err error) {
	var videoDao dao.Video
	videoList, err := videoDao.GetVideoListByAuthorID(user_id)
	if err != nil {
		return videos, err
	}
	videosTemp := make([]server.VideoPageInfo, 0, len(*videoList))
	err = vi.ProcessVideos(&videosTemp, videoList, user_id)
	if err != nil {
		log.Println("video_serverImp->Feed err:", err.Error())
		return videos, errors.New("获取视频失败")
	}
	videos = &videosTemp
	return
}
