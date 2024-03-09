package dto

import (
	"time"

	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
)

type SaveCampaignResponse struct {
	ID               int       `json:"id"`
	UserID           int       `json:"user_id"`
	Name             string    `json:"name"`
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	Perks            string    `json:"perks"`
	GoalAmount       int       `json:"goal_amount"`
	BackerCount      int       `json:"backer_count"`
	CurrentAmount    int       `json:"current_amount"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func FormatSaveCampaignResponse(campaign *entity.CampaignModels) *SaveCampaignResponse {
	formatted := &SaveCampaignResponse{
		ID:               campaign.ID,
		UserID:           campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		Perks:            campaign.Perks,
		GoalAmount:       campaign.GoalAmount,
		BackerCount:      campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		CreatedAt:        campaign.CreatedAt,
		UpdatedAt:        campaign.UpdatedAt,
	}
	return formatted
}
