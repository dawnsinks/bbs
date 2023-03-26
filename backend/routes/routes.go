package routes

import (
	"bbs/controller"
	"bbs/logger"
	middlewares "bbs/middlewares/jwt"
	"bbs/middlewares/ratelimit"
	"bbs/settings"
	"github.com/gin-gonic/gin"
	"time"
)

func Init() *gin.Engine {
	if settings.Conf.AppConfig.Mode == "dev" {
		gin.SetMode("debug")
	} else {
		gin.SetMode("release")
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.Use(ratelimit.LimitMiddleware(2*time.Second, 10000))

	v1 := r.Group("/api/v1")
	{
		v1.POST("/signup", controller.SignUpHandler)
		v1.POST("/login", controller.LoginHandler)
	}
	g := v1.Use(middlewares.JWTAuthMiddleware())
	{
		g.GET("/community", controller.CommunityHandler)
		g.GET("/community/:id", controller.CommunityDetailHandler)
		g.POST("/post", controller.CreatePostHandler)
		g.POST("/posts", controller.GetPostListHandler)
		g.POST("/posts2", controller.GetPostListHandler2)
		g.POST("/vote", controller.PostVoteHandler)
	}

	return r
}
