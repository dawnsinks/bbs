package controller

import (
	"bbs/logic"
	"bbs/models"
	"github.com/gin-gonic/gin"
)

func PostVoteHandler(c *gin.Context) {
	p := new(models.PostVoteData)
	if err := c.ShouldBindJSON(&p); err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	userId, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}
	err = logic.PostVote(userId, p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
