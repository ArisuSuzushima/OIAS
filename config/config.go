package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigFile("config.json")
	viper.SetDefault("db.host", dbConfig.Host)
	viper.SetDefault("db.port", dbConfig.Port)
	viper.SetDefault("db.user", dbConfig.User)
	viper.SetDefault("db.password", dbConfig.Password)
	viper.SetDefault("db.database", dbConfig.Database)
	conf := &config{}
	if err := viper.ReadInConfig(); err != nil {
		logrus.Warning("Can't read config, trying to modify!")
		if err := viper.WriteConfig(); err != nil {
			logrus.Fatal("Error writing config: ", err)
		}
	}
	if err := viper.Unmarshal(conf); err != nil {
		log.Fatal(err)
	}
}

func Get(key string) any {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("Error reading config file, ", err)
	}
	return viper.Get(key)
}

func Set(key string, value any) {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("Error reading config file, ", err)
	}
	viper.Set(key, value)
}
