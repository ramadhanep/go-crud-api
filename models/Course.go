package models

import "time"

type Course struct {
	ID               uint           `gorm:"size:100;primary_key" json:"id"`
	CourseCategoryID uint           `gorm:"foreignKey:ID" json:"course_category_id"`
	Name             string         `gorm:"size:100" json:"name"`
	Description      string         `gorm:"size:100;type:text" json:"description"`
	PricePerHour     int            `gorm:"size:100" json:"price_per_hour"`
	Avatar           string         `gorm:"type:text" json:"avatar"`
	CreatedAt        time.Time      `json:"-"`
	UpdatedAt        time.Time      `json:"-"`
	DeletedAt        *time.Time     `json:"-"`
	CourseCategory   CourseCategory `json:"course_category"`
}
