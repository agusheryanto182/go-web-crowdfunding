package entity

import (
	"time"

	"github.com/leekchan/accounting"
)

type CampaignModels struct {
	ID               int                   `gorm:"column:id;type:INT;primaryKey" json:"id"`
	UserID           int                   `gorm:"column:user_id" json:"user_id"`
	Name             string                `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	ShortDescription string                `gorm:"column:short_description;type:VARCHAR(255)" json:"short_description"`
	Description      string                `gorm:"column:description;type:TEXT" json:"description"`
	Perks            string                `gorm:"column:perks;type:VARCHAR(255)" json:"perks"`
	GoalAmount       int                   `gorm:"column:goal_amount;type:INT" json:"goal_amount"`
	BackerCount      int                   `gorm:"column:backer_count;type:INT" json:"backer_count"`
	CurrentAmount    int                   `gorm:"column:current_amount;type:INT" json:"current_amount"`
	CreatedAt        time.Time             `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time             `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
	CampaignImages   []CampaignImageModels `gorm:"foreignKey:CampaignID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"campaign_images"`
	User             UserModels            `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"users"`
}

func (CampaignModels) TableName() string {
	return "campaigns"
}

func (c CampaignModels) GoalAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(c.GoalAmount)
}

func (c CampaignModels) CurrentAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(c.CurrentAmount)
}

type CampaignImageModels struct {
	ID         int            `gorm:"column:id;type:INT;primaryKey" json:"id"`
	CampaignID int            `gorm:"column:campaign_id;type:int" json:"campaign_id"`
	Campaign   CampaignModels `gorm:"foreignKey:CampaignID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	FileName   string         `gorm:"column:file_name;type:VARCHAR(255)" json:"file_name"`
	IsPrimary  int            `gorm:"column:is_primary;type:int;default:0" json:"is_primary"`
	CreatedAt  time.Time      `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}

func (CampaignImageModels) TableName() string {
	return "campaign_images"
}
