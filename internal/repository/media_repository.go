package repository

import (
	"TwClone/internal/database"
	"TwClone/internal/entity"
	"context"
)

type MediaRepositoryImpl struct{}

func (r MediaRepositoryImpl) Create(ctx context.Context, media *entity.Media) error {
	return database.DB.WithContext(ctx).Create(media).Error
}

func (r MediaRepositoryImpl) FindByTweetID(ctx context.Context, tweetID int64) ([]*entity.Media, error) {
	var medias []*entity.Media
	result := database.DB.WithContext(ctx).Where("tweet_id = ?", tweetID).Find(&medias)
	if result.Error != nil {
		return nil, result.Error
	}
	return medias, nil
}

func (r MediaRepositoryImpl) FindByID(ctx context.Context, id int64) (*entity.Media, error) {
	var media entity.Media
	result := database.DB.WithContext(ctx).First(&media, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &media, nil
}
