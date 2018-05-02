package models

type UserGame struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	UserId   int64  `gorm:"not null; default:0; type:int; index" json:"userId"`
	Chance   int64  `gorm:"not null; default:0; type:int" json:"chance"`
	Success  int64  `gorm:"not null; default:0; type:int" json:"success"`
	Score    int64  `gorm:"not null; default:0; type:int" json:"score"`
	GameNum  int64  `gorm:"not null; default:0; type:int" json:"game_num"`
	LoginDay int64  `gorm:"not null; default:0; type:int" json:"loginDay"`
	Goods    int64  `gorm:"not null; default:0; type:int" json:"goods"`
	Sign     string `gorm:"not null; default:''; type:varchar(256)" json:"sign"`
	GameSign string `gorm:"not null; default:''; type:varchar(256)" json:"game_sign"`
	CreateAt int64  `gorm:"not null; default:0; type:int" json:"createAt"`
}

// todo: 用户游戏表
func (UserGame) TableName() string {
	return "user_game"
}

type Question struct {
	ID       int64   `gorm:"primary_key" json:"id"`
	Num1     int64   `gorm:"not null; default:0; type:int" json:"num1"`
	Operator string  `gorm:"not null; default:''; type:varchar(128)" json:"operator"`
	Num2     int64   `gorm:"not null; default:0; type:int" json:"num2"`
	Result   int64   `gorm:"not null; default:0; type:int" json:"result"`
	Success  int64   `gorm:"not null; default:0; type:int" json:"success"`
	Seconds  float64 `gorm:"not null; default:0; type:float" json:"seconds"`
	Level    int64   `gorm:"not null; default:0; type:int; index" json:"level"`
}

// todo: 问题表
func (Question) TableName() string {
	return "question"
}

type UserShare struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	UserId   int64  `gorm:"not null; default:0; type:int; index" json:"userId"`
	OpenGid  string `gorm:"not null; default:0; type:varchar(256); index; column:open_gid" json:"openGid"`
	Num      int64  `gorm:"not null; default:0; type:int" json:"num"`
	Date     int64  `gorm:"not null; default:''; type:int" json:"date"`
	CreateAt int64  `gorm:"not null; default:0; type:int" json:"createAt"`
}

// todo: 用户分享数据
func (UserShare) TableName() string {
	return "user_share"
}

type GiftResult struct {
	ID         int64  `gorm:"primary_key" json:"id"`
	UserId     int64  `gorm:"not null; default:0; type:int; index" json:"userId"`
	GiftId     int64  `gorm:"not null; default:0; type:int; index" json:"giftId"`
	GiftImgUrl string `gorm:"not null; default:''; type:varchar(256)" json:"gift_img_url"`
	GiftName   string `gorm:"not null; default:''; type:varchar(128)" json:"gift_name"`
	RealName   string `gorm:"not null; default:''; type:varchar(128)" json:"real_name"`
	Phone      string `gorm:"not null; default:''; type:varchar(128)" json:"phone"`
	Address    string `gorm:"not null; default:''; type:varchar(256)" json:"address"`
	CreateAt   int64  `gorm:"not null; default:0; type:int" json:"create_at"`
}

// todo: 礼物中奖表
func (GiftResult) TableName() string {
	return "gift_result"
}