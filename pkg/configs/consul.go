package configs

import (
	"github.com/goccy/go-json"
	"go.uber.org/zap"
)

type AppConfig struct {
	DbConnectionString string
}

var (
	appConfig AppConfig
)

func SetConfigs(conf []byte) {
	ac := AppConfig{}
	err := json.Unmarshal(conf, &ac)
	if err != nil {
		zap.L().Fatal("Error while unmarshalling configs", zap.Error(err))
		panic("Config not found!")
	}
	appConfig = ac
}

func GetConfigs() *AppConfig {
	return &appConfig
}
