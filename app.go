package main

import (
	"OMG_ITS_ALLNET_SERVER/config"
	"OMG_ITS_ALLNET_SERVER/logger"
	"OMG_ITS_ALLNET_SERVER/server/aimedb"
	"OMG_ITS_ALLNET_SERVER/server/servlet"
)

func main() {
	config.EnvInit()
	logger.Init()
	aimedb.NewServer()
	servlet.RunServer()
}
