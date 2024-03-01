package entity

import (
	"os/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	Campaign   Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       user.User
}
