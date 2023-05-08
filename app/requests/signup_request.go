package requests //请求验证容错

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"` //valid是一个常用的第三方验证库，可以定义结构体字段的验证规则，如果不符合，验证库会返回相应的错误信息
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func SignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	//request包下的SignupPhoneExist仅定义格式：后续可换函数名
	//传入请求结构体，返回ValidateStruct函数
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}
	return validate(data, rules, messages)
}

func SignupEmailExist(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}
	// 配置初始化
	return validate(data, rules, messages)
}
