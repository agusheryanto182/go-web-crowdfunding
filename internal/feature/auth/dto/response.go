package dto

import (
	"time"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
)

type UserResponse struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Occupation string    `json:"occupation"`
	Email      string    `json:"email"`
	Avatar     string    `json:"avatar,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FormatCreateUserResponse(user *entity.UserModels) UserResponse {
	createdUser := UserResponse{}
	createdUser.ID = user.ID
	createdUser.Name = user.Name
	createdUser.Occupation = user.Occupation
	createdUser.Email = user.Email
	createdUser.Avatar = user.Avatar
	createdUser.CreatedAt = user.CreatedAt
	createdUser.UpdatedAt = user.UpdatedAt

	return createdUser
}

type VerifyOTPResponse struct {
	AccessToken string `json:"access_token"`
}
