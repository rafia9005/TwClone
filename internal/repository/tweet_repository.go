package repository

import (
	"TwClone/internal/database"
	"TwClone/internal/entity"
	"context"
	"strings"
)

type TweetRepositoryImpl struct{}

func (r TweetRepositoryImpl) Create(ctx context.Context, tweet *entity.Tweet) error {
	result := database.DB.WithContext(ctx).Create(tweet)

	if result.Error != nil {
		errMsg := result.Error.Error()
		if strings.Contains(errMsg, "duplicate key") || strings.Contains(errMsg, "unique constraint") {
			return ErrDuplicate
		}
		return result.Error
	}
	return nil
}

func (r TweetRepositoryImpl) FindAll(ctx context.Context) ([]*entity.Tweet, error) {
	var tweets []*entity.Tweet
	result := database.DB.WithContext(ctx).Find(&tweets)
	if result.Error != nil {
		return nil, result.Error
	}
	return tweets, nil
}
