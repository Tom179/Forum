package topic

import (
	"goWeb/app/models"
	"goWeb/app/models/category"
	"goWeb/app/models/user"
	"goWeb/pkg/database"
)

type Topic struct {
	models.BaseModel

	Title      string `json:"title,omitempty" gorm:"type:varchar(255);not null;index"` //标题
	Body       string `json:"body,omitempty" gorm:"type:longtext;not null"`            //正文
	UserID     string `json:"user_id,omitempty" gorm:"type:bigint;not null;index"`     //创建者
	CategoryID string `json:"category_id,omitempty" gorm:"type:bigint;not null;index"` //分类

	// 通过 user_id 关联用户，有什么用？？？？
	User user.User `json:"user"` //
	// 通过 category_id 关联分类
	Category category.Category `json:"category"`

	models.CommonTimestampsField
}

func (topic *Topic) Create() {
	database.DB.Create(&topic)
}
func (topic *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&topic)
	return result.RowsAffected
}

func (topic *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topic) //查找结构体对象内的id自动进行匹配
	return result.RowsAffected
}
