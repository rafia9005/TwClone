package repository

import (
	"TWclone/internal/database"
	"TWclone/internal/entity"
	"context"
)

type MentionRepositoryImpl struct{}

func (r MentionRepositoryImpl) Create(ctx context.Context, mention *entity.Mention) error {
	return database.DB.WithContext(ctx).Create(mention).Error
}

func (r MentionRepositoryImpl) FindByTweetID(ctx context.Context, tweetID int64) ([]*entity.Mention, error) {
	var mentions []*entity.Mention
	result := database.DB.WithContext(ctx).Where("tweet_id = ?", tweetID).Find(&mentions)
	if result.Error != nil {
		return nil, result.Error
	}
	return mentions, nil
}

func (r MentionRepositoryImpl) FindByUserID(ctx context.Context, userID int64) ([]*entity.Mention, error) {
	var mentions []*entity.Mention
	result := database.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&mentions)
	if result.Error != nil {
		return nil, result.Error
	}
	return mentions, nil
}
