package user_login

import (
	"douyin/dao"
	"douyin/middleware"
	"douyin/models"
	"errors"
	"fmt"
)

// PostUserLogin 注册用户并得到token和id
func PostUserLogin(username, password string) (*LoginResponse, error) {
	return NewPostUserLoginFlow(username, password).Do()
}

func NewPostUserLoginFlow(username, password string) *PostUserLoginFlow {
	return &PostUserLoginFlow{username: username, password: password}
}

type PostUserLoginFlow struct {
	username string
	password string

	data   *LoginResponse
	userid int64
	token  string
}

func (q *PostUserLoginFlow) Do() (*LoginResponse, error) {
	//对参数进行合法性验证
	fmt.Println("111")
	if err := q.checkNum(); err != nil {
		return nil, err
	}

	fmt.Println("222")
	//更新数据到数据库
	if err := q.updateData(); err != nil {
		return nil, err
	}

	fmt.Println("333")
	//打包response
	if err := q.packResponse(); err != nil {
		return nil, err
	}
	return q.data, nil
}

func (q *PostUserLoginFlow) checkNum() error {
	if q.username == "" {
		return errors.New("用户名为空")
	}
	if len(q.username) > MaxUsernameLength {
		return errors.New("用户名长度超出限制")
	}
	if q.password == "" {
		return errors.New("密码为空")
	}
	return nil
}

func (q *PostUserLoginFlow) updateData() error {

	//准备好userInfo,默认name为username
	userLogin := models.UserLogin{Username: q.username, Password: q.password}
	userinfo := models.UserInfo{User: &userLogin, Name: q.username}

	//判断用户名是否已经存在
	userLoginDAO := dao.NewUserLoginDao()
	println("s001")
	if userLoginDAO.IsUserExistByUsername(q.username) {
		println("s002")
		return errors.New("用户名已存在")
	}
	println("s003")
	//更新操作，由于userLogin属于userInfo，故更新userInfo即可，且由于传入的是指针，所以插入的数据内容也是清楚的
	userInfoDAO := dao.NewUserInfoDAO()
	println("s004")
	err := userInfoDAO.AddUserInfo(&userinfo)
	if err != nil {
		println("s005")
		return err
	}

	//颁发token
	token, err := middleware.ReleaseToken(userLogin)
	if err != nil {
		return err
	}
	q.token = token
	q.userid = userinfo.Id

	println("s005")
	return nil
}

func (q *PostUserLoginFlow) packResponse() error {
	q.data = &LoginResponse{
		UserId: q.userid,
		Token:  q.token,
	}
	return nil
}
