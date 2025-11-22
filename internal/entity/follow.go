package entity

import "time"

// Follow represents a following relationship between two users.
type Follow struct {
	FollowerID  int64     `gorm:"primaryKey;index" json:"follower_id"`
	FollowingID int64     `gorm:"primaryKey;index" json:"following_id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}
