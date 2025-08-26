package config

import (
	"os"
	"encoding/json"
	"path/filepath"
)

type Config struct {
	DB_URL          string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUsername = username
	return write(*cfg)
}

const configFileName string = ".gatorconfig.json"

func Read() (cfg Config, err error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	configFile, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}
	defer func() {
		closeErr := configFile.Close()
		if err == nil {
			err = closeErr
		}
	}()
	err = json.NewDecoder(configFile).Decode(&cfg)
	if err != nil {
		return Config{}, err
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

func write(cfg Config) (err error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	configFile, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := configFile.Close()
		if err == nil {
			err = closeErr
		}
	}()

	err = json.NewEncoder(configFile).Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
