package entity

// Hashtag stores unique hashtag values.
type Hashtag struct {
	ID      int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	TagName string `gorm:"size:255;uniqueIndex;not null" json:"tag_name"`
}
