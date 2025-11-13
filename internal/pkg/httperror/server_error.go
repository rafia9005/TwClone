package httperror

import (
	"errors"
	"net/http"

	"TWclone/internal/pkg/constant"
)

func NewServerError() *ResponseError {
	msg := constant.InternalServerErrorMessage

	err := errors.New(msg)

	return NewResponseError(err, http.StatusInternalServerError, msg)
}
