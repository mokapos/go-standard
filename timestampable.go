package root

import (
	"time"
)

// Timestampable is an incomplete struct that should be included in any entity structs.
type Timestampable struct {
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}
