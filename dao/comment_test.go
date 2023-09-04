package dao

import (
	"douyin/models"
	"testing"
	"time"
)

func TestAddCommentAndUpdateCount(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	comment := &models.Comment{
		UserInfoId: 1,
		VideoId:    1,
		Content:    "This is a test comment.",
		CreatedAt:  time.Now(),
		CreateDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	// 调用被测试的函数
	err := commentDao.AddCommentAndUpdateCount(comment)
	if err != nil {
		t.Errorf("Error adding comment and updating count: %v", err)
	}

	// 进行断言或其他验证
	// 比如验证添加的评论和更新的视频评论数是否与预期一致
}

func TestDeleteCommentAndUpdateCountById(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	commentId := int64(1)
	videoId := int64(1)

	// 调用被测试的函数
	err := commentDao.DeleteCommentAndUpdateCountById(commentId, videoId)
	if err != nil {
		t.Errorf("Error deleting comment and updating count: %v", err)
	}

	// 进行断言或其他验证
	// 比如验证评论是否被正确删除，视频评论数是否被正确更新
}

func TestQueryCommentById(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	commentId := int64(1)
	comment := &models.Comment{}

	// 调用被测试的函数
	err := commentDao.QueryCommentById(commentId, comment)
	if err != nil {
		t.Errorf("Error querying comment by ID: %v", err)
	}

	// 进行断言或其他验证
	// 比如验证查询结果是否符合预期
}

func TestQueryCommentListByVideoId(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	videoId := int64(1)
	var comments []*models.Comment

	// 调用被测试的函数
	err := commentDao.QueryCommentListByVideoId(videoId, &comments)
	if err != nil {
		t.Errorf("Error querying comment list by video ID: %v", err)
	}

	// 进行断言或其他验证
	// 比如验证查询结果是否符合预期
}


