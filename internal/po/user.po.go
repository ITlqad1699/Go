package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid;type:varchar(255);primaryKey;not null"`
	UserName string    `gorm:"column:user_name;type:varchar(255);nullable"`
	Email    string    `gorm:"column:email;type:varchar(255);nullable"`
	IsActive bool      `gorm:"column:is_active;type:boolean"`
	Roles    []ROLE    `gorm:"many2many:user_roles;"`
}

func (u *User) TableName() string {
	return "TBL_USER"
}
