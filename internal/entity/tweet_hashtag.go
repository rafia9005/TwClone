package entity

// TweetHashtag is the many-to-many join between tweets and hashtags.
type TweetHashtag struct {
	TweetID   int64 `gorm:"primaryKey;index" json:"tweet_id"`
	HashtagID int64 `gorm:"primaryKey;index" json:"hashtag_id"`
}
