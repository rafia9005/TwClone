package provider

import (
	"fmt"

	"TWclone/internal/config"
	"TWclone/internal/controller"

	// register generated swagger docs
	_ "TWclone/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// Register all entity controllers
	followCtrl := controller.NewFollowController()
	followCtrl.Route(base)

	likeCtrl := controller.NewLikeController()
	likeCtrl.Route(base)

	hashtagCtrl := controller.NewHashtagController()
	hashtagCtrl.Route(base)

	tweetHashtagCtrl := controller.NewTweetHashtagController()
	tweetHashtagCtrl.Route(base)

	mentionCtrl := controller.NewMentionController()
	mentionCtrl.Route(base)

	mediaCtrl := controller.NewMediaController()
	mediaCtrl.Route(base)

	notifCtrl := controller.NewNotificationController()
	notifCtrl.Route(base)

	router.NoRoute(appCtrl.RouteNotFound)
	router.NoMethod(appCtrl.MethodNotAllowed)

	// Swagger UI (if docs are generated with swag). Run `swag init` in repo root
	// to generate the docs under ./docs. The swagger UI will be available at
	// /swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
