package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper
}

func Load(configPaths []string) (*Config, error) {
	m := &Config{v: viper.New()}
	return m, nil
}
