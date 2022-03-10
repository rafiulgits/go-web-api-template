package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
	ListenPort    int        `json:"listenPort"`
	DebugEnv      bool       `json:"debug"`
	DBMigration   bool       `json:"dbMigration"`
	SeedDBEntries bool       `json:"seedDbEntries"`
	AppAPIKey     string     `json:"appAPIKey"`
	DBConfig      DBConfig   `json:"db"`
	LogConfig     LogConfig  `json:"logging"`
	CorsConfig    CorsConfig `json:"cors"`
}

const filename = "appsettings.json"

var appConfig *AppConfig

func loadConfig() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	filePath := fmt.Sprintf("%s/%s", workingDir, filename)
	stream, err := os.Open(filePath)
	if err != nil {

		panic(err)
	}
	parseErr := json.NewDecoder(stream).Decode(&appConfig)
	if parseErr != nil {
		panic(parseErr)
	}
}

func LoadConfig() *AppConfig {
	var loadOnce sync.Once
	loadOnce.Do(loadConfig)
	return appConfig
}

func GetAppConfig() *AppConfig {
	return appConfig
}
