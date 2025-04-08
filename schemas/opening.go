package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role        string
	Company     string
	Salary      int64
	Location    string
	Remote      bool
	Link        string
	Description string
}

type OpeningResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
	Role        string    `json:"role"`
	Company     string    `json:"company"`
	Salary      int64     `json:"salary"`
	Location    string    `json:"location"`
	Remote      bool      `json:"remote"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
}
