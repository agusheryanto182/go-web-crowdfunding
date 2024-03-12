package dto

import (
	"github.com/agusheryanto182/go-web-crowdfunding/internal/entity"
)

type ResponseImages struct {
	ID       int    `json:"id"`
	FileName string `json:"file_name"`
}

type SaveCampaignResponse struct {
	ID                  int              `json:"id"`
	UserID              int              `json:"user_id"`
	Name                string           `json:"name"`
	ShortDescription    string           `json:"short_description"`
	Description         string           `json:"description"`
	Perks               string           `json:"perks"`
	GoalAmount          int              `json:"goal_amount"`
	BackerCount         int              `json:"backer_count"`
	CurrentAmount       int              `json:"current_amount"`
	CampaignImageModels []ResponseImages `json:"campaign_image"`
}

func FormatSaveCampaignResponse(campaign *entity.CampaignModels) *SaveCampaignResponse {
	var imageFormatted []ResponseImages
	formatted := &SaveCampaignResponse{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		Perks:            campaign.Perks,
		GoalAmount:       campaign.GoalAmount,
		BackerCount:      campaign.BackerCount,
		CurrentAmount:    campaign.CurrentAmount,
	}
	for _, image := range campaign.CampaignImages {
		format := ResponseImages{
			ID:       image.ID,
			FileName: image.FileName,
		}
		imageFormatted = append(imageFormatted, format)
		formatted.CampaignImageModels = imageFormatted
	}

	return formatted
}

func FormatCampaignsResponse(campaign []*entity.CampaignModels) []*SaveCampaignResponse {
	var format []*SaveCampaignResponse
	for _, campaigns := range campaign {
		formatted := FormatSaveCampaignResponse(campaigns)
		format = append(format, formatted)
	}
	return format
}
