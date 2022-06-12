package serverimp

import (
	"Tiktok/dao"
	"Tiktok/service/server"
)

type Comment_ServerImp struct {
}

var comDao dao.Comment

func (co *Comment_ServerImp) CountByVideoID(id uint64) (count int64, err error) {
	var comDao dao.Comment
	count, err = comDao.CountByVideoID(id)
	if err != nil {
		return count, err
	}
	return count, err
}

func (co *Comment_ServerImp) SendComment(comment *dao.Comment) (comInfo server.CommentInfo, err error) {
	var user_serverImp User_ServerImp
	comment, err = comment.CreateComment(comment)
	if err != nil {
		return comInfo, err
	}
	//获取视频评论用户信息，用于组装返回体
	userPageInfo, err := user_serverImp.GetUserPageInfo(comment.User_ID)
	if err != nil {
		return comInfo, err
	}
	//组装返回体
	comInfo = server.CommentInfo{
		ID:         comment.ID,
		UserInfo:   *userPageInfo,
		Content:    comment.Content,
		CreateDate: comment.Create_Date.Format("01-02"),
	}
	return comInfo, err
}

func (co *Comment_ServerImp) DelCommentByID(comment_id uint64) (err error) {
	err = comDao.DeleteCommentByID(comment_id)
	if err != nil {
		return err
	}
	return err
}

func (co *Comment_ServerImp) GetCommentList(video_id uint64, user_id uint64) (comList *[]server.CommentInfo, err error) {
	var user_serverImp User_ServerImp
	commentList, err := comDao.GetCommentListByVideoID(video_id)
	if err != nil {
		return comList, err
	}
	com := make([]server.CommentInfo, 0, len(*commentList))
	for _, comment := range *commentList {
		userIfno, _ := user_serverImp.GetUserPageInfo(comment.User_ID)
		comtemp := server.CommentInfo{
			ID:         comment.ID,
			UserInfo:   *userIfno,
			Content:    comment.Content,
			CreateDate: comment.Create_Date.Format("01-02"),
		}
		com = append(com, comtemp)
	}
	comList = &com
	return comList, err
}
