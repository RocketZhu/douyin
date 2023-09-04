package chat

import (
	"douyin/dao"
	"douyin/models"
	"errors"
	"time"
)

func PostChat(userId, toUserId int64, actionType int32, content string) error {
	if !dao.NewUserInfoDAO().IsUserExistById(toUserId) {
		return errors.New("该用户不存在")
	}
	if actionType != 1 {
		return errors.New("未定义操作")
	}
	if userId == toUserId {
		return errors.New("不能给自己发送消息")
	}
	var err error
	currentTime := time.Now()
	messageTime := currentTime.Format("2006-01-02 15:04:05")
	message := &models.Message{
		ToUserId:   toUserId,
		FromUserId: userId,
		Content:    content,
		CreatTime:  messageTime,
	}
	switch actionType {
	case 1:
		err = dao.NewMessageDAO().AddMessage(message)
	default:
		return errors.New("未定义操作")
	}
	return err
}
