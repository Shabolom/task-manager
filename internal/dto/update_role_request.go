package dto

type UpdateRoleRequest struct {
	RoleID      int64  `json:"role_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
