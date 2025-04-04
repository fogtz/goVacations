package schemas

import (
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
