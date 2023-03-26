package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	AuthorID    int64  `json:"author_id" gorm:"author_id"`
	CommunityID int64  `json:"community_id" gorm:"community_id"`
	Status      int32  `json:"status" gorm:"status"`
	Title       string `json:"title" gorm:"title"`
	Content     string `json:"content" gorm:"content"`
}
