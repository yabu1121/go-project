package model

import "github.com/google/uuid"

type Report struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TargetMonth string `json:"target_month" gorm:"not null"`
	TotalExpense int64 `json:"total_expense" gorm:"not null"`
	TotalIncome int64 `json:"total_income" gorm:"not null"`
	TotalSavings int64 `json:"total_savings" gorm:"not null"`
	CategoryBreakdown map[string]int64 `json:"category_breakdown" gorm:"type:jsonb"`

	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	User User `json:"user" gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE"`
}