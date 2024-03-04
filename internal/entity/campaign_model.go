package entity

import (
	"time"

	"github.com/leekchan/accounting"
)

type CampaignModels struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImageModels
	User             UserModels
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
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (CampaignModels) TableName() string {
	return "campaigns"
}

func (CampaignImageModels) TableName() string {
	return "campaign_images"
}
