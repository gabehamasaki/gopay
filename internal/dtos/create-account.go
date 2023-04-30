package dtos

import (
	"net/http"

	"github.com/gabehamasaki/gopay/internal/helpers"
)

type CreateAccountRequestDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
	Pass  string `json:"pass"`
	Type  string `json:"type"`
}

func (a CreateAccountRequestDTO) Validate() error {

	if a.Name == "" && a.Email == "" && a.Pass == "" && a.Type == "" && a.Cpf == "" {
		return &helpers.HttpError{
			StatusCode: http.StatusBadRequest,
			Message:    "request body is empty or malformed",
			Op:         "create-account-validation",
		}
	}

	if a.Name == "" {
		return &helpers.HttpError{
			StatusCode: http.StatusBadRequest,
			Message:    "name is empty or malformed",
			Op:         "create-account-validation",
		}
	}

	if a.Cpf == "" {
		return &helpers.HttpError{
			StatusCode: http.StatusBadRequest,
			Message:    "cpf is empty or malformed",
			Op:         "create-account-validation",
		}
	}

	if a.Email == "" {
		return &helpers.HttpError{
			StatusCode: http.StatusBadRequest,
			Message:    "email is empty or malformed",
			Op:         "create-account-validation",
		}
	}

	if a.Pass == "" {
		return &helpers.HttpError{
			StatusCode: http.StatusBadRequest,
			Message:    "password is empty or malformed",
			Op:         "create-account-validation",
		}
	}

	if a.Type == "" {
		return &helpers.HttpError{
			StatusCode: http.StatusBadRequest,
			Message:    "type is empty or malformed",
			Op:         "create-account-validation",
		}
	}

	return nil
}
