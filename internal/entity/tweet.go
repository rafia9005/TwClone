package entity

import "time"

// Tweet represents a post (tweet) created by a user.
type Tweet struct {
	ID               int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID           int64     `gorm:"not null;index" json:"user_id"`
	Content          string    `gorm:"size:280;not null" json:"content"`
	ReplyToTweetID   *int64    `gorm:"index" json:"reply_to_tweet_id,omitempty"`
	RetweetedTweetID *int64    `gorm:"index" json:"retweeted_tweet_id,omitempty"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
