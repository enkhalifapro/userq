package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Load viper config. `name` argument used to construct config file name `paths` to search them.
func Load(name string, paths ...string) error {
	// set config file name & extension.
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")

	// search config in specified paths.
	for _, p := range append(paths, "./etc/") {
		if p != "" {
			viper.AddConfigPath(p)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read config: %v", err)
	}
	return nil
}
