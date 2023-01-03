package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Entries struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Content   string    `gorm:"type:text" json:"content"`
	UserID    uuid.UUID `json:"userId" gorm:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Deleted   bool `gorm:"type:bool" json:"deleted"`
}

func (Entries) TableName() string {
	return "entries"
}

type UpdateContent struct {
	Content string    `json:"content"`
	UserID  uuid.UUID `json:"UserId" gorm:"type:uuid"`
}
