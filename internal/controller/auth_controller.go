package controller

import (
	"net/http"

	"TWclone/internal/config"
	"TWclone/internal/dto"
	"TWclone/internal/entity"
	"TWclone/internal/pkg/utils/encryptutils"
	"TWclone/internal/pkg/utils/jwtutils"
	"TWclone/internal/repository"

	"github.com/gin-gonic/gin"
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

func (c *AuthController) Route(r gin.IRouter) {
	g := r.Group("/auth")
	g.POST("/login", c.Login)
	g.POST("/register", c.Register)
}

type loginRequest struct {
	Email    string `json:"email" binding:"omitempty,email"`
	Username string `json:"username" binding:"omitempty"`
	Password string `json:"password" binding:"required"`
}

type registerRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name"`
	Username string `json:"username" binding:"required"`
	Avatar   string `json:"avatar"`
	Banner   string `json:"banner"`
	Bio      string `json:"bio"`
	Password string `json:"password" binding:"required"`
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "loginRequest")})
		return
	}

	var user *entity.User
	var err error
	if req.Email != "" {
		user, err = c.userRepo.FindByEmail(ctx, req.Email)
	} else if req.Username != "" {
		user, err = c.userRepo.FindByUsername(ctx, req.Username)
	} else {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "email or username required"})
		return
	}

	if err != nil {
		if err == repository.ErrRecordNotFound {
			ctx.JSON(http.StatusUnauthorized, dto.WebResponse[any]{Message: "invalid credentials"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to lookup user"})
		return
	}

	if !c.encryptor.Check(req.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, dto.WebResponse[any]{Message: "invalid credentials"})
		return
	}

	var token string
	if c.jwt != nil {
		t, err := c.jwt.Sign(user.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: "failed to generate token"})
			return
		}
		token = t
	}

	userDTO := dto.FromEntity(user)

	ctx.JSON(http.StatusOK, dto.WebResponse[any]{Data: gin.H{"token": token, "user": userDTO}})
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "registerRequest")})
		return
	}

	// hash password before storing
	hashed, err := c.encryptor.Hash(req.Password)
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

	if err := c.userRepo.Create(ctx, user); err != nil {
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
