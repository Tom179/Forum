package category

import (
	"goWeb/app/models"
	"goWeb/pkg/database"
)

type Category struct {
	models.BaseModel

	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	models.CommonTimestampsField
}

func (Category) TableName() string {
	return "category"
}

func (category *Category) Create() {
	database.DB.Create(&category)
}

func (category *Category) Save() int64 {
	res := database.DB.Save(&category)
	database.DB.Scopes()
	return res.RowsAffected
}
