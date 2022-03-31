package types

type Repository[T any] interface {
	Save(element T) (err error)
	GetById(id any, element *T) (data *T, err error)
}
