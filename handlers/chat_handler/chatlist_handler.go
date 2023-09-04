package chat_handler

import (
	"douyin/service/chat"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ChatListResponse struct {
	StatusCode  int32  `json:"status_code"`
	StatusMsg   string `json:"status_msg,omitempty"`
	MessageList *chat.List
}

func ChatListHandler(c *gin.Context) {
	rawPreMessageTime, ok := c.GetQuery("pre_msg_time")
	if !ok {
		c.JSON(http.StatusOK, ChatListResponse{
			StatusCode: 1,
			StatusMsg:  "preMessage解析错误",
		})
		return
	}
	//获取ToUserId
	rawToUserId, _ := c.Get("to_user_id")
	ToUserId, ok := rawToUserId.(int64)
	if !ok {
		c.JSON(http.StatusOK, ChatListResponse{
			StatusCode: 1,
			StatusMsg:  "to_user_id解析错误",
		})
		return
	}

	//获取from_user_id，并改称为user_id
	rawUserId, _ := c.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		c.JSON(http.StatusOK, ChatListResponse{
			StatusCode: 1,
			StatusMsg:  "user_id解析错误",
		})
		return
	}

	ChatList, err := chat.QueryChatList(userId, ToUserId, rawPreMessageTime)
	if err != nil {
		c.JSON(http.StatusOK, ChatListResponse{
			StatusCode:  0,
			MessageList: ChatList,
		})
	} else {
		c.JSON(http.StatusOK, ChatListResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	}
}
