package VideoHandler

import (
	"douyin/middleware"
	"douyin/service/video"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	*video.FeedVideoList
}

func FeedVideoListHandler(c *gin.Context) {
	p := NewFeedVideoList(c)
	token, ok := c.GetQuery("token")
	//无登录状态
	if !ok {
		err := p.DoNoToken()
		if err != nil {
			p.FeedVideoListError(err.Error())
		}
		return
	}

	//有登录状态
	err := p.DoHasToken(token)
	if err != nil {
		p.FeedVideoListError(err.Error())
	}
}

type FeedVideoList struct {
	*gin.Context
}

func NewFeedVideoList(c *gin.Context) *FeedVideoList {
	return &FeedVideoList{Context: c}
}

// DoNoToken 未登录的视频流推送处理
func (p *FeedVideoList) DoNoToken() error {
	rawTimestamp := p.Query("latest_time")
	var latestTime time.Time
	intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
	if err == nil {
		latestTime = time.Unix(0, intTime*1e6) //注意：前端传来的时间戳是以ms为单位的
	}
	videoList, err := video.QueryFeedVideoList(0, latestTime)
	if err != nil {
		return err
	}
	p.FeedVideoListOk(videoList)
	return nil
}

// DoHasToken 如果是登录状态，则生成UserId字段
func (p *FeedVideoList) DoHasToken(token string) error {
	//解析成功
	if claim, err := middleware.ParseToken(token); err == nil {
		//token超时
		if time.Now().Unix() > claim.ExpiresAt {
			return errors.New("token超时")
		}
		rawTimestamp := p.Query("latest_time")
		var latestTime time.Time
		intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
		if err != nil {
			latestTime = time.Unix(0, intTime*1e6) //注意：前端传来的时间戳是以ms为单位的
		}
		//调用service层接口
		videoList, err := video.QueryFeedVideoList(claim.Id, latestTime)
		if err != nil {
			return err
		}
		p.FeedVideoListOk(videoList)
		return nil
	}
	//解析失败
	return errors.New("token不正确")
}

func (p *FeedVideoList) FeedVideoListError(msg string) {
	p.JSON(http.StatusOK, FeedResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func (p *FeedVideoList) FeedVideoListOk(videoList *video.FeedVideoList) {
	p.JSON(http.StatusOK, FeedResponse{
		StatusCode:    0,
		FeedVideoList: videoList,
	},
	)
}
