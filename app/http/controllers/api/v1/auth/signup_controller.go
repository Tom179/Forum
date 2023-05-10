package auth //注册

import (
	"github.com/gin-gonic/gin"
	v1 "goWeb/app/http/controllers/api/v1"
	"goWeb/app/models/user"
	"goWeb/app/requests"
	"goWeb/app/response"
)

type SignupController struct { //?继承自baseAPIController 有什么用？
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) { //处理函数:需要三个验证

	request := requests.SignupPhoneExistRequest{}

	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok { //验证格式（c，..定义格式）
		return
	}

	/*c.JSON(http.StatusOK, gin.H{ //如果格式无误且成功获取到请求，查库
		"exist": user.IsPhoneExist(request.Phone),
	})*/
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})

}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{} //创建结构体
	/*	if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"errors": err,
			}) //
			fmt.Println(err)
			return
		}
		errs := requests.SignupEmailExist(&request, c) //验证请求格式

		if len(errs) > 0 { //不能采用errs！=nil因为就算没错，也是一个长度为0的空切片，而不是nil
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"errors": errs,
			})
			fmt.Println(errs)
			return
		}*/ //写到了👇以下函数中

	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}

	/*c.JSON(http.StatusOK, gin.H{ //如果成功获取到请求，验证,传入一个string
		"exist": user.IsEmailExist(request.Email), //查库
	})*/
	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})

}
