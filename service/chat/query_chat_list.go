package chat

import (
	"douyin/dao"
	"douyin/models"
	"errors"
	"fmt"
)

type List struct {
	Messages []*models.Message `json:"message_list"`
}

func QueryChatList(userId, ToUserId int64, preMessageTime string) (*List, error) {
	if !dao.NewUserInfoDAO().IsUserExistById(int64(userId)) {
		return nil, fmt.Errorf("用户%d不存在", userId)
	}

	if !dao.NewUserInfoDAO().IsUserExistById(int64(ToUserId)) {
		return nil, fmt.Errorf("对方用户%d不存在", ToUserId)
	}
	//向dao层请求消息列表
	var ChatList []*models.Message
	err := dao.NewMessageDAO().QueryChatListByUsers(userId, ToUserId, preMessageTime, &ChatList)
	if err != nil {
		return nil, errors.New("暂无聊天记录")
	}

	return &List{Messages: ChatList}, nil
}
