package models

import "gorm.io/gorm"

type Course struct {
	IDCourse    uint   `gorm:"primaryKey;column:id_course"`
	Title       string `gorm:"type:varchar(150);not null"`
	Description string `gorm:"type:text"`
	TeacherID   uint   `gorm:"not null"`

	Teacher User `gorm:"foreignKey:TeacherID"`
	gorm.Model
}
