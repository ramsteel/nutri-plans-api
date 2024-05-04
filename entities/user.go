package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	AuthID    uuid.UUID      `json:"id" gorm:"column:id;primaryKey;type:uuid"`
	Auth      Auth           `json:"auth"`
	FirstName string         `json:"first_name" gorm:"type:varchar(30)"`
	LastName  string         `json:"last_name" gorm:"type:varchar(30)"`
	Dob       time.Time      `json:"dob" gorm:"type:date"`
	Gender    string         `json:"gender" gorm:"type:char;size:1"`
	CountryID uint           `json:"-"`
	Country   Country        `json:"country"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
