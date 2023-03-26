package controller

import (
	"bbs/logic"
	"bbs/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	var p models.SignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	if err := logic.SignUp(&p); err != nil {
		zap.L().Error("用户注册失败", zap.Error(err))
		ResponseError(c, CodeUserExist)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

func LoginHandler(c *gin.Context) {
	var p models.Login
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	if token, err := logic.Login(&p); err != nil {
		zap.L().Error("用户登录失败", zap.Error(err))
		ResponseError(c, CodeUserNotExist)
		return
	} else {
		ResponseSuccess(c, token)
	}
}
