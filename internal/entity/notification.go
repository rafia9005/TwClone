package entity

import "time"

// Notification represents a notification sent to a user.
type Notification struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	RecipientID int64     `gorm:"index;not null" json:"recipient_id"`
	SenderID    *int64    `json:"sender_id,omitempty"`
	Type        string    `gorm:"size:50" json:"type"`
	TweetID     *int64    `gorm:"index" json:"tweet_id,omitempty"`
	IsRead      bool      `gorm:"default:false" json:"is_read"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
