package usecases

import (
	"fmt"
	"net/http"

	"github.com/gabehamasaki/gopay/internal/dtos"
	"github.com/gabehamasaki/gopay/internal/helpers"
	"github.com/gabehamasaki/gopay/internal/models"
	"github.com/gabehamasaki/gopay/internal/services"
)

func CreateAccount(request *dtos.CreateAccountRequestDTO, service services.Service[models.Account]) (*models.Account, error) {

	if exist, _ := service.FindByEmail(request.Email); exist != nil {
		return nil, &helpers.HttpError{
			Message:    "an account with that email already exists",
			Op:         "account-use-cases",
			StatusCode: http.StatusFound,
		}
	}

	if exist, _ := service.FindByCpf(request.Cpf); exist != nil {
		return nil, &helpers.HttpError{
			Message:    "an account with that cpf already exists",
			Op:         "account-use-cases",
			StatusCode: http.StatusFound,
		}
	}

	account := &models.Account{}
	account.Parse(request)
	account.GenerateId()

	err := service.Create(account)
	if err != nil {
		re, _ := err.(*helpers.HttpError)
		return nil, &helpers.HttpError{
			Message:    re.Message,
			Op:         fmt.Sprintf("account-use-cases: %s", re.Op),
			StatusCode: re.StatusCode,
		}
	}

	return account, nil
}
