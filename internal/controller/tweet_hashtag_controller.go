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

type TweetHashtagController struct {
	repo repository.TweetHashtagRepositoryImpl
}

func NewTweetHashtagController() *TweetHashtagController {
	return &TweetHashtagController{repo: repository.TweetHashtagRepositoryImpl{}}
}

func (c *TweetHashtagController) Route(r gin.IRouter) {
	g := r.Group("/tweet-hashtags")
	g.POST("", c.Create)
	g.GET("/tweet/:tweet_id", c.ByTweet)
	g.GET("/hashtag/:hashtag_id", c.ByHashtag)
}

func (c *TweetHashtagController) Create(ctx *gin.Context) {
	var th entity.TweetHashtag
	if err := ctx.ShouldBindJSON(&th); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "TweetHashtag")})
		return
	}
	if err := c.repo.Create(context.Background(), &th); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, th)
}

func (c *TweetHashtagController) ByTweet(ctx *gin.Context) {
	tweetID, _ := strconv.ParseInt(ctx.Param("tweet_id"), 10, 64)
	ths, err := c.repo.FindByTweetID(context.Background(), tweetID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ths)
}

func (c *TweetHashtagController) ByHashtag(ctx *gin.Context) {
	hashtagID, _ := strconv.ParseInt(ctx.Param("hashtag_id"), 10, 64)
	ths, err := c.repo.FindByHashtagID(context.Background(), hashtagID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ths)
}
