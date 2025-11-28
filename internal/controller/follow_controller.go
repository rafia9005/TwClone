package controller

import (
	"context"
	"net/http"
	"strconv"

	"TwClone/internal/dto"
	"TwClone/internal/entity"
	"TwClone/internal/middleware"
	"TwClone/internal/repository"

	"github.com/labstack/echo/v4"
)

type FollowController struct {
	repo repository.FollowRepositoryImpl
}

func NewFollowController() *FollowController {
	return &FollowController{repo: repository.FollowRepositoryImpl{}}
}

func (c *FollowController) Route(g *echo.Group) {
	fg := g.Group("/follows", middleware.AuthMiddleware())
	fg.POST("", c.Create)
	fg.DELETE("", c.Delete)
	fg.GET("/followers/:id", c.Followers)
	fg.GET("/following/:id", c.Following)
}

func (c *FollowController) Create(ctx echo.Context) error {
	var follow entity.Follow
	if err := ctx.Bind(&follow); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "Follow")})
	}
	if err := c.repo.Create(context.Background(), &follow); err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, follow)
}

// CreateFollow godoc
// @Summary Create follow
// @Description Follow a user
// @Tags follows
// @Accept json
// @Produce json
// @Param follow body entity.Follow true "Follow payload"
// @Success 201 {object} entity.Follow
// @Failure 400 {object} dto.WebResponse
// @Router /api/v1/follows [post]

// DeleteFollow godoc
// @Summary Unfollow
// @Description Unfollow a user
// @Tags follows
// @Accept json
// @Produce json
// @Param follower_id query int true "Follower ID"
// @Param following_id query int true "Following ID"
// @Success 200 {object} dto.WebResponse
// @Failure 400 {object} dto.WebResponse
// @Router /api/v1/follows [delete]

// GetFollowers godoc
// @Summary Get followers
// @Description Get followers of a user
// @Tags follows
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} entity.Follow
// @Router /api/v1/follows/followers/{id} [get]

// GetFollowing godoc
// @Summary Get following
// @Description Get users followed by a user
// @Tags follows
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} entity.Follow
// @Router /api/v1/follows/following/{id} [get]

func (c *FollowController) Delete(ctx echo.Context) error {
	followerID, _ := strconv.ParseInt(ctx.QueryParam("follower_id"), 10, 64)
	followingID, _ := strconv.ParseInt(ctx.QueryParam("following_id"), 10, 64)
	if err := c.repo.Delete(context.Background(), followerID, followingID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "unfollowed"})
}

func (c *FollowController) Followers(ctx echo.Context) error {
	userID, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	follows, err := c.repo.FindFollowers(context.Background(), userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, follows)
}

func (c *FollowController) Following(ctx echo.Context) error {
	userID, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	follows, err := c.repo.FindFollowing(context.Background(), userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, follows)
}
