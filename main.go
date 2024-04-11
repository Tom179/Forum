package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"goWeb/bootstrap"
	_ "goWeb/config"
	"goWeb/pkg/config" //btsConfig和config的区别？这两个配置分别配置了什么信息？
)

func init() {
	// 为了调用root/config的init()函数，有必要吗？不是导包就行？
	//btsConfig.Initialize()
}

func main() {
	//map1 := make(map[string]string)
	//map1["key1"] = "what"
	//fmt.Println(map1."key1")

	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()

	config.InitConfig(env) //根据命令行传入的参数加载配置。（导入config包时调用init函数，新建viper对象，设置读取配置文件的类型、后缀等）
	/*	fmt.Println("读取配置文件")

		prt := config.InternalGet("app.port") //证明viper中的多级配置，不论是yaml和map键值都可以通过'.'的方式来获取多级配置
		fmt.Println("读取到的app.port为", prt)
	*/

	bootstrap.SetupLogger() // 初始化 Logger

	r := gin.New()
	bootstrap.SetupDB()     //初始化数据库
	bootstrap.SetupRedis()  //初始化redis
	bootstrap.SetupRoute(r) //初始化路由，包括中间件

	/*r.GET("/test_auth", MiddleWares.AuthJWT(), func(c *gin.Context) { //中间件测试
		userModel := auth.CurrentUser(c)
		response.Data(c, userModel)
	    //中间件测试：登录身份，哪些路由需要授权才能访问，直接为其添加中间件即可
	})*/
	/*r.GET("/test_guest", MiddleWares.GuestJWT(), func(c *gin.Context) {
		c.String(http.StatusOK, "Hello guest")
	})*/ //中间件测试：游客身份，哪些路由只有游客才能访问，直接为其添加中间件即可。
	//fmt.Println(captcha.NewCaptcha().VerifyCaptcha("07d7ADH7Fn7cvV4Vf9nM", "743529")) //检测验证码函数

	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		// 错误处理，端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
