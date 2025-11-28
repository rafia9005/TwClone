package controller

import (
	"TwClone/internal/dto"
	"TwClone/internal/entity"
	"TwClone/internal/repository"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NotificationController struct {
	repo repository.NotificationRepositoryImpl
}

func NewNotificationController() *NotificationController {
	return &NotificationController{repo: repository.NotificationRepositoryImpl{}}
}

func (c *NotificationController) Route(g *echo.Group) {
	ng := g.Group("/notifications")
	ng.POST("", c.Create)
	ng.GET("/recipient/:recipient_id", c.ByRecipient)
	ng.PUT("/:id/read", c.MarkAsRead)
}

// CreateNotification godoc
// @Summary Create notification
// @Description Create a notification for a recipient
// @Tags notifications
// @Accept json
// @Produce json
// @Param notification body entity.Notification true "Notification payload"
// @Success 201 {object} entity.Notification
// @Failure 400 {object} dto.WebResponse
// @Router /api/v1/notifications [post]
func (c *NotificationController) Create(ctx echo.Context) error {
	var notif entity.Notification
	if err := ctx.Bind(&notif); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.WebResponse[any]{Message: "invalid request", Errors: extractFieldErrors(err, "Notification")})
	}
	if err := c.repo.Create(context.Background(), &notif); err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusCreated, notif)
}

// GetNotificationsByRecipient godoc
// @Summary Notifications by recipient
// @Description Get notifications for a recipient
// @Tags notifications
// @Accept json
// @Produce json
// @Param recipient_id path int true "Recipient ID"
// @Success 200 {array} entity.Notification
// @Router /api/v1/notifications/recipient/{recipient_id} [get]
func (c *NotificationController) ByRecipient(ctx echo.Context) error {
	recipientID, _ := strconv.ParseInt(ctx.Param("recipient_id"), 10, 64)
	notifs, err := c.repo.FindByRecipientID(context.Background(), recipientID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, notifs)
}

// MarkNotificationAsRead godoc
// @Summary Mark notification as read
// @Description Mark a notification as read by id
// @Tags notifications
// @Accept json
// @Produce json
// @Param id path int true "Notification ID"
// @Success 200 {object} dto.WebResponse
// @Router /api/v1/notifications/{id}/read [put]
func (c *NotificationController) MarkAsRead(ctx echo.Context) error {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err := c.repo.MarkAsRead(context.Background(), id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, dto.WebResponse[any]{Message: err.Error()})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "marked as read"})
}
