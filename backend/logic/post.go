package logic

import (
	"bbs/dao/mysql"
	"bbs/dao/redis"
	"bbs/models"
	snowflake "bbs/pkg/sf"
)

func CreatePost(post *models.Post) error {
	post.ID = uint(snowflake.GenId())

	err := mysql.CreatePost(post)
	if err != nil {
		return err
	}
	return nil
}

func GetPostById(pid int64) (data *models.Post, err error) {
	return mysql.GetPost(pid)
}

func GetPosts(page, size int64) (posts *[]models.Post, err error) {
	posts, err = mysql.GetPosts(page, size)
	return
}

func GetPosts2(p *models.ParamPostList) (data []*models.Post, err error) {
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return nil, err
	}

	data = make([]*models.Post, 0, len(ids))

	err = mysql.GetPostsByIds(data, ids)
	if err != nil {
		return nil, err
	}

	return
}
