
package main
 
import (
	"fmt"
	"time"
	"main/config"
	"github.com/spf13/viper"
)
 
func main() {
 
	// init config
	_, err := config.Init()
	if err != nil {
		fmt.Println(err)
	}
	// 注意：只能在 init 之后再次通过 viper.Get 方法读取配置，否则不生效
	for {	
		cfg := &config.Config{
			Name:     viper.GetString("name"),
			Host:     viper.GetString("host"),
			Username: viper.GetString("username"),
			Password: viper.GetString("password"),
		}
		fmt.Println(cfg.Name)
		fmt.Println(cfg.Host)
		fmt.Println(cfg.Username)
		fmt.Println(viper.GetBool("aa"))
		time.Sleep(4 * time.Second)
	}
 
}