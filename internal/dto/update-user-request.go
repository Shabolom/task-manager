package dto

import "time"

type UpdateUserRequest struct {
	RoleID      *int64     `json:"role_id,omitempty"`
	Email       *string    `json:"email,omitempty"`
	FullName    *string    `json:"full_name,omitempty"`
	IsActive    *bool      `json:"is_active,omitempty"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
}

type UpdateUserPasswordRequest struct {
	Password string `json:"password" validate:"required,min=6"`
}
