package middleware

import (
	"net/http"
	"time"

	"TwClone/internal/pkg/httperror"
	"TwClone/internal/pkg/logger"
	"github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"
)

func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			start := time.Now()
			path := ctx.Request().URL.Path

			err := next(ctx)

			params := map[string]any{
				"status_code": ctx.Response().Status,
				"client_ip":   ctx.RealIP(),
				"method":      ctx.Request().Method,
				"latency":     time.Since(start).String(),
				"path":        path,
			}

			if err == nil {
				logger.Log.WithFields(params).Info("incoming request")
				return nil
			}

			logErrors(ctx, params, err)
			return err
		}
	}
}

func logErrors(ctx echo.Context, params map[string]any, err error) {
	errorsList := []error{}
	switch e := err.(type) {
	case validator.ValidationErrors:
		params["status_code"] = http.StatusBadRequest
		errorsList = append(errorsList, err)
	case *httperror.ResponseError:
		params["status_code"] = e.GetCode()
		errorsList = append(errorsList, e.OriginalError())
	default:
		params["status_code"] = http.StatusInternalServerError
		errorsList = append(errorsList, err)
	}

	params["errors"] = errorsList
	logger.Log.WithFields(params).Error("got error")
}
