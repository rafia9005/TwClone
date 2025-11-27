package controller

import (
	"net/http"
	"strconv"

	"TWclone/internal/dto"
	"TWclone/internal/entity"
	"TWclone/internal/pkg/utils/encryptutils"
	"TWclone/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// extractFieldErrors parses validator.ValidationErrors to []dto.FieldError
func extractFieldErrors(err error, structPrefix string) []dto.FieldError {
	var fieldErrors []dto.FieldError
	if verrs, ok := err.(validator.ValidationErrors); ok {
		for _, verr := range verrs {
			field := verr.Field()
			// Remove struct prefix if present
			if structPrefix != "" && len(field) > len(structPrefix) && field[:len(structPrefix)] == structPrefix {
				field = field[len(structPrefix):]
			}
			fieldErrors = append(fieldErrors, dto.FieldError{
				Field:   field,
				Message: verr.Error(),
			})
		}
	} else {
		// fallback: try to parse error string
		msg := err.Error()
		field := "body"
		fieldErrors = append(fieldErrors, dto.FieldError{Field: field, Message: msg})
	}
	return fieldErrors
}

// UserController handles user CRUD.
type UserController struct {
	repo repository.UserRepositoryImpl
}

func NewUserController() *UserController {
	return &UserController{repo: repository.UserRepositoryImpl{}}
}

func (c *UserController) Route(r gin.IRouter) {
	g := r.Group("/users")
	g.POST("", c.Create)
	g.GET("", c.FindAll)
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
func (c *UserController) Create(ctx *gin.Context) {
	var req createUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "createUserReq")})
		return
	}

	// hash password before storing
	enc := encryptutils.NewBcryptEncryptor(10)
	hashed, err := enc.Hash(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to hash password"})
		return
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

	if err := c.repo.Create(ctx, user); err != nil {
		if err.Error() == "duplicate_email" {
			ctx.JSON(http.StatusConflict, dto.WebResponse[any]{
				Message: "validation error",
				Errors:  []dto.FieldError{{Field: "email", Message: "email already exists"}},
			})
			return
		}
		if err.Error() == "duplicate_username" {
			ctx.JSON(http.StatusConflict, dto.WebResponse[any]{
				Message: "validation error",
				Errors:  []dto.FieldError{{Field: "username", Message: "username already exists"}},
			})
			return
		}
		if err == repository.ErrDuplicate {
			ctx.JSON(http.StatusConflict, dto.WebResponse[any]{
				Message: "validation error",
				Errors:  []dto.FieldError{{Field: "unknown", Message: "email or username already exists"}},
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed create user"})
		return
	}

	ctx.JSON(http.StatusCreated, dto.WebResponse[dto.UserResponse]{Message: "created", Data: dto.FromEntity(user)})
}

// ListUsers godoc
// @Summary List users
// @Description Get list of users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} dto.WebResponse
// @Router /api/v1/users [get]
func (c *UserController) FindAll(ctx *gin.Context) {
	users, err := c.repo.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to fetch users"})
		return
	}
	// convert to response DTOs to avoid leaking password
	resp := make([]dto.UserResponse, 0, len(users))
	for _, u := range users {
		resp = append(resp, dto.FromEntity(u))
	}
	ctx.JSON(http.StatusOK, dto.WebResponse[any]{Data: resp})
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
func (c *UserController) FindByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid id"})
		return
	}

	user, err := c.repo.FindByID(ctx, id)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, dto.WebResponse[any]{Message: "user not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to fetch user"})
		return
	}
	ctx.JSON(http.StatusOK, dto.WebResponse[dto.UserResponse]{Data: dto.FromEntity(user)})
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
func (c *UserController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid id"})
		return
	}

	var req updateUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "updateUserReq")})
		return
	}

	user, err := c.repo.FindByID(ctx, id)
	if err != nil {
		if err == repository.ErrRecordNotFound {
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
		enc := encryptutils.NewBcryptEncryptor(10)
		hashed, err := enc.Hash(*req.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to hash password"})
			return
		}
		user.Password = hashed
	}

	if err := c.repo.Update(ctx, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to update user"})
		return
	}
	ctx.JSON(http.StatusOK, dto.WebResponse[dto.UserResponse]{Message: "updated", Data: dto.FromEntity(user)})
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

// GetUserToken godoc
// @Summary Get current user
// @Description Get current authenticated user's info via token
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} dto.WebResponse
// @Failure 401 {object} dto.WebResponse
// @Router /api/v1/users/token [get]
func (c *UserController) UserToken(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, dto.WebResponse[any]{Message: "unauthorized"})
		return
	}

	user, err := c.repo.FindByID(ctx, userID.(int64))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to fetch user"})
		return
	}

	ctx.JSON(http.StatusOK, dto.WebResponse[dto.UserResponse]{Data: dto.FromEntity(user)})
}
