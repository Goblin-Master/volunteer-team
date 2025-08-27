package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"runtime"
	"time"
)

func GetRootPath(myPath string) string {
	_, fileName, _, ok := runtime.Caller(0)
	if !ok {
		panic("Something wrong with getting root path")
	}
	absPath, err := filepath.Abs(fileName)
	rootPath := filepath.Dir(filepath.Dir(absPath))
	if err != nil {
		panic(any(err))
	}
	return filepath.Join(rootPath, myPath)
}
func LoadConfig() {
	// 初始化时间为东八区的时间
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
	viper.SetConfigName("config")           // 配置文件名称（无扩展名）
	viper.SetConfigType("yaml")             // 配置类型
	viper.AddConfigPath(GetRootPath("etc")) // 配置文件路径
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("配置读取失败: %w", err))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("配置解析失败: %w", err))
	}
	return
}
