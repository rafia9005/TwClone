package middleware

import (
	"strings"

	"TWclone/internal/constant"
	"TWclone/internal/pkg/httperror"
	"TWclone/internal/pkg/logger"
	"TWclone/internal/pkg/utils/jwtutils"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtUtil jwtutils.JwtUtil
}

func NewAuthMiddleware(jwtUtil jwtutils.JwtUtil) *AuthMiddleware {
	return &AuthMiddleware{
		jwtUtil: jwtUtil,
	}
}

func (m *AuthMiddleware) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := m.parseAccessToken(ctx)
		if err != nil {
			// attach a helpful error to the context for logging/response
			// parseAccessToken will return an UnauthorizedError when token not found
			// but we also log diagnostic flags inside parseAccessToken
			ctx.Error(err)
			ctx.Abort()
			return
		}

		claims, err := m.jwtUtil.Parse(accessToken)
		if err != nil {
			// Log parse error (do NOT include token value). This helps understand
			// if the token is malformed, expired, or signed with a different secret.
			logger.Log.WithField("source", "auth_middleware").Errorf("jwt parse error: %v", err)
			ctx.Error(httperror.NewUnauthorizedError())
			ctx.Abort()
			return
		}

		ctx.Set(constant.CTX_USER_ID, claims.UserID)
		ctx.Next()
	}
}

func (m *AuthMiddleware) parseAccessToken(ctx *gin.Context) (string, error) {
	// Try Authorization header first
	accessToken := ctx.GetHeader("Authorization")
	headerPresent := false
	if accessToken != "" {
		headerPresent = true
		// header present
		splitToken := strings.Split(accessToken, " ")
		if len(splitToken) == 2 && splitToken[0] == "Bearer" {
			// do NOT log token value
			logger.Log.Info("auth: Authorization header present (Bearer token)")
			return splitToken[1], nil
		}
		// bad header format
		logger.Log.Info("auth: Authorization header malformed")
		return "", httperror.NewUnauthorizedError()
	}

	// Fallback: try cookie named accessToken
	cookiePresent := false
	if cookie, err := ctx.Cookie("accessToken"); err == nil && cookie != "" {
		cookiePresent = true
		// cookie present (do not log cookie value)
		logger.Log.Infof("auth: cookie present=%v headerPresent=%v", cookiePresent, headerPresent)
		return cookie, nil
	}

	// If neither header nor cookie present, log flags for diagnostics
	logger.Log.Infof("auth: token not found - headerPresent=%v cookiePresent=%v", headerPresent, cookiePresent)

	// No token found
	return "", httperror.NewUnauthorizedError()
}
