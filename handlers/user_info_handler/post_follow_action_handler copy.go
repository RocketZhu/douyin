package user_info_handler

import (
	"douyin/models"
	"douyin/service/user_info"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserResponse struct {
	StatusCode int32            `json:"status_code"`
	StatusMsg  string           `json:"status_msg,omitempty"`
	User       *models.UserInfo `json:"user"`
}

type ProxyPostFollowAction struct {
	*gin.Context

	userId     int64
	followId   int64
	actionType int
}

func PostFollowActionHandler(c *gin.Context) {
	p := &ProxyPostFollowAction{Context: c}
	var err error
	if err = p.prepareNum(); err != nil {
		p.SendError(err.Error())
		return
	}
	if err = p.startAction(); err != nil {
		//当错误为model层发生的，那么就是重复键值的插入了
		if errors.Is(err, user_info.ErrIvdAct) || errors.Is(err, user_info.ErrIvdFolUsr) {
			p.SendError(err.Error())
		} else {
			p.SendError("请勿重复关注")
		}
		return
	}
	p.SendOk("操作成功")
}

func (p *ProxyPostFollowAction) prepareNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId

	//解析需要关注的id
	followId := p.Query("to_user_id")
	parseInt, err := strconv.ParseInt(followId, 10, 64)
	if err != nil {
		return err
	}
	p.followId = parseInt

	//解析action_type
	actionType := p.Query("action_type")
	parseInt, err = strconv.ParseInt(actionType, 10, 32)
	if err != nil {
		return err
	}
	p.actionType = int(parseInt)
	return nil
}

func (p *ProxyPostFollowAction) startAction() error {
	err := user_info.PostFollowAction(p.userId, p.followId, p.actionType)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProxyPostFollowAction) SendError(msg string) {
	p.JSON(http.StatusOK, gin.H{
		"StatusCode": 1,
		"status_msg": msg,
	})
}

func (p *ProxyPostFollowAction) SendOk(msg string) {
	p.JSON(http.StatusOK, gin.H{
		"StatusCode": 0,
		"status_msg": msg,
	})
}
