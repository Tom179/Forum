// Package middlewares Gin 中间件
package MiddleWares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goWeb/app/models/user"
	"goWeb/jwt"
	"goWeb/pkg/response"
)

func AuthJWT() gin.HandlerFunc { //从jwt中解析出用户信息写入c *Context中（一个公共的地方）
	return func(c *gin.Context) {

		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParserToken(c)

		// JWT 解析失败，有错误发生
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("错误，请查看 %v 相关的接口认证文档", "goWeb"))
			return
		}

		// JWT 解析成功，设置用户信息
		userModel := user.Get(claims.UserID) //从token中获得用户ID!!!
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，用户可能已删除")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)

		c.Next()
	}
}
