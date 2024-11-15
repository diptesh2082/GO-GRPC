package models

import (
	"time"
	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	UserID       uint
	PlanName     string
	Price        float64
	StartDate    time.Time
	EndDate      time.Time
	IsActive     bool
}
