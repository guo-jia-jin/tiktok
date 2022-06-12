package server

type Relation_Server interface {
	//当前用户是否关注目标用户
	IsFollow(user_id uint64, tag_id uint64) (bool, error)
	//根据当前用户id查询其关注数
	GetFollowCount(user_id uint64) (int64, error)
	//根据当前用户id查询其的粉丝数
	GetFollowerCount(user_id uint64) (int64, error)

	//当前用户关注目标用户动作
	FollowAction(user_id uint64, tag_id uint64) error
	//当前用户取关目标用户动作
	UnFollowAction(user_id uint64, tag_id uint64) error
	//获取当前用户关注列表
	GetFollowList(user_id uint64) (*[]RelationPageUser, error)
	//获取当前用户粉丝列表
	GetFollowerList(user_id uint64) (*[]RelationPageUser, error)
}

type RelationPageUser struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	Follow_Count   int64  `json:"follow_count"`
	Follower_Count int64  `json:"follower_count"`
	Is_Follow      bool   `json:"is_follow"`
}
