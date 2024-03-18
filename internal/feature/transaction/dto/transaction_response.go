package dto

import (
	"time"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
)

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Campaign  string    `json:"campaign"`
	Amount    int       `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction *entity.TransactionModels) *CampaignTransactionFormatter {
	formatter := &CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.Users.Name
	formatter.Campaign = transaction.Campaigns.Name
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}

func FormatCampaignTransactions(transactions []*entity.TransactionModels) []*CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return nil
	}

	var transactionsFormatter []*CampaignTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}
	return transactionsFormatter
}

type TransactionFormatter struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

func FormatTransaction(transaction *entity.TransactionModels) TransactionFormatter {
	formatter := TransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.CampaignID = transaction.CampaignID
	formatter.UserID = transaction.UserID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.Code = transaction.Code
	formatter.PaymentURL = transaction.PaymentURL
	return formatter
}
