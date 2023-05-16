package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goWeb/app/http/MiddleWares"
	"goWeb/bootstrap"
	"net/http"
)

func main() {
	r := gin.New()
	bootstrap.SetupDB() //初始化数据库
	bootstrap.SetupRedis()
	bootstrap.SetupRoute(r) //初始化路由，包括中间件

	/*r.GET("/test_auth", MiddleWares.AuthJWT(), func(c *gin.Context) { //中间件测试
		userModel := auth.CurrentUser(c)
		response.Data(c, userModel)
	    //中间件测试：登录身份，哪些路由需要授权才能访问，直接为其添加中间件即可
	})*/
	r.GET("/test_guest", MiddleWares.GuestJWT(), func(c *gin.Context) {
		c.String(http.StatusOK, "Hello guest")
	}) //中间件测试：游客身份，哪些路由只有游客才能访问，直接为其添加中间件即可。
	//fmt.Println(captcha.NewCaptcha().VerifyCaptcha("07d7ADH7Fn7cvV4Vf9nM", "743529")) //检测验证码函数

	err := r.Run(":3000") //最好写到配置文件中端口（数据库连接，密钥......都不要写死）
	if err != nil {
		fmt.Println(err.Error())
	}

}
