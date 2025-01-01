package config

import (
	"encoding/json"
	"os"

	"github.com/mohdjishin/sportsphere/constants"
	"github.com/mohdjishin/sportsphere/env"
)

type ServerConfig struct {
	Addr         string `json:"addr"`
	ReadTimeout  int
	WriteTimeout int
	Port         string `json:"port"`
	LogLevel     string `json:"log_level"`
}

var (
	Config *ServerConfig
)

func Init() {
	pathToEnv := env.GetEnv(constants.CONFIG_PATH, "config/config.json")
	fileData, err := os.ReadFile(pathToEnv)
	// configuration is a basic requirement, hence we panic if it's not found
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileData, &Config)
	if err != nil {
		panic(err)
	}
}
