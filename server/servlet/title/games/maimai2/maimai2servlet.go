package maimai2

import (
	"github.com/gin-gonic/gin"
)

var endpoints = []string{
	"GetGameRankingApi",
	"GetUserCharacterApi",
	"GetUserItemApi",
	"GetUserPortraitApi",
	"GetUserRatingApi",
	"UploadUserPhotoApi",
	"UploadUserPlaylogApi",
	"UploadUserPortraitApi",
	"UpsertUserAllApi",
	"CMGetUserCardApi",
	"CMGetUserCardPrintErrorApi",
	"CMGetUserDataApi",
	"CMGetUserItemApi",
	"CMUpsertUserPrintApi",
	"GetUserFavoriteItemApi",
	"GetServerAnnouncementApi",
}

func AddMaimai2Routes(rg *gin.RouterGroup) {
	maimai2 := rg.Group("/Maimai2Servlet")
	for _, endpoint := range endpoints {
		maimai2.GET("/"+endpoint, func(c *gin.Context) {
			c.String(200, "Maimai2 "+endpoint)
		})
	}
}
