package dto

import "TwClone/internal/entity"

// UserResponse is the API representation of a user (no password included).
type UserResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar,omitempty"`
	Banner    string `json:"banner,omitempty"`
	Bio       string `json:"bio,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// FromEntity converts an entity.User to UserResponse. Time formatting is RFC3339.
func FromEntity(u *entity.User) UserResponse {
	var createdAt, updatedAt string
	if !u.CreatedAt.IsZero() {
		createdAt = u.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	}
	if !u.UpdatedAt.IsZero() {
		updatedAt = u.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	}

	return UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		Username:  u.Username,
		Avatar:    u.Avatar,
		Banner:    u.Banner,
		Bio:       u.Bio,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
