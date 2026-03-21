package dto

type UpdateBoardRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description,omitempty"`
}
