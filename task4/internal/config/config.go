package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sync"
)

var (
	once     sync.Once
	instance *Configuration
)

type Configuration struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`

	MySQL struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`

	Auth struct {
		JwtSecret string `yaml:"jwt_secret"`
		TokenExp  int64  `yaml:"token_exp"`
	} `yaml:"auth"`
}

func GetConfig() *Configuration {
	return instance
}

func InitConfig(Path string) {
	once.Do(func() {
		if err := LoadConfig(Path); err != nil {
			log.Fatalf("yaml配置加载失败: %v", err)
		}
	})
}

func LoadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("加载全局配置文件失败: %w", err)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(instance); err != nil {
		return fmt.Errorf("解析全局配置文件失败: %w", err)
	}

	if instance.Server.Port == "" {
		instance.Server.Port = "8080"
	}
	if instance.MySQL.Port == "" {
		instance.MySQL.Port = "3306"
	}

	return nil
}
