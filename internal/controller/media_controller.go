package controller

import (
	"TwClone/internal/dto"
	"TwClone/internal/entity"
	"TwClone/internal/repository"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MediaController struct {
	repo repository.MediaRepositoryImpl
}

func NewMediaController() *MediaController {
	return &MediaController{repo: repository.MediaRepositoryImpl{}}
}

func (c *MediaController) Route(g *echo.Group) {
	mg := g.Group("/media")
	mg.POST("", c.Create)
	mg.GET("/tweet/:tweet_id", c.ByTweet)
	mg.GET("/:id", c.ByID)
}

func (c *MediaController) Create(ctx echo.Context) error {
	var media entity.Media
	if err := ctx.Bind(&media); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "Media")})
	}
	if err := c.repo.Create(context.Background(), &media); err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, media)
}

// CreateMedia godoc
// @Summary Create media
// @Description Upload or register media for a tweet
// @Tags media
// @Accept json
// @Produce json
// @Param media body entity.Media true "Media payload"
// @Success 201 {object} entity.Media
// @Failure 400 {object} dto.WebResponse
// @Router /api/v1/media [post]

// GetMediaByTweet godoc
// @Summary Media by tweet
// @Description Get media for a tweet
// @Tags media
// @Accept json
// @Produce json
// @Param tweet_id path int true "Tweet ID"
// @Success 200 {array} entity.Media
// @Router /api/v1/media/tweet/{tweet_id} [get]

// GetMediaByID godoc
// @Summary Get media
// @Description Get media by id
// @Tags media
// @Accept json
// @Produce json
// @Param id path int true "Media ID"
// @Success 200 {object} entity.Media
// @Router /api/v1/media/{id} [get]

func (c *MediaController) ByTweet(ctx echo.Context) error {
	tweetID, _ := strconv.ParseInt(ctx.Param("tweet_id"), 10, 64)
	media, err := c.repo.FindByTweetID(context.Background(), tweetID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, media)
}

func (c *MediaController) ByID(ctx echo.Context) error {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	media, err := c.repo.FindByID(context.Background(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, dto.WebResponse[any]{Message: "not found"})
	}
	return ctx.JSON(http.StatusOK, media)
}
