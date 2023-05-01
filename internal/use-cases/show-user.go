package usecases

import (
	"net/http"

	"github.com/gabehamasaki/gopay/internal/helpers"
	"github.com/gabehamasaki/gopay/internal/models"
	"github.com/gabehamasaki/gopay/internal/services"
)

func ShowAccount(id string, service services.Service[models.Account]) (*models.Account, error) {

	if id == "" {
		return nil, &helpers.HttpError{
			Message: "query parameters id is required",
			StatusCode: http.StatusBadRequest,
			Op: "show-account-controller",
		}
	}

	account, err := service.FindOne(id)
	if err != nil {
		return nil, err
	}

	return account, nil
}