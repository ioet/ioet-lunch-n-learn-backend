package repositories

type Base[T any] interface {
	GetAll() ([]T, error)
	GetByID(ID string) (T, error)
	Create(data T) (T, error)
	Update(ID string, data T) (T, error)
	Delete(ID string) error
}
