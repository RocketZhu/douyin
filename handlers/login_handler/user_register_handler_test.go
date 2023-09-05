package UserLoginHandler

import (
	"douyin/middleware"
	"douyin/models"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestUserRegisterHandler(t *testing.T) {
	//创建测试数据库
	testdb := InitTestDB()
	tx := testdb.Begin()
	defer tx.Rollback()

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
}
func InitTestDB() *gorm.DB {
	mysqlHost := "127.0.0.1"
	mysqlPort := "3306"
	mysqlUser := "root"
	mysqlPassword := "Wszyc20021019"
	databaseName := "test"
	testDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlUser, mysqlPassword, mysqlHost, mysqlPort, databaseName)

	// 连接测试数据库
	testDB, err := gorm.Open(mysql.Open(testDSN), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database: " + err.Error())
	}

	// 在这里可以添加数据库初始化结构等操作
	// 如果需要创建测试表，可以在这里执行 AutoMigrate
	err = testDB.AutoMigrate(&models.UserInfo{}, &models.UserLogin{})
	if err != nil {
		panic(err)
	}
	return testDB
}
