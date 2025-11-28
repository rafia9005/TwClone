package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"TwClone/internal/constant"
	"TwClone/internal/dto"
	pkgconstant "TwClone/internal/pkg/constant"
	"TwClone/internal/pkg/httperror"
	"TwClone/internal/pkg/utils/validationutils"
	"github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"
)

func ErrorHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			err := next(ctx)
			if err == nil {
				return nil
			}

			switch e := err.(type) {
			case validator.ValidationErrors:
				return handleValidationError(ctx, e)
			case *json.SyntaxError:
				return handleJsonSyntaxError(ctx)
			case *json.UnmarshalTypeError:
				return handleJsonUnmarshalTypeError(ctx, e)
			case *time.ParseError:
				return handleParseTimeError(ctx, e)
			case *httperror.ResponseError:
				return ctx.JSON(e.GetCode(), dto.WebResponse[any]{
					Message: e.DisplayMessage(),
				})
			default:
				if errors.Is(e, io.EOF) {
					return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{
						Message: pkgconstant.EOFErrorMessage,
					})
				}
				return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{
					Message: pkgconstant.InternalServerErrorMessage,
				})
			}
		}
	}
}

func handleJsonSyntaxError(ctx echo.Context) error {
	return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{
		Message: pkgconstant.JsonSyntaxErrorMessage,
	})
}

func handleJsonUnmarshalTypeError(ctx echo.Context, err *json.UnmarshalTypeError) error {
	return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{
		Message: fmt.Sprintf(pkgconstant.JsonUnmarshallTypeErrorMessage, err.Field),
	})
}

func handleParseTimeError(ctx echo.Context, err *time.ParseError) error {
	return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{
		Message: fmt.Sprintf("please send time in format of %s, got: %s", constant.ConvertGoTimeLayoutToReadable(err.Layout), err.Value),
	})
}

func handleValidationError(ctx echo.Context, err validator.ValidationErrors) error {
	ve := []dto.FieldError{}

	for _, fe := range err {
		ve = append(ve, dto.FieldError{
			Field:   fe.Field(),
			Message: validationutils.TagToMsg(fe),
		})
	}

	return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{
		Message: pkgconstant.ValidationErrorMessage,
		Errors:  ve,
	})
}
