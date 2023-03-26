package logic

import (
	"bbs/dao/mysql"
	"bbs/models"
	"bbs/pkg/jwt"
	snowflake "bbs/pkg/sf"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"go.uber.org/zap"
)

var secret = "dawnsink"

func encryptPassword(p string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(p)))
}

func SignUp(p *models.SignUp) (err error) {
	if exist := mysql.CheckUserExist(p.Username); exist {
		return errors.New("用户已存在")
	}
	var u = models.User{
		UserId:   snowflake.GenId(),
		UserName: p.Username,
		Password: encryptPassword(p.Password),
	}

	if err := mysql.InsertUser(&u); err != nil {
		zap.L().Error("用户注册失败", zap.Error(err))
	}
	return
}

func Login(p *models.Login) (token *models.Token, err error) {
	if exist := mysql.CheckUserExist(p.Username); !exist {
		return nil, errors.New("用户名错误")
	}

	user := models.User{
		UserName: p.Username,
		Password: encryptPassword(p.Password),
	}

	if err := mysql.Login(&user); err != nil {
		zap.L().Error("用户登录失败", zap.Error(err))
		return nil, nil
	}
	a, r, err := jwt.GenToken(uint64(user.UserId), user.UserName)
	token = &models.Token{
		UserId:       user.UserId,
		UserName:     p.Username,
		AccessToken:  a,
		RefreshToken: r,
	}
	return token, nil
}
