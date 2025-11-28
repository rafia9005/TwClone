package repository

import (
	"TwClone/internal/database"
	"TwClone/internal/entity"
	"context"
)

type HashtagRepositoryImpl struct{}

func (r HashtagRepositoryImpl) Create(ctx context.Context, hashtag *entity.Hashtag) error {
	return database.DB.WithContext(ctx).Create(hashtag).Error
}

func (r HashtagRepositoryImpl) FindByTag(ctx context.Context, tag string) (*entity.Hashtag, error) {
	var hashtag entity.Hashtag
	result := database.DB.WithContext(ctx).Where("tag_name = ?", tag).First(&hashtag)
	if result.Error != nil {
		return nil, result.Error
	}
	return &hashtag, nil
}

func (r HashtagRepositoryImpl) FindAll(ctx context.Context) ([]*entity.Hashtag, error) {
	var hashtags []*entity.Hashtag
	result := database.DB.WithContext(ctx).Find(&hashtags)
	if result.Error != nil {
		return nil, result.Error
	}
	return hashtags, nil
}
