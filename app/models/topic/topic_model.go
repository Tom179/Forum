package topic

import (
	"goWeb/app/models"
	"goWeb/app/models/category"
	"goWeb/app/models/user"
)

type Topic struct {
	models.BaseModel

	Title      string `json:"title,omitempty" `      //标题
	Body       string `json:"body,omitempty" `       //正文
	UserID     string `json:"user_id,omitempty"`     //创建者
	CategoryID string `json:"category_id,omitempty"` //分类

	// 通过 user_id 关联用户
	User user.User `json:"user"` //关联用户

	// 通过 category_id 关联分类
	Category category.Category `json:"category"` //关联分类

	models.CommonTimestampsField
}
