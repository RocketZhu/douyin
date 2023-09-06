package user_info_handler

import (
	"douyin/dao"
	"douyin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfoResponse struct {
	StatusCode int32            `json:"status_code"`          // 状态码
	StatusMsg  string           `json:"status_msg,omitempty"` // 状态消息
	User       *models.UserInfo `json:"user"`                 // 用户信息
}

func UserInfoHandler(c *gin.Context) {
	rawId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusOK, UserInfoResponse{
			StatusCode: 1,
			StatusMsg:  "用户ID解析错误",
		})
		return
	}

	UserId, ok := rawId.(int64) // 将用户ID转换为int64类型
	if !ok {
		c.JSON(http.StatusOK, UserInfoResponse{
			StatusCode: 1,
			StatusMsg:  "用户ID解析失败",
		})
		return
	}
	UserInfoDAO := dao.NewUserInfoDAO() // 创建用户信息数据访问对象

	var userInfo models.UserInfo                                             // 创建用户信息对象
	if err := UserInfoDAO.QueryUserInfoById(UserId, &userInfo); err != nil { // 根据用户ID查询用户信息
		c.JSON(http.StatusOK, UserInfoResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(), // 查询用户信息失败的错误消息
		})
		return
	}

	c.JSON(http.StatusOK, UserInfoResponse{
		StatusCode: 0,
		User:       &userInfo, // 返回查询到的用户信息
	})
}
