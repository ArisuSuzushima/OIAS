package logger

import (
	"OMG_ITS_ALLNET_SERVER/config"
	log "github.com/sirupsen/logrus"
	"os"
)

func Init() {
	FileLogger := config.Get("log.file").(bool)
	DebugMode := config.Get("log.debug").(bool)

	if FileLogger {
		file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Panicln("Error opening log file: ", err)
		}
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(file)
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		log.SetLevel(log.WarnLevel)
	}

	if DebugMode {
		log.SetLevel(log.DebugLevel)
	}
}
