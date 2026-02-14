package model

import (
	"github.com/google/uuid"
)

type Category struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name string    `json:"name" gorm:"not null"`

	// *にしたらnullableってこと
	UserID *uuid.UUID `json:"user_id" gorm:"type:uuid;index"`
	User   User       `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
}
