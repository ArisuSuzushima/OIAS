package servlet

import (
	"OMG_ITS_ALLNET_SERVER/server/servlet/allnet_auth"
	"OMG_ITS_ALLNET_SERVER/server/servlet/title"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func RunServer() {
	getRoutes()
	router.Run(":80")
}

func getRoutes() {
	// ALL.NET General system routes
	allnet := router.Group("/sys")
	allnet_auth.AddAllnetRoutes(allnet)

	// Game Title Server
	game := router.Group("/")
	title.AddTitleRoutes(game)
}
