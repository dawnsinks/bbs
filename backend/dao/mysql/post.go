package mysql

import (
	"bbs/models"
)

func CreatePost(p *models.Post) error {
	return db.Create(&p).Error

}

func GetPost(pid int64) (p *models.Post, err error) {
	p = new(models.Post)
	err = db.Find(&p).Error
	return p, err
}

func GetPosts(page, size int64) (posts *[]models.Post, err error) {
	err = db.Offset(int((page - 1) * size)).Limit(int(size)).Find(&posts).Error
	return posts, err
}

func GetPostsByIds(data []*models.Post, ids []string) error {
	return db.Find(&data, ids).Error
}
