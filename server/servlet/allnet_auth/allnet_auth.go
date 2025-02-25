package allnet_auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddAllnetRoutes(rg *gin.RouterGroup) {

	// /sys/servlet/Alive - ALL.Net Servlet Alive Route
	rg.Group("/servlet").GET("/Alive", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// /sys/test - Aqua Test Route
	rg.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Server running")
	})

}
