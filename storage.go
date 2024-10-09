package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/adrg/xdg"
)

type Config string

type ConfigFileStorage struct {
	filePath string
}

func InitNewConfigFileStorage(folderName string) (*ConfigFileStorage, error) {
	filePath, err := xdg.ConfigFile(folderName)
	if err != nil {
		return nil, fmt.Errorf("could not resolve path for config file: %w", err)
	}

	return &ConfigFileStorage{
		filePath: filePath,
	}, nil
}

/**
* How to use runtime binding:
* ConfigStore.Get("settings.json", "null").then((response) => {
*     const data = JSON.parse(response);
* });
*/
func (s *ConfigFileStorage) Get(fileName string, defaultValue string) (Config, error) {
	_, err := os.Stat(s.filePath + string(os.PathSeparator) + fileName)
	if os.IsNotExist(err) {
		return Config(defaultValue), nil
	}
	buf, err := fs.ReadFile(os.DirFS(s.filePath), fileName)
	if err != nil {
		return Config(defaultValue), fmt.Errorf("could not read the configuration file: %w", err)
	}
	return Config(buf), nil
}

/**
* How to use runtime binding:
* ConfigStore.Set("settings.json", JSON.stringify({"key1": "value1", "key2": "value2"});
*/
func (s *ConfigFileStorage) Set(fileName string, value Config) error {
	err := os.MkdirAll(s.filePath, 0755)
	if err != nil {
		return fmt.Errorf("could not create the configuration directory: %w", err)
	}
	err = os.WriteFile(s.filePath+string(os.PathSeparator)+fileName, []byte(value), 0644)
	if err != nil {
		return fmt.Errorf("could not write the configuration file: %w", err)
	}
	return nil
}