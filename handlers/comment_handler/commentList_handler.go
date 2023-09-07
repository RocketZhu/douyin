package comment_handler

import (
	"douyin/service/comment"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	*comment.List
}

type ProxyCommentListHandler struct {
	*gin.Context

	videoId int64
	userId  int64
}

func QueryCommentListHandler(c *gin.Context) {
	p := &ProxyCommentListHandler{Context: c}
	//解析参数
	if err := p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}

	//正式调用
	commentList, err := comment.QueryCommentList(p.userId, p.videoId)
	if err != nil {
		p.SendError(err.Error())
		return
	}

	//成功返回
	p.SendOk(commentList)
}

func (p *ProxyCommentListHandler) parseNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId

	rawVideoId := p.Query("video_id")
	videoId, err := strconv.ParseInt(rawVideoId, 10, 64)
	if err != nil {
		return err
	}
	p.videoId = videoId

	return nil
}

func (p *ProxyCommentListHandler) SendError(msg string) {
	p.JSON(http.StatusOK, ListResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func (p *ProxyCommentListHandler) SendOk(commentList *comment.List) {
	p.JSON(http.StatusOK, ListResponse{
		StatusCode: 0,
		List:       commentList,
	})
}
