package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Config struct {
	ServerAddr string
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

	viper.SetDefault("server_addr", "localhost:5132")
	viper.SetDefault("token", "")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Fprintf(os.Stderr, "Error reading config file: %s\n", err)
		}
	}
}

// GetServerAddr returns the configured server address in the form of host:port
func GetServerAddr() string {
	return viper.GetString("server_addr")
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

	configFile := filepath.Join(configDir, "config.yaml")
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
