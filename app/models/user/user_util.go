package user

import "goWeb/DataBase" //注意：两个包不能相互依赖

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool { //仅传入一个字符串string
	var count int64
	DataBase.DB.Model(User{}).Where("email = ?", email).Count(&count) //连接的表对应在model结构体带的TableName方法中，这里为"users"
	return count > 0
} //DB.Model(结构体{})绑定了整个表的结构，无需传入任何值

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	DataBase.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
