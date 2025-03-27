package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type (
	Config struct {
		App  App  `mapstructure:"app" validate:"required"`
		Log  Log  `mapstructure:"log" validate:"required"`
		HTTP HTTP `mapstructure:"http" validate:"required"`
	}
	App struct {
		Name string `mapstructure:"name" validate:"required"`
		Env  string `mapstructure:"env" validate:"required"`
	}
	Log struct {
		File string `mapstructure:"file" validate:"required"`
	}
	HTTP struct {
		Port string `mapstructure:"port" validate:"required"`
	}
)

func New(fileName string) (*Config, error) {
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("#config_e1. Error read /configs yaml file: %w", err)
	}

	cfg := &Config{}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("#config_e2. Error unmarshal /configs yaml file: %w", err)
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		return nil, fmt.Errorf("#config_e3. Error validate /configs yaml file: %w", err)
	}

	return cfg, nil
}

func GetAppConfigFile() string {
	result, exists := os.LookupEnv("APP_CONFIG_FILE")
	if !exists || result == "" {
		result = "local"
	}
	return result
}
