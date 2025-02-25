package title

import (
	"OMG_ITS_ALLNET_SERVER/server/servlet/title/games/maimai2"
	"github.com/gin-gonic/gin"
)

func AddTitleRoutes(rg *gin.RouterGroup) {
	maimai2.AddMaimai2Routes(rg)
}
