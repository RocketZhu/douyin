package user_info_handler

import (
	"douyin/service/user_info"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FriendsListResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	*user_info.FriendsList
}

type ProxyQueryFriendsList struct {
	*gin.Context

	userId int64

	*user_info.FriendsList
}

func QueryFriendsListHandler(c *gin.Context) {
	p := &ProxyQueryFriendsList{Context: c}
	var err error
	if err = p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}
	if err = p.prepareData(); err != nil {
		p.SendError(err.Error())
		return
	}
	p.SendOk("请求成功")
}

func (p *ProxyQueryFriendsList) parseNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

func (p *ProxyQueryFriendsList) prepareData() error {
	list, err := user_info.QueryFriendsList(p.userId)
	if err != nil {
		return err
	}
	p.FriendsList = list
	return nil
}

func (p *ProxyQueryFriendsList) SendError(msg string) {
	p.JSON(http.StatusOK, FriendsListResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func (p *ProxyQueryFriendsList) SendOk(msg string) {
	p.JSON(http.StatusOK, FriendsListResponse{
		StatusCode:  0,
		StatusMsg:   msg,
		FriendsList: p.FriendsList,
	})
}
