package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AddTopicRequest struct {
	Title      string `json:"title,omitempty" valid:"title"`
	Body       string `json:"body,omitempty" valid:"body"`
	CategoryID string `json:"category_id,omitempty" valid:"category_id"`
}
type TopicListRequest struct {
	PerPage     string `json:"per_page" valid:"per_page"`
	CurrentPage string `json:"current_page" valid:"current_page"`
}

type TopicSearchRequest struct {
	Keyword string `json:"keyword"`
}

func TopicSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"title":       []string{"required", "min_cn:3", "max_cn:40"},
		"body":        []string{"required", "min_cn:10", "max_cn:50000"},
		"category_id": []string{"required"}, //自定义验证规则查库
	}
	messages := govalidator.MapData{
		"title": []string{
			"required:帖子标题为必填项",
			"min_cn:标题长度需大于 3",
			"max_cn:标题长度需小于 40",
		},
		"body": []string{
			"required:帖子内容为必填项",
			"min_cn:长度需大于 10",
		},
		"category_id": []string{
			"required:帖子分类为必填项",
			//"exists:帖子分类未找到",
		},
	}
	return validate(data, rules, messages)
}

func TopicList(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"per_page":     []string{"required"},
		"current_page": []string{"required"},
	}
	messages := govalidator.MapData{
		"per_page": []string{
			"required:每页多少条数据必填",
			//"numeric:参数类型为数字",
		},
		"current_page": []string{
			"required:当前页号必填",
			//"numeric:参数类型为数字",
		},
	}
	return validate(data, rules, messages)
}
