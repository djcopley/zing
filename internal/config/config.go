package config

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func getConfigDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "zing"), nil
}

// InitConfig initializes the viper configuration
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	zingConfigDir, err := getConfigDir()
	if err == nil {
		viper.AddConfigPath(zingConfigDir)
	}
	viper.AddConfigPath(".")

	viper.SetDefault("server_addr", "localhost")
	viper.SetDefault("server_port", 50051)
	viper.SetDefault("token", "")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			log.Printf("Error reading config file: %s\n", err)
		}
	}
}

// EnsureConfigFile ensures that the config file exists
func EnsureConfigFile() error {
	if viper.ConfigFileUsed() != "" {
		return nil
	}

	configDir, err := getConfigDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	configFile := filepath.Join(configDir, "config.toml")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		f, err := os.Create(configFile)
		if err != nil {
			return err
		}
		f.Close()
	}

	viper.SetConfigFile(configFile)
	return viper.WriteConfig()
}

// GetServerAddr returns the configured server address in the form of host:port
func GetServerAddr() string {
	addr := viper.GetString("server_addr")
	port := viper.GetInt("server_port")
	return fmt.Sprintf("%s:%d", addr, port)
}

// SetServerAddr stores the server address
func SetServerAddr(addr string) error {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	portNum, err := net.LookupPort("tcp", port)
	if err != nil {
		return err
	}
	viper.Set("server_addr", host)
	viper.Set("server_port", portNum)
	return viper.WriteConfig()
}

// GetToken returns the stored authentication token
func GetToken() string {
	return viper.GetString("token")
}

// SetToken stores the authentication token
func SetToken(token string) error {
	viper.Set("token", token)
	return viper.WriteConfig()
}
