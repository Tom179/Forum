package user

import "goWeb/app/models"

type User struct { //结构体大写！！
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"` //"-"这是在指示 JSON 解析器忽略字段 。后面接口返回用户数据时候，这三个字段都会被隐藏。

	models.CommonTimestampsField
}

//在Go语言中，可以通过组合（Composition）的方式来复用已有的结构体，并在新的结构体中添加自定义的字段和方法。
//这样做可以避免代码冗余，并提高代码的可维护性和可复用性。
//在这段代码中，User结构体组合了两个已有的模型结构体：
//models.BaseModel和models.CommonTimestampsField。
//这两个模型结构体都定义了一些公共的字段和方法，用于存储和管理模型的基本信息，
//比如创建时间、更新时间、ID等。通过组合这两个模型结构体，User结构体继承了它们的公共字段和方法，
//并添加了自己的定制字段，比如Name、Email、Phone和Password。
