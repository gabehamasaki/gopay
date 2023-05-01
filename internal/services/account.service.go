package services

import (
	"net/http"
	"strings"

	"github.com/gabehamasaki/gopay/internal/helpers"
	"github.com/gabehamasaki/gopay/internal/models"
)

type AccountService struct {
}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (s *AccountService) Create(a *models.Account) error {
	err := db.Create(a).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return &helpers.HttpError{
				Message:    "cpf or email already exists",
				StatusCode: http.StatusFound,
				Op:         "account-create-service",
			}
		}
		return &helpers.HttpError{
			Message:    "Could not be able to create account",
			StatusCode: http.StatusInternalServerError,
			Op:         "account-create-service",
		}
	}
	return nil
}

func (s *AccountService) FindOne(id string) (*models.Account, error) {
	dest := &models.Account{}

	err := db.First(&dest, "id = ?", id).Error
	if err != nil {
		return nil, &helpers.HttpError{
			Message:    "Could not be able to find account",
			StatusCode: http.StatusInternalServerError,
			Op:         "account-find-one-service",
		}
	}

	return dest, nil
}

func (s *AccountService) FindByEmail(email string) (*models.Account, error) {

	dest := &models.Account{}

	err := db.First(&dest, "email = ?", email).Error
	if err != nil {
		return nil, &helpers.HttpError{
			Message:    "Could not be able to find account",
			StatusCode: http.StatusInternalServerError,
			Op:         "account-find-one-service",
		}
	}

	return dest, nil
}

func (s *AccountService) FindByCpf(cpf string) (*models.Account, error) {
	dest := &models.Account{}

	err := db.First(&dest, "cpf = ?", cpf).Error
	if err != nil {
		return nil, &helpers.HttpError{
			Message:    "Could not be able to find account",
			StatusCode: http.StatusInternalServerError,
			Op:         "account-find-one-service",
		}
	}

	return dest, nil
}

func (s *AccountService) FindMany(dest *[]models.Account) error {
	err := db.Find(&dest, &models.Account{}).Error
	if err != nil {
		return &helpers.HttpError{
			Message:    "Could not be able to find all account's",
			StatusCode: http.StatusInternalServerError,
			Op:         "account-find-many-service",
		}
	}

	return nil
}

func (s *AccountService) Save(a *models.Account) error {

	if err := db.Save(&a).Error; err != nil {
		return &helpers.HttpError{
			Message:    "Could not be able to save this account",
			StatusCode: http.StatusInternalServerError,
			Op:         "account-save-service",
		}
	}

	return nil
}

func (s *AccountService) Delete(id string) error {

	acc, err := s.FindOne(id)
	if err != nil {
		return &helpers.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
			Op:         "account-delete-service",
		}
	}

	if err := db.Delete(&acc).Error; err != nil {
		return &helpers.HttpError{
			Message:    "Could not be able to delete this account",
			StatusCode: http.StatusInternalServerError,
			Op:         "account-delete-service",
		}
	}

	return nil
}
