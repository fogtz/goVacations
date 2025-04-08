package handler

import (
	"fmt"

	"github.com/fogtz/goVacations.git/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

type CreateVacancyRequest struct {
	Role        string `json:"role"`
	Company     string `json:"company"`
	Salary      int64  `json:"salary"`
	Location    string `json:"location"`
	Remote      *bool  `json:"remote"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

func paramRequired(param string, reqType string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, reqType)
}

func (request *CreateVacancyRequest) Validate() error {
	if request.Role == "" && request.Company == "" && request.Salary <= 0 && request.Location == "" && request.Link == "" && request.Description == "" && request.Remote == nil {
		return fmt.Errorf("malformed request: all params are required")
	}
	if request.Role == "" {
		return paramRequired("role", "string")
	}
	if request.Company == "" {
		return paramRequired("company", "string")
	}
	if request.Salary <= 0 {
		return paramRequired("salary", "int")
	}
	if request.Location == "" {
		return paramRequired("location", "string")
	}
	if request.Link == "" {
		return paramRequired("link", "string")
	}
	if request.Description == "" {
		return paramRequired("description", "string")
	}
	if request.Remote == nil {
		return paramRequired("remote", "bool")
	}
	return nil
}

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetDB()
}
