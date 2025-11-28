package controller

import (
	"context"
	"net/http"

	"TwClone/internal/dto"
	"TwClone/internal/entity"
	"TwClone/internal/repository"

	"github.com/labstack/echo/v4"
)

type HashtagController struct {
	repo repository.HashtagRepositoryImpl
}

func NewHashtagController() *HashtagController {
	return &HashtagController{repo: repository.HashtagRepositoryImpl{}}
}

func (c *HashtagController) Route(g *echo.Group) {
	hg := g.Group("/hashtags")
	hg.POST("", c.Create)
	hg.GET("", c.FindAll)
	hg.GET("/:tag", c.FindByTag)
}

func (c *HashtagController) Create(ctx echo.Context) error {
	var hashtag entity.Hashtag
	if err := ctx.Bind(&hashtag); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "Hashtag")})
	}
	if err := c.repo.Create(context.Background(), &hashtag); err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, hashtag)
}

// CreateHashtag godoc
// @Summary Create hashtag
// @Description Create a new hashtag
// @Tags hashtags
// @Accept json
// @Produce json
// @Param hashtag body entity.Hashtag true "Hashtag payload"
// @Success 201 {object} entity.Hashtag
// @Failure 400 {object} dto.WebResponse
// @Router /api/v1/hashtags [post]

// ListHashtags godoc
// @Summary List hashtags
// @Description Get all hashtags
// @Tags hashtags
// @Accept json
// @Produce json
// @Success 200 {array} entity.Hashtag
// @Router /api/v1/hashtags [get]

// GetHashtagByTag godoc
// @Summary Get hashtag
// @Description Get hashtag by tag
// @Tags hashtags
// @Accept json
// @Produce json
// @Param tag path string true "Tag"
// @Success 200 {object} entity.Hashtag
// @Router /api/v1/hashtags/{tag} [get]

func (c *HashtagController) FindByTag(ctx echo.Context) error {
	tag := ctx.Param("tag")
	hashtag, err := c.repo.FindByTag(context.Background(), tag)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, dto.WebResponse[any]{Message: "not found"})
	}
	return ctx.JSON(http.StatusOK, hashtag)
}

func (c *HashtagController) FindAll(ctx echo.Context) error {
	hashtags, err := c.repo.FindAll(context.Background())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, hashtags)
}
