package services

import (
	"net/http"

	"github.com/gabehamasaki/gopay/internal/helpers"
	"github.com/gabehamasaki/gopay/internal/models"
)

type AccountService[T any] struct {
}

func (s *AccountService[T]) Create(a *models.Account) error {
	a.GenerateId()
	err := db.Create(a).Error
	if err != nil {
		return &helpers.HttpError{
			Message:    "Could not be able to create account",
			StatusCode: http.StatusInternalServerError,
			Op:         "account-create-service",
		}
	}
	return nil
}

func (s *AccountService[T]) FindOne(id string, dest *models.Account) error {

	err := db.First(&dest, "id = ?", id).Error
	if err != nil {
		return &helpers.HttpError{
			Message:    "Could not be able to find account",
			StatusCode: http.StatusInternalServerError,
			Op:         "account-find-one-service",
		}
	}

	return nil
}

func (s *AccountService[T]) FindMany(dest *[]models.Account) error {
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

func (s *AccountService[T]) Save(a *models.Account) error {

	if err := db.Save(&a).Error; err != nil {
		return &helpers.HttpError{
			Message:    "Could not be able to save this account",
			StatusCode: http.StatusInternalServerError,
			Op:         "account-save-service",
		}
	}

	return nil
}

func (s *AccountService[T]) Delete(id string) error {

	acc := new(models.Account)

	if err := s.FindOne(id, acc); err != nil {
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
