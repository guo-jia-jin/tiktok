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
	"sync"
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
	var relation_serverImp Relation_ServerImp
	var favorite_serverImp Favorite_ServerImp
	//创建协程组，当这一组的协程全部完成，才会结束本方法
	var wg sync.WaitGroup
	wg.Add(4)                      //等待协程数
	errChan := make(chan error, 3) //记录主要协程发生错误的协程
	//获取user用于组装返回体的协程
	var user *dao.User
	go func() {
		defer wg.Done()
		user, err = user.GetUserByID(user_id)
		if err != nil {
			log.Println("user_severImp->GetUserPageInfo->in goroutine-1 happened err:", err.Error())
			errChan <- err
		}
	}()
	//获取userInfo用于组装返回体的协程
	var userInfo *dao.UserInfo
	go func() {
		defer wg.Done()
		userInfo, err = userInfo.GetUserInfoByID(user_id)
		if err != nil {
			log.Println("user_severImp->GetUserPageInfo->in goroutine-2 happened err:", err.Error())
			errChan <- err
		} else {
			//生成userInfo中的可访问url
			var server_utils Server_Utils
			userInfo.Avatar = server_utils.UrlUnParse(userInfo.Avatar)
			userInfo.Background = server_utils.UrlUnParse(userInfo.Background)
		}
	}()
	//获取follow_conut和follower_count的协程
	var follow_count int64
	var follower_count int64
	go func() {
		defer wg.Done()
		var e error
		follow_count, e = relation_serverImp.GetFollowCount(user_id)
		follower_count, e = relation_serverImp.GetFollowerCount(user_id)
		if e != nil {
			log.Println("user_severImp->GetUserPageInfo->in goroutine-3 happened err:", err.Error())
		}
	}()
	var favorite_count int64
	var total_favorite int64
	//获取favorite_count和total_favorite的协程
	go func() {
		defer wg.Done()
		var e error
		favorite_count, e = favorite_serverImp.UserFavoriteCount(user_id)
		total_favorite, e = favorite_serverImp.TotalUserFavorite(user_id)
		if e != nil {
			log.Println("user_severImp->GetUserPageInfo->in goroutine-4 happened err:", err.Error())
		}
	}()
	//等待上方4个协程结束
	go func() {
		defer close(errChan)
		//log.Println("wait")
		wg.Wait()
	}()
	for e := range errChan {
		if e != nil {
			return nil, e
		}
	}
	//组装数据
	userPageInfo = &server.UserPageInfo{
		UserInfo:       *userInfo,
		Name:           user.Name,
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
