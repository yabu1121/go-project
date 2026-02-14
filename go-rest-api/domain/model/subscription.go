package model

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID               uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ServiceName      string    `json:"service_name" gorm:"not null"`
	MonthlyFee       int64     `json:"monthly_fee" gorm:"not null"`

	//毎月何日に引き落とされるか
	BillingCycleDays int       `json:"billing_cycle_days" gorm:"not null"`
	NextBillingDate  time.Time `json:"next_billing_date" gorm:"not null;index"`
	IsActive         bool      `json:"is_active" gorm:"not null;default:true"`

	UserId uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	User   User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`

	CategoryId uuid.UUID `json:"category_id" gorm:"type:uuid;not null;index"`
	Category   Category  `json:"category" gorm:"foreignKey:CategoryId; constraint:OnDelete:CASCADE"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
