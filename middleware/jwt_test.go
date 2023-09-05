package middleware

import (
	"douyin/dao"
	"douyin/models"
	"fmt"
	"testing"
)

func TestReleaseToken(t *testing.T) {
	testdb := dao.InitTestDB()
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
	userInfoDAO := dao.NewUserInfoDAO()
	err := userInfoDAO.AddUserInfo(userinfo, tx)
	if err != nil {
		t.Errorf("Add userinfo failed")
	}
	token, err := ReleaseToken(*userlogin)
	if err != nil {
		t.Errorf("Release token failed")
	}
	fmt.Printf("token: %v\n", token)
}
