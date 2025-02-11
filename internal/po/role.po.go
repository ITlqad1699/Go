package po

import (
	"gorm.io/gorm"
)

type ROLE struct {
	gorm.Model
	ID       int64  `gorm:"column:id;type:int;primaryKey;not null; autoIncrement:true; comment: 'Primary key is ID'"`
	RoleName string `gorm:"column:role_name;type:varchar(36);nullable"`
	RoleNote string `gorm:"column:role_note;type:text;nullable"`
}

func (r *ROLE) TableName() string {
	return "TBL_ROLE"
}
