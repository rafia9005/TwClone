package repository

import (
	"TwClone/internal/database"
	"TwClone/internal/entity"
	"context"
)

type LikeRepositoryImpl struct{}

func (r LikeRepositoryImpl) Create(ctx context.Context, like *entity.Like) error {
	return database.DB.WithContext(ctx).Create(like).Error
}

func (r LikeRepositoryImpl) Delete(ctx context.Context, userID, tweetID int64) error {
	return database.DB.WithContext(ctx).Where("user_id = ? AND tweet_id = ?", userID, tweetID).Delete(&entity.Like{}).Error
}

func (r LikeRepositoryImpl) FindByTweet(ctx context.Context, tweetID int64) ([]*entity.Like, error) {
	var likes []*entity.Like
	result := database.DB.WithContext(ctx).Where("tweet_id = ?", tweetID).Find(&likes)
	if result.Error != nil {
		return nil, result.Error
	}
	return likes, nil
}

func (r LikeRepositoryImpl) FindByUser(ctx context.Context, userID int64) ([]*entity.Like, error) {
	var likes []*entity.Like
	result := database.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&likes)
	if result.Error != nil {
		return nil, result.Error
	}
	return likes, nil
}
