package controller

import (
	"TWclone/internal/dto"
	"TWclone/internal/entity"
	"TWclone/internal/repository"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HashtagController struct {
	repo repository.HashtagRepositoryImpl
}

func NewHashtagController() *HashtagController {
	return &HashtagController{repo: repository.HashtagRepositoryImpl{}}
}

func (c *HashtagController) Route(r gin.IRouter) {
	g := r.Group("/hashtags")
	g.POST("", c.Create)
	g.GET("", c.FindAll)
	g.GET("/:tag", c.FindByTag)
}

func (c *HashtagController) Create(ctx *gin.Context) {
	var hashtag entity.Hashtag
	if err := ctx.ShouldBindJSON(&hashtag); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "Hashtag")})
		return
	}
	if err := c.repo.Create(context.Background(), &hashtag); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, hashtag)
}

func (c *HashtagController) FindByTag(ctx *gin.Context) {
	tag := ctx.Param("tag")
	hashtag, err := c.repo.FindByTag(context.Background(), tag)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	ctx.JSON(http.StatusOK, hashtag)
}

func (c *HashtagController) FindAll(ctx *gin.Context) {
	hashtags, err := c.repo.FindAll(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, hashtags)
}
