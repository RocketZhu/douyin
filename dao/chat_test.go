package dao

import (
	"douyin/models"
	"testing"
	"time"
)

func TestAddMessage(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	message := &models.Message{
		FromUserId: 1,
		ToUserId:   2,
		Content:    "Hello, this is a test message.",
		CreatTime:  time.Now().Format("2006-01-02 15:04:05"),
	}

	// 调用被测试的函数
	err := messageDao.AddMessage(message)
	if err != nil {
		t.Errorf("Error adding message: %v", err)
	}
}

func TestQueryChatListByUsers(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	userId := int64(1)
	toUserId := int64(2)
	preMessageTime := "2023-01-01 00:00:00"

	var messages []*models.Message

	// 调用被测试的函数
	err := messageDao.QueryChatListByUsers(userId, toUserId, preMessageTime, &messages)
	if err != nil {
		t.Errorf("Error querying chat list: %v", err)
	}
}

func setupDB() {
	// 初始化数据库连接
	InitDB()
}
