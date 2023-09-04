// dao/video_test.go

package dao

import (
	"douyin/models"
	"testing"
	"time"
)

func TestQueryVideoByVideoId(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	videoId := int64(1)
	video := &models.Video{}

	// 调用被测试的函数
	err := videoDAO.QueryVideoByVideoId(videoId, video)
	if err != nil {
		t.Errorf("Error querying video by ID: %v", err)
	}

	expectedVideo := &models.Video{
		// 设置 Video 的各个字段，以便与预期值进行比较
	}

	if video.Id != expectedVideo.Id ||
		video.UserInfoId != expectedVideo.UserInfoId ||
		video.PlayUrl != expectedVideo.PlayUrl {
		t.Errorf("Video mismatch.\nExpected: %+v\nActual: %+v", expectedVideo, video)
	}
}

func TestQueryVideoCountByUserId(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	userId := int64(1)
	count := new(int64)

	// 调用被测试的函数
	err := videoDAO.QueryVideoCountByUserId(userId, count)
	if err != nil {
		t.Errorf("Error querying video count by user ID: %v", err)
	}

	// 进行断言或其他验证
	// 比如验证查询结果是否符合预期
}

func TestQueryVideoListByUserId(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	userId := int64(1)
	videoList := []*models.Video{}

	// 调用被测试的函数
	err := videoDAO.QueryVideoListByUserId(userId, &videoList)
	if err != nil {
		t.Errorf("Error querying video list by user ID: %v", err)
	}

	// 进行断言或其他验证
	// 比如验证查询结果是否符合预期
}

func TestPlusOneFavorByUserIdAndVideoId(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	userId := int64(1)
	videoId := int64(1)

	// 调用被测试的函数
	err := videoDAO.PlusOneFavorByUserIdAndVideoId(userId, videoId)
	if err != nil {
		t.Errorf("Error adding one favor by user ID and video ID: %v", err)
	}

	// 进行断言或其他验证
	// 比如验证赞的数量是否增加，用户是否正确记录点赞的视频
}

func TestMinusOneFavorByUserIdAndVideoId(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	userId := int64(1)
	videoId := int64(1)

	// 调用被测试的函数
	err := videoDAO.MinusOneFavorByUserIdAndVideoId(userId, videoId)
	if err != nil {
		t.Errorf("Error subtracting one favor by user ID and video ID: %v", err)
	}

	// 进行断言或其他验证
	// 比如验证赞的数量是否减少，用户是否正确记录取消点赞的视频
}

func TestQueryFavorVideoListByUserId(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	userId := int64(1)
	videoList := []*models.Video{}

	// 调用被测试的函数
	err := videoDAO.QueryFavorVideoListByUserId(userId, &videoList)
	if err != nil {
		t.Errorf("Error querying favor video list by user ID: %v", err)
	}

	// 进行断言或其他验证
	// 比如验证查询结果是否符合预期
}
func TestQueryVideoListByLimitAndTime(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	limit := 10
	latestTime := time.Now()
	videoList := []*models.Video{}

	// 调用被测试的函数
	err := videoDAO.QueryVideoListByLimitAndTime(limit, latestTime, &videoList)
	if err != nil {
		t.Errorf("Error querying video list by limit and time: %v", err)
	}

}

func TestIsVideoExistById(t *testing.T) {
	// 初始化数据库连接和事务
	setupDB()

	// 准备测试数据
	videoId := int64(1)

	// 调用被测试的函数
	exists := videoDAO.IsVideoExistById(videoId)

	expectedExists := true // 预期用户存在
	if exists != expectedExists {
		t.Errorf("Expected user existence: %v, but got: %v", expectedExists, exists)
	}
}
