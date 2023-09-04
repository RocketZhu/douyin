package middleware

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

// 计算字符串的SHA1哈希值
func CalculateSHA1(s string) string {
	hasher := sha1.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

// 将计算后的哈希值其添加到请求上下文中
func SHAMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取参数 "password"，如果不存在，则从POST表单获取
		password := c.DefaultQuery("password", c.PostForm("password"))

		// 计算密码的 SHA1 哈希值，并添加到请求上下文中
		if password != "" {
			sha1Password := CalculateSHA1(password)
			c.Set("password", sha1Password)
		}
		c.Next()
	}
}