package VideoHandler

import (
	"douyin/service/video"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FavorVideoListResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	*video.FavorList
}

type FavorVideoListHandler struct {
	*gin.Context
	userId int64
}

func QueryFavorVideoListHandler(c *gin.Context) {
	p := &FavorVideoListHandler{Context: c}
	//解析参数
	if err := p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}

	//正式调用
	favorVideoList, err := video.QueryFavorVideoList(p.userId)
	if err != nil {
		p.SendError(err.Error())
		return
	}

	//成功返回
	p.SendOk(favorVideoList)
}

func (p *FavorVideoListHandler) parseNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

func (p *FavorVideoListHandler) SendError(msg string) {
	p.JSON(http.StatusOK, FavorVideoListResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

func (p *FavorVideoListHandler) SendOk(favorList *video.FavorList) {
	p.JSON(http.StatusOK, FavorVideoListResponse{
		StatusCode: 0,
		FavorList:  favorList,
	})
}
