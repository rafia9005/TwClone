package middleware

import (
	"strings"

	"TwClone/internal/constant"
	"TwClone/internal/pkg/httperror"
	"TwClone/internal/pkg/logger"
	"TwClone/internal/pkg/utils/jwtutils"

	echo "github.com/labstack/echo/v4"
)

type AuthMiddlewareImpl struct {
	jwtUtil jwtutils.JwtUtil
}

func NewAuthMiddleware(jwtUtil jwtutils.JwtUtil) *AuthMiddlewareImpl {
	return &AuthMiddlewareImpl{
		jwtUtil: jwtUtil,
	}
}

// package-level default jwt util used by the convenience AuthMiddleware() function.
var defaultJwt jwtutils.JwtUtil

// SetDefaultJwtUtil sets the default JwtUtil used by AuthMiddleware().
func SetDefaultJwtUtil(j jwtutils.JwtUtil) {
	defaultJwt = j
}

// AuthMiddleware is a convenience function returning an echo middleware using the
// package-level default JwtUtil. Call SetDefaultJwtUtil(...) during application
// initialization (for example in server.RegisterMiddleware) to provide a JwtUtil.
func AuthMiddleware() echo.MiddlewareFunc {
	if defaultJwt == nil {
		// If not initialized, return a middleware that rejects requests.
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(ctx echo.Context) error {
				return httperror.NewUnauthorizedError()
			}
		}
	}
	return NewAuthMiddleware(defaultJwt).Authorization()
}

func (m *AuthMiddlewareImpl) Authorization() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			accessToken, err := m.parseAccessToken(ctx)
			if err != nil {
				logger.Log.Debugf("auth: parseAccessToken error: %v", err)
				return err
			}

			claims, err := m.jwtUtil.Parse(accessToken)
			if err != nil {
				// log parse error for debugging without exposing full token
				tokenPreview := accessToken
				if len(tokenPreview) > 8 {
					tokenPreview = tokenPreview[:8]
				}
				logger.Log.Debugf("auth: jwt parse failed (token preview=%s): %v", tokenPreview, err)
				return httperror.NewUnauthorizedError()
			}

			ctx.Set(constant.CTX_USER_ID, claims.UserID)
			return next(ctx)
		}
	}
}

func (m *AuthMiddlewareImpl) parseAccessToken(ctx echo.Context) (string, error) {
	accessToken := ctx.Request().Header.Get("Authorization")
	if accessToken == "" || len(accessToken) == 0 {
		// Fallback: try cookie named "accessToken" (frontend may send token in cookie)
		if ck, err := ctx.Cookie("accessToken"); err == nil && ck != nil && ck.Value != "" {
			return ck.Value, nil
		}
		return "", httperror.NewUnauthorizedError()
	}

	splitToken := strings.Split(accessToken, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return "", httperror.NewUnauthorizedError()
	}
	return splitToken[1], nil
}
