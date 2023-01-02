package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Email         string    `gorm:"size:255;not null;unique" json:"email"`
	Username      string    `gorm:"size:255;not null;unique" json:"username"`
	Password      string    `gorm:"size:255;not null;" json:"-"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Deleted       gorm.DeletedAt
	DeletedStatus bool `gorm:"type:bool" json:"deletestatus"`
	StatusUser    bool `gorm:"type:bool" json:"status"`
	Entries       []Entries
}

func (User) TableName() string {
	return "users"
}
