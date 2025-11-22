package entity

import "time"

// Mention records when a user is mentioned in a tweet.
type Mention struct {
	TweetID   int64     `gorm:"primaryKey;index" json:"tweet_id"`
	UserID    int64     `gorm:"primaryKey;index" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
