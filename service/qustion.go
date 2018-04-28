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
	GiftShow              int64 = 1
	MaxShare              int64 = 10
	IndexTitle                  = "挑战成功赢娃娃"
	IndexShareTitle             = "[%s@我]发起挑战，会算1+2=？么..."
	NewShareTxt                 = "邀请好友一起来挑战"
	IndexShareImg               = "https://ylll111.xyz/jjds_wawa.png"
	IndexHtmlTitle1             = "邀请好友一起来挑战"
	IndexBarTitle               = "加减大师"
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
)

// todo: init user info
func (q *Question) Index(ctx context.Context, req *proto.IndexRequest, rsp *proto.IndexResponse) error {
	rsp.Status = &proto.RspStatus{Code: RSP_SUCCESS}
	rsp.Config = &proto.IndexConfig{}
	rsp.Config.IndexTitle = IndexTitle
	rsp.Config.IndexShareTitle = IndexShareTitle
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
	rsp.Config.GameFailedShareTitle = GameFailedShareTitle
	rsp.Config.GameSuccessShareTitle = GameSuccessShareTitle
	rsp.Config.GameFailedShareImg = GameFailedShareImg
	rsp.Config.GameSuccessShareImg = GameSuccessShareImg
	rsp.Config.FailedShareText1 = FailedShareText1
	rsp.Config.FailedShareText2 = FailedShareText2
	rsp.Config.FailedShareText3 = FailedShareText3
	rsp.Config.FailedShareText4 = FailedShareText4
	rsp.Config.UserShareTitle = UserShareTitle
	rsp.Config.UserShareImg = UserShareImg
	rsp.Config.UserBarTitle = UserBarTitle
	rsp.Config.GiftBarTitle = GiftBarTitle

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
	if err := models.UpdateUserInfo(user); err != nil {
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
		log.Printf("")
	}
	if req.RightNums > user.Score {
		user.Score = req.RightNums
		if err := models.UpdateUserInfo(user); err != nil {
			return err
		}
	}
	return nil
}
