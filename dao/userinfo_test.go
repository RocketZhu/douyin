package dao

import (
	"douyin/models"
	"testing"
)

func TestQueryUserInfoById(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	userId := int64(1)
	userinfo := &models.UserInfo{}

	// 调用被测试的函数
	err := userInfoDAO.QueryUserInfoById(userId, userinfo)
	if err != nil {
		t.Errorf("Error querying user info by ID: %v", err)
	}

	// 进行断言或其他验证
	// 比如验证查询结果是否符合预期
}

func TestAddUserInfo(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	userinfo := &models.UserInfo{
		Name:          "Test User",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}

	// 调用被测试的函数
	err := userInfoDAO.AddUserInfo(userinfo)
	if err != nil {
		t.Errorf("Error adding user info: %v", err)
	}
}

func TestIsUserExistById(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	userId := int64(1)

	// 调用被测试的函数
	exists := userInfoDAO.IsUserExistById(userId)

	// 进行断言或其他验证
	expectedExists := true  // 预期用户存在
	if exists != expectedExists {
		t.Errorf("Expected user existence: %v, but got: %v", expectedExists, exists)
	}
}

func TestGetFollowListByUserId(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	userId := int64(1)
	var userList []*models.UserInfo

	// 调用被测试的函数
	err := userInfoDAO.GetFollowListByUserId(userId, &userList)
	if err != nil {
		t.Errorf("Error getting follow list by user ID: %v", err)
	}
}

// 其他测试函数的类似代码...

