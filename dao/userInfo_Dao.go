package dao

import (
	"Tiktok/global"
	"log"
)

type UserInfo struct {
	UserID     uint64 `json:"id"`               //用户唯一ID
	Avatar     string `json:"avatar"`           //用户头像url
	Background string `json:"background_image"` //用户主页背景url
	Signature  string `json:"signature"`        //用户签名
}

func (UserInfo) TableName() string {
	return "userinfos"
}

func (us *UserInfo) GetUserInfoByID(userId uint64) (userInfo *UserInfo, err error) {
	err = global.GVA_DB.Where(map[string]interface{}{"user_id": userId}).First(&userInfo).Error
	if err != nil {
		log.Println("UserInfo_Dao->GetUserInfoByID: find user_id = ", userId, " failed")
		log.Println("err:", err.Error())
		return nil, err
	}
	return userInfo, err
}

func (us *UserInfo) CreateUserInfo(userInfo *UserInfo) (err error) {
	err = global.GVA_DB.Model(UserInfo{}).Create(&userInfo).Error
	if err != nil {
		log.Println("UserInfo_Dao->CreateUserInfo: created faild")
		log.Println("err:", err.Error())
		return err
	}
	return err
}
