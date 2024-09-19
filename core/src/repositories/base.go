package repositories

import (
	"github.com/ioet/ioet-lunch-n-learn-backend/core/src/models/user"
)

type Base[T any] interface {
	GetAll() ([]T, error)
	GetByID(ID string) (T, error)
	Create(data T) (T, error)
	Update(ID string, data T) (T, error)
	Delete(ID string) error
}

type BaseUserRepository interface {
	GetAll() ([]user.User, error)
	GetByID(ID string) (user.User, error)
	Create(data user.User) (user.User, error)
	Update(ID string, data user.User) (user.User, error)
	Delete(ID string) error
	GetAllByHouseID(houseID string) ([]user.User, error)
}
