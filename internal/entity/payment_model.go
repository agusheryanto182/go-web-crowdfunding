package entity

type PaymentModels struct {
	ID     int
	Amount int
}

func (PaymentModels) TableName() string {
	return "payments"
}
