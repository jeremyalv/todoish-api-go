package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Username     string `mapstructure:"DB_USERNAME"`
		Password     string `mapstructure:"DB_PASSWORD"`
		RootPassword string `mapstructure:"DB_ROOT_PASSWORD"`
		DatabaseName string `mapstructure:"DB_NAME"`
		Host         string `mapstructure:"DB_HOST"`
		Port         string `mapstructure:"DB_PORT"`
		URL          string
		MaxIdleConnections int    `mapstructure:"MAX_IDLE_CONNECTIONS"`
		MaxOpenConnections int    `mapstructure:"MAX_OPEN_CONNECTIONS"`
	}
)

func LoadConfig() *Config {
	cfg := &Config{}
	
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("failed to read config from file:", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Panic("failed to load environment variable:", err)
	}

	// Set MySQL Connection URL
	cfg.URL = fmt.Sprintf("%s:%s@tcp(%s)/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.DatabaseName,
	)

	return cfg
}