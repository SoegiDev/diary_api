package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoles struct {
	UserId    uuid.UUID `gorm:"primaryKey"`
	RoleId    uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (UserRoles) TableName() string {
	return "user_roles"
}
