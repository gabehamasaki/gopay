package controllers

import (
	"net/http"

	"github.com/gabehamasaki/gopay/internal/dtos"
	"github.com/gabehamasaki/gopay/internal/helpers"
	"github.com/gabehamasaki/gopay/internal/models"
	"github.com/gabehamasaki/gopay/internal/services"
	usecases "github.com/gabehamasaki/gopay/internal/use-cases"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	Service services.Service[models.Account]
}

func (a *AccountController) Create(c *gin.Context) {
	request := &dtos.CreateAccountRequestDTO{}

	err := c.Bind(request)
	if err != nil {
		helpers.SendError(c, &helpers.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
			Op:         "account-controller",
		})
		return
	}

	err = request.Validate()
	if err != nil {
		helpers.SendError(c, err)
		return
	}
	var acc *models.Account
	acc, err = usecases.CreateAccount(request, a.Service)
	if err != nil {
		if err != nil {
			helpers.SendError(c, err)
			return
		}
	}

	helpers.SendSuccess(c, "account-controller", acc)
}

func (a *AccountController) Show(c *gin.Context) {

}

func (a *AccountController) Update(c *gin.Context) {

}

func (a *AccountController) Delete(c *gin.Context) {

}
