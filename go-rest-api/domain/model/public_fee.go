package model

import (
	"time"
	
	"github.com/google/uuid"
)

type PublicFee struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ServiceName string `json:"service_name" gorm:"not null"`
	FeeType string `json:"fee_type" gorm:"not null"`
	MonthlyFee int64 `json:"monthly_fee" gorm:"not null"`
	UsageMonth string `json:"usage_month" gorm:"not null"`
	BillingDate time.Time `json:"billing_date" gorm:"not null"`

	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	User User `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`

	CategoryID uuid.UUID `json:"category_id" gorm:"type:uuid;not null;index"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID; constraint:OnDelete:CASCADE"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}