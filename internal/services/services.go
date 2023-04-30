package services

type Service[T interface{}] interface {
	Create(t *T) error
	FindOne(id string, dest *T) error
	FindMany(dest []*T) error
	Save(t *T) error
	Delete(id string) error
}
