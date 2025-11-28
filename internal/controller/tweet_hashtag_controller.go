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

type TweetHashtagController struct {
	repo repository.TweetHashtagRepositoryImpl
}

func NewTweetHashtagController() *TweetHashtagController {
	return &TweetHashtagController{repo: repository.TweetHashtagRepositoryImpl{}}
}

func (c *TweetHashtagController) Route(g *echo.Group) {
	thg := g.Group("/tweet-hashtags")
	thg.POST("", c.Create)
	thg.GET("/tweet/:tweet_id", c.ByTweet)
	thg.GET("/hashtag/:hashtag_id", c.ByHashtag)
}

func (c *TweetHashtagController) Create(ctx echo.Context) error {
	var th entity.TweetHashtag
	if err := ctx.Bind(&th); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "TweetHashtag")})
	}
	if err := c.repo.Create(context.Background(), &th); err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, th)
}

// CreateTweetHashtag godoc
// @Summary Create tweet-hashtag relation
// @Description Attach hashtag to a tweet
// @Tags tweet-hashtags
// @Accept json
// @Produce json
// @Param th body entity.TweetHashtag true "TweetHashtag payload"
// @Success 201 {object} entity.TweetHashtag
// @Failure 400 {object} dto.WebResponse
// @Router /api/v1/tweet-hashtags [post]

// GetTweetHashtagsByTweet godoc
// @Summary Hashtags by tweet
// @Description Get hashtags for a tweet
// @Tags tweet-hashtags
// @Accept json
// @Produce json
// @Param tweet_id path int true "Tweet ID"
// @Success 200 {array} entity.TweetHashtag
// @Router /api/v1/tweet-hashtags/tweet/{tweet_id} [get]

// GetTweetHashtagsByHashtag godoc
// @Summary Tweets by hashtag
// @Description Get tweet-hashtag entries by hashtag id
// @Tags tweet-hashtags
// @Accept json
// @Produce json
// @Param hashtag_id path int true "Hashtag ID"
// @Success 200 {array} entity.TweetHashtag
// @Router /api/v1/tweet-hashtags/hashtag/{hashtag_id} [get]

func (c *TweetHashtagController) ByTweet(ctx echo.Context) error {
	tweetID, _ := strconv.ParseInt(ctx.Param("tweet_id"), 10, 64)
	ths, err := c.repo.FindByTweetID(context.Background(), tweetID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, ths)
}

func (c *TweetHashtagController) ByHashtag(ctx echo.Context) error {
	hashtagID, _ := strconv.ParseInt(ctx.Param("hashtag_id"), 10, 64)
	ths, err := c.repo.FindByHashtagID(context.Background(), hashtagID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, ths)
}
