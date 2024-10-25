package config

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
type Config struct {
	PostgresName     string `mapstructure:"POSTGRES_NAME"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
}

// LoadConfig reads configuration from environment variables.
func LoadConfig(path, file string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(file)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	if err := validateConfig(config); err != nil {
		return Config{}, err
	}

	return config, nil
}

// validateConfig checks if all required fields in the Config struct are set.
func validateConfig(config Config) error {
	val := reflect.ValueOf(config)
	typ := reflect.TypeOf(config)

	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).String() == "" {
			return fmt.Errorf("missing required configuration: %s", typ.Field(i).Name)
		}
	}
	return nil
}
