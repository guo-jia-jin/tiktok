package dao

import (
	"Tiktok/global"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint64    //用户唯一id
	Name       string    //用户名
	Password   string    //用户密码
	Created_At time.Time //用户创建时间
}

func (User) TableName() string {
	return "users"
}

func (u *User) CreateUser(userInfo *User) (user *User, err error) {
	err = global.GVA_DB.Model(User{}).Create(&userInfo).Error
	if err != nil {
		log.Println("User_Dao-CreateUser: failed")
		log.Println("err:", err.Error())
		return nil, err
	}
	user = userInfo
	return user, err
}

func (u *User) GetUserByName(userName string) (user *User, err error) {
	err = global.GVA_DB.Where(map[string]interface{}{"name": userName}).First(&user).Error
	if err != nil {
		log.Println("User_Dao-GetUserByName: find name = ", userName, " failed")
		log.Println("err:", err)
		return nil, err
	}
	return user, err
}

func (u *User) GetUserByID(userId uint64) (user *User, err error) {
	err = global.GVA_DB.Where(map[string]interface{}{"id": userId}).First(&user).Error
	if err != nil {
		log.Println("User_Dao-GetUserByID: find id = ", userId, " failed")
		log.Println("err:", err)
		return nil, err
	}
	return user, err
}

func (u *User) IsExistByName(userName string) (isExist bool) {
	user := User{}
	if !errors.Is(global.GVA_DB.Where("name = ?", userName).First(&user).Error, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}
