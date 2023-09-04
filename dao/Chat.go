package dao

import (
	"douyin/models"
	"errors"
	"gorm.io/gorm"
	"time"
)

type MessageDao struct {
}

var messageDao MessageDao

func NewMessageDAO() *MessageDao {
	return &messageDao
}

func (m *MessageDao) AddMessage(message *models.Message) error {
	if message == nil {
		return errors.New("AddMessage messsage 空指针")
	}
	//执行事务
	return DB.Transaction(func(tx *gorm.DB) error {
		//添加评论数据
		if err := tx.Create(message).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

func (m *MessageDao) QueryChatListByUsers(userId, ToUserId int64, preMessageTime string, messages *[]*models.Message) error {
	if messages == nil {
		return errors.New("messages空指针")
	}

	layout := "2006-01-02 15:04:05" // 使用与数据库时间格式相匹配的 layout
	parsedMessageTime, err := time.Parse(layout, preMessageTime)
	if err != nil {
		return err
	}
	if err := DB.Where("to_user_id=? AND from_user_id=? AND created_at < ?", ToUserId, userId, parsedMessageTime).
		Order("created_at DESC").Limit(30).Find(messages).Error; err != nil {
		return err
	}
	return nil
}
