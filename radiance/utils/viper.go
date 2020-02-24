package utils

import (
	"errors"

	"github.com/spf13/viper"
)

//SetConfigViper --
func SetConfigViper(config, path string) error {
	viper.SetConfigName(config)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		return errors.New("The system cannot find the path specified")
	}
	viper.SetEnvPrefix("DMP")
	viper.AutomaticEnv()
	return err
}
