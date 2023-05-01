package usecases

import (
	"net/http"

	"github.com/gabehamasaki/gopay/internal/dtos"
	"github.com/gabehamasaki/gopay/internal/helpers"
	"github.com/gabehamasaki/gopay/internal/models"
	"github.com/gabehamasaki/gopay/internal/services"
)

func UpdateAccount(id string, request *dtos.UpdateAccountRequestDTO ,service services.Service[models.Account]) (*models.Account, error) {

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

	if request.Name != "" {
		account.Name = request.Name
	}
	if request.Cpf != "" {
		account.Cpf = request.Cpf
	}
	if request.Email != "" {
		account.Email = request.Email
	}
	if request.Pass != "" {
		account.Pass = request.Pass
		account.EncryptPass()
	}
	if request.Type != "" {
		account.Type = request.Type
	}
	if request.Activated != nil {
		account.Activated = *request.Activated
	}
	if request.Balance > 0.0 {
		account.Balance = request.Balance
	}

	err = service.Save(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}