package auth //注册

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "goWeb/app/http/controllers/api/v1"
	"goWeb/app/models/user"
	"goWeb/app/requests"
	"net/http"
)

type SignupController struct { //?继承自baseAPIController 有什么用？
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) { //处理函数:需要三个验证

	request := requests.SignupPhoneExistRequest{}

	if err := c.ShouldBindJSON(&request); err != nil { //从请求c中绑定参数到request结构体中
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}
	//用govalidator验证request↓
	errs := requests.ValidateSignupPhoneExist(&request, c)
	// errs 返回长度等于零即通过，大于 0 即有错误发生
	if len(errs) > 0 {
		// 验证失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}
	//验证request

	c.JSON(http.StatusOK, gin.H{ //如果成功获取到请求，验证
		"exist": user.IsPhoneExist(request.Phone),
	})

}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{} //创建结构体
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": err,
		}) //
		fmt.Println(err)
		return
	}
	errs := requests.ValidateSignupEmailExist(&request, c)
	if len(errs) > 0 { //不能采用errs！=nil因为就算没错，也是一个长度为0的空切片，而不是nil
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		fmt.Println(errs)
		return
	}
	c.JSON(http.StatusOK, gin.H{ //如果成功获取到请求，验证,传入一个string
		"exist": user.IsEmailExist(request.Email),
	})

}
