package model

import (
	"time"

	"github.com/google/uuid"
)

type NotificationLog struct {
	ID                 uuid.UUID    `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	SubscriptionRemind bool         `json:"subscription_remind" gorm:"not null;default:false"`
	PublicFeeRemind    bool         `json:"public_fee_remind" gorm:"not null;default:false"`
	SendAt             time.Time    `json:"send_at" gorm:"autoCreateTime"`
	Status             string       `json:"status" gorm:"not null;default:'send'"`

	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	User   User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
}