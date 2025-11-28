package controller

import (
	"net/http"

	"TwClone/internal/config"
	"TwClone/internal/dto"
	"TwClone/internal/entity"
	"TwClone/internal/pkg/utils/encryptutils"
	"TwClone/internal/pkg/utils/jwtutils"
	"TwClone/internal/repository"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	userRepo  repository.UserRepositoryImpl
	encryptor encryptutils.BcryptEncryptor
	jwt       jwtutils.JwtUtil
}

func NewAuthController(cfg *config.Config) *AuthController {
	// create bcrypt encryptor with configured cost (fallback to 10)
	cost := 10
	if cfg != nil && cfg.App != nil && cfg.App.BCryptCost != 0 {
		cost = cfg.App.BCryptCost
	}
	enc := encryptutils.NewBcryptEncryptor(cost)

	var ju jwtutils.JwtUtil
	if cfg != nil && cfg.Jwt != nil {
		ju = jwtutils.NewJwtUtil(cfg.Jwt)
	}

	return &AuthController{
		userRepo:  repository.UserRepositoryImpl{},
		encryptor: enc,
		jwt:       ju,
	}
}

func (c *AuthController) Route(g *echo.Group) {
	ag := g.Group("/auth")
	ag.POST("/login", c.Login)
	ag.POST("/register", c.Register)
}

type loginRequest struct {
	Email    string `json:"email" validate:"omitempty,email"`
	Username string `json:"username" validate:"omitempty"`
	Password string `json:"password" validate:"required"`
}

type registerRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name"`
	Username string `json:"username" validate:"required"`
	Avatar   string `json:"avatar"`
	Banner   string `json:"banner"`
	Bio      string `json:"bio"`
	Password string `json:"password" validate:"required"`
}

// Login godoc
// @Summary Login
// @Description Authenticate user with email or username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param login body loginRequest true "Login payload"
// @Success 200 {object} dto.WebResponse
// @Failure 400 {object} dto.WebResponse
// @Failure 401 {object} dto.WebResponse
// @Router /api/v1/auth/login [post]
func (c *AuthController) Login(ctx echo.Context) error {
	var req loginRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "loginRequest")})
	}

	var user *entity.User
	var err error
	if req.Email != "" {
		user, err = c.userRepo.FindByEmail(ctx.Request().Context(), req.Email)
	} else if req.Username != "" {
		user, err = c.userRepo.FindByUsername(ctx.Request().Context(), req.Username)
	} else {
		return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "email or username required"})
	}

	if err != nil {
		if err == repository.ErrRecordNotFound {
			return ctx.JSON(http.StatusUnauthorized, dto.WebResponse[any]{Message: "invalid credentials"})
		}
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to lookup user"})
	}

	if !c.encryptor.Check(req.Password, user.Password) {
		return ctx.JSON(http.StatusUnauthorized, dto.WebResponse[any]{Message: "invalid credentials"})
	}

	var token string
	if c.jwt != nil {
		t, err := c.jwt.Sign(user.ID)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to generate token"})
		}
		token = t

	}

	userDTO := dto.FromEntity(user)

	return ctx.JSON(http.StatusOK, dto.WebResponse[any]{Data: echo.Map{"token": token, "user": userDTO}})
}

// Register godoc
// @Summary Register
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param register body registerRequest true "Register payload"
// @Success 201 {object} dto.WebResponse
// @Failure 400 {object} dto.WebResponse
// @Failure 409 {object} dto.WebResponse
// @Router /api/v1/auth/register [post]
func (c *AuthController) Register(ctx echo.Context) error {
	var req registerRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "registerRequest")})
	}

	// hash password before storing
	hashed, err := c.encryptor.Hash(req.Password)
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

	if err := c.userRepo.Create(ctx.Request().Context(), user); err != nil {
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
