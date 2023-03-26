package mysql

import (
	"bbs/models"
	"bbs/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() (err error) {
	url := settings.DBURL()
	db, err = gorm.Open(mysql.Open(url))
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.User{}, &models.Community{})

	if err != nil {
		return err
	}
	return
}
