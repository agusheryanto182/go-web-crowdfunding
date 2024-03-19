package entity

import (
	"time"
)

type AssistantModel struct {
	ID        uint64     `gorm:"column:id;type:int;primaryKey" json:"id"`
	UserID    uint64     `json:"user_id" form:"user_id"`
	Role      string     `json:"role" form:"role"`
	Text      string     `json:"text" form:"text"`
	Users     UserModels `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"users"`
	CreatedAt time.Time  `json:"created_at" form:"created_at"`
}
