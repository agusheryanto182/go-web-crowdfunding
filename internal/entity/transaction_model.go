package entity

import (
	"os/user"
	"time"
)

type TransactionModels struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	Campaign   CampaignModels
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       user.User
}

func (TransactionModels) TableName() string {
	return "transactions"
}
