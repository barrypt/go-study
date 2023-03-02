package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/constant"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/extension"
	"github.com/spf13/viper"
)

type ContentSwg struct {
	ShowDoc     bool   `json:"showDoc"`
	RoutePrefix string `json:"routePrefix"`
	Version     string `json:"version"`
}

type Swagger struct {
	SwaggerConfig *ContentSwg `json:"SwaggerConfig"`
}

func main() {
	c := &config.AppConfig{
		AppID:          "beta.infocore.h5api",
		Cluster:        "dev",
		IP:             "http://apollo.config.beta",
		NamespaceName:  "test.yml",
		IsBackupConfig: false,
		Secret:         "",
	}
	extension.AddFormatParser(constant.YML, &Parser{})
	agollo.SetBackupFileHandler(&FileHandler{})

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})

	viper.SetConfigType("yaml")
	viper.SetConfigName("2")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("C:\\Users\\49688\\Desktop")
	if err:= viper.ReadInConfig();err!=nil{
		fmt.Println("err",err)
	}
	var gg = viper.GetInt("A.b")
	var tt = viper.AllSettings()
	fmt.Println("ff", gg)
	fmt.Println("tt", tt)

	key1, _ := client.GetConfigCache("test.yml").Get("content")
	fmt.Println("key11", key1)

	if err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	split := strings.Split(c.NamespaceName, ",")
	for _, n := range split {
		checkKey(n, client)
	}

}

func checkKey(namespace string, client agollo.Client) {
	cache := client.GetConfigCache(namespace)
	count := 0
	cache.Range(func(key, value interface{}) bool {

		va := value.(string)

		var sw Swagger

		json.Unmarshal([]byte(va), &sw)

		fmt.Println("va", sw.SwaggerConfig.RoutePrefix)

		switch res := value.(type) {
		case string:
			json.Unmarshal([]byte(va), &sw)
			fmt.Println("va", sw.SwaggerConfig.RoutePrefix)
			fmt.Println("res", res)
		case *Swagger:
			ff := value.(*Swagger)
			fmt.Println("ff", ff.SwaggerConfig)
		case *ContentSwg:
			fmt.Printf("%+v\n", 111)
		}
		fmt.Println("key : ", key, ", value :", value)
		count++
		return true
	})
	if count < 1 {
		panic("config key can not be null")
	}
}

// FileHandler 默认备份文件读写
type FileHandler struct {
}

// WriteConfigFile write config to file
func (fileHandler *FileHandler) WriteConfigFile(config *config.ApolloConfig, configPath string) error {
	//fmt.Println(config.Configurations)
	return nil
}

// GetConfigFile get real config file
func (fileHandler *FileHandler) GetConfigFile(configDir string, appID string, namespace string) string {
	return ""
}

// LoadConfigFile load config from file
func (fileHandler *FileHandler) LoadConfigFile(configDir string, appID string, namespace string) (*config.ApolloConfig, error) {
	return &config.ApolloConfig{}, nil
}

// Parser properties转换器
type Parser struct {
}

// Parse 内存内容=>yml文件转换器
func (d *Parser) Parse(configContent interface{}) (map[string]interface{}, error) {
	fmt.Println("parse", configContent)

	m := make(map[string]interface{})
	m["content"] = configContent
	return m, nil
}
