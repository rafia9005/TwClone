package controller

import (
	"net/http"
	"strconv"

	"TwClone/internal/dto"
	"TwClone/internal/entity"
	"TwClone/internal/middleware"
	"TwClone/internal/pkg/utils/encryptutils"
	"TwClone/internal/repository"

	"github.com/labstack/echo/v4"
)

// UserController handles user CRUD.
type UserController struct {
	repo repository.UserRepositoryImpl
}

func NewUserController() *UserController {
	return &UserController{repo: repository.UserRepositoryImpl{}}
}

func (c *UserController) Route(g *echo.Group) {
	ug := g.Group("/users", middleware.AuthMiddleware())
	ug.POST("", c.Create)
	ug.GET("", c.FindAll)
	ug.GET("/token", c.UserToken)
	ug.GET("/:id", c.FindByID)
	ug.PUT("/:id", c.Update)
	ug.DELETE("/:id", c.Delete)
}

type createUserReq struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name"`
	Username string `json:"username" validate:"required"`
	Avatar   string `json:"avatar"`
	Banner   string `json:"banner"`
	Bio      string `json:"bio"`
	Password string `json:"password" validate:"required"`
}

type updateUserReq struct {
	Email    *string `json:"email"`
	Name     *string `json:"name"`
	Avatar   *string `json:"avatar"`
	Banner   *string `json:"banner"`
	Bio      *string `json:"bio"`
	Password *string `json:"password"`
}

// CreateUser godoc
// @Summary Create user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body createUserReq true "Create user payload"
// @Success 201 {object} dto.WebResponse
// @Failure 400 {object} dto.WebResponse
// @Failure 409 {object} dto.WebResponse
// @Router /api/v1/users [post]
func (c *UserController) Create(ctx echo.Context) error {
	var req createUserReq
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "createUserReq")})
	}

	// hash password before storing
	enc := encryptutils.NewBcryptEncryptor(10)
	hashed, err := enc.Hash(req.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to hash password"})
	}

	user := &entity.User{
		Email:    req.Email,
		Name:     req.Name,
		Username: req.Username,
		Avatar:   req.Avatar,
		Banner:   req.Banner,
		Bio:      req.Bio,
		Password: hashed,
	}

	if err := c.repo.Create(ctx.Request().Context(), user); err != nil {
		if err.Error() == "duplicate_email" {
			return ctx.JSON(http.StatusConflict, dto.WebResponse[any]{
				Message: "validation error",
				Errors:  []dto.FieldError{{Field: "email", Message: "email already exists"}},
			})
		}
		if err.Error() == "duplicate_username" {
			return ctx.JSON(http.StatusConflict, dto.WebResponse[any]{
				Message: "validation error",
				Errors:  []dto.FieldError{{Field: "username", Message: "username already exists"}},
			})
		}
		if err == repository.ErrDuplicate {
			return ctx.JSON(http.StatusConflict, dto.WebResponse[any]{
				Message: "validation error",
				Errors:  []dto.FieldError{{Field: "unknown", Message: "email or username already exists"}},
			})
		}
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed create user"})
	}

	return ctx.JSON(http.StatusCreated, dto.WebResponse[dto.UserResponse]{Message: "created", Data: dto.FromEntity(user)})
}

// ListUsers godoc
// @Summary List users
// @Description Get list of users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} dto.WebResponse
// @Router /api/v1/users [get]
func (c *UserController) FindAll(ctx echo.Context) error {
	users, err := c.repo.FindAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to fetch users"})
	}
	// convert to response DTOs to avoid leaking password
	resp := make([]dto.UserResponse, 0, len(users))
	for _, u := range users {
		resp = append(resp, dto.FromEntity(u))
	}
	return ctx.JSON(http.StatusOK, dto.WebResponse[any]{Data: resp})
}

// GetUser godoc
// @Summary Get user by id
// @Description Get a user by its ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.WebResponse
// @Failure 400 {object} dto.WebResponse
// @Failure 404 {object} dto.WebResponse
// @Router /api/v1/users/{id} [get]
func (c *UserController) FindByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid id"})
	}

	user, err := c.repo.FindByID(ctx.Request().Context(), id)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			return ctx.JSON(http.StatusNotFound, dto.WebResponse[any]{Message: "user not found"})
		}
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to fetch user"})
	}
	return ctx.JSON(http.StatusOK, dto.WebResponse[dto.UserResponse]{Data: dto.FromEntity(user)})
}

// UpdateUser godoc
// @Summary Update user
// @Description Update a user's information
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body updateUserReq true "Update payload"
// @Success 200 {object} dto.WebResponse
// @Failure 400 {object} dto.WebResponse
// @Failure 404 {object} dto.WebResponse
// @Router /api/v1/users/{id} [put]
func (c *UserController) Update(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid id"})
	}

	var req updateUserReq
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "updateUserReq")})
	}

	user, err := c.repo.FindByID(ctx.Request().Context(), id)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			return ctx.JSON(http.StatusNotFound, dto.WebResponse[any]{Message: "user not found"})
		}
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to fetch user"})
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
		enc := encryptutils.NewBcryptEncryptor(10)
		hashed, err := enc.Hash(*req.Password)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to hash password"})
		}
		user.Password = hashed
	}

	if err := c.repo.Update(ctx.Request().Context(), user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to update user"})
	}
	return ctx.JSON(http.StatusOK, dto.WebResponse[dto.UserResponse]{Message: "updated", Data: dto.FromEntity(user)})
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 {object} dto.WebResponse
// @Failure 400 {object} dto.WebResponse
// @Failure 404 {object} dto.WebResponse
// @Router /api/v1/users/{id} [delete]
func (c *UserController) Delete(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid id"})
	}

	if err := c.repo.Delete(ctx.Request().Context(), id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to delete user"})
	}
	return ctx.NoContent(http.StatusNoContent)
}

// GetUserToken godoc
// @Summary Get current user
// @Description Get current authenticated user's info via token
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} dto.WebResponse
// @Failure 401 {object} dto.WebResponse
// @Router /api/v1/users/token [get]
func (c *UserController) UserToken(ctx echo.Context) error {
	// read user id set by auth middleware using the package constant key
	uid := ctx.Get("ctx-user-id")
	if uid == nil {
		return ctx.JSON(http.StatusUnauthorized, dto.WebResponse[any]{Message: "unauthorized"})
	}

	var userID int64
	switch v := uid.(type) {
	case int64:
		userID = v
	case int:
		userID = int64(v)
	default:
		return ctx.JSON(http.StatusUnauthorized, dto.WebResponse[any]{Message: "unauthorized"})
	}

	user, err := c.repo.FindByID(ctx.Request().Context(), userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to fetch user"})
	}

	return ctx.JSON(http.StatusOK, dto.WebResponse[dto.UserResponse]{Data: dto.FromEntity(user)})
}
