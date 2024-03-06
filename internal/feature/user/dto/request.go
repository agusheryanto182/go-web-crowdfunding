package dto

type UpdateUserRequest struct {
	Name       string `form:"name" json:"name" validate:"required,min=1,max=100"`
	Occupation string `form:"occupation" json:"occupation" validate:"required, max=100"`
	Email      string `form:"email" json:"email" validate:"required, email"`
	Password   string `form:"name" json:"password" validate:"required"`
}

type UpdateAvatarRequest struct {
	Avatar string `form:"avatar"`
}
