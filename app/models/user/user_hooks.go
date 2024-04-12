package user

import (
	"goWeb/pkg/Hash"

	"gorm.io/gorm"
)

// BeforeSave GORM 的模型钩子，在创建和更新模型Create、Update、Save前被自动调用!!!!!!!
// 因为我们使用的是模型钩子，所以原有的注册逻辑不需要修改
func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {
	if !Hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = Hash.BcryptHash(userModel.Password)
	}
	return
}
