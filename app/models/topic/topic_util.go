package topic

import (
	"fmt"
	"goWeb/pkg/database"
	"goWeb/pkg/logger"
	"gorm.io/gorm/clause"
)

func Get(idstr string) (topic Topic) {
	database.DB.Preload(clause.Associations).Where("id", idstr).First(&topic)
	return
}

func GetList(perPage int, currentPage int) (Topics []Topic) { //把分页逻辑抽取？
	//database.DB.Offset(currentPage - 1).Limit(perPage).Find(&Topics)
	var totalCount int64
	if err := database.DB.Model(&Topic{}).Count(&totalCount).Error; err != nil {
		logger.Warn("获取Topic总记录数失败")
	}
	if currentPage*perPage > int(totalCount) {
		logger.Error("页数超出数据范围")
	} else { //正常,进行分页查询
		offset := (currentPage - 1) * perPage
		database.DB.Limit(perPage).Offset(offset).Find(&Topics)
	}
	return
}

func Search(keyword string) (Topics []Topic) {
	result := database.DB.Where("title like ? or body like ?", "%"+keyword+"%", "%"+keyword+"%").Find(&Topics)
	if result.Error != nil {
		fmt.Println("错误信息为:", result.Error)
	}
	return
}
