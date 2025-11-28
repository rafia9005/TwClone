package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"TwClone/internal/config"
	imw "TwClone/internal/middleware"
	"TwClone/internal/pkg/logger"
	"TwClone/internal/pkg/utils/jwtutils"
	"TwClone/internal/pkg/utils/validationutils"
	"TwClone/internal/provider"

	"github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"
	emw "github.com/labstack/echo/v4/middleware"
)

type HttpServer struct {
	cfg    *config.Config
	server *http.Server
}

type EchoValidator struct {
	validator *validator.Validate
}

func (e *EchoValidator) Validate(i interface{}) error {
	return e.validator.Struct(i)
}

func NewHttpServer(cfg *config.Config) *HttpServer {
	router := echo.New()

	// register validator
	v := validator.New()
	v.RegisterTagNameFunc(validationutils.TagNameFormatter)
	v.RegisterCustomTypeFunc(validationutils.DecimalType)
	router.Validator = &EchoValidator{validator: v}

	RegisterMiddleware(router, cfg)

	provider.BootstrapHttp(cfg, router)

	return &HttpServer{
		cfg: cfg,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port),
			Handler: router,
		},
	}
}

func (s *HttpServer) Start() {
	logger.Log.Info("Running HTTP server on port:", s.cfg.HttpServer.Port)
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Log.Fatal("Error while HTTP server listening:", err)
	}
	logger.Log.Info("HTTP server is not receiving new requests...")
}

func (s *HttpServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.cfg.HttpServer.GracePeriod)*time.Second)
	defer cancel()

	logger.Log.Info("Attempting to shut down the HTTP server...")
	if err := s.server.Shutdown(ctx); err != nil {
		logger.Log.Fatal("Error shutting down HTTP server:", err)
	}
	logger.Log.Info("HTTP server shut down gracefully")
}

func RegisterMiddleware(router *echo.Echo, cfg *config.Config) {
	// use echo middleware plus internal middleware adapted to echo
	router.Use(
		emw.Gzip(),
		emw.CORSWithConfig(emw.CORSConfig{
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			// When AllowCredentials is true, Access-Control-Allow-Origin must be a specific origin
			// (not '*'). Use AllowOriginFunc to permit localhost origins (any port) during dev.
			// Allow all origins (useful in development). When AllowCredentials is true,
			// Echo will echo the request Origin back in Access-Control-Allow-Origin when
			// this function returns true.
			AllowOriginFunc: func(origin string) (bool, error) {
				return true, nil
			},
			AllowCredentials: true,
		}),
		emw.Recover(),
	)

	// internal middleware functions should return echo.MiddlewareFunc
	// initialize default jwt util for middleware that needs it
	if cfg != nil && cfg.Jwt != nil {
		// Log non-sensitive JWT config to help debugging token issues (issuer and allowed algs)
		logger.Log.Infof("jwt config loaded: issuer=%s, allowed_algs=%v", cfg.Jwt.Issuer, cfg.Jwt.AllowedAlgs)
		imw.SetDefaultJwtUtil(jwtutils.NewJwtUtil(cfg.Jwt))
	}
	router.Use(imw.Logger())
	router.Use(imw.ErrorHandler())
	router.Use(imw.RequestTimeout(cfg))
}
