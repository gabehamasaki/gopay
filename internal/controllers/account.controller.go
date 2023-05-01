package controllers

import (
	"net/http"

	"github.com/gabehamasaki/gopay/config"
	"github.com/gabehamasaki/gopay/internal/dtos"
	"github.com/gabehamasaki/gopay/internal/helpers"
	"github.com/gabehamasaki/gopay/internal/models"
	"github.com/gabehamasaki/gopay/internal/services"
	usecases "github.com/gabehamasaki/gopay/internal/use-cases"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	Service services.Service[models.Account]
	Logger *config.Logger
}

func (a *AccountController) Create(c *gin.Context) {
	request := &dtos.CreateAccountRequestDTO{}

	if err := c.Bind(request); err != nil {
		a.Logger.Errorf("error creating account: %v", err)
		helpers.SendError(c, &helpers.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
			Op:         "create-account-controller",
		})
		return
	}

	if err := request.Validate(); err != nil {
		a.Logger.Errorf("error creating account: %v", err.Error())
		helpers.SendError(c, err)
		return
	}

	account, err := usecases.CreateAccount(request, a.Service)
	if err != nil {
			a.Logger.Errorf("error creating account: %v", err.Error())
			helpers.SendError(c, err)
			return
	}

	helpers.SendSuccess(c, "create-account-controller", account)
}

func (a *AccountController) Show(c *gin.Context) {

	id := c.Query("id")
	account, err := usecases.ShowAccount(id, a.Service)
	if err != nil {
		a.Logger.Errorf("error show account: %v", err)
		helpers.SendError(c, err)
		return
	}

	helpers.SendSuccess(c, "show-account-controller", account)
}

func (a *AccountController) Update(c *gin.Context) {

}

func (a *AccountController) Delete(c *gin.Context) {

}
