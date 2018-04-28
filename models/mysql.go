package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Amniversary/real-game-question/config"
	"fmt"
	"log"
)

var db *gorm.DB

func NewMysql(c *config.Config) {
	openDb(c)
}

func openDb(c *config.Config) {
	db1, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&loc=Local",
		c.User,
		c.Pass,
		c.Host,
		c.DBName,
	))
	if err != nil {
		log.Printf("init database err: [%v]", err)
		return
	}
	if c.IfShowSql {
		db1.LogMode(true)
	}

	db = db1
	db.DB().SetMaxIdleConns(0)
	initTable()
}

func initTable() {
	db.AutoMigrate(new(Question), new(UserShare), new(UserGame))
}