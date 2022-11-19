package middleware

import (
	"androidmarket/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type MyClaim struct {
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}

// Mykey 密钥
var Mykey = []byte("hellothisishellozmc")

// Uid uid
var Uid string

// CheckToken jwt中间件
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从header中获取x-token
		tokenstring := c.Request.Header.Get("x-token")
		if tokenstring == "" {
			response.FailLogin(c, "您未登录！")
			// 中断
			c.Abort()
			return
		}
		// 解密
		token, _ := jwt.ParseWithClaims(tokenstring, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
			return Mykey, nil
		})
		// ？
		// 非法访问
		//if token == nil {
		//	response.FailWithDetail(c, nil, "非法访问！请登录！")
		//	// 中断
		//	c.Abort()
		//	return
		//}
		if claim, ok := token.Claims.(*MyClaim); ok && token.Valid {
			Uid = claim.Uid
		} else {
			response.FailWithDetail(c, nil, "令牌已失效，请重新登录！")
			// 中断
			c.Abort()
			return
		}
		c.Next()
	}
}
