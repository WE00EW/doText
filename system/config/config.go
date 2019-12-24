package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
	"strings"
)


func init() {
	//配置优先级: 命令行参数 > 环境变量 > 配置中心｜本地配置文件

	//默认使用全局viper实例
	v := viper.GetViper()
	useEnv(v)
	useDefault(v)
	//disabled := useApollo(v)

	// 远程配置中心被禁用
	//if disabled {

		useLocalConfigFiles(v)
	//	source = Local
	//} else {
	//	source = Remote
	//}

}

// useEnv 允许通过环境变量进行设置
func useEnv(v *viper.Viper) {
	// 通过环境变量设置时，“host.path” 应当设置为 “HOST_PATH”
	envReplacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(envReplacer)
	v.AutomaticEnv()
}

func useDefault(v *viper.Viper) {
	v.SetDefault("host.path", "/api")
	v.SetDefault("host.port", 80)
	v.SetDefault("log.level", "info")
	v.SetDefault("log.path", "./logs")
	v.SetDefault("log.filename", "app-%Y%m%d.log")
}

// 使用本地配置文件
func useLocalConfigFiles(v *viper.Viper) {
	//log.Infof("从本地获取配置...\n")

	err :=useConfigFiles(v,getConfigDirs(),"config","yaml")

	// 加载本地配置文件失败,panic
	if err !=nil {
		panic(err)
	}
}

func useConfigFiles(v *viper.Viper,dirs []string,cfgName string,cfgExt string) error {
	v.SetConfigName(cfgName)
	v.SetConfigType(cfgExt)

	for _, dir := range dirs {
		if len(dir) == 0 {
			continue
		}
		//log.Infof("Finding config file (%s) in: %s\r\n",cfgName, dir)
		v.AddConfigPath(dir)
	}

	err := v.ReadInConfig() // Find and read the config file
	if err != nil {
		// Handle errors reading the config file
		return fmt.Errorf("从如下位置加载 \"%s\" 失败！\r\n%s\r\n\r\n错误信息:%s\r\n\r\n提示：您可以通过设置环境变量来确定配置文件名称和位置: ${CFG_DIR}/${CFG_FILE}[.${CFG_ENV}].yaml.\r\n对于apollo配置文件，使用：${CFG_DIR}/${CFG_FILE_APOLLO}[.${CFG_ENV}].yaml", cfgName+"."+cfgExt, strings.Join(dirs, "\r\n- "),err.Error())
	}

	//cfgfile := v.ConfigFileUsed()
	//log.Infof("Config file found: %s\r\n", cfgfile)

	//isRelease := IsReleaseMode()
	//log.Infof("IsReleaseMode=%v\r\n", isRelease)

	return nil
}


// 获取配置文件路径
func getConfigDirs() []string {
	// 添加多个配置文件路径，以先找到的为准
	rootDir := RootDir()
	configDir := filepath.FromSlash(path.Join(rootDir, "config"))
	envDir := os.Getenv("CFG_DIR")
	//log.Infof("CFG_DIR=%s\r\n", envDir)

	homeDir := filepath.FromSlash(path.Join(HomeDir(), ".fip"))
	etcDirs := []string{envDir, configDir, rootDir, homeDir}
	return etcDirs
}

// 获取配置文件名
//func getConfigName() string {
//	//cfgName := strings.TrimSpace(os.Getenv("CFG_FILE"))
//	//if len(cfgName) == 0 {
//	//	cfgName = defaultCFGName
//	//}
//	//
//	//log.Infof("CFG_FILE=%s\r\n", cfgName)
//	//
//	//// 根据环境变量加载不同的配置文件
//	//env := strings.TrimSpace(os.Getenv("CFG_ENV"))
//	//
//	//if len(env) == 0 {
//	//	env = defaultCFGEnv
//	//}
//	//
//	//cfgFileName := cfgName + "." + env
//	//
//	//log.Infof("CFG_ENV=%s\r	\n", env)
//	//return cfgFileName
//}