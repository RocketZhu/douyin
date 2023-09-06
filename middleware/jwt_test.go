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
	//token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiZXhwIjoxNjk0NTg0NzY1LCJpYXQiOjE2OTM5Nzk5NjUsImlzcyI6ImRvdXlpbiJ9.IAsK9W8Eh7m_4DYcL1rbYsivAW4f_dfAQoci1qImJiw
}

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiZXhwIjoxNjk0NTg0NzY1LCJpYXQiOjE2OTM5Nzk5NjUsImlzcyI6ImRvdXlpbiJ9.IAsK9W8Eh7m_4DYcL1rbYsivAW4f_dfAQoci1qImJiw"
	claims, err := ParseToken(token)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Printf("claims: %v\n", claims.Id)

}
