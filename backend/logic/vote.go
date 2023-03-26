package logic

import (
	"bbs/dao/redis"
	"bbs/models"
	"strconv"
)

func PostVote(userId int64, p *models.PostVoteData) error {
	return redis.VoteForPost(strconv.Itoa(int(userId)), p.PostId, float64(p.Direction))
}
