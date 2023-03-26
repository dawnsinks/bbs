package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId       int64  `gorm:"user_id" json:"user_id,string"`
	UserName     string `gorm:"user_name" json:"user_name"`
	Password     string `gorm:"password" json:"password"`
	Email        string `gorm:"email" json:"email"`
	Gender       int    `gorm:"gender" json:"gender"`
	Avatar       string `gorm:"avatar" json:"avatar"`
	Introduction string `gorm:"Introduction" json:"introduction"`
}
