package dao

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "root:123456@tcp(localhost:3306)/dousheng_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println("connect fail:", err.Error())
	}
	return db
}
func TestCountByVideoID(t *testing.T) {
	DB := Init()
	var commentList *[]Comment
	err := DB.Model(Comment{}).Where(map[string]interface{}{"video_id": 16}).
		Find(&commentList).Order("create_date desc").Error
	if err != nil {
		log.Println("Comment_Dao-GetCommentListByVideoID:failed")
		log.Println("err:", err.Error())
	}
	log.Println("Comment_Dao-GetCommentListByVideoID:success")
	fmt.Printf("commentList: %v\n", *commentList)
}
