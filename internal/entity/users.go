package entity

import "time"

// User represents a user in the system. Includes GORM tags for migrations.
// If your DB column names differ, adjust the `gorm:"column:..."` tags.
type User struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" db:"id" json:"id"`
	Email     string    `gorm:"size:255;uniqueIndex;not null" db:"email" json:"email"`
	Name      string    `gorm:"size:255" db:"name" json:"name"`
	Username  string    `gorm:"size:100;uniqueIndex;not null" db:"username" json:"username"`
	Avatar    string    `gorm:"size:1024" db:"avatar" json:"avatar,omitempty"`
	Banner    string    `gorm:"size:1024" db:"banner" json:"banner,omitempty"`
	Bio       string    `gorm:"type:text" db:"bio" json:"bio,omitempty"`
	Password  string    `gorm:"size:255;not null" db:"password" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" db:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" db:"updated_at" json:"updated_at"`
}
