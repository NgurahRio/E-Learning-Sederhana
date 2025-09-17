package models

import "gorm.io/gorm"

type User struct {
	IDUser       uint   `gorm:"primaryKey;column:id_user"`
	Name         string `gorm:"type:varchar(100);not null"`
	Email        string `gorm:"type:varchar(100);uniqueIndex;not null"`
	PasswordHash string `gorm:"type:varchar(255);not null"`
	RoleID       uint   `gorm:"not null"`
	PhotoURL     string `gorm:"type:varchar(255)"`

	Role Role `gorm:"foreignKey:RoleID"`
	gorm.Model
}
