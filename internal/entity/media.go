package entity

import "time"

// Media stores metadata for media attached to a tweet.
type Media struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	TweetID   int64     `gorm:"index;not null" json:"tweet_id"`
	MediaType string    `gorm:"size:50" json:"media_type"`
	URL       string    `gorm:"size:1024" json:"url"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
