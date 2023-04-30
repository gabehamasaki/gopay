package services

import (
	"net/http"

	"github.com/gabehamasaki/gopay/internal/helpers"
	"github.com/gabehamasaki/gopay/internal/models"
	"gorm.io/gorm"
)

type AccountService struct {
	db *gorm.DB
}

func (s *AccountService) Create(a *models.Account) error {
	a.GenerateId()
	err := s.db.Create(a).Error
	if err != nil {
		return &helpers.HttpError{
			Message:    "Could not be able to create account",
			StatusCode: http.StatusInternalServerError,
			Op:         "account-create-service",
		}
	}
	return nil
}

func (s *AccountService) FindOne(id string, dest *models.Account) error {

	err := s.db.First(&dest, "id = ?", id).Error
	if err != nil {
		return &helpers.HttpError{
			Message:    "Could not be able to find account",
			StatusCode: http.StatusInternalServerError,
			Op:         "account-find-one-service",
		}
	}

	return nil
}

func (s *AccountService) FindMany(dest *[]models.Account) error {
	err := s.db.Find(&dest, &models.Account{}).Error
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

	if err := s.db.Save(&a).Error; err != nil {
		return &helpers.HttpError{
			Message:    "Could not be able to save this account",
			StatusCode: http.StatusInternalServerError,
			Op:         "account-save-service",
		}
	}

	return nil
}

func (s *AccountService) Delete(id string) error {

	acc := new(models.Account)

	if err := s.FindOne(id, acc); err != nil {
		return &helpers.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
			Op:         "account-delete-service",
		}
	}

	if err := s.db.Delete(&acc).Error; err != nil {
		return &helpers.HttpError{
			Message:    "Could not be able to delete this account",
			StatusCode: http.StatusInternalServerError,
			Op:         "account-delete-service",
		}
	}

	return nil
}
