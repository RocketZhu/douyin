package user_info

import (
	"douyin/dao"
	"douyin/models"
)

type FriendsList struct {
	UserList []*models.UserInfo `json:"user_list"`
}

func QueryFriendsList(userId int64) (*FriendsList, error) {
	return NewQueryFriendsListFlow(userId).Do()
}

type QueryFriendsListFlow struct {
	userId int64

	userList []*models.UserInfo

	*FriendsList
}

func NewQueryFriendsListFlow(userId int64) *QueryFriendsListFlow {
	return &QueryFriendsListFlow{userId: userId}
}

func (q *QueryFriendsListFlow) Do() (*FriendsList, error) {
	var err error
	if err = q.checkNum(); err != nil {
		return nil, err
	}
	if err = q.prepareData(); err != nil {
		return nil, err
	}
	if err = q.packData(); err != nil {
		return nil, err
	}

	return q.FriendsList, nil
}

func (q *QueryFriendsListFlow) checkNum() error {
	if !dao.NewUserInfoDAO().IsUserExistById(q.userId) {
		return ErrUserNotExist
	}
	return nil
}

func (q *QueryFriendsListFlow) prepareData() error {
	var userList []*models.UserInfo
	err := dao.NewUserInfoDAO().GetFriendsListByUserId(q.userId, &userList)
	if err != nil {
		return err
	}
	for i, _ := range userList {
		userList[i].IsFollow = true //当前用户的关注列表，故isFollow定为true
	}
	q.userList = userList
	return nil
}

func (q *QueryFriendsListFlow) packData() error {
	q.FriendsList = &FriendsList{UserList: q.userList}

	return nil
}
