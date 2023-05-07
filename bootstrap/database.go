package bootstrap

import (
	"fmt"
	"goWeb/DB"

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
		"121.0.0.1",
		"3306",
		"tb1", //表名
		"utf8mb4",
	)
	dbConfig = mysql.New(mysql.Config{
		DSN: dsn,
	})

	// 连接数据库，并设置 GORM 的日志模式
	DB.Connect(dbConfig, logger.Default.LogMode(logger.Info))

	// 设置最大连接数
	DB.SQLDB.SetMaxOpenConns(10)
	// 设置最大空闲连接数
	DB.SQLDB.SetMaxIdleConns(5)
	// 设置每个链接的过期时间
	DB.SQLDB.SetConnMaxLifetime(time.Duration(24) * time.Hour) //一天

}
