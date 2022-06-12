package serverimp

import (
	"Tiktok/dao"
	"Tiktok/handler/request"
	"Tiktok/middleware"
	"Tiktok/service/server"
	"errors"
	"log"
	"math/rand"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type User_ServerImp struct {
}

func (us *User_ServerImp) Register(userRegInfo *request.UserRequest) (user *dao.User, err error) {
	userTemp := dao.User{}
	isExist := user.IsExistByName(userRegInfo.UserName)
	if isExist {
		return user, errors.New("User name already exists")
	} else {
		userTemp.Name = userRegInfo.UserName
		userTemp.Password = middleware.BcryptHash(userRegInfo.Password)
		userTemp.Created_At = time.Now()
		user, err = user.CreateUser(&userTemp)

		if err != nil {
			log.Println("user_severImp -> Register err1:", err.Error())
			return user, errors.New("创建用户失败")
		}
		err := us.CreateUserInfo(user.ID)
		if err != nil {
			log.Println("user_severImp -> Register err2:", err.Error())
			return user, errors.New("创建用户相关信息失败")
		}
	}
	return user, err
}

func (us *User_ServerImp) Login(userLogInfo *request.UserRequest) (user *dao.User, err error) {
	user, err = user.GetUserByName(userLogInfo.UserName)
	//fmt.Printf("user: %v\n", *user)
	//fmt.Printf("userLogInfo: %v\n", *userLogInfo)
	if err == nil {
		if ok := middleware.BcryptCheck(userLogInfo.Password, user.Password); !ok {
			return nil, errors.New("用户名或者密码不正确")
		}
	} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("该用户不存在")
	} else {
		log.Println("user_serverImp -> Login err:", err.Error())
		return nil, errors.New("登录异常")
	}
	return user, err
}

func (us *User_ServerImp) GetUserPageInfo(user_id uint64) (userPageInfo *server.UserPageInfo, err error) {
	var user *dao.User
	var relation_serverImp Relation_ServerImp
	var favorite_serverImp Favorite_ServerImp
	//获取user
	user, err = user.GetUserByID(user_id)
	if err != nil {
		return userPageInfo, err
	}
	//通过user获取用户名
	userName := user.Name
	var userInfo *dao.UserInfo
	//获取userInfo用于组装返回体
	userInfo, err = userInfo.GetUserInfoByID(user_id)
	//生成userInfo中的可访问url
	var server_utils Server_Utils
	userInfo.Avatar = server_utils.UrlUnParse(userInfo.Avatar)
	userInfo.Background = server_utils.UrlUnParse(userInfo.Background)
	if err != nil {
		return userPageInfo, err
	}
	//获取follow_conut和follower_count
	follow_count, err := relation_serverImp.GetFollowCount(user_id)
	follower_count, err := relation_serverImp.GetFollowerCount(user_id)
	favorite_count, err := favorite_serverImp.UserFavoriteCount(user_id)
	total_favorite, err := favorite_serverImp.TotalUserFavorite(user_id)
	//组装数据
	userPageInfo = &server.UserPageInfo{
		UserInfo:       *userInfo,
		Name:           userName,
		Follow_Count:   follow_count,
		Follower_Count: follower_count,
		FavoriteCount:  favorite_count,
		TotalFavorited: total_favorite,
	}
	return userPageInfo, err
}

func (us *User_ServerImp) CreateUserInfo(user_id uint64) (err error) {
	var userInfo dao.UserInfo
	//设置默认头像和背景
	rand.Seed(time.Now().Unix())
	random := rand.Intn(2) + 1
	name := "default" + strconv.Itoa(random)
	avatarsUrl := "/static/avatars/" + name + ".jpg"
	background := "/static/backgrounds/" + "default1bk.jpg"
	//组装用户信息
	userInfo.UserID = user_id
	userInfo.Avatar = avatarsUrl
	userInfo.Background = background
	err = userInfo.CreateUserInfo(&userInfo)
	if err != nil {
		return err
	}
	return err
}
