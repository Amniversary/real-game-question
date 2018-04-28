package service

import (
	"context"
	"log"

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
	GiftShow        int64 = 1
	MaxShare        int64 = 10
	IndexTitle            = "挑战成功赢娃娃"
	IndexShareTitle       = "约群朋友一起来挑战"
	NewShareTxt           = "邀请好友一起来挑战"
	IndexShareImg         = ""
	IndexHtmlTitle1       = ""
	IndexHtmlTitle2       = ""
	IndexBarTitle         = ""
	GameRule1             = "登录自动获得一次挑战机会，每天可以分享到10个不同的群，获得10次免费机会。"
	GameRule2             = "由于微信版本更新，可能会导致小程序的某些功能出现异常，请退出微信后重新打开。"
	GameRule3             = "挑战过程中答题时间会随着挑战进行而缩短。"
	GameRule4             = "挑战过程中，在规定时间内连续判断对40道题简单的数字加减题，挑战成功。"
	GameRule5             = "挑战成功后，您可以免费挑选娃娃，（如果挑选的娃娃缺货，将随机发货），填写领取信息，客服按照申请顺序发货，娃娃包邮！"
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
	rsp.Config.IndexHtmlTitle2 = IndexHtmlTitle2
	rsp.Config.IndexBarTitle = IndexBarTitle
	rsp.Config.GameRule1 = GameRule1
	rsp.Config.GameRule2 = GameRule2
	rsp.Config.GameRule3 = GameRule3
	rsp.Config.GameRule4 = GameRule4
	rsp.Config.GameRule5 = GameRule5
	rsp.Config.MaxShare = MaxShare
	rsp.Config.GiftShow = GiftShow
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
	rsp.TodayShares = models.GetUserShareCount(req.UserId, now.BeginningOfDay().Unix())
	rsp.Chance = user.Chance
	rsp.Score = user.Score
	rsp.Success = user.Success
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
	data, err := models.GetQuestionList()
	if err != nil {
		log.Printf("get question list query err: [%v]", err)
		return err
	}
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
	if req.Score > user.Score {
		user.Score = req.Score
		if err := models.UpdateUserInfo(user); err != nil {
			return err
		}
	}
	return nil
}
