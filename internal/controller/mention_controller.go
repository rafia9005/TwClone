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

type MentionController struct {
	repo repository.MentionRepositoryImpl
}

func NewMentionController() *MentionController {
	return &MentionController{repo: repository.MentionRepositoryImpl{}}
}

func (c *MentionController) Route(r gin.IRouter) {
	g := r.Group("/mentions")
	g.POST("", c.Create)
	g.GET("/tweet/:tweet_id", c.ByTweet)
	g.GET("/user/:user_id", c.ByUser)
}

func (c *MentionController) Create(ctx *gin.Context) {
	var mention entity.Mention
	if err := ctx.ShouldBindJSON(&mention); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "Mention")})
		return
	}
	if err := c.repo.Create(context.Background(), &mention); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, mention)
}

// CreateMention godoc
// @Summary Create mention
// @Description Create a mention for a tweet
// @Tags mentions
// @Accept json
// @Produce json
// @Param mention body entity.Mention true "Mention payload"
// @Success 201 {object} entity.Mention
// @Failure 400 {object} dto.WebResponse
// @Router /api/v1/mentions [post]

// GetMentionsByTweet godoc
// @Summary Mentions by tweet
// @Description Get mentions for a tweet
// @Tags mentions
// @Accept json
// @Produce json
// @Param tweet_id path int true "Tweet ID"
// @Success 200 {array} entity.Mention
// @Router /api/v1/mentions/tweet/{tweet_id} [get]

// GetMentionsByUser godoc
// @Summary Mentions by user
// @Description Get mentions of a user
// @Tags mentions
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} entity.Mention
// @Router /api/v1/mentions/user/{user_id} [get]

func (c *MentionController) ByTweet(ctx *gin.Context) {
	tweetID, _ := strconv.ParseInt(ctx.Param("tweet_id"), 10, 64)
	mentions, err := c.repo.FindByTweetID(context.Background(), tweetID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, mentions)
}

func (c *MentionController) ByUser(ctx *gin.Context) {
	userID, _ := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	mentions, err := c.repo.FindByUserID(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, mentions)
}
