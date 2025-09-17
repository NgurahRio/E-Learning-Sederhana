package models

import "gorm.io/gorm"

type StudentCourse struct {
	IDStudentCourse uint `gorm:"primaryKey;column:id_student_course"`
	StudentID       uint `gorm:"not null"`
	CourseID        uint `gorm:"not null"`

	Student User   `gorm:"foreignKey:StudentID"`
	Course  Course `gorm:"foreignKey:CourseID"`

	gorm.Model
}
