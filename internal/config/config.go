package config

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func getConfigDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "zing"), nil
}

type Config struct {
	v *viper.Viper
}

// NewConfig initializes the viper configuration
func NewConfig() *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")

	zingConfigDir, err := getConfigDir()
	if err == nil {
		v.AddConfigPath(zingConfigDir)
	}
	v.AddConfigPath(".")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.SetDefault("server_addr", "localhost")
	v.SetDefault("server_port", 50051)
	v.SetDefault("token", "")

	v.SetDefault("redis.addr", "")
	v.SetDefault("redis.username", "")
	v.SetDefault("redis.password", "")
	v.SetDefault("redis.db", 0)
	v.SetDefault("redis.tls", false)

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			log.Printf("Error reading config file: %s\n", err)
		}
	}

	return &Config{v: v}
}

func (c *Config) ConfigFileUsed() string {
	return c.v.ConfigFileUsed()
}

// GetServerAddr returns the configured server address in the form of host:port
func (c *Config) GetServerAddr() string {
	addr := c.v.GetString("server_addr")
	port := c.v.GetInt("server_port")
	return fmt.Sprintf("%s:%d", addr, port)
}

// SetServerAddr stores the server address
func (c *Config) SetServerAddr(addr string) error {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}
	portNum, err := net.LookupPort("tcp", port)
	if err != nil {
		return err
	}
	c.v.Set("server_addr", host)
	c.v.Set("server_port", portNum)
	return c.v.WriteConfig()
}

// GetToken returns the stored authentication token
func (c *Config) GetToken() string {
	return c.v.GetString("token")
}

// SetToken stores the authentication token
func (c *Config) SetToken(token string) error {
	c.v.Set("token", token)
	return c.v.WriteConfig()
}

// Redis configuration helpers

// GetRedisAddr returns the Redis address in the form host:port
func (c *Config) GetRedisAddr() string {
	return c.v.GetString("redis.addr2")
}

// GetRedisUsername returns the Redis username, if any (for ACLs)
func (c *Config) GetRedisUsername() string {
	return c.v.GetString("redis.username")
}

// GetRedisPassword returns the Redis password
func (c *Config) GetRedisPassword() string {
	return c.v.GetString("redis.password")
}

// GetRedisDB returns the Redis database number
func (c *Config) GetRedisDB() int {
	return c.v.GetInt("redis.db")
}

// GetRedisTLS indicates whether TLS should be used for Redis connections
func (c *Config) GetRedisTLS() bool {
	return c.v.GetBool("redis.tls")
}
