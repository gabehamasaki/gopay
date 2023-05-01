package dtos

import (
	"net/http"

	"github.com/gabehamasaki/gopay/internal/helpers"
)

type UpdateAccountRequestDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
	Pass  string `json:"pass"`
	Type  string `json:"type"`
	Balance   float64 `json:"balance"`
	Activated *bool    `json:"activated"`
}

func (a UpdateAccountRequestDTO) Validate() error {

	if a.Name != "" || a.Email != "" || a.Pass != "" || a.Type != "" || a.Cpf != "" || a.Balance > 0.0 || 
	a.Activated != nil {
		return nil
	}

	return &helpers.HttpError{
		StatusCode: http.StatusBadRequest,
		Message:    "request body is empty or malformed",
		Op:         "update-account-validation",
	}
}
