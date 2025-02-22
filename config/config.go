package config

import (
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigFile("config.json")
	viper.SetDefault("log", logConfig)
	viper.SetDefault("db", dbConfig)
	conf := &config{}
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Can't read config, trying to modify!")
		if err := viper.WriteConfig(); err != nil {
			log.Fatal("Error writing config: ", err)
		}
	}
	if err := viper.Unmarshal(conf); err != nil {
		log.Fatal(err)
	}
}

func Get(key string) any {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file, ", err)
	}
	return viper.Get(key)
}

func Set(key string, value any) {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file, ", err)
	}
	viper.Set(key, value)
}
