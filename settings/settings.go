package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config") // 指定配置文件名称(不需要制定配置文件的扩展名)
	viper.SetConfigType("yaml")   // 指定配置文件类型(专用于配置文件没有扩展名的情况)
	viper.AddConfigPath(".")      // 指定查找配置文件的路径(这里使用相对路径)
	err = viper.ReadInConfig()    // 读取配置信息
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return err
		// panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
	})
	return nil
}
