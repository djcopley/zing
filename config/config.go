package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Config behavior:
// 1. Config located at $HOME/.config/zing/config.toml

// Fork in the road: one config table? or different client/server config tables?
// Start with one config table until there is a strong need for multiple - if maintaining canonical server port number
// with dual sources of truth becomes cumbersome, we can pull that configuration into this file.

type Config struct {
	ServerAddr string
	ServerPort int
	Token      string
}

// InitConfig initializes the viper configuration
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	homeDir, err := os.UserHomeDir()
	if err == nil {
		viper.AddConfigPath(filepath.Join(homeDir, ".zing"))
	}
	viper.AddConfigPath(".")

	viper.SetDefault("server_addr", "")
	viper.SetDefault("server_port", 5132)
	viper.SetDefault("token", "")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError *viper.ConfigFileNotFoundError
		if !errors.As(err, configFileNotFoundError) {
			log.Printf("Error reading config file: %s\n", err)
		}
	}
}

// GetServerAddr returns the configured server address in the form of host:port
func GetServerAddr() string {
	addr := viper.GetString("server_addr")
	port := viper.GetInt("server_port")
	return fmt.Sprintf("%s:%d", addr, port)
}

// SetServerAddr stores the server address
func SetServerAddr(addr string) error {
	viper.Set("server_addr", addr)
	return viper.WriteConfig()
}

// GetHost returns the host part of the server address (for backward compatibility)
func GetHost() string {
	addr := GetServerAddr()
	host, _, _ := splitServerAddr(addr)
	return host
}

// GetPort returns the port part of the server address (for backward compatibility)
func GetPort() int {
	addr := GetServerAddr()
	_, port, _ := splitServerAddr(addr)
	return port
}

// splitServerAddr splits a server address into host and port parts
func splitServerAddr(addr string) (host string, port int, err error) {
	host = "localhost"
	port = 6132

	if addr == "" {
		return
	}

	parts := strings.Split(addr, ":")
	if len(parts) > 0 && parts[0] != "" {
		host = parts[0]
	}

	if len(parts) > 1 && parts[1] != "" {
		portInt, err := strconv.Atoi(parts[1])
		if err == nil {
			port = portInt
		}
	}

	return
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

// EnsureConfigFile ensures that the config file exists
func EnsureConfigFile() error {
	if viper.ConfigFileUsed() != "" {
		return nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir := filepath.Join(homeDir, ".zing")
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
