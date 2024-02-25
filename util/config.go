package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Name string `mapstructure:"Name"`
}

func ConfigurationFile() (config Config, err error) {

	viper.AddConfigPath(".")
	viper.SetConfigName("app") // name of config file (without extension)
	viper.SetConfigType("env") // REQUIRED if the config file does not have the extension in the name

	/*AutomaticEnv makes Viper check if environment variables match any of the existing
	keys (config, default or flags). If matching env vars are found, they are loaded into Viper.*/
	viper.AutomaticEnv()

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&config)
	return
}
