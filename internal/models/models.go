package models

import (
	"time"

	"github.com/google/uuid"
)

type baseModel struct {
	Id        uuid.UUID `gorm:"primaryKey; not null" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *baseModel) GenerateId() {
	b.Id = uuid.New()
}
