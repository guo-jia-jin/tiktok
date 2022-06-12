package server

type Favorite_Server interface {
	//根据当前视频id和用户id是否点赞该视频
	IsFavorite(video_id uint64, user_id uint64) (bool error)
	//根据视频id获取此视频的点赞数
	VideoFvaoriteCount(video_id uint64) (int64, error)
	//根据用户id获取此用户被点赞数
	TotalUserBeFavorited(user_id uint64) (int64, error)
	//根据用户id获取这个用户点赞视频数量
	UserFavoriteCount(user_id uint64) (int64, error)

	//点赞动作
	FavoriteAction(user_id uint64, video_id uint64, action_type uint16) error
	//获取当前用户的所有点赞视频
	GetUserFavoriteList(user_id uint64, cur_id uint64) (*[]VideoPageInfo, error)
}
