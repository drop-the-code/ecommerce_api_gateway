package config

import "github.com/spf13/viper"

func Environment() *viper.Viper {
	v1 := viper.New()
	v1.SetConfigFile(".env")
	v1.ReadInConfig()
	return v1
}
