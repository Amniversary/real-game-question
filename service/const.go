package service

const (
	Empty            int64 = 0
	RSP_SUCCESS            = 1000
	RSP_ERROR              = 1
	THIS_APPID       int64 = 2
	USER_SERVER_NAME       = "RealApiGateway"
)

const (
	GET_USER_SHARE_INFO_MSG = "获取用户分享信息失败"
	GET_USER_INFO_MSG       = "获取用户信息失败"
	SYSTEM_ERRPR_MSG        = "系统错误"

	SHARE_SUCCESS_MSG  = "分享成功，挑战次数+1"
	SHARE_REPEATED_MSG = "本群今日已经分享过了"
	USER_GET_GIFT_NO_LIMIT = "用户领取次数已经用完"
)
