package repository

import (
	"TwClone/internal/database"
	"TwClone/internal/entity"
	"context"
)

type FollowRepositoryImpl struct{}

func (r FollowRepositoryImpl) Create(ctx context.Context, follow *entity.Follow) error {
	return database.DB.WithContext(ctx).Create(follow).Error
}

func (r FollowRepositoryImpl) Delete(ctx context.Context, followerID, followingID int64) error {
	return database.DB.WithContext(ctx).Where("follower_id = ? AND following_id = ?", followerID, followingID).Delete(&entity.Follow{}).Error
}

func (r FollowRepositoryImpl) FindFollowers(ctx context.Context, userID int64) ([]*entity.Follow, error) {
	var follows []*entity.Follow
	result := database.DB.WithContext(ctx).Where("following_id = ?", userID).Find(&follows)
	if result.Error != nil {
		return nil, result.Error
	}
	return follows, nil
}

func (r FollowRepositoryImpl) FindFollowing(ctx context.Context, userID int64) ([]*entity.Follow, error) {
	var follows []*entity.Follow
	result := database.DB.WithContext(ctx).Where("follower_id = ?", userID).Find(&follows)
	if result.Error != nil {
		return nil, result.Error
	}
	return follows, nil
}
