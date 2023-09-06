package UserLoginHandler // Package handlers 包名为handlers，用于存放处理函数

import (
	"douyin/service/user_login"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRegisterResponse struct { // 定义用户注册响应的数据结构
	StatusCode int32                     `json:"status_code"`          // 状态码
	StatusMsg  string                    `json:"status_msg,omitempty"` // 状态消息
	Login      *user_login.LoginResponse // 登录响应
}

func UserRegisterHandler(c *gin.Context) { // 用户注册处理函数，接收一个gin.Context类型的参数
	username := c.Query("username")            // 获取请求中的username参数
	passwordValue, exists := c.Get("password") // 获取请求中的password参数
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password parameter is missing",
		})
		return
	}
	password, ok := passwordValue.(string)
	if !ok {
		fmt.Println("Password parameter is not a string")
		return
	}
	

	RegisterResponse, err := user_login.PostUserLogin(username, password) // 调用登录函数进行登录操作

	if err != nil { // 如果发生错误
		c.JSON(http.StatusOK, UserRegisterResponse{ // 返回一个包含错误信息的JSON响应
			StatusCode: 1,           // 状态码为1
			StatusMsg:  err.Error(), // 状态消息为错误信息
		})
		return // 结束函数执行
	}

	c.JSON(http.StatusOK, UserRegisterResponse{ // 返回一个包含登录响应的JSON响应
		StatusCode: 0,                // 状态码为0
		Login:      RegisterResponse, // 登录响应为RegisterResponse
	})
}
