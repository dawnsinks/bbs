package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const ContextUserIdKey = "UserID"

var UserNotLogin = errors.New("用户未登录")

func GetCurrentUser(c *gin.Context) (UserID int64, err error) {
	uid, ok := c.Get(ContextUserIdKey)
	if !ok {
		return 0, UserNotLogin
	}
	UserID, ok = uid.(int64)
	if !ok {
		return 0, UserNotLogin
	}
	return
}

func getPageInfo(c *gin.Context) (page, size int64) {
	pageStr := c.Query("offset")
	SizeStr := c.Query("limit")
	var err error
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(SizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
