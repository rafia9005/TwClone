package entity

import "time"

// Like represents a user liking a tweet.
type Like struct {
	UserID    int64     `gorm:"primaryKey;index" json:"user_id"`
	TweetID   int64     `gorm:"primaryKey;index" json:"tweet_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
