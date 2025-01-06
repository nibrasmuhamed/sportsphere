package config

import (
	"encoding/json"
	"os"

	"github.com/mohdjishin/sportsphere/pkg/constants"
	"github.com/mohdjishin/sportsphere/pkg/helper"
)

type ServerConfig struct {
	Addr         string `json:"addr"`
	ReadTimeout  int
	WriteTimeout int
	Port         string `json:"port"`
	LogLevel     string `json:"log_level"`
	DatabaseName string `json:"database_name"`
	DatabaseType string `json:"database_type"`
}

var (
	config *ServerConfig
)

func Init() {
	pathToEnv := helper.GetEnv(constants.CONFIG_PATH, "config/config.json")
	fileData, err := os.ReadFile(pathToEnv)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		panic(err)
	}
}

func Get() *ServerConfig {
	return config
}
