package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds the configuration for the application
type Config struct {
	Port string `mapstructure:"PORT"`
}

// LoadConfig reads the configuration from a file and environment variables
func LoadConfig() (Config, error) {

	viper.SetConfigName("service") // name of config file (without extension)
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")       // optionally look for config in the working directory
	viper.AutomaticEnv()           // read environment variables that match

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("error reading config file: %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("unable to decode into struct: %s", err)
	}

	return config, nil
}
