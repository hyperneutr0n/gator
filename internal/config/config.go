package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DB_URL							string		`json:"db_url"`
	CurrentUsername			string		`json:"current_user_name"`
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUsername = username
	return write(*cfg)
}

const configFileName string = ".gatorconfig.json"

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, nil
	}

	configFile, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}
	defer configFile.Close()

	cfg := Config{}
	err = json.NewDecoder(configFile).Decode(&cfg)
	if err != nil {
		return Config{}, nil
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configFileName), nil
}

func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	configFile, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer configFile.Close()

	err = json.NewEncoder(configFile).Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}