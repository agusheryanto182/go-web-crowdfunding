package entity

import (
	"time"
)

type TransactionModels struct {
	ID         int            `gorm:"column:id;type:int;primaryKey" json:"id"`
	CampaignID int            `gorm:"column:campaign_id;type:int" json:"campaign_id"`
	UserID     int            `gorm:"column:user_id" json:"user_id"`
	Amount     int            `gorm:"column:amount;type:int" json:"amount"`
	Status     string         `gorm:"column:status;type:VARCHAR(255)" json:"status"`
	Code       string         `gorm:"column:code;type:VARCHAR(255)" json:"code"`
	PaymentURL string         `gorm:"column:payment_url;type:VARCHAR(255)" json:"payment_url"`
	Campaigns  CampaignModels `gorm:"foreignKey:CampaignID" json:"campaigns"`
	Users      UserModels     `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"users"`
	CreatedAt  time.Time      `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}

func (TransactionModels) TableName() string {
	return "transactions"
}
