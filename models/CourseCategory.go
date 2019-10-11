package models

import "time"

type CourseCategory struct {
	ID        uint       `gorm:"size:11;primary_key" json:"id"`
	Name      string     `gorm:"size:100" json:"name"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
