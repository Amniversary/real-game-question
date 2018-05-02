package service

import (
	"context"
	"log"
	"fmt"
	"time"

	"github.com/Amniversary/real-game-question/models"
	"github.com/micro/go-micro/client"
	"github.com/jinzhu/now"
	proto "github.com/Amniversary/real-game-question/proto"
	userrpc "github.com/reechou/real-api-gateway/gateway/proto"
)

type Question struct {
	Client client.Client
}

var (
	GiftList = []*proto.GiftInfo{
		&proto.GiftInfo{Id: 1, Name: "美国兔邦尼长耳兔公仔", Img: "https://ylll111.xyz/23.png"},
		&proto.GiftInfo{Id: 2, Name: "调皮国宝大熊猫", Img: "https://ylll111.xyz/22.png"},
		&proto.GiftInfo{Id: 3, Name: "可爱小黄鸡公仔", Img: "https://ylll111.xyz/21.png"},
		&proto.GiftInfo{Id: 4, Name: "韩国布朗熊公仔", Img: "https://ylll111.xyz/8.png"},
		&proto.GiftInfo{Id: 5, Name: "Anthony安东尼不二兔", Img: "https://ylll111.xyz/7.png"},
		&proto.GiftInfo{Id: 6, Name: "可爱软萌饼干猫", Img: "https://ylll111.xyz/2.png"},
		&proto.GiftInfo{Id: 7, Name: "闪电皮卡丘公仔", Img: "https://ylll111.xyz/18.png"},
	}

	RankInfo = []*proto.RankInfo{
		&proto.RankInfo{PlayTimes: 743, NickName: "Gavean", AvatarUrl: "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTLWveUcua5t8zvZDoX9cSrOFgfb4pNel6pF0ia2xCfaGJXedDJpwdxTKawJiaEMicDXav96BYAZnRUWQ/0"},
		&proto.RankInfo{PlayTimes: 671, NickName: "东巴拉", AvatarUrl: "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTJUqvktkiaB0poEJISuA4f6VcGjFBliarSl1DjbLJibVIvw9QQ7RzPhySwrKvRo3TKdxODNCst1RWPIQ/0"},
		&proto.RankInfo{PlayTimes: 411, NickName: "伟仔", AvatarUrl: "https://wx.qlogo.cn/mmopen/vi_32/DYAIOgq83eprCjIwMn5f9CyL4z4lJK8NIbymn4XKjnf17lGkIRiawgS1nSpOuPU8sAiaGhulibqKduTOICiaibmWPMA/0"},
		&proto.RankInfo{PlayTimes: 287, NickName: "漫步人生", AvatarUrl: "https://wx.qlogo.cn/mmopen/vi_32/DYAIOgq83eq8Pu4XpTRQ9hRAWISSKaDnZKTCzwSPJVicsIq6seXic5Lc8j9yHg1xh9WMyOaQCpIgydnDbYwDKLlA/0"},
		&proto.RankInfo{PlayTimes: 267, NickName: "心简单，世界就简单°", AvatarUrl: "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTLmiaNPYbLOk7GFGmmx1Hw6VuuD1COmoGibbd8JibnyXibpIbpuAu7DeLwPz3ddDeArC6iaLmgYZ3uhicmA/0"},
		&proto.RankInfo{PlayTimes: 231, NickName: "白月光", AvatarUrl: "https://wx.qlogo.cn/mmopen/vi_32/DYAIOgq83eoor0YwdtnmTEwr6HyQJyqHFxuiaQbPPLUCwKiaBx5ibozdbPbM5J2TaQrttZB1CKlHALM6mEWgia1fJQ/0"},
		&proto.RankInfo{PlayTimes: 168, NickName: "晓风残月", AvatarUrl: "https://wx.qlogo.cn/mmopen/vi_32/DYAIOgq83eo5HP1ibenk8NrOaK196lSoaHQxtQGDMNoHiagyQgtuxlQ38AWBHrjAgOz9iaibsViaTJur8P7lAfIMLBQ/0"},
		&proto.RankInfo{PlayTimes: 140, NickName: "酿腐豆的酋长", AvatarUrl: "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIEONynt1pibqxP0FeC2s1viaiaEI6JxcjBiaczyvCLsFdVvekWwobXwKzr0WO5B8pGmwPSToGlBTDnvw/0"},
	}

	RankList = []*proto.RankList{
		&proto.RankList{Goods: 23, NickName: "熊二爱西瓜", AvatarUrl: "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTLz2MxhMpmo27nI3KfznJPgczf1Jl160HpSGaaSWGoYEcN2Ypx1o0cibLSiaVqaodZ980CnSG87mCVw/0"},
		&proto.RankList{Goods: 21, NickName: "红豆小姐", AvatarUrl: "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTL63dLLFyh9NCWcTBG6k5CeygMmhlqzYoBA9Pn8bo9ukUtRZsozAmSIKYstscAo5sXl3CWCHSzfiaA/0"},
		&proto.RankList{Goods: 18, NickName: "落伍小女人", AvatarUrl: "https://wx.qlogo.cn/mmopen/vi_32/DYAIOgq83eq1r0iaCPgAaoSbicNaaNTawp9VBVlSJHhQCwEib60zIenmaHxZRO6Ez4Okv92DPXOBuibnjVuibmicqWJg/0"},
		&proto.RankList{Goods: 17, NickName: "阿呆的阿傻", AvatarUrl: "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTKmXoYDRv98a26XwybLETD2sIEgTrCzw0NokqbyQlpLXOTGt9oKFApHoQPyF0TNqA4tflnIc1MiaWQ/0"},
		&proto.RankList{Goods: 17, NickName: "朱春芽", AvatarUrl: "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTJfDyRaMfF1JbfYvALt5GwMceia2WxdOOg34MbR75RzKCEhkicaTewiauAicIagJA1tHSASzK1SiaMsHpQ/0"},
		&proto.RankList{Goods: 14, NickName: "🍁詩億語🍑🍉", AvatarUrl: "https://wx.qlogo.cn/mmopen/vi_32/DYAIOgq83epgkfwib2s9d2GJfbl6TNCahTLo2JZDKvGaBA4jS42KtI1NdHA0dkoibIwicBiaNNgNXRAteoib2jc2GZQ/0"},
		&proto.RankList{Goods: 11, NickName: "Gex", AvatarUrl: "https://wx.qlogo.cn/mmopen/vi_32/Zd0jstSaf4uxHHvj7rsZPskHD1H0PH5FsEJE0CdibT7xDEicbRJAkq96sL8bqZQzgOLIicCM0GPKjjNHibgor6I8yg/0"},
		&proto.RankList{Goods: 10, NickName: "七七", AvatarUrl: "https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTK8UYbedASKg1eXSPTDjwf4WbjibRKCDX3ERtCUlUJZgjRUNBIwYUCJUelYADXf78azicWBk2nZFiarA/0"},
	}

	GiftShow              int64 = 1
	MaxShare              int64 = 10
	ChallengeTimes        int64 = 50000
	IndexTitle                  = "挑战成功赢娃娃"
	IndexShareTitle             = "[%s@我]发起挑战，会算1+2=？么..."
	NewShareTxt                 = "邀请好友一起来挑战"
	IndexShareImg               = "https://ylll111.xyz/jjds_wawa.png"
	IndexHtmlTitle1             = "邀请好友一起来挑战"
	IndexBarTitle               = "加减"
	IndexShareTxt               = "约群朋友一起来挑战"
	GameRule1                   = "登录自动获得一次挑战机会，每天可以分享到10个不同的群，获得10次免费机会。"
	GameRule2                   = "由于微信版本更新，可能会导致小程序的某些功能出现异常，请退出微信后重新打开。"
	GameRule3                   = "挑战过程中答题时间会随着挑战进行而缩短。"
	GameRule4                   = "挑战过程中，在规定时间内连续判断对40道题简单的数字加减题，挑战成功。"
	GameRule5                   = "挑战成功后，您可以免费挑选娃娃，（如果挑选的娃娃缺货，将随机发货），填写领取信息，客服按照申请顺序发货，娃娃包邮！"
	UserHtmlTitle1              = "邀请好友一起来挑战"
	IndexPersonShareTitle       = "分享成功，分享到群能获得再次挑战机会"
	GamePersonShareTitle        = "分享成功，分享到群能获得再次挑战机会"
	UserPersonShareTitle        = "分享成功，分享到群能获得再次挑战机会"
	GameFailedShareTitle        = "[%s@我]发起挑战，会算1+2=？么..."
	GameSuccessShareTitle       = "[%s@我]发起挑战，会算1+2=？么..."
	GameFailedShareImg          = "https://ylll111.xyz/jjds_wawa.png"
	GameSuccessShareImg         = "https://ylll111.xyz/jjds_wawa.png"
	FailedShareText1            = "差一点就能成功了"
	FailedShareText2            = "加油，你可以的"
	FailedShareText3            = "再来一次"
	FailedShareText4            = "获得更多挑战机会"
	SuccessShareText1           = "你很厉害哦"
	SuccessShareText2           = "12个赞送给你"
	SuccessShareText3           = "获得娃娃+1"
	SuccessShareText4           = "请前往我的战绩查看"
	SuccessShareText5           = "炫耀一下"
	GameBarTitle                = "加减大师"
	UserShareTitle              = "[%s@我]发起挑战，会算1+2=？么..."
	UserShareImg                = "https://ylll111.xyz/jjds_wawa.png"
	UserBarTitle                = "加减大师"
	GiftBarTitle                = "加减大师"
	WzIf                        = "0"
	TzType                      = "1"
	FailedClick                 = "2"
)

// todo: init user info
func (q *Question) Index(ctx context.Context, req *proto.IndexRequest, rsp *proto.IndexResponse) error {
	rsp.Status = &proto.RspStatus{Code: RSP_SUCCESS}
	rsp.Config = &proto.IndexConfig{}
	rsp.Config.IndexTitle = IndexTitle
	rsp.Config.IndexShareTitle = fmt.Sprintf(IndexShareTitle, req.Name)
	rsp.Config.NewShareTxt = NewShareTxt
	rsp.Config.IndexShareImg = IndexShareImg
	rsp.Config.IndexHtmlTitle1 = IndexHtmlTitle1
	rsp.Config.IndexBarTitle = IndexBarTitle
	rsp.Config.GameRule1 = GameRule1
	rsp.Config.GameRule2 = GameRule2
	rsp.Config.GameRule3 = GameRule3
	rsp.Config.GameRule4 = GameRule4
	rsp.Config.GameRule5 = GameRule5
	rsp.Config.MaxShare = MaxShare
	rsp.Config.GiftShow = GiftShow
	rsp.Config.SuccessShareText1 = SuccessShareText1
	rsp.Config.SuccessShareText2 = SuccessShareText2
	rsp.Config.SuccessShareText3 = SuccessShareText3
	rsp.Config.SuccessShareText4 = SuccessShareText4
	rsp.Config.SuccessShareText5 = SuccessShareText5
	rsp.Config.UserHtmlTitle1 = UserHtmlTitle1
	rsp.Config.IndexPersonShareTitle = IndexPersonShareTitle
	rsp.Config.GameBarTitle = GameBarTitle
	rsp.Config.WzIf = WzIf
	rsp.Config.TzType = TzType
	rsp.Config.GamePersonShareTitle = GamePersonShareTitle
	rsp.Config.UserPersonShareTitle = UserPersonShareTitle
	rsp.Config.GameFailedShareTitle = fmt.Sprintf(GameFailedShareTitle, req.Name)
	rsp.Config.GameSuccessShareTitle = fmt.Sprintf(GameSuccessShareTitle, req.Name)
	rsp.Config.GameFailedShareImg = GameFailedShareImg
	rsp.Config.GameSuccessShareImg = GameSuccessShareImg
	rsp.Config.FailedShareText1 = FailedShareText1
	rsp.Config.FailedShareText2 = FailedShareText2
	rsp.Config.FailedShareText3 = FailedShareText3
	rsp.Config.FailedShareText4 = FailedShareText4
	rsp.Config.UserShareTitle = fmt.Sprintf(UserShareTitle, req.Name)
	rsp.Config.UserShareImg = UserShareImg
	rsp.Config.UserBarTitle = UserBarTitle
	rsp.Config.GiftBarTitle = GiftBarTitle
	rsp.Config.FailedClick = FailedClick
	rsp.Config.IndexShareTxt = IndexShareTxt
	rsp.GiftInfo = GiftList
	rsp.Rank = &proto.Rank{RankList: RankList, RankInfo: RankInfo}
	times := models.GetChallengeTimes()
	rsp.ChallengeTimes = times + ChallengeTimes
	// get user info
	userClient := userrpc.NewUserService(USER_SERVER_NAME, q.Client)
	userInfo, err := userClient.GetUserList(ctx, &userrpc.GetUserListRequest{
		AppId:      THIS_APPID,
		UserIdList: []int64{req.UserId},
	})
	if err != nil || len(userInfo.UserList) == 0 {
		log.Printf("[rpc] get user list err: [%v]", err)
		rsp.Status.Code = RSP_ERROR
		rsp.Status.Msg = GET_USER_INFO_MSG
		return nil
	}
	nowStart := now.BeginningOfDay().Unix()
	// get user game info
	user := &models.UserGame{UserId: req.UserId}
	has := models.GetUser(user)
	if !has {
		log.Printf("user info is empty")
	}
	log.Printf("userInfo: [%v]", user)
	if !has {
		user.Chance = 1
		user.LoginDay = nowStart
		user.Sign = fmt.Sprintf("%d_%d", time.Now().UnixNano(), user.UserId)
		if err := models.CreateUser(user); err != nil {
			log.Printf("create user game info err: [%v]", err)
			return err
		}
	} else {
		if nowStart > user.LoginDay {
			user.Chance += 1
			user.LoginDay = nowStart
			if err := models.UpdateUserInfo(user); err != nil {
				log.Printf("update user game info err: [%v]", err)
				return err
			}
		}
	}
	rsp.PlayTime = user.GameNum
	rsp.Ss = 1
	rsp.Sign = user.Sign
	rsp.TodayShares = models.GetUserShareCount(req.UserId, now.BeginningOfDay().Unix())
	rsp.Chance = user.Chance
	rsp.Score = user.Score
	rsp.Goods = user.Goods
	return nil
}

// todo: 获取题目列表
func (q *Question) GetQuestionList(ctx context.Context, req *proto.GetQuestionRequest, rsp *proto.GetQuestionResponse) error {
	rsp.Status = &proto.RspStatus{Code: RSP_SUCCESS}
	if req.Sign == "" || req.UserId == 0 {
		log.Printf("params can't be empty: [%v]", req)
		rsp.Status.Code = RSP_ERROR
		rsp.Status.Msg = "params can't be empty."
		return nil
	}
	user := &models.UserGame{UserId: req.UserId}
	has := models.GetUser(user)
	if !has {
		rsp.Status.Code = RSP_ERROR
		rsp.Status.Msg = GET_USER_INFO_MSG
		return nil
	}
	user.GameNum += 1
	user.Chance -= 1
	user.GameSign = fmt.Sprintf("%d_%d_%d", time.Now().UnixNano(), user.UserId, user.GameNum)
	if err := models.UpdateUserGameInfo(user); err != nil {
		return err
	}
	data, err := models.GetQuestionList()
	if err != nil {
		log.Printf("get question list query err: [%v]", err)
		return err
	}
	rsp.GameStatusSign = user.GameSign
	rsp.Data = data
	return nil
}

// todo: 获取用户分享数据
func (q *Question) GetUserShare(ctx context.Context, req *proto.GetUserShareRequest, rsp *proto.GetUserShareResponse) error {
	rsp.Status = &proto.RspStatus{Code: RSP_SUCCESS}
	if req.UserId == 0 || req.EncryptedData == "" || req.Iv == "" {
		log.Printf("params can't be empty: [%v]", req)
		rsp.Status.Code = RSP_ERROR
		rsp.Status.Msg = "params can't be empty."
		return nil
	}

	user := &models.UserGame{UserId: req.UserId}
	has := models.GetUser(user)
	if !has {
		rsp.Status.Code = RSP_ERROR
		rsp.Status.Msg = GET_USER_INFO_MSG
		return nil
	}
	// todo: rpc - get user info
	userClient := userrpc.NewUserService(USER_SERVER_NAME, q.Client)
	userShare, err := userClient.GetShareInfo(ctx, &userrpc.GetShareInfoRequest{
		AppId:         THIS_APPID,
		UserId:        req.UserId,
		EncryptedData: req.EncryptedData,
		Iv:            req.Iv,
	})
	if err != nil || userShare.OpenGid == "" {
		log.Printf("get user share info err: [%v]", err)
		rsp.Status.Code = RSP_ERROR
		rsp.Status.Msg = GET_USER_SHARE_INFO_MSG
		return nil
	}
	nowTime := now.BeginningOfDay().Unix()
	isShareSuccess := false
	shareInfo := &models.UserShare{UserId: req.UserId, OpenGid: userShare.OpenGid}
	Ok := models.GetUserShare(shareInfo)
	if !Ok {
		isShareSuccess = true
		shareInfo.Num = 1
		shareInfo.Date = nowTime
		if err := models.CreateUserShare(shareInfo); err != nil {
			log.Printf("create user share info err: [%v]", err)
			return err
		}
	} else {
		if nowTime > shareInfo.Date {
			isShareSuccess = true
			shareInfo.Date = nowTime
		}
		shareInfo.Num += 1
		if err := models.UpdateUserShare(shareInfo); err != nil {
			return err
		}
	}
	rsp.Repeat = 0
	if isShareSuccess {
		user.Chance += 1
		if err := models.UpdateUserInfo(user); err != nil {
			return err
		}
		rsp.Repeat = 1
	}
	rsp.ErrorCode = 1
	rsp.Chance = user.Chance
	rsp.TodayShares = models.GetUserShareCount(user.UserId, nowTime)
	return nil
}

// todo: 游戏结果记录
func (q *Question) UploadResult(ctx context.Context, req *proto.UploadResultRequest, rsp *proto.UploadResultResponse) error {
	rsp.Status = &proto.RspStatus{Code: RSP_SUCCESS}
	if req.UserId == 0 {
		log.Printf("params can't be empty: [%v]", req)
		rsp.Status.Code = RSP_ERROR
		rsp.Status.Msg = "params can't be empty."
		return nil
	}
	user := &models.UserGame{UserId: req.UserId}
	has := models.GetUser(user)
	if !has {
		rsp.Status.Code = RSP_ERROR
		rsp.Status.Msg = GET_USER_INFO_MSG
		return nil
	}
	if req.GameStatusSign != user.GameSign {
		log.Printf("game stop game sign:[%v] != req game sign:[%v]", user.GameSign, req.GameStatusSign)
		return fmt.Errorf("sign key err.")
	}
	if req.RightNums > user.Score {
		user.Score = req.RightNums
		if err := models.UpdateUserInfo(user); err != nil {
			return err
		}
	}
	return nil
}
