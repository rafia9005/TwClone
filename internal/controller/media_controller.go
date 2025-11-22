package controller

import (
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
	g.POST("/", c.Create)
	g.GET("/tweet/:tweet_id", c.ByTweet)
	g.GET("/:id", c.ByID)
}

func (c *MediaController) Create(ctx *gin.Context) {
	var media entity.Media
	if err := ctx.ShouldBindJSON(&media); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.repo.Create(context.Background(), &media); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, media)
}

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
