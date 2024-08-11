package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func Load() error {
	rootDir, err := findRootDir()
	if err != nil {
		return err
	}
	envPath := filepath.Join(rootDir, "configs", ".env")
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		return fmt.Errorf("could not find .env file in %s", envPath)
	}

	viper.SetConfigFile(envPath)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	AppConfig = Config{
		Server: ServerConfig{
			Port: viper.GetInt("API_PORT"),
		},
		Database: DatabaseConfig{
			Port:     viper.GetInt("DB_PORT"),
			Host:     viper.GetString("DB_HOST"),
			User:     viper.GetString("DB_USER"),
			Name:     viper.GetString("DB_NAME"),
			Password: viper.GetString("DB_PASS"),
		},
	}
	return nil
}

func findRootDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if currentDir == "/" {
			break
		}

		if _, err := os.Stat(currentDir + "/go.mod"); err == nil {
			return currentDir, nil
		}

		currentDir = currentDir[:len(currentDir)-1]
	}

	return "", fmt.Errorf("could not find root directory")
}
