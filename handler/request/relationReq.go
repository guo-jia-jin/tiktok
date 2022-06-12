package request

type RelationActionReq struct {
	Token       string `form:"token" binding:"required"`
	To_User_ID  uint   `form:"to_user_id" binding:"required"`
	Action_Type uint   `form:"action_type" binding:"required"`
}
