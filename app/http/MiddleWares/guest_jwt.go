package MiddleWares

import (
	"goWeb/pkg/jwt"
	"goWeb/pkg/response"

	"github.com/gin-gonic/gin"
)

// GuestJWT 强制使用游客身份访问
func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) > 0 { //Authorization中有东西

			_, err := jwt.NewJWT().ParserToken(c)
			if err == nil { // 解析 token 成功，说明登录成功了,返回提示
				response.Unauthorized(c, "请使用游客身份访问")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
