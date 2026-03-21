package dto

import "time"

type Task struct {
	ID          int64      `db:"id" json:"id"`
	BoardID     int64      `db:"board_id" json:"board_id"`
	StatusID    int64      `db:"status_id" json:"status_id"`
	Title       string     `db:"title" json:"title"`
	Description string     `db:"description" json:"description"`
	Priority    string     `db:"priority" json:"priority"`
	DueDate     *time.Time `db:"due_date" json:"due_date,omitempty"`
	CreatorId   int64      `db:"creator_id" json:"creator_id,omitempty"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
