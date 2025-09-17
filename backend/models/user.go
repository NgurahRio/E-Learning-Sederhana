package models

type User struct {
	IDUser   uint   `gorm:"primaryKey;column:id_user"`
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	RoleID   uint   `gorm:"column:role_id;not null"`

	Role    Role     `gorm:"foreignKey:RoleID;references:IDRole"`
	Courses []Course `gorm:"foreignKey:TeacherID;references:IDUser" json:"-"` // relasi teacher → courses
}
