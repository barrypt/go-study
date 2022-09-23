
package config
 
import (
	"log"
 
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)
 
type Config struct {
	Name     string
	Host     string
	Username string
	Password string
}
 
func Init() (*Config, error) {
	viper.AddConfigPath("conf")   // 设置配置文件路径
	viper.SetConfigName("config") // 设置配置文件名
	viper.SetConfigType("yaml")   // 设置配置文件类型格式为YAML
 
	// 初始化配置文件
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return &Config{}, err
	}
	// 监控配置文件变化并热加载程序，即不重启程序进程就可以加载最新的配置
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
 
	c := &Config{
		Name:     viper.GetString("name"),
		Host:     viper.GetString("host"),
		Username: viper.GetString("username"),
		Password: viper.GetString("password"),
	}
 
	return c, nil
}
