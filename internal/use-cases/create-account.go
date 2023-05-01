package usecases

import (
	"fmt"

	"github.com/gabehamasaki/gopay/internal/dtos"
	"github.com/gabehamasaki/gopay/internal/helpers"
	"github.com/gabehamasaki/gopay/internal/models"
	"github.com/gabehamasaki/gopay/internal/services"
)

func CreateAccount(request *dtos.CreateAccountRequestDTO, service services.Service[models.Account]) (*models.Account, error) {


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
