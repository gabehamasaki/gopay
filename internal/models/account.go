package models

import "github.com/gabehamasaki/gopay/internal/dtos"

type Account struct {
	BaseModel
	Name      string  `gorm:"not null" json:"name"`
	Email     string  `gorm:"not null;unique" json:"email"`
	Cpf       string  `gorm:"not null;unique" json:"cpf"`
	Pass      string  `gorm:"not null" json:"pass"`
	Type      string  `gorm:"default:'normal'" json:"type"`
	Balance   float64 `gorm:"default:0" json:"balance"`
	Activated bool    `gorm:"not null,default:false" json:"activated"`
}

func (a *Account) Parse(request *dtos.CreateAccountRequestDTO) {
	a.Name = request.Name
	a.Email = request.Email
	a.Pass = request.Pass
	a.Type = request.Type
	a.Cpf = request.Cpf
	a.Balance = 0
	a.Activated = false
}
