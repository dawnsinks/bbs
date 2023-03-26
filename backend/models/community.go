package models

import "gorm.io/gorm"

type Community struct {
	gorm.Model
	CommunityId   string `json:"community_id" gorm:"community_id"`
	CommunityName string `json:"community_name" gorm:"community_name"`
	Introduction  string `json:"introduction" gorm:"introduction"`
}
