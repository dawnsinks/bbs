package controller

import (
	"bbs/logic"
	"bbs/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)
	if err := c.ShouldBindJSON(&p); err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	userId, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}
	p.AuthorID = userId
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("Create Post failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)

}

func GetPostDetailHandler(c *gin.Context) {
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	data, err := logic.GetPostById(pid)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func GetPostListHandler(c *gin.Context) {
	page, size := getPageInfo(c)

	posts, err := logic.GetPosts(page, size)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, posts)
}

func GetPostListHandler2(c *gin.Context) {
	p := models.ParamPostList{
		CommunityID: 0,
		Page:        1,
		Size:        10,
		Order:       models.OrderTime,
	}
	if err := c.ShouldBindQuery(&p); err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	data, err := logic.GetPosts2(&p)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
