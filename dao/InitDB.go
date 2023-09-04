package dao

import (
	"context"
	"douyin/models"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MakeDSN() string {
	mysqlHost := "127.0.0.1"
	mysqlPort := "3306"
	mysqlUser := "root"
	mysqlPassword := "Wszyc20021019"
	databaseName := "douyin"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlUser, mysqlPassword, mysqlHost, mysqlPort, databaseName)

	return dsn
}
func InitDB() {
	dsn := MakeDSN()
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true, //缓存预编译命令
		SkipDefaultTransaction: true, //禁用默认事务操作
	})
	if err != nil {
		log.Println(err)
	}
	err = DB.AutoMigrate(&models.UserInfo{}, &models.Video{}, &models.Comment{}, &models.UserLogin{})
	if err != nil {
		panic(err)
	}
	log.Println("connect success:")

}

var ctx = context.Background()
var rdb *redis.Client

func init() {
	redisHost := "127.0.0.1"
	redisPort := 6379
	rdb = redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf("%s:%v", redisHost, redisPort),
			DB:   0,
		})
}
