package models

import (
	"github.com/gabehamasaki/gopay/internal/dtos"
	"golang.org/x/crypto/bcrypt"
)

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

func (a *Account) EncryptPass() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(a.Pass), 14)
	a.Pass = string(bytes)
}

func (a *Account) CheckPass(pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Pass), []byte(pass))
	return err == nil
}

func (a *Account) Parse(request *dtos.CreateAccountRequestDTO) {
	a.Name = request.Name
	a.Email = request.Email
	a.Pass = request.Pass
	a.Type = request.Type
	a.Cpf = request.Cpf
	a.Balance = 0
	a.Activated = false

	a.EncryptPass()
}
