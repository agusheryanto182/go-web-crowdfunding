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

type UploadImageFormatter struct {
	Avatar string `json:"avatar"`
}

func UpdateAvatarResponse(user *entity.UserModels) UploadImageFormatter {
	response := UploadImageFormatter{}
	response.Avatar = user.Avatar
	return response
}

func FormatterUser(user *entity.UserModels) *UserResponse {
	format := &UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Avatar:     user.Avatar,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
	return format
}

func FormatterUsers(users []*entity.UserModels) []*UserResponse {
	var format []*UserResponse
	for _, user := range users {
		formattedUsers := FormatterUser(user)
		format = append(format, formattedUsers)
	}
	return format
}
