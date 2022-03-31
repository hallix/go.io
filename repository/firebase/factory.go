package firebase

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/hallix/go.io/types"
)

type repository[T any] struct{}

var client *firestore.Client
var collection *firestore.CollectionRef
var cxt *context.Context

func CreateRepository[T any](appCxt context.Context, projectId, datastoreName string, entity T) types.Repository[T] {

	var err error
	client, err = firestore.NewClient(appCxt, projectId)

	if err != nil {
		fmt.Print(err)
		panic(15)
	}

	collection = client.Collection(datastoreName)
	cxt = &appCxt

	return repository[T]{}
}

func (repo repository[T]) Save(element T) (err error) {

	_, _, err = collection.Add(*cxt, element)

	return err

}

func (repo repository[T]) GetById(id any, element *T) (data *T, err error) {

	stringId := fmt.Sprint(id)

	docSnapShot, err := collection.Doc(stringId).Get(*cxt)

	if err != nil {
		panic(err)
	}

	err = docSnapShot.DataTo(element)

	return element, err
}
