package chat_handler

import (
	"douyin/service/chat"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type postChatFlow struct {
	*gin.Context
	userId     int64
	toUserId   int64
	actionType int32
	content    string
}
type chatResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func PostMessage(c *gin.Context) {
	m := &postChatFlow{Context: c}
	if err := m.parseNum(); err != nil {
		m.chatError(err.Error())
		return
	}
	err := chat.PostChat(m.userId, m.toUserId, m.actionType, m.content)
	if err != nil {
		m.chatError(err.Error())
	}
	m.chatOk()
}

// 获取并解析数据
func (m *postChatFlow) parseNum() error {
	rawUserId, _ := m.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		m.chatError("userID解析错误")
	}
	m.userId = userId
	rawToUserId := m.Query("to_user_id")
	toUserId, err := strconv.ParseInt(rawToUserId, 10, 64)
	if err != nil {
		return err
	}
	m.toUserId = toUserId
	rawActionType := m.Query("action_type")
	actionTypeInt64, err := strconv.ParseInt(rawActionType, 10, 32)
	actionType := int32(actionTypeInt64)
	if err != nil {
		return err
	}
	m.actionType = actionType
	m.content = m.Query("content")
	return nil
}

func (m *postChatFlow) chatOk() {
	m.JSON(http.StatusOK, chatResponse{
		StatusCode: 0,
	})
}
func (m *postChatFlow) chatError(msg string) {
	m.JSON(http.StatusOK, chatResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}
