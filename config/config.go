package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config stores the global settings
type Config struct {
	ApiName    string `json:"api_name,required"`
	ApiVersion string `json:"api_version,required"`
	ApiAddr    string `json:"api_addr,required"`
	DbAddr     string `json:"db_addr,required"`
	DbName     string `json:"db_name,required"`
}

var instance *Config = nil

// init always runs only once, irresepective of number of times this package is imported
func init() {
	instance = loadConfig("config.json")
}

// GetInstance returns the singleton config instance
func GetInstance() *Config {
	return instance
}

func loadConfig(path string) *Config {
	f, err := os.Open(path)
	if err != nil {
		log.Print(err.Error())
	}
	defer f.Close()
	var conf Config
	err = json.NewDecoder(f).Decode(&conf)
	if err != nil {
		log.Print(err.Error())
	}

	return &conf
}
