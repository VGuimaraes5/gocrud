package model

import (
	"time"
)

type User struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deteled_at,omitempty"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Role      *string    `json:"role"`
}
