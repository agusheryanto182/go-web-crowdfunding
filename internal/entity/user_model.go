package entity

import "time"

type UserModels struct {
	ID         int       `gorm:"column:id;type:INT;primaryKey" json:"id"`
	Name       string    `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	Occupation string    `gorm:"column:occupation;type:VARCHAR(255)" json:"occupation"`
	Email      string    `gorm:"column:email;type:VARCHAR(255)" json:"email"`
	Password   string    `gorm:"column:password;type:VARCHAR(255)" json:"password"`
	Avatar     string    `gorm:"column:avatar;type:VARCHAR(255)" json:"avatar"`
	Role       string    `gorm:"column:role;type:VARCHAR(255)" json:"role"`
	IsVerified bool      `gorm:"column:is_verified;default:false" json:"is_verified"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

type OTPModels struct {
	ID         uint       `gorm:"primaryKey;autoIncrement" json:"id" `
	UserID     int        `gorm:"index;unique" json:"user_id" `
	User       UserModels `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"user" `
	OTP        string     `gorm:"column:otp;type:varchar(255)" json:"otp"`
	ExpiredOTP int64      `gorm:"column:expired_otp;type:bigint" json:"expired_otp" `
}

func (UserModels) TableName() string {
	return "users"
}

func (OTPModels) TableName() string {
	return "otp"
}
