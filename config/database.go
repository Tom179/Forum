package config

import (
	"goWeb/pkg/config"
)

func init() { //这样的话config中每个模块的init都可以新增一套配置

	config.Add("database", func() map[string]interface{} {
		return map[string]interface{}{

			// 默认数据库
			"connection": config.Env("DB_CONNECTION", "mysql"),

			"mysql": map[string]interface{}{ //嵌套结构

				// 数据库连接信息
				"host":     config.Env("DB_HOST", "127.0.0.1"),
				"port":     config.Env("DB_PORT", "3306"),
				"database": config.Env("DB_DATABASE", "goweb"),
				"username": config.Env("DB_USERNAME", ""),
				"password": config.Env("DB_PASSWORD", ""),
				"charset":  "utf8mb4",

				// 连接池配置
				"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 25),  //最大空闲连接数，这里设置的默认值为 100，表示数据库连接池中最多可以保持的空闲连接数为 100
				"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 100), //最大打开连接数，表示数据库连接池中最多可以同时打开的连接数为 25
				"max_life_seconds":     config.Env("DB_MAX_LIFE_SECONDS", 5*60),
			},
			"sqlite": map[string]interface{}{
				"database": config.Env("DB_SQL_FILE", "database/database.db"),
			},
		}
	})
}
