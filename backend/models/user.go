package models

type User struct {
	IDUser   uint   `gorm:"primaryKey;column:id_user"`
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"` // ini dipakai utk hash password
	RoleID   uint   `gorm:"column:role_id;not null"`

	Role Role `gorm:"foreignKey:RoleID;references:IDRole"`
}

func (User) TableName() string {
	return "users"
}
