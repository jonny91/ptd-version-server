package controlls

import (
	"PTDVersionServer/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleVersion(context *gin.Context) {
	platform, _ := context.GetQuery("platform")
	appVersion, _ := context.GetQuery("version")

	v, ok := config.Versions.Data[platform]
	if !ok {
		context.String(http.StatusBadRequest, "")
	} else {
		updateConfig := v.UpdateConfig[appVersion]

		context.JSON(http.StatusOK, gin.H{
			"platform": v.Platform,
			"v":        v.V,
			"url":      v.Url,
			"force":    updateConfig,
		})
	}
}

func HandleReload(ctx *gin.Context) {
	config.ReadAll()
	ctx.String(http.StatusOK, "success!")
}
