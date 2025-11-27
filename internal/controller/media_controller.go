package controller

import (
	"TWclone/internal/dto"
	"TWclone/internal/entity"
	"TWclone/internal/repository"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MediaController struct {
	repo repository.MediaRepositoryImpl
}

func NewMediaController() *MediaController {
	return &MediaController{repo: repository.MediaRepositoryImpl{}}
}

func (c *MediaController) Route(r gin.IRouter) {
	g := r.Group("/media")
	g.POST("", c.Create)
	g.GET("/tweet/:tweet_id", c.ByTweet)
	g.GET("/:id", c.ByID)
}

func (c *MediaController) Create(ctx *gin.Context) {
	var media entity.Media
	if err := ctx.ShouldBindJSON(&media); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "Media")})
		return
	}
	if err := c.repo.Create(context.Background(), &media); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, media)
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

func (c *MediaController) ByTweet(ctx *gin.Context) {
	tweetID, _ := strconv.ParseInt(ctx.Param("tweet_id"), 10, 64)
	media, err := c.repo.FindByTweetID(context.Background(), tweetID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, media)
}

func (c *MediaController) ByID(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	media, err := c.repo.FindByID(context.Background(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	ctx.JSON(http.StatusOK, media)
}
