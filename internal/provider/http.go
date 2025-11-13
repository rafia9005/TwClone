package provider

import (
	"TWclone/internal/config"
	"TWclone/internal/controller"

	"github.com/gin-gonic/gin"
)

func BootstrapHttp(cfg *config.Config, router *gin.Engine) {

	appController := controller.NewAppController()
	appController.Route(router)
}
