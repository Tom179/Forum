package bootstrap

import (
	"fmt"
	"goWeb/DataBase"
	"goWeb/app/models/user"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {

	var dbConfig gorm.Dialector
	// 构建 DSN 信息
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
		"root",
		"Qy85891607",
		"127.0.0.1",
		"3306",
		"goweb", //库名，gorm不能自动创库，只能自动创表
		"utf8mb4",
	)
	dbConfig = mysql.New(mysql.Config{
		DSN: dsn,
	})

	// 连接数据库，并设置 GORM 的日志模式
	DataBase.Connect(dbConfig, logger.Default.LogMode(logger.Info)) //给DB赋值，之后再访问DataBase包的时候可以直接使用DB
	DataBase.DB.AutoMigrate(&user.User{})                           //自动建表，填入多个参数同样能一次性建多个表

	DataBase.SQLDB.SetMaxOpenConns(10)                               // 设置最大连接数
	DataBase.SQLDB.SetMaxIdleConns(5)                                // 设置最大空闲连接数
	DataBase.SQLDB.SetConnMaxLifetime(time.Duration(24) * time.Hour) // 设置每个链接的过期时间//一天

}
