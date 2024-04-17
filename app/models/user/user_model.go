package user

import (
	"goWeb/app/models"
	"goWeb/pkg/Hash"
	"goWeb/pkg/database"
)

type User struct { //结构体大写！！
	models.BaseModel
	Name     string `json:"name,omitempty" gorm:"type:varchar(255);not null;index"`
	Email    string `json:"-" gorm:"type:varchar(255);not null"`
	Phone    string `json:"-" gorm:"type:varchar(255);not null"`
	Password string `json:"-" gorm:"type:varchar(255);not null"` //"-"这是在指示 JSON 解析器忽略字段 。后面接口返回用户数据时候，这三个字段都会被隐藏。

	models.CommonTimestampsField

	//在Go语言中，可以通过组合的方式来复用已有的结构体，并在新的结构体中添加自定义的字段和方法。
}

func (User) TableName() string { //使用*User和User都能绑定方法。如果需要修改结构体的状态或者避免额外的内存复制，建议使用指针类型作为接收者；如果只需要读取结构体而不需要修改，可以使用非指针普通类型类型作为接收者。
	return "user"
}

//如果模型中没有定义TableName方法，GORM会自动根据模型名生成对应的表名，生成规则为将模型名转为复数形式，然后在末尾添加"s"
//↑所以该模型对应的是users的表

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return Hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}

func (userModel *User) Create() {
	database.DB.Create(&userModel)
}
