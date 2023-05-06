package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goWeb/bootstrap"
)

func main() {
	r := gin.New()
	bootstrap.SetupRoute(r) //写入中间件，路由的注册
	err := r.Run(":3000")   //默认8080
	if err != nil {
		fmt.Println(err.Error())
	}
}
