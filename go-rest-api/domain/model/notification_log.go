package model

import (
	"time"

	"github.com/google/uuid"
)

type NotificationLog struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	// 何に関する通知かを文字列で持つ (例: "subscription", "public_fee")
	Type      string    `json:"type" gorm:"not null;index"`
	// 対象となるデータのID (Nullableにしておき、柔軟に)
	TargetID  *uuid.UUID `json:"target_id" gorm:"type:uuid"`
	
	SendAt    time.Time `json:"send_at" gorm:"autoCreateTime"`
	Status    string    `json:"status" gorm:"not null;default:'sent'"`

	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	User      User      `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}