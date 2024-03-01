package dto

import "time"

type UserResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Occupation     string    `json:"occupation"`
	Email          string    `json:"email"`
	AvatarFileName string    `json:"avatar_file_name,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
