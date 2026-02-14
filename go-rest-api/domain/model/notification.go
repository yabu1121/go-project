package model

import "github.com/google/uuid"

type Notification struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`	
	EnableSubscriptionNotification bool `json:"enable_subscription_notification" gorm:"not null;default:true"`
	EnablePublicFeeNotification bool `json:"enable_public_fee_notification" gorm:"not null;default:true"`
	ReminderDay int `json:"reminder_day" gorm:"not null;default:1"`

	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	User User `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
}