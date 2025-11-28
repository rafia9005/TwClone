package provider

import (
	"net/http"

	"TwClone/internal/config"
	"TwClone/internal/controller"
	"TwClone/internal/pkg/utils/jwtutils"

	echo "github.com/labstack/echo/v4"
)

func BootstrapHttp(cfg *config.Config, router *echo.Echo) {
	// App-level routes
	appController := controller.NewAppController()
	appController.Route(router)

	// API v1 group for all controllers
	api := router.Group("/api/v1")

	// Register controllers (in-place constructors)
	controller.NewAuthController(cfg).Route(api)
	controller.NewUserController().Route(api)
	controller.NewLikeController().Route(api)
	controller.NewFollowController().Route(api)
	controller.NewHashtagController().Route(api)
	controller.NewMediaController().Route(api)
	controller.NewMentionController().Route(api)
	controller.NewNotificationController().Route(api)
	controller.NewTweetHashtagController().Route(api)

	// Ensure OPTIONS preflight requests are handled even if a specific route isn't matched.
	// Echo's CORS middleware will attach the necessary CORS headers; this handler returns 200 for OPTIONS.
	router.OPTIONS("/*", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	// Temporary debug endpoint to parse JWT from Authorization header using server config.
	// Helps debugging mismatches between token issuer/alg/secret.
	router.GET("/internal/debug/jwt", func(c echo.Context) error {
		if cfg == nil || cfg.Jwt == nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "jwt config not available on server"})
		}

		auth := c.Request().Header.Get("Authorization")
		if auth == "" {
			// try cookie fallback
			if ck, err := c.Cookie("accessToken"); err == nil && ck != nil {
				auth = "Bearer " + ck.Value
			}
		}
		if auth == "" {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "authorization header or accessToken cookie not provided"})
		}

		// naive parse
		var token string
		if len(auth) > 7 && auth[:7] == "Bearer " {
			token = auth[7:]
		} else {
			token = auth
		}

		ju := jwtutils.NewJwtUtil(cfg.Jwt)
		claims, err := ju.Parse(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "failed to parse token", "detail": err.Error()})
		}

		return c.JSON(http.StatusOK, echo.Map{"claims": claims})
	})
}
