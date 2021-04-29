package config

import (
	"shortURL/pkg/logger"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var _config *Config

// Config ...
type Config struct {
	Logger   LoggerConfig   `mapstructure:"logger"`
	HTTP     HTTPConfig     `mapstructure:"http"`
	Redis    RedisConfig    `mapstructure:"redis"`
	MySQL    MySQLConfig    `mapstructure:"mysql"`
	ShortURL ShortURLConfig `mapstructure:"shorturl"`
}

type LoggerConfig struct {
	Level  string           `mapstructure:"level"`
	Format logger.LogFormat `mapstructure:"format"`
}

type HTTPConfig struct {
	Port uint16 `mapstructure:"port"`
}

type MySQLConfig struct {
	DSN string `mapstructure:"dsn"`
}

type ShortURLConfig struct {
	BasePath   string        `mapstructure:"base_path"`
	ExpireTime time.Duration `mapstructure:"expire_time"`
}

type RedisConfig struct {
	Host         string        `mapstructure:"host"`
	Port         uint          `mapstructure:"port"`
	DB           int           `mapstructure:"db"`
	Password     string        `mapstructure:"password"`
	MinIdleConns int           `mapstructure:"min_idle_conns"`
	MaxPoolSize  int           `mapstructure:"max_pool_size"`
	DialTimeout  time.Duration `mapstructure:"dial_timeout"`
}

// NewInjection ...
func (c *Config) NewInjection() *Config {
	return c
}

// Get get global config
func Get() *Config {
	return _config
}

// Set global config
func Set(config *Config) {
	_config = config
}

func New() (*Config, error) {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	configPath := viper.GetString("CONFIG_PATH")
	if configPath == "" {
		configPath = viper.GetString("PROJ_DIR") + "/deployment/config"
	}

	configName := viper.GetString("CONFIG_NAME")
	if configName == "" {
		configName = "default"
	}

	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return &config, err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return &config, err
	}

	_config = &config
	return &config, nil
}
