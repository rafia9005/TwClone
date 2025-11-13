package controller

import (
	"net/http"

	"TWclone/internal/dto"

	"github.com/gin-gonic/gin"
)

type AppController struct {
}

func NewAppController() *AppController {
	return &AppController{}
}

func (c *AppController) Route(r *gin.Engine) {
	r.GET("/", c.Index)
	r.NoRoute(c.RouteNotFound)
	r.NoMethod(c.MethodNotAllowed)
	r.GET("/health", c.Health)
}

func (c *AppController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, dto.WebResponse[any]{
		Message: "status UP",
	})
}

func (c *AppController) RouteNotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, dto.WebResponse[any]{
		Message: "route not found",
	})
}

func (c *AppController) MethodNotAllowed(ctx *gin.Context) {
	ctx.JSON(http.StatusMethodNotAllowed, dto.WebResponse[any]{
		Message: "method not allowed",
	})
}

func (c *AppController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, dto.WebResponse[any]{
		Message: "TWClone API's is ready to use!",
	})
}
