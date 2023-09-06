package VideoHandler

import (
	"douyin/config"
	"douyin/service/video"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

var (
	videoFormat = map[string]struct{}{
		".mp4":  {},
		".avi":  {},
		".wmv":  {},
		".flv":  {},
		".mpeg": {},
		".mov":  {},
	}
	//pictureFormat = map[string]struct{}{
	//	".jpg": {},
	//	".bmp": {},
	//	".png": {},
	//	".svg": {},
	//}
)

// PublishVideoHandler 发布视频，并截取一帧画面作为封面
func PublishVideoHandler(c *gin.Context) {
	//准备参数
	rawId, _ := c.Get("user_id")

	userId, ok := rawId.(int64)
	if !ok {
		PublishVideoError(c, "解析UserId出错")
		return
	}

	title := c.PostForm("title")

	form, err := c.MultipartForm()
	if err != nil {
		PublishVideoError(c, err.Error())
		return
	}

	//多文件上传
	files := form.File["data"]
	for _, file := range files {
		suffix := filepath.Ext(file.Filename)  //得到后缀
		if _, ok := videoFormat[suffix]; !ok { //判断是否为视频格式
			PublishVideoError(c, "不支持的视频格式")
			continue
		}
		name := util.NewFileName(userId) //根据userId得到唯一的文件名
		filename := name + suffix
		savePath := filepath.Join(config.ServerConfig.StaticSourcePath, filename)
		err = c.SaveUploadedFile(file, savePath)
		if err != nil {
			PublishVideoError(c, err.Error())
			continue
		}
		//截取一帧画面作为封面
		err = util.SaveImageFromVideo(name, false)
		if err != nil {
			PublishVideoError(c, err.Error())
			continue
		}
		//数据库持久化
		err := video.PostVideo(userId, filename, name+util.GetDefaultImageSuffix(), title)
		if err != nil {
			PublishVideoError(c, err.Error())
			continue
		}
		PublishVideoOk(c, file.Filename+"上传成功")
	}
}

func PublishVideoError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"StatusCode": 1,
		"StatusMsg":  msg,
	})
}

func PublishVideoOk(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"StatusCode": 0,
		"StatusMsg":  msg,
	})
}
