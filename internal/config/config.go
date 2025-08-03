package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"


type Config struct {
	DBURL string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}


func Read() (Config, error) {
	var cfg Config
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return cfg, err
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		return cfg, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}


func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(cfg)
}


func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configFilePath := filepath.Join(home, configFileName)
	return configFilePath, nil
}	


func write(cfg *Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	configJson, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(configFilePath, configJson, 0664)
	if err != nil {
		return err
	}
	return nil
}
