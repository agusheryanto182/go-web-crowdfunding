package dto

type RegisterUserRequest struct {
	Name            string `json:"name" validate:"required,min=1,max=100"`
	Occupation      string `json:"occupation" validate:"required,max=100"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,eqfield=PasswordConfirm"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`
}

type VerifyOTPRequest struct {
	Email string `form:"email" json:"email" validate:"required,email"`
	OTP   string `form:"otp" json:"otp" validate:"required"`
}

type SignInUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"`
}
