package bootstrap

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"novel_learning/app/dao"
	"novel_learning/global/my_errors"
	"novel_learning/global/variable"
	zerolog "novel_learning/log"
	"os"
)

func checkRequiredFolders() {
	// 1, 检查项目配置文件
	if _, err := os.Stat(variable.BasePath + "/config/config.yml"); err != nil {
		log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
	}
	// 2, 检查数据库配置文件
	if _, err := os.Stat(variable.BasePath + "/config/database.yml"); err != nil{
		log.Fatal(my_errors.ErrorsConfigGormNotExists + err.Error())
	}
	// 3, 检查日志日志保存目录是否存在
	if _, err := os.Stat(variable.BasePath + "/storage/logs"); err != nil {
		log.Fatal(my_errors.ErrorsStorageLogsNotExists + err.Error())
	}
}

func initConfigYml() {
	viper.AddConfigPath(variable.BasePath + "/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
	}
	viper.SetConfigName("database")
	// 加载sql初始化文件，合并配置文件
	if err := viper.MergeInConfig(); err != nil {
		log.Fatal(my_errors.ErrorsConfigYamlNotExists + err.Error())
	}

}


/* **********
整个项目初始化
* **********/
func init() {
	fmt.Println("初始化ing ...")
	// 检查项目必须的非编译目录是否存在，避免编译后调用的时候缺少相关目录
	checkRequiredFolders()
	// 出书画配置文件相关的配置
	initConfigYml()
	fmt.Println("初始化完成！")
	// 初始化日志配置
	zerolog.SetUp()
	// 初始化数据库库配置
	dao.SetUp()
}