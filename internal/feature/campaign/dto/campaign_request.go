package dto

type CreateRequestCampaign struct {
	UserID           int    `json:"user_id"`
	Name             string `json:"name" validate:"required"`
	ShortDescription string `json:"short_description" validate:"required"`
	Description      string `json:"description" validate:"required"`
	Perks            string `json:"perks" validate:"required"`
	GoalAmount       int    `json:"goal_amount" validate:"required"`
}

type UpdateRequestCampaign struct {
	ID               int
	Name             string `json:"name" validate:"required"`
	ShortDescription string `json:"short_description" validate:"required"`
	Description      string `json:"description" validate:"required"`
	Perks            string `json:"perks" validate:"required"`
	GoalAmount       int    `json:"goal_amount" validate:"required"`
}

type CreateRequestCampaignImage struct {
	CampaignID int    `json:"campaign_id"`
	FileName   string `json:"file_name"`
	IsPrimary  bool   `json:"is_primary"`
}

type SetPrimaryImageRequest struct {
	ID         int  `json:"id"`
	CampaignID int  `json:"campaign_id"`
	IsPrimary  bool `json:"is_primary"`
}
