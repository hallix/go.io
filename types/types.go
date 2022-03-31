package types

type Repository[T any] interface {
	Save(element T)
	GetById(id any, element *T) (data *T, err error)
}
