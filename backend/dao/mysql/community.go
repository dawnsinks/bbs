package mysql

import "bbs/models"

func GetCommunityList(list *[]models.Community) (err error) {
	return db.Find(&list).Error
}

func GetCommunity(id int64) (data *models.Community, err error) {
	data = new(models.Community)
	err = db.Find(&data, "CommunityId = ?", id).Error
	return
}
