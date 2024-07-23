package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetEnv(Key string) string {
	_ = godotenv.Load()
	return os.Getenv(Key)
}

func SetEnvConf(ConfKey string, ConfSub string) {
	var Config = []string{ConfKey, ConfSub}
	var EnvKey = strings.Join(Config, "_")
	env := GetEnv(EnvKey)
	var orig = Get(ConfKey + "." + ConfSub)
	if env != "" {
		switch orig.(type) {
		case string:
			Set(ConfKey, env)
		case int:
			conv, err := strconv.Atoi(env)
			if err != nil {
				log.Panic("[Config] Error converting string to int: ", err)
			}
			Set(ConfKey, conv)
		case bool:
			switch env {
			case "true":
				Set(ConfKey, true)
			case "false":
				Set(ConfKey, false)
			}
		case []string:
			trim := strings.TrimSpace(env)
			trim = strings.TrimPrefix(trim, "[")
			trim = strings.TrimSuffix(trim, "]")
			var envArray []string
			for _, v := range strings.Split(trim, ",") {
				trimSub := strings.TrimPrefix(v, "\"")
				trimSub = strings.TrimSuffix(trimSub, "\"")
				envArray = append(envArray, trimSub)
			}
			Set(ConfKey, envArray)
		}
	}
}

func EnvInit() {
	var dict = map[string][]string{
		"db": {"host", "port", "user", "password", "database"},
	}
	for key, value := range dict {
		for _, sub := range value {
			SetEnvConf(key, sub)
		}
	}
}

func TZ() string {
	tz := GetEnv("TZ")
	if tz == "" {
		return "Asia/Shanghai"
	}
	return tz
}
