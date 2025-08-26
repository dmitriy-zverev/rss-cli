package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	DBUrl           string `json:string`
	CurrentUserName string `json:string`
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	configFile, err := os.OpenFile(configPath, os.O_RDONLY, 0755)
	if err != nil {
		return Config{}, err
	}

	var config Config
	decoder := json.NewDecoder(configFile)
	if err = decoder.Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(name string) error {
	if name == "" {
		return errors.New("error: cannot set empty name")
	}

	c.CurrentUserName = name
	if err := write(*c); err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := homeDir + "/" + CONFIG_FILE_NAME
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		configFile, err := os.Create(configPath)
		if err != nil {
			return "", err
		}

		configFile.Write([]byte(CONFIG_FILE_PLACEHOLDER))
	}

	return configPath, nil
}

func write(config Config) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := homeDir + "/" + CONFIG_FILE_NAME
	configFile, err := os.OpenFile(configPath, os.O_RDWR, 0755)
	if err != nil {
		return err
	}

	configJsonData, err := json.Marshal(config)
	if err != nil {
		return err
	}

	configFile.Write(configJsonData)

	return nil
}
