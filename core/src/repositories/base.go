package repositories

import "github.com/google/uuid"

type Base[T any] interface {
	GetAll() ([]T, error)
	GetByID(ID uuid.UUID) (T, error)
	Create(data T) (T, error)
	Update(ID uuid.UUID, data T) (T, error)
	Delete(ID uuid.UUID) error
	CheckExist(ID uuid.UUID) (bool, error)
}
