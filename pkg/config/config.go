package config //配置信息
import (
	"fmt"
	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
	"goWeb/pkg/helpers"
	"os"
)

var viper *viperlib.Viper

// ConfigFunc 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// ConfigFuncs 先加载到此数组，loadConfig 再动态生成配置信息
var ConfigFuncs map[string]ConfigFunc //一个键为ConfigFunc的map，而ConfigFunc函数返回的又是一个（）的map
func init() {

	viper = viperlib.New()
	viper.AddConfigPath(".")   //环境变量配置文件查找的路径，相对于 main.go
	viper.SetConfigType("env") //直接写后缀没有.

	/*	viper.SetEnvPrefix("appenv") //告诉 viper 在读取环境变量时，只解析带有特定前缀的环境变量，并忽略其他环境变量

		viper.AutomaticEnv()*/

	ConfigFuncs = make(map[string]ConfigFunc)
}
func InitConfig(env string) {
	// 1. 加载环境变量
	loadEnv(env) //已经读取了.env文件，写入到了viper中
	// 2. 注册配置信息
	loadConfig() //把configFuncsMap写到viper配置信息中
	////////////////加载了2个配置，直接写在.env文件配置和
	allSettings := viper.AllSettings()
	fmt.Println("遍历并输出所有配置信息:", allSettings)
	for key, value := range allSettings {
		fmt.Printf("------%s: %v\n", key, value)
	}
}

func loadConfig() { //把map写入到viper中
	fmt.Println("configFunc为", ConfigFuncs)
	for name, fn := range ConfigFuncs { //遍历
		viper.Set(name, fn()) //viper.Set()函数并不会写入到配置文件中，而是内存中
	}
}

func loadEnv(envSuffix string) {

	// 默认加载 .env 文件，如果有传参 --env=name 的话，加载 .env.name 文件
	envPath := ".env"
	if len(envSuffix) > 0 {
		filepath := ".env." + envSuffix
		if _, err := os.Stat(filepath); err == nil {
			// 如 .env.testing 或 .env.stage
			envPath = filepath
		} else {
			panic(err)
		}
	}

	// 加载 env
	viper.SetConfigName(envPath)                 //环境变量的文件名
	if err := viper.ReadInConfig(); err != nil { //读取配置文件
		panic(err)
	}

	// 监控 .env 文件，变更时马上重新加载
	viper.WatchConfig()
}

// Env 读取环境变量，支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 { //有变长参数
		return InternalGet(envName, defaultValue[0])
	}
	return InternalGet(envName)
}

// Add 新增配置项
func Add(name string, configFn ConfigFunc) { //将函数赋值给ConfigFuncsMap
	ConfigFuncs[name] = configFn
}

// Get 获取配置项
// 第一个参数 path 允许使用点式获取，如：app.name
// 第二个参数允许传参默认值
func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue...)
}

func InternalGet(path string, defaultValue ...interface{}) interface{} { //读取viper配置信息，为什么不直接用viper.get()而是封装多层？
	// config 或者环境变量不存在的情况
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) { //有点不明白什么情况下会Set(path),这个“path”是没有意义的，它知识一个形参的名字，具体要看实参传入的类型是什么。比如，path==之前写入viper的configFuncMap的键时，就相等了噻
		if len(defaultValue) > 0 {
			return defaultValue[0] //?
		}
		return nil
	}
	//fmt.Println("InternalGet()内部获得的app.port为", viper.Get("app.port"), "访问", path)
	return viper.Get(path)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(InternalGet(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(InternalGet(path, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(InternalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(InternalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(InternalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(InternalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
