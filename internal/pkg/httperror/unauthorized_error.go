package httperror

import (
	"errors"
	"net/http"

	"TWclone/internal/pkg/constant"
)

func NewUnauthorizedError() *ResponseError {
	msg := constant.UnauthorizedErrorMessage

	err := errors.New(msg)

	return NewResponseError(err, http.StatusUnauthorized, msg)
}
