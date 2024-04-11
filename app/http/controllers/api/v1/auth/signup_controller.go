package auth //注册

import (
	"github.com/gin-gonic/gin"
	v1 "goWeb/app/http/controllers/api/v1"
	"goWeb/app/models/user"
	"goWeb/app/requests"
	"goWeb/pkg/jwt"
	"goWeb/pkg/response"
)

type SignupController struct { //?继承自baseAPIController 有什么用？
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) { //处理函数:需要三个验证
	//panic("这是panic测试")
	request := requests.SignupPhoneExistRequest{}

	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok { //验证格式（c，..定义格式）
		return
	}

	/*c.JSON(http.StatusOK, gin.H{ //如果格式无误且成功获取到请求，查库
		"exist": user.IsPhoneExist(request.Phone),
	})*/
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone), //验证函数
	})

}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{} //创建结构体

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

// SignupUsingEmail 使用 Email + 验证码进行注册
func (sc *SignupController) SignupUsingEmail(c *gin.Context) {

	// 1. 验证表单
	request := requests.SignupUsingEmailRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingEmail); !ok {
		return
	}

	if user.IsEmailExist(request.Email) {
		response.Unauthorized(c, "邮箱已存在，注册失败")
		return
	}

	// 2. 验证成功，创建数据
	userModel := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	userModel.Create()

	if userModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)
		response.CreatedJSON(c, gin.H{
			"token": token,
			"data":  userModel,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}
