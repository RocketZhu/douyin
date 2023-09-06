package dao

import (
	"douyin/models"
	"errors"
	"sync"

	"gorm.io/gorm"
)

type UserLoginDAO struct {
}

var (
	userLoginDao  *UserLoginDAO
	userLoginOnce sync.Once
)

// NewUserLoginDao 创建一个新的 UserLoginDAO 实例
func NewUserLoginDao() *UserLoginDAO {
	userLoginOnce.Do(func() {
		userLoginDao = &UserLoginDAO{}
	})
	return userLoginDao
}

// QueryUserLogin 根据用户名和密码查询用户登录信息
func (u *UserLoginDAO) QueryUserLogin(username, password string, login *models.UserLogin, DB *gorm.DB) error {
	if login == nil {
		return errors.New("结构体指针为空")
	}
	DB.Where("username=? and password=?", username, password).First(login)
	if login.Id == 0 {
		return errors.New("用户不存在，账号或密码出错")
	}
	return nil
}

// IsUserExistByUsername 检查用户名是否存在
func (u *UserLoginDAO) IsUserExistByUsername(username string, DB *gorm.DB) bool {
	var userLogin models.UserLogin
	DB.Where("username=?", username).First(&userLogin)
	return userLogin.Id != 0
}
