package VideoHandler

import (
	"douyin/service/video"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	*video.List
}

type VideoList struct {
	c *gin.Context
}

func QueryVideoListHandler(c *gin.Context) {
	p := &VideoList{c: c}
	rawId, _ := c.Get("user_id")
	err := p.DoQueryVideoListByUserId(rawId)
	if err != nil {
		p.QueryVideoListError(err.Error())
	}
}

// DoQueryVideoListByUserId 根据userId字段进行查询
func (p *VideoList) DoQueryVideoListByUserId(rawId interface{}) error {
	userId, ok := rawId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	//与结构体
	videoList, err := video.QueryVideoListByUserId(userId)
	if err != nil {
		return err
	}

	p.QueryVideoListOk(videoList)
	return nil
}

func (p *VideoList) QueryVideoListError(msg string) {
	p.c.JSON(http.StatusOK, ListResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func (p *VideoList) QueryVideoListOk(videoList *video.List) {
	p.c.JSON(http.StatusOK, ListResponse{
		StatusCode: 0,
		List:       videoList,
	})
}
