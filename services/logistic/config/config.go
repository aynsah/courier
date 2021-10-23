package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Port string `json:"port"`

	DBType     string `json:"dbType"`
	DBAddress  string `json:"dbAddress"`
	DBUsername string `json:"dbUsername"`
	DBPassword string `json:"dbPassword"`
	DBName     string `json:"dbName"`
}

var Config *Configuration

func LoadConfig() error {
	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}

	Config = new(Configuration)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	return nil
}
