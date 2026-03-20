package models

import "time"

type User struct {
	ID        uint       `json:"id" json:",omitempty"`
	UUID      string     `json:"uuid" json:",omitempty"`
	CreatedAt time.Time  `json:"created_at" json:",omitempty"` // Adicione isso
	UpdatedAt time.Time  `json:"updated_at" json:",omitempty"` // Adicione isso
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
