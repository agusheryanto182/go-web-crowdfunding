package payment

import "github.com/agusheryanto182/go-web-crowdfunding/internal/entity"

type PaymentServiceInterface interface {
	GetPaymentURL(transaction entity.TransactionModels, user entity.UserModels) (string, error)
}
