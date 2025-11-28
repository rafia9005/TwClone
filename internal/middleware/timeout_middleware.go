package middleware

import (
	"context"
	"time"

	"TwClone/internal/config"
	"TwClone/internal/pkg/httperror"
	echo "github.com/labstack/echo/v4"
)

func RequestTimeout(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			timeoutCtx, cancel := context.WithTimeout(
				ctx.Request().Context(),
				time.Duration(cfg.HttpServer.RequestTimeoutPeriod)*time.Second,
			)
			defer cancel()

			// attach new request context
			req := ctx.Request()
			req = req.WithContext(timeoutCtx)
			ctx.SetRequest(req)

			done := make(chan error, 1)

			go func() {
				defer func() {
					if r := recover(); r != nil {
						if err, ok := r.(error); ok {
							done <- err
							return
						}
						done <- nil
						return
					}
				}()
				done <- next(ctx)
			}()

			select {
			case <-timeoutCtx.Done():
				return httperror.NewTimeoutError()
			case err := <-done:
				return err
			}
		}
	}
}
