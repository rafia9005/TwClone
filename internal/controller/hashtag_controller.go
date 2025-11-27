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

// CreateHashtag godoc
// @Summary Create hashtag
// @Description Create a new hashtag
// @Tags hashtags
// @Accept json
// @Produce json
// @Param hashtag body entity.Hashtag true "Hashtag payload"
// @Success 201 {object} entity.Hashtag
// @Failure 400 {object} dto.WebResponse
// @Router /api/v1/hashtags [post]

// ListHashtags godoc
// @Summary List hashtags
// @Description Get all hashtags
// @Tags hashtags
// @Accept json
// @Produce json
// @Success 200 {array} entity.Hashtag
// @Router /api/v1/hashtags [get]

// GetHashtagByTag godoc
// @Summary Get hashtag
// @Description Get hashtag by tag
// @Tags hashtags
// @Accept json
// @Produce json
// @Param tag path string true "Tag"
// @Success 200 {object} entity.Hashtag
// @Router /api/v1/hashtags/{tag} [get]

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
