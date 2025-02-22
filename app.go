package main

import (
	"OMG_ITS_ALLNET_SERVER/config"
	"OMG_ITS_ALLNET_SERVER/logger"
	"OMG_ITS_ALLNET_SERVER/server/aimedb"
	"OMG_ITS_ALLNET_SERVER/server/title"
)

func main() {
	config.EnvInit()
	logger.Init()
	aimedb.NewServer()
	title.Run()
}
