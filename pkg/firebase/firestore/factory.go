package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/hallix/go.io/pkg/types"
)

type repository[T any] struct{}

var client *firestore.Client
var collection *firestore.CollectionRef
var cxt *context.Context

func CreateRepository[T any](appCxt context.Context, projectId, datastoreName string) types.Repository[T] {

	var err error
	client, err = firestore.NewClient(appCxt, projectId)

	if err != nil {
		panic(err)
	}

	collection = client.Collection(datastoreName)
	cxt = &appCxt

	return repository[T]{}
}
