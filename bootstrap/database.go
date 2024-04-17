package bootstrap

import (
	"fmt"
	"goWeb/app/models/category"
	"goWeb/app/models/topic"
	"goWeb/app/models/user"
	"goWeb/pkg/config"
	"goWeb/pkg/database"
	"goWeb/pkg/logger"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() { //请换为用配置文件的形式

	var dbConfig gorm.Dialector
	// 构建 DSN 信息
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
		config.Get("database.mysql.username"),
		config.Get("database.mysql.password"),
		config.Get("database.mysql.host"),
		config.Get("database.mysql.port"),
		config.Get("database.mysql.database"),
		config.Get("database.mysql.charset"),
	) //构建字符串

	dbConfig = mysql.New(mysql.Config{ //?和mysql.open(dsn)有什么不一样？
		DSN: dsn,
	})

	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dbConfig, logger.NewGormLogger()) /*logger.Default.LogMode(logger.Info)更换自定义*/ //给DB赋值，之后再访问DataBase包的时候可以直接使用DB
	database.DB.AutoMigrate(&user.User{})              //自动建表，填入多个参数同样能一次性建多个表
	database.DB.AutoMigrate(&category.Category{}, &topic.Topic{})
	//多个表每次都写成AotoMigrate这样是不是不好

	database.SQLDB.SetMaxOpenConns(10)                               // 设置最大连接数
	database.SQLDB.SetMaxIdleConns(5)                                // 设置最大空闲连接数
	database.SQLDB.SetConnMaxLifetime(time.Duration(24) * time.Hour) // 设置每个链接的过期时间//一天

}
