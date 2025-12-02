package controller

import (
	"context"
	"net/http"
	"strconv"

	"TwClone/internal/dto"
	"TwClone/internal/entity"
	"TwClone/internal/repository"

	"github.com/labstack/echo/v4"
)

type LikeController struct {
	repo repository.LikeRepositoryImpl
}

func NewLikeController() *LikeController {
	return &LikeController{repo: repository.LikeRepositoryImpl{}}
}

func (c *LikeController) Route(g *echo.Group) {
	lg := g.Group("/likes")
	lg.POST("", c.Create)
	lg.DELETE("", c.Delete)
	lg.GET("/tweet/:tweet_id", c.ByTweet)
	lg.GET("/user/:user_id", c.ByUser)
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
func (c *LikeController) Create(ctx echo.Context) error {
	var like entity.Like
	if err := ctx.Bind(&like); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "Like")})
	}
	if err := c.repo.Create(context.Background(), &like); err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, like)
}

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
func (c *LikeController) Delete(ctx echo.Context) error {
	userID, _ := strconv.ParseInt(ctx.QueryParam("user_id"), 10, 64)
	tweetID, _ := strconv.ParseInt(ctx.QueryParam("tweet_id"), 10, 64)
	if err := c.repo.Delete(context.Background(), userID, tweetID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "unliked"})
}

// GetLikesByTweet godoc
// @Summary Likes by tweet
// @Description Get likes for a tweet
// @Tags likes
// @Accept json
// @Produce json
// @Param tweet_id path int true "Tweet ID"
// @Success 200 {array} entity.Like
// @Router /api/v1/likes/tweet/{tweet_id} [get]
func (c *LikeController) ByTweet(ctx echo.Context) error {
	tweetID, _ := strconv.ParseInt(ctx.Param("tweet_id"), 10, 64)
	likes, err := c.repo.FindByTweet(context.Background(), tweetID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, likes)
}

// GetLikesByUser godoc
// @Summary Likes by user
// @Description Get likes by a user
// @Tags likes
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} entity.Like
// @Router /api/v1/likes/user/{user_id} [get]
func (c *LikeController) ByUser(ctx echo.Context) error {
	userID, _ := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	likes, err := c.repo.FindByUser(context.Background(), userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, likes)
}
