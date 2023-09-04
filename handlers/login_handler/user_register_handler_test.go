package UserLoginHandler

import (
	"douyin/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUserRegisterHandler(t *testing.T) {
	// 创建一个虚拟的 HTTP 请求
	req, _ := http.NewRequest("GET", "/douyin/user/register/?username=first&password=111111", nil)
	// 创建一个响应记录器
	recorder := httptest.NewRecorder()

	// 创建一个 Gin 上下文
	c, _ := gin.CreateTestContext(recorder)
	c.Request = req

	// 获取参数 "password"，如果不存在，则从POST表单获取
	password := c.DefaultQuery("password", c.PostForm("password"))

	// 计算密码的 SHA1 哈希值，并添加到请求上下文中
	if password != "" {
		sha1Password := middleware.CalculateSHA1(password)
		c.Set("password", sha1Password)
	}

	// 调用被测试的处理函数
	UserRegisterHandler(c)

	// 检查响应是否符合预期
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	// 进一步检查响应内容是否符合预期
	// 你可以根据 UserRegisterResponse 结构来解析 JSON 响应并进行断言
	// 例如，检查 StatusCode 和 Login 字段的值是否符合预期

	// 例如：
	// var response UserRegisterResponse
	// err := json.Unmarshal(recorder.Body.Bytes(), &response)
	// if err != nil {
	//     t.Errorf("Error decoding JSON: %v", err)
	// }
	// if response.StatusCode != 0 {
	//     t.Errorf("Expected StatusCode 0, but got %d", response.StatusCode)
	// }
	// 进一步检查其他字段

}
