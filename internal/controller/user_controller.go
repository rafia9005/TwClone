package controller

import (
	"net/http"
	"strconv"

	"TWclone/internal/dto"
	"TWclone/internal/entity"
	"TWclone/internal/repository"

	"github.com/gin-gonic/gin"
)

// UserController handles user CRUD.
type UserController struct {
	repo repository.UserRepositoryImpl
}

func NewUserController() *UserController {
	return &UserController{repo: repository.UserRepositoryImpl{}}
}

func (c *UserController) Route(r gin.IRouter) {
	g := r.Group("/users")
	g.POST("/", c.Create)
	g.GET("/", c.FindAll)
	g.GET("/:id", c.FindByID)
	g.PUT("/:id", c.Update)
	g.DELETE("/:id", c.Delete)
}

type createUserReq struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name"`
	Username string `json:"username" binding:"required"`
	Avatar   string `json:"avatar"`
	Banner   string `json:"banner"`
	Bio      string `json:"bio"`
	Password string `json:"password" binding:"required"`
}

type updateUserReq struct {
	Email    *string `json:"email"`
	Name     *string `json:"name"`
	Avatar   *string `json:"avatar"`
	Banner   *string `json:"banner"`
	Bio      *string `json:"bio"`
	Password *string `json:"password"`
}

func (c *UserController) Create(ctx *gin.Context) {
	var req createUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: []dto.FieldError{{Field: "body", Message: err.Error()}}})
		return
	}

	user := &entity.User{
		Email:    req.Email,
		Name:     req.Name,
		Username: req.Username,
		Avatar:   req.Avatar,
		Banner:   req.Banner,
		Bio:      req.Bio,
		Password: req.Password,
	}

	if err := c.repo.Create(ctx, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed create user"})
		return
	}

	ctx.JSON(http.StatusCreated, dto.WebResponse[*entity.User]{Message: "created", Data: user})
}

func (c *UserController) FindAll(ctx *gin.Context) {
	users, err := c.repo.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to fetch users"})
		return
	}
	ctx.JSON(http.StatusOK, dto.WebResponse[any]{Data: users})
}

func (c *UserController) FindByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid id"})
		return
	}

	user, err := c.repo.FindByID(ctx, id)
	if err != nil {
		if err == repository.ERR_RECORD_NOT_FOUND {
			ctx.JSON(http.StatusNotFound, dto.WebResponse[any]{Message: "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to fetch user"})
		return
	}
	ctx.JSON(http.StatusOK, dto.WebResponse[*entity.User]{Data: user})
}

func (c *UserController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid id"})
		return
	}

	var req updateUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: []dto.FieldError{{Field: "body", Message: err.Error()}}})
		return
	}

	user, err := c.repo.FindByID(ctx, id)
	if err != nil {
		if err == repository.ERR_RECORD_NOT_FOUND {
			ctx.JSON(http.StatusNotFound, dto.WebResponse[any]{Message: "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to fetch user"})
		return
	}

	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Avatar != nil {
		user.Avatar = *req.Avatar
	}
	if req.Banner != nil {
		user.Banner = *req.Banner
	}
	if req.Bio != nil {
		user.Bio = *req.Bio
	}
	if req.Password != nil {
		user.Password = *req.Password
	}

	if err := c.repo.Update(ctx, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to update user"})
		return
	}
	ctx.JSON(http.StatusOK, dto.WebResponse[*entity.User]{Message: "updated", Data: user})
}

func (c *UserController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid id"})
		return
	}

	if err := c.repo.Delete(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to delete user"})
		return
	}
	ctx.Status(http.StatusNoContent)
}
