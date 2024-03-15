package dto

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
)

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" validate:"required"`
	User entity.UserModels
}

type CreateTransactionInput struct {
	Amount     int `json:"amount" validate:"required"`
	CampaignID int `json:"campaign_id" validate:"required"`
	User       entity.UserModels
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
