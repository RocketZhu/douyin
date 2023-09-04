package video

import (
	"douyin/dao"
	"douyin/models"
	"errors"
)

type List struct {
	Videos []*models.Video `json:"video_list,omitempty"`
}

// QueryVideoListByUserIdFlow 用于封装关于查询点赞视频列表的流程和操作。
type QueryVideoListByUserIdFlow struct {
	userId int64
	videos []*models.Video //视频列表

	videoList *List // 响应的视频列表数据
}

// QueryVideoListByUserId 根据用户ID查询视频列表
func QueryVideoListByUserId(userId int64) (*List, error) {
	flow := &QueryVideoListByUserIdFlow{userId: userId}
	if err := flow.checkNum(); err != nil {
		return nil, err
	}
	if err := flow.packData(); err != nil {
		return nil, err
	}
	return flow.videoList, nil

}

// checkNum 检查用户是否存在
func (flow *QueryVideoListByUserIdFlow) checkNum() error {
	//检查userId是否存在
	if !dao.NewUserInfoDAO().IsUserExistById(flow.userId) {
		return errors.New("用户不存在")
	}

	return nil
}

// packData 用于构建响应数据，包括查询视频列表、查询用户信息、填充视频列表信息。
// 注意：Video由于在数据库中没有存储作者信息，所以需要手动使用userid进行填充
func (flow *QueryVideoListByUserIdFlow) packData() error {
	err := dao.NewVideoDAO().QueryVideoListByUserId(flow.userId, &flow.videos)
	if err != nil {
		return err
	}
	//作者信息查询
	var userInfo models.UserInfo
	err = dao.NewUserInfoDAO().QueryUserInfoById(flow.userId, &userInfo)
	p := dao.NewIndexMap()
	if err != nil {
		return err
	}
	//填充信息(Author和IsFavorite字段
	for i := range flow.videos {
		flow.videos[i].Author = userInfo
		flow.videos[i].IsFavorite = p.GetFavorState(flow.userId, flow.videos[i].Id)
	}

	flow.videoList = &List{Videos: flow.videos}

	return nil
}
