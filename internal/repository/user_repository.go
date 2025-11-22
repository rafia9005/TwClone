package repository

// package repository

import (
	"context"
	"errors"
	"strings"

	"TWclone/internal/database"
	"TWclone/internal/entity"

	"gorm.io/gorm"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	// ErrDuplicate is returned when a unique constraint (duplicate key) is violated.
	ErrDuplicate = errors.New("duplicate key")
)

type UserRepositoryImpl struct{}

// Create inserts a new user.
func (r UserRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	result := database.DB.WithContext(ctx).Create(user)
	if result.Error != nil {
		errMsg := result.Error.Error()
		if strings.Contains(errMsg, "duplicate key") || strings.Contains(errMsg, "unique constraint") {
			// Try to detect which field is duplicated
			if strings.Contains(errMsg, "email") {
				return errors.New("duplicate_email")
			}
			if strings.Contains(errMsg, "username") {
				return errors.New("duplicate_username")
			}
			return ErrDuplicate
		}
		return result.Error
	}
	return nil
}

// Delete removes a user by id.
func (r UserRepositoryImpl) Delete(ctx context.Context, id int64) error {
	return database.DB.WithContext(ctx).Delete(&entity.User{}, id).Error
}

// FindAll returns all users.
func (r UserRepositoryImpl) FindAll(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	result := database.DB.WithContext(ctx).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// FindByEmail finds a user by email.
func (r UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	result := database.DB.WithContext(ctx).Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, result.Error
	}
	return &user, nil
}

// FindByUsername finds a user by username.
func (r UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	result := database.DB.WithContext(ctx).Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, result.Error
	}
	return &user, nil
}

// FindByID finds a user by id.
func (r UserRepositoryImpl) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	var user entity.User
	result := database.DB.WithContext(ctx).First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, result.Error
	}
	return &user, nil
}

// Update modifies an existing user.
func (r UserRepositoryImpl) Update(ctx context.Context, user *entity.User) error {
	return database.DB.WithContext(ctx).Save(user).Error
}
