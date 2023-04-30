package models

type Account struct {
	baseModel
	Name      string  `gorm:"not null" json:"name"`
	Email     string  `gorm:"not null" json:"email"`
	Cpf       string  `gorm:"not null" json:"cpf"`
	Pass      string  `gorm:"not null" json:"pass"`
	Type      string  `gorm:"default:'normal'" json:"type"`
	Balance   float64 `gorm:"default:0" json:"balance"`
	Activated bool    `gorm:"not null" json:"activated"`
}
