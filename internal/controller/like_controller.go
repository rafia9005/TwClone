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
