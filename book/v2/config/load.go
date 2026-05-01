package config

import (
	"os"

	"github.com/caarlos0/env/v6" // 推荐使用的第三方库
	"gopkg.in/yaml.v3"
)

// 全局变量，通过函数对我提供访问
var config *Config

// C 如果没有配置文件怎么办？就使用默认配置，方便开发者
func C() *Config {
	if config == nil {
		config = Default()
	}

	return config
}

// LoadConfigFromYaml 加载配置，把外部配置读到 config全局变量里 yaml文件 -> config
func LoadConfigFromYaml(configPath string) error {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	// 默认值
	config = C()
	return yaml.Unmarshal(content, config)
}

// LoadConfigFromEnv 从环境变量读取配置
func LoadConfigFromEnv() error {
	config = C()

	return env.Parse(config)
}
