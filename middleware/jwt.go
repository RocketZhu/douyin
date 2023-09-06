package middleware

import (
	"douyin/models"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type MyClaims struct {
	Id                 int64 `json:"Id"` // 用户ID
	jwt.StandardClaims       // JWT标准声明
}

type CommonResponse struct {
	StatusCode int32 `json:"status_code"`
	//存储返回的状态码，如404，200等信息
	StatusMsg string `json:"status_msg,omitempty"`
}

var MySecret = []byte("TiktokSecretKey") // JWT签名密钥

const TokenExpireDuration = time.Hour * 24 * 7 // Token过期时间

func ReleaseToken(user models.UserLogin) (string, error) {
	c := MyClaims{
		Id: user.Id, // 设置用户ID
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 设置Token过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "douyin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c) // 创建JWT令牌
	return token.SignedString(MySecret)                   // 使用密钥对令牌进行签名并返回
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil // 返回密钥用于验证签名
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*MyClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token") // 无效的令牌
	}

	return claims, nil // 返回解析后的声明
}

func JWTMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//无论注册或登录，都先向客户端请求token
		tokenStr := c.DefaultQuery("token", c.PostForm("token"))
		//用户不存在
		if tokenStr == "" {
			c.JSON(http.StatusOK, CommonResponse{StatusCode: 401, StatusMsg: "用户不存在"})
			c.Abort() //阻止执行
			return
		}
		//验证token
		tokenStruck, err := ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusOK, CommonResponse{
				StatusCode: 403,
				StatusMsg:  "token不正确",
			})
			c.Abort() //阻止执行
			return
		}
		//token超时
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			c.JSON(http.StatusOK, CommonResponse{
				StatusCode: 402,
				StatusMsg:  "token过期",
			})
			c.Abort() //阻止执行
			return
		}
		c.Set("user_id", tokenStruck.Id)
		c.Next()
	}
}
