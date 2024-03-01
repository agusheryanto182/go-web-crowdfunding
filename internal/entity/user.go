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
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp DEFAULT CURRENT_TIMESTAMP" json:"updated_at"`
}
