package controller

import (
	"TWclone/internal/dto"
	"TWclone/internal/entity"
	"TWclone/internal/repository"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	repo repository.NotificationRepositoryImpl
}

func NewNotificationController() *NotificationController {
	return &NotificationController{repo: repository.NotificationRepositoryImpl{}}
}

func (c *NotificationController) Route(r gin.IRouter) {
	g := r.Group("/notifications")
	g.POST("", c.Create)
	g.GET("/recipient/:recipient_id", c.ByRecipient)
	g.PUT("/:id/read", c.MarkAsRead)
}

func (c *NotificationController) Create(ctx *gin.Context) {
	var notif entity.Notification
	if err := ctx.ShouldBindJSON(&notif); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "Notification")})
		return
	}
	if err := c.repo.Create(context.Background(), &notif); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, notif)
}

func (c *NotificationController) ByRecipient(ctx *gin.Context) {
	recipientID, _ := strconv.ParseInt(ctx.Param("recipient_id"), 10, 64)
	notifs, err := c.repo.FindByRecipientID(context.Background(), recipientID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, notifs)
}

func (c *NotificationController) MarkAsRead(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := c.repo.MarkAsRead(context.Background(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "marked as read"})
}
