package models

type Role struct {
	IDRole   uint   `gorm:"primaryKey;column:id_role"`
	RoleName string `gorm:"type:varchar(50);unique;not null"`

	Users []User `gorm:"foreignKey:RoleID;references:IDRole"`
}

func (Role) TableName() string {
	return "roles"
}
