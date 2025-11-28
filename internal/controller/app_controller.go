package controller

import (
	"net/http"

	"TwClone/internal/dto"

	"github.com/go-playground/validator/v10"
	echo "github.com/labstack/echo/v4"
)

type AppController struct {
}

func extractFieldErrors(err error, structPrefix string) []dto.FieldError {
	var fieldErrors []dto.FieldError
	if verrs, ok := err.(validator.ValidationErrors); ok {
		for _, verr := range verrs {
			field := verr.Field()
			// Remove struct prefix if present
			if structPrefix != "" && len(field) > len(structPrefix) && field[:len(structPrefix)] == structPrefix {
				field = field[len(structPrefix):]
			}
			fieldErrors = append(fieldErrors, dto.FieldError{
				Field:   field,
				Message: verr.Error(),
			})
		}
	} else {
		// fallback: try to parse error string
		msg := err.Error()
		field := "body"
		fieldErrors = append(fieldErrors, dto.FieldError{Field: field, Message: msg})
	}
	return fieldErrors
}

func NewAppController() *AppController {
	return &AppController{}
}

func (c *AppController) Route(r *echo.Echo) {
	r.GET("/health", c.Health)
}

func (c *AppController) Health(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, dto.WebResponse[any]{
		Message: "status UP",
	})
}

func (c *AppController) RouteNotFound(ctx echo.Context) error {
	return ctx.JSON(http.StatusNotFound, dto.WebResponse[any]{
		Message: "route not found",
	})
}

func (c *AppController) MethodNotAllowed(ctx echo.Context) error {
	return ctx.JSON(http.StatusMethodNotAllowed, dto.WebResponse[any]{
		Message: "method not allowed",
	})
}
