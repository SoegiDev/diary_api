package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Deleted   bool `gorm:"type:bool;default:false" json:"deleted"`
	IsActive  bool `gorm:"type:bool;default:false" json:"IsActive"`
}

type InputRole struct {
	Name string `gorm:"size:20" json:"name"`
}

func (Role) TableName() string {
	return "roles"
}
