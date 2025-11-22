package repository

import (
	"TWclone/internal/database"
	"TWclone/internal/entity"
	"context"
)

type NotificationRepositoryImpl struct{}

func (r NotificationRepositoryImpl) Create(ctx context.Context, notif *entity.Notification) error {
	return database.DB.WithContext(ctx).Create(notif).Error
}

func (r NotificationRepositoryImpl) FindByRecipientID(ctx context.Context, recipientID int64) ([]*entity.Notification, error) {
	var notifs []*entity.Notification
	result := database.DB.WithContext(ctx).Where("recipient_id = ?", recipientID).Find(&notifs)
	if result.Error != nil {
		return nil, result.Error
	}
	return notifs, nil
}

func (r NotificationRepositoryImpl) MarkAsRead(ctx context.Context, id int64) error {
	return database.DB.WithContext(ctx).Model(&entity.Notification{}).Where("id = ?", id).Update("is_read", true).Error
}
