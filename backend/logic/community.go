package logic

import (
	"bbs/dao/mysql"
	"bbs/models"
)

func GetCommunityList() (list []models.Community, err error) {
	err = mysql.GetCommunityList(&list)
	return list, err
}

func GetCommunity(id int64) (data *models.Community, err error) {
	return mysql.GetCommunity(id)
}
