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

type LikeController struct {
	repo repository.LikeRepositoryImpl
}

func NewLikeController() *LikeController {
	return &LikeController{repo: repository.LikeRepositoryImpl{}}
}

func (c *LikeController) Route(r gin.IRouter) {
	g := r.Group("/likes")
	g.POST("", c.Create)
	g.DELETE("", c.Delete)
	g.GET("/tweet/:tweet_id", c.ByTweet)
	g.GET("/user/:user_id", c.ByUser)
}

func (c *LikeController) Create(ctx *gin.Context) {
	var like entity.Like
	if err := ctx.ShouldBindJSON(&like); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "Like")})
		return
	}
	if err := c.repo.Create(context.Background(), &like); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, like)
}

// CreateLike godoc
// @Summary Create like
// @Description Like a tweet
// @Tags likes
// @Accept json
// @Produce json
// @Param like body entity.Like true "Like payload"
// @Success 201 {object} entity.Like
// @Failure 400 {object} dto.WebResponse
// @Router /api/v1/likes [post]

// DeleteLike godoc
// @Summary Remove like
// @Description Remove like from a tweet
// @Tags likes
// @Accept json
// @Produce json
// @Param user_id query int true "User ID"
// @Param tweet_id query int true "Tweet ID"
// @Success 200 {object} dto.WebResponse
// @Router /api/v1/likes [delete]

// GetLikesByTweet godoc
// @Summary Likes by tweet
// @Description Get likes for a tweet
// @Tags likes
// @Accept json
// @Produce json
// @Param tweet_id path int true "Tweet ID"
// @Success 200 {array} entity.Like
// @Router /api/v1/likes/tweet/{tweet_id} [get]

// GetLikesByUser godoc
// @Summary Likes by user
// @Description Get likes by a user
// @Tags likes
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} entity.Like
// @Router /api/v1/likes/user/{user_id} [get]

func (c *LikeController) Delete(ctx *gin.Context) {
	userID, _ := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	tweetID, _ := strconv.ParseInt(ctx.Query("tweet_id"), 10, 64)
	if err := c.repo.Delete(context.Background(), userID, tweetID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "unliked"})
}

func (c *LikeController) ByTweet(ctx *gin.Context) {
	tweetID, _ := strconv.ParseInt(ctx.Param("tweet_id"), 10, 64)
	likes, err := c.repo.FindByTweet(context.Background(), tweetID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, likes)
}

func (c *LikeController) ByUser(ctx *gin.Context) {
	userID, _ := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	likes, err := c.repo.FindByUser(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, likes)
}
