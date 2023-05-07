package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goWeb/bootstrap"
)

func main() {
	r := gin.New()
	bootstrap.SetupDB()
	bootstrap.SetupRoute(r) //写入中间件，路由的注册，初始化路由绑定
	err := r.Run(":3000")   //端口，数据库连接，密钥......都不要写死，写到配置文件中
	if err != nil {
		fmt.Println(err.Error())
	}

}
