package video

import (
	"douyin/dao"
	"douyin/models"
	"douyin/util"
)

// PostVideo 投稿视频
func PostVideo(userId int64, videoName, coverName, title string) error {
	//准备好是视频和封面的地址
	videoName = util.GetFileUrl(videoName)
	coverName = util.GetFileUrl(coverName)
	//将视频和封面地址组合添加到数据库
	video := &models.Video{
		UserInfoId: userId,
		PlayUrl:    videoName,
		CoverUrl:   coverName,
		Title:      title,
	}
	if err := dao.NewVideoDAO().AddVideo(video); err != nil {
		return err
	}
	return nil

}
