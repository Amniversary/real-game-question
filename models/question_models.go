package models

import (
	"log"
	"time"

	proto "github.com/Amniversary/real-game-question/proto"
)

func GetUserShare(userShare *UserShare) bool {
	if err := db.Where("user_id = ? and open_gid = ?", userShare.UserId, userShare.OpenGid).First(&userShare).Error; err != nil {
		log.Printf("query first user share info err: [%v]", err)
	}
	if userShare.ID == 0 {
		return false
	}
	return true
}

func GetUserShareCount(userId int64, nowTime int64) int64 {
	var count int64
	err := db.Model(&UserShare{}).Where("user_id = ? and date = ?", userId, nowTime).Count(&count).Error
	if err != nil {
		log.Printf("get user share count err: [%v]", err)
		return 0
	}
	return count
}

func CreateUserShare(userShare *UserShare) error {
	userShare.CreateAt = time.Now().Unix()
	if err := db.Create(&userShare).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserShare(userShare *UserShare) error {
	err := db.Model(&UserShare{}).Where("user_id = ? and open_gid = ?", userShare.UserId, userShare.OpenGid).Update(&userShare).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUser(user *UserGame) (bool) {
	if err := db.Where("user_id = ?", user.UserId).First(&user).Error; err != nil {
		log.Printf("getUser query first err: [%v]", err)
	}
	if user.ID == 0 {
		return false
	}
	return true
}

func CreateUser(user *UserGame) error {
	user.CreateAt = time.Now().Unix()
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserInfo(user *UserGame) error {
	if err := db.Model(&UserGame{}).Where("user_id = ?", user.UserId).Update(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetQuestionList() ([]*proto.QuestionList, error) {
	var data []*proto.QuestionList
	for i := 1; i < 5; i++ {
		var info []*proto.QuestionList
		err := db.Table("question").
			Select("`num1`, `operator`, `num2`, `result`, `success`, `seconds`").
			Where("`level` = ?", i).
			Limit(10).Order("rand()").Find(&info).Error
		if err != nil {
			return nil, err
		}
		for _, v := range info {
			data = append(data, v)
		}
	}
	return data, nil
}
