package user

import (
	"goWeb/Hash"
	"goWeb/app/models"
	"goWeb/pkg/database"
)

type User struct { //结构体大写！！
	models.BaseModel
	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"` //"-"这是在指示 JSON 解析器忽略字段 。后面接口返回用户数据时候，这三个字段都会被隐藏。

	models.CommonTimestampsField

	//在Go语言中，可以通过组合的方式来复用已有的结构体，并在新的结构体中添加自定义的字段和方法。
	//这样做可以避免代码冗余，并提高代码的可维护性和可复用性。
	//在这段代码中，User结构体组合了两个已有的模型结构体：
	//这两个模型结构体都定义了一些公共的字段和方法，User结构体继承了它们的公共字段和方法，并添加了自己的定制字段，比如Name、Email、Phone和Password。
}

func (User) TableName() string {
	return "tbtbtb"
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
