package controller

import (
	"TWclone/internal/entity"
	"TWclone/internal/repository"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FollowController struct {
	repo repository.FollowRepositoryImpl
}

func NewFollowController() *FollowController {
	return &FollowController{repo: repository.FollowRepositoryImpl{}}
}

func (c *FollowController) Route(r gin.IRouter) {
	g := r.Group("/follows")
	g.POST("/", c.Create)
	g.DELETE("/", c.Delete)
	g.GET("/followers/:id", c.Followers)
	g.GET("/following/:id", c.Following)
}

func (c *FollowController) Create(ctx *gin.Context) {
	var follow entity.Follow
	if err := ctx.ShouldBindJSON(&follow); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.repo.Create(context.Background(), &follow); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, follow)
}

func (c *FollowController) Delete(ctx *gin.Context) {
	followerID, _ := strconv.ParseInt(ctx.Query("follower_id"), 10, 64)
	followingID, _ := strconv.ParseInt(ctx.Query("following_id"), 10, 64)
	if err := c.repo.Delete(context.Background(), followerID, followingID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "unfollowed"})
}

func (c *FollowController) Followers(ctx *gin.Context) {
	userID, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	follows, err := c.repo.FindFollowers(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, follows)
}

func (c *FollowController) Following(ctx *gin.Context) {
	userID, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	follows, err := c.repo.FindFollowing(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, follows)
}
