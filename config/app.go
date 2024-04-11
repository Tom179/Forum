// Package config 站点配置信息
package config

import (
	"goWeb/pkg/config"
)

func init() {
	config.Add("app", func() map[string]interface{} { //Add():将函数（值）赋值给configFuncMap中，name(键)。这个值里面的函数做的是返回一个map[string]interface{}类型，也就是（不同键对应的不同场景，每个函数返回一套不同的配置）一组配置信息：
		return map[string]interface{}{

			// 应用名称
			"name": config.Env("APP_NAME", "goWeb"), //配置信息默认值

			// 当前环境，用以区分多环境，一般为 local, stage, production, test
			"env": config.Env("APP_ENV", "production"),

			// 是否进入调试模式
			"debug": config.Env("APP_DEBUG", false),
			"新配置":   config.Env("asdfoaw", "这是默认信息"), //Env函数就是去配置里面找，如果有就返回已有的信息，没有就返回默认信息
			// 应用服务端口
			"port": config.Env("APP_PORT", "3000"),

			// 加密会话、JWT 加密
			"key": config.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),

			// 用以生成链接
			"url": config.Env("APP_URL", "http://localhost:3000"),

			// 设置时区，JWT 里会使用，日志记录里也会使用到
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
