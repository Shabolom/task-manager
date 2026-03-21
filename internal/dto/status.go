package dto

import "time"

type Status struct {
	ID        int64      `db:"id" json:"id"`
	BoardID   int64      `db:"board_id" json:"board_id"`
	Name      string     `db:"name" json:"name"`
	Color     string     `db:"color" json:"color"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
