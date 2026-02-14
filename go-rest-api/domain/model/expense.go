package model

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Amount      int64     `json:"amount" gorm:"not null;check:amount > 0"`
	Description string    `json:"description" gorm:"type:text"`
	SpentAt     time.Time `json:"spent_at" gorm:"not null;index"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`

	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	User   User      `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`

	CategoryID uuid.UUID `json:"category_id" gorm:"type:uuid;not null;index"`
	Category   Category  `json:"category" gorm:"foreignKey:CategoryID; constraint:OnDelete:CASCADE"`
}
