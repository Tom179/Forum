package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goWeb/bootstrap"
	"goWeb/captcha"
)

func main() {
	r := gin.New()
	bootstrap.SetupDB() //初始化数据库
	bootstrap.SetupRedis()
	bootstrap.SetupRoute(r) //初始化路由，包括中间件

	fmt.Println(captcha.NewCaptcha().VerifyCaptcha("07d7ADH7Fn7cvV4Vf9nM", "743529", true))

	err := r.Run(":3000") //最好写到配置文件中端口（数据库连接，密钥......都不要写死）
	if err != nil {
		fmt.Println(err.Error())
	}

}
