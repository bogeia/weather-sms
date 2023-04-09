package config

import (
	"github.com/spf13/viper"
)

// Load ...
func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var c *Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return c, nil
}
