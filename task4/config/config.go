package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	// 根据 config.yaml 调整字段
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`

	Logging struct {
		Level string `yaml:"level"`
	} `yaml:"logging"`

	Jwt struct {
		JWTSecret string `yaml:"secret"`
	} `yaml:"jwt"`
}

var Cfg Config

// Load 从指定路径读取 YAML 配置；若 path 为空则按优先级使用：ENV CONFIG_PATH -> ./config.yaml -> ./config/config.yaml -> ../config.yaml
func Load(path string) error {
	if path == "" {
		if p := os.Getenv("CONFIG_PATH"); p != "" {
			path = p
		} else {
			candidates := []string{"config.yaml", filepath.Join("config", "config.yaml"), filepath.Join("..", "config.yaml")}
			for _, c := range candidates {
				if _, err := os.Stat(c); err == nil {
					path = c
					break
				}
			}
			if path == "" {
				wd, _ := os.Getwd()
				for _, c := range candidates {
					p := filepath.Join(wd, c)
					if _, err := os.Stat(p); err == nil {
						path = p
						break
					}
				}
			}
		}
	}

	if path == "" {
		return os.ErrNotExist
	}

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("读取配置文件失败")
		return err
	}
	return yaml.Unmarshal(data, &Cfg)
}
