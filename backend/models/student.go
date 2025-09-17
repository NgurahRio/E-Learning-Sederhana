package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name    string   `json:"name"`
	Email   string   `gorm:"uniqueIndex" json:"email"`
	Courses []Course `gorm:"many2many:student_courses;" json:"courses,omitempty"`
}
