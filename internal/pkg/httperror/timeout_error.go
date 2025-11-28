package httperror

import (
	"errors"
	"net/http"

	"TwClone/internal/pkg/constant"
)

func NewTimeoutError() *ResponseError {
	msg := constant.RequestTimeoutErrorMessage

	err := errors.New(msg)

	return NewResponseError(err, http.StatusRequestTimeout, msg)
}
