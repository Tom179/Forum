package category

import "goWeb/pkg/database"

func Get(id string) (cat Category) {

	database.DB.Where("id", id).First(&cat)
	return cat
}

func All() (categorys []Category) {
	database.DB.Find(&categorys)
	return
}
