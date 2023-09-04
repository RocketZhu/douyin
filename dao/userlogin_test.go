package dao

import (
	"douyin/models"
	"testing"
)

func TestQueryUserLogin(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	username := "testuser"
	password := "testpassword"
	login := &models.UserLogin{}

	// 调用被测试的函数
	err := userLoginDao.QueryUserLogin(username, password, login)
	if err != nil {
		t.Errorf("Error querying user login: %v", err)
	}
}

func TestIsUserExistByUsername(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	username := "testuser"

	// 调用被测试的函数
	exists := userLoginDao.IsUserExistByUsername(username)

	expectedExists := true  // 预期用户存在
	if exists != expectedExists {
		t.Errorf("Expected user existence: %v, but got: %v", expectedExists, exists)
	}
}


