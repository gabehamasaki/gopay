package services

import (
	"github.com/gabehamasaki/gopay/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Service[T any] interface {
	Create(t *T) error
	FindOne(id string, dest *T) error
	FindMany(dest *[]T) error
	Save(t *T) error
	Delete(id string) error
}

func Initialize() {
	db = config.GetDatabase()

}
