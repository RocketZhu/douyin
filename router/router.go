package router

import (
	"douyin/config"
	"douyin/dao"
	"douyin/handlers/chat_handler"
	"douyin/handlers/comment_handler"
	loginHandler "douyin/handlers/login_handler"
	userInfoHandler "douyin/handlers/user_info_handler"
	videoHandler "douyin/handlers/video_handler"
	"douyin/middleware"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	dao.InitDB()
	r := gin.Default()
	r.Static("static", config.ServerConfig.StaticSourcePath)
	baseGroup := r.Group("/douyin")

	// 基础接口
	baseGroup.GET("/feed/", videoHandler.FeedVideoListHandler)
	baseGroup.GET("/user/", middleware.JWTMiddleWare(), userInfoHandler.UserInfoHandler)
	baseGroup.POST("/user/login/", middleware.SHAMiddleware(), loginHandler.UserLoginHandler)
	baseGroup.POST("/user/register/", middleware.SHAMiddleware(), loginHandler.UserRegisterHandler)
	baseGroup.POST("/publish/action/", middleware.JWTMiddleWare(), videoHandler.PublishVideoHandler)
	baseGroup.GET("/publish/list/", middleware.GetIdWithoutAuth(), videoHandler.QueryVideoListHandler)

	GroupOne := baseGroup.Group("/favorite", middleware.JWTMiddleWare())
	GroupOne.POST("/favorite/action/", videoHandler.PostFavorHandler)
	GroupOne.GET("/favorite/list/", videoHandler.QueryFavorVideoListHandler)

	GroupTwo := baseGroup.Group("/comment", middleware.JWTMiddleWare())
	GroupTwo.POST("/comment/action/", comment_handler.PostCommentHandler)
	GroupTwo.GET("/comment/list/", comment_handler.QueryCommentListHandler)

	GroupThree := baseGroup.Group("/relation", middleware.JWTMiddleWare())
	GroupThree.POST("/relation/action/", userInfoHandler.PostFollowActionHandler)
	GroupThree.GET("/relation/follow/list/", userInfoHandler.QueryFollowListHandler)
	GroupThree.GET("/relation/follower/list/", userInfoHandler.QueryFollowerHandler)
	GroupThree.GET("/relation/friend/list/")

	GroupFour := baseGroup.Group("/message", middleware.JWTMiddleWare())
	GroupFour.GET("/message/chat/", chat_handler.ChatListHandler)
	GroupFour.POST("/message/action/", chat_handler.PostMessage)
	return r
}
