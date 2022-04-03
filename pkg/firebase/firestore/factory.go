package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/hallix/go.io/pkg/types"
)

type repository[T any] struct {
	client     *firestore.Client
	collection *firestore.CollectionRef
	context    *context.Context
}

func CreateRepository[T any](context context.Context, projectId, datastoreName string) (types.Repository[T], error) {

	client, err := firestore.NewClient(context, projectId)

	if err != nil {
		return nil, err
	}

	collection := client.Collection(datastoreName)

	return &repository[T]{client, collection, &context}, nil
}
