package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery()) //日志和错误回复中间件

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})

	})

	r.Run(":8080") //默认8080
}
