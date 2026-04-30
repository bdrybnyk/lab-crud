package main

import "github.com/spf13/viper"

type Config struct {
	DBUrl      string `mapstructure:"DB_URL"`
	ServerPort string `mapstructure:"SERVER_PORT"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()
	err := viper.Unmarshal(&cfg)
	return cfg, err
}
