package ginutils

import (
	"net/http"

	"TwClone/internal/constant"
	"TwClone/internal/dto"
	echo "github.com/labstack/echo/v4"
)

func ResponseOK[T any](ctx echo.Context, data T) error {
	return ResponseJSON(ctx, http.StatusOK, constant.ResponseSuccessMessage, data, nil)
}

func ResponseOKPlain(ctx echo.Context) error {
	return ResponseOK[any](ctx, nil)
}

func ResponseOKPagination[T any](ctx echo.Context, data T, paging *dto.PageMetaData) error {
	return ResponseJSON(ctx, http.StatusOK, constant.ResponseSuccessMessage, data, paging)
}

func ResponseCreated[T any](ctx echo.Context, data T) error {
	return ResponseJSON(ctx, http.StatusCreated, constant.ResponseSuccessMessage, data, nil)
}

func ResponseCreatedPlain(ctx echo.Context) error {
	return ResponseCreated[any](ctx, nil)
}

func ResponseJSON[T any](ctx echo.Context, statusCode int, message string, data T, paging *dto.PageMetaData) error {
	return ctx.JSON(statusCode, dto.WebResponse[T]{
		Message: message,
		Data:    data,
		Paging:  paging,
	})
}
