package video

import (
	"douyin/dao"
	"douyin/models"
	"errors"
)

type FavorList struct {
	Videos []*models.Video `json:"video_list"`
}

// QueryFavorVideoListFlow 用于封装关于查询点赞视频列表的流程和操作。
type QueryFavorVideoListFlow struct {
	userId int64

	videos []*models.Video

	videoList *FavorList
}

func QueryFavorVideoList(userId int64) (*FavorList, error) {
	flow := &QueryFavorVideoListFlow{userId: userId}
	if err := flow.checkNum(); err != nil {
		return nil, err
	}
	if err := flow.prepareData(); err != nil {
		return nil, err
	}
	if err := flow.packData(); err != nil {
		return nil, err
	}
	return flow.videoList, nil
}

// flow *QueryFavorVideoListFlow 结构体指针，用于存储查询点赞视频列表的流程中所需要的数据和状态
func (flow *QueryFavorVideoListFlow) checkNum() error {
	if !dao.NewUserInfoDAO().IsUserExistById(flow.userId) {
		return errors.New("用户状态异常")
	}
	return nil
}

func (flow *QueryFavorVideoListFlow) prepareData() error {
	err := dao.NewVideoDAO().QueryFavorVideoListByUserId(flow.userId, &flow.videos)
	if err != nil {
		return err
	}
	//填充信息(Author和IsFavorite字段)，因为是点赞列表，所以所有的都是点赞状态
	for i := range flow.videos {
		//作者信息查询
		var userInfo models.UserInfo
		err = dao.NewUserInfoDAO().QueryUserInfoById(flow.videos[i].UserInfoId, &userInfo)
		if err == nil { //如果查询未出错，则更新作者信息
			flow.videos[i].Author = userInfo
		}
		flow.videos[i].IsFavorite = true
	}
	return nil
}

func (flow *QueryFavorVideoListFlow) packData() error {
	flow.videoList = &FavorList{Videos: flow.videos}
	return nil
}
