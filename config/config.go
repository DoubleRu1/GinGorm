package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Myapp struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
}

var Appconfig *Config

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("ReadInConfig Error %v\n", err)
	}
	Appconfig = &Config{}
	if err := viper.Unmarshal(Appconfig); err != nil {
		log.Fatalf("UnmarshalConfig Error %v\n", err)
	}
	initDB()
}
