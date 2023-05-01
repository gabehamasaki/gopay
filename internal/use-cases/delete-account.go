package usecases

import (
	"net/http"

	"github.com/gabehamasaki/gopay/internal/helpers"
	"github.com/gabehamasaki/gopay/internal/models"
	"github.com/gabehamasaki/gopay/internal/services"
)

func DeleteAccount(id string, service services.Service[models.Account]) error {

	if id == "" {
		return &helpers.HttpError{
			Message: "query parameters id is required",
			StatusCode: http.StatusBadRequest,
			Op: "show-account-controller",
		}
	}

	err := service.Delete(id)
	if err != nil {
		return err
	}

	return nil
}