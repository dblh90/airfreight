package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type App struct {
	DBConfig *DbConfig `yaml:"db"`
	Port     string    `yaml:"port"`
}

type AppConfig struct {
	App App `yaml:"app"`
}

type DbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"name"`
}

func LoadConfig(configPath string) (*AppConfig, error) {
	yfile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	config := AppConfig{}
	err = yaml.Unmarshal(yfile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
