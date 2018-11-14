package master

import (
	"encoding/json"
	"io/ioutil"
)

// Config 程序配置
type Config struct {
	APIPort         int `json:"apiPort"`
	APIReadTimeout  int `json:"apiReadTimeput"`
	APIWriteTimeout int `json:"apiWriteTimeout"`
}

var (
	// GConfig 单例
	GConfig *Config
)

// InitConfig 加载配置
func InitConfig(filename string) (err error) {
	var (
		conf Config
	)

	// 读取配置文件
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// 做JSON反序列化
	err = json.Unmarshal(content, &conf)
	if err != nil {
		return err
	}

	GConfig = &conf

	return
}
