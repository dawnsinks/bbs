package mysql

import (
	"bbs/models"
	"errors"
	"gorm.io/gorm"
)

func CheckUserExist(uname string) bool {
	var u models.User

	result := db.First(&u, "user_name = ?", uname)

	if result.Error == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func InsertUser(u *models.User) (err error) {
	return db.Create(&u).Error
}

func Login(u *models.User) (err error) {
	result := db.First(&u, "user_name = ? and password = ?", u.UserName, u.Password)
	if result.Error == gorm.ErrRecordNotFound {
		return errors.New("用户名或密码错误")
	}
	return
}
