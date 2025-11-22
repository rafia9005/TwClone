package repository

import (
	"TWclone/internal/database"
	"TWclone/internal/entity"
	"context"
)

type TweetHashtagRepositoryImpl struct{}

func (r TweetHashtagRepositoryImpl) Create(ctx context.Context, th *entity.TweetHashtag) error {
	return database.DB.WithContext(ctx).Create(th).Error
}

func (r TweetHashtagRepositoryImpl) FindByTweetID(ctx context.Context, tweetID int64) ([]*entity.TweetHashtag, error) {
	var ths []*entity.TweetHashtag
	result := database.DB.WithContext(ctx).Where("tweet_id = ?", tweetID).Find(&ths)
	if result.Error != nil {
		return nil, result.Error
	}
	return ths, nil
}

func (r TweetHashtagRepositoryImpl) FindByHashtagID(ctx context.Context, hashtagID int64) ([]*entity.TweetHashtag, error) {
	var ths []*entity.TweetHashtag
	result := database.DB.WithContext(ctx).Where("hashtag_id = ?", hashtagID).Find(&ths)
	if result.Error != nil {
		return nil, result.Error
	}
	return ths, nil
}
