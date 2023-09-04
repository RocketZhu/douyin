package video

import (
	"douyin/dao"
	"errors"
)

func PostFavorState(userId, videoId, actionType int64) error {
	p := &PostFavorStateFlow{
		userId:     userId,
		videoId:    videoId,
		actionType: actionType,
	}
	if !dao.NewUserInfoDAO().IsUserExistById(p.userId) {
		return errors.New("用户不存在")
	}
	if p.actionType != 1 && p.actionType != 2 {
		return errors.New("未定义的行为")
	}
	switch p.actionType {
	case 1:
		return p.PlusOperation()
	case 2:
		return p.MinusOperation()
	default:
		return errors.New("未定义的操作")
	}
}

type PostFavorStateFlow struct {
	userId     int64
	videoId    int64
	actionType int64
}

// PlusOperation 点赞操作
func (p *PostFavorStateFlow) PlusOperation() error {
	//视频点赞数目+1
	err := dao.NewVideoDAO().PlusOneFavorByUserIdAndVideoId(p.userId, p.videoId)
	if err != nil {
		return errors.New("不要重复点赞")
	}
	//对应的用户是否点赞的映射状态更新
	dao.NewIndexMap().UpdateFavorState(p.userId, p.videoId, true)
	return nil
}

// MinusOperation 取消点赞
func (p *PostFavorStateFlow) MinusOperation() error {
	//视频点赞数目-1
	err := dao.NewVideoDAO().MinusOneFavorByUserIdAndVideoId(p.userId, p.videoId)
	if err != nil {
		return errors.New("点赞数目已经为0")
	}
	//对应的用户是否点赞的映射状态更新
	dao.NewIndexMap().UpdateFavorState(p.userId, p.videoId, false)
	return nil
}
