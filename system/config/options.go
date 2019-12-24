package config


import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
)

// IsReleaseMode 是否为生产模式
func IsReleaseMode() bool {
	env := os.Getenv("CFG_ENV")
	return env != "debug" && env != "development" && env != "dev"
}

// RootDir 应用所在根目录
func RootDir() string {
	wd, _ := os.Getwd()
	return wd
}

// HomeDir 当前用户的 $HOME 目录
func HomeDir() string {
	home, _ := homedir.Dir()
	return home
}

// HostPath 获取配置的虚拟目录
func HostPath() string {
	return viper.GetString("host.path")
}

// HostPort 获取配置的服务端口
func HostPort() int {
	return viper.GetInt("host.port")
}

// LogLevel 日志级别
func LogLevel() string {
	return viper.GetString("log.level")
}

// LogPath 日志存放目录
func LogPath() string {
	return viper.GetString("log.path")
}

// LogFilePrefix 日志文件前缀
func LogFilePrefix(cate string) string {
	return viper.GetString("log.prefix." + cate)
}
