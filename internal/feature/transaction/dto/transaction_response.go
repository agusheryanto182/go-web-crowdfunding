package dto

import (
	"time"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
)

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction *entity.TransactionModels) *CampaignTransactionFormatter {
	formatter := &CampaignTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.Users.Name
	formatter.Amount = transaction.Amount
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
