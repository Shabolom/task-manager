package dto

import "time"

type TaskUpdateRequest struct {
	TaskID      int64      `json:"-"`
	BoardID     *int64     `json:"board_id,omitempty"`
	StatusID    *int64     `json:"status_id,omitempty"`
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	Position    *int       `json:"position,omitempty"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}
