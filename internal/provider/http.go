package provider

import (
	"fmt"

	"TWclone/internal/config"
	"TWclone/internal/controller"

	"github.com/gin-gonic/gin"
)

func BootstrapHttp(cfg *config.Config, router *gin.Engine) {

	version := "1"
	if cfg != nil && cfg.App != nil && cfg.App.Version != "" {
		version = cfg.App.Version
	}
	basePath := fmt.Sprintf("/api/v%s", version)

	base := router.Group(basePath)

	appCtrl := controller.NewAppController()
	appCtrl.Route(base)

	userCtrl := controller.NewUserController()
	userCtrl.Route(base)

	authCtrl := controller.NewAuthController(cfg)
	authCtrl.Route(base)

	router.NoRoute(appCtrl.RouteNotFound)
	router.NoMethod(appCtrl.MethodNotAllowed)
}
