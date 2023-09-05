package dao

import (
	"douyin/models"
	"fmt"
	"testing"
)

func TestQueryUserLogin(t *testing.T) {
	testdb := InitTestDB()
	tx := testdb.Begin()
	defer tx.Rollback()


	// 准备测试数据
	username := "first"
	password := "111111"
	userlogin := &models.UserLogin{
		Username: username,
		Password: password,
	}
	userinfo := &models.UserInfo{User: userlogin, Name: username}
	err := userInfoDAO.AddUserInfo(userinfo, tx)
	if err != nil {
		t.Errorf("add userinfo failed")
	}
	// 调用被测试的函数
	login := &models.UserLogin{}
	err = userLoginDao.QueryUserLogin(username, password, login, tx)
	if err != nil {
		t.Errorf("Error querying user login: %v", err)
	}
	fmt.Printf("login: %+v\n", *login)
}

func TestIsUserExistByUsername(t *testing.T) {
	testdb := InitTestDB()
	tx := testdb.Begin()
	defer tx.Rollback()
	// 准备测试数据
	username := "first"
	password := "111111"
	userlogin := &models.UserLogin{
		Username: username,
		Password: password,
	}
	userinfo := &models.UserInfo{User: userlogin, Name: username}
	err := userInfoDAO.AddUserInfo(userinfo, tx)
	if err != nil {
		t.Errorf("add userinfo failed")
	}
	// 调用被测试的函数
	userLoginDao = NewUserLoginDao()
	exists := userLoginDao.IsUserExistByUsername(userlogin.Username, tx)

	// 查询并打印两个表的数据
	var userInfos []models.UserInfo
	var userLogins []models.UserLogin

	// 查询 user_infos 表的数据
	tx.Find(&userInfos)
	fmt.Printf("userInfos: %+v\n", userInfos)

	// 查询 user_logins 表的数据
	tx.Find(&userLogins)
	fmt.Printf("userLogins: %+v\n", userLogins)

	expectedExists := true // 预期用户存在
	if exists != expectedExists {
		t.Errorf("Expected user existence: %v, but got: %v", expectedExists, exists)
	}
}
