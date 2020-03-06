package config

import (
	"encoding/json"
	"errors"
	"os"
	. "sling_cafe/log"
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
	err := validateConfig()
	if err != nil {
		Log.Error(err.Error())
	}
}

// GetInstance returns the singleton config instance
func GetInstance() *Config {
	return instance
}

func validateConfig() error {
	if instance.ApiName == "" {
		return errors.New("api_name field is empty")
	}
	if instance.ApiVersion == "" {
		return errors.New("api_version field is empty")
	}
	if instance.ApiAddr == "" {
		return errors.New("api_addr field is empty")
	}
	if instance.DbAddr == "" {
		return errors.New("db_addr field is empty")
	}
	if instance.DbName == "" {
		return errors.New("db_name field is empty")
	}

	return nil
}

func loadConfig(path string) *Config {
	f, err := os.Open(path)
	if err != nil {
		Log.Error(err.Error())
	}
	defer f.Close()
	var conf Config

	err = json.NewDecoder(f).Decode(&conf)
	if err != nil {
		Log.Error(err.Error())
	}

	return &conf
}
