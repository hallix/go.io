# go.io
go.io is a library that provides easy to use APIs and other Go constructs for performing CRUD operations on various persistance layers both storage and in memory.
## Supported persistance layers
1. Firestore
2. Redis(Coming soon)
3. MySQL(Coming soon)
4. Postgres(Coming soon)

### Examples
#### Firestore

1. Call your preferred persistance layer repository factory function which will return an implementation of the Repository interface:
##### Repository Interface
```
type Repository[T any] interface {
	Save(element T) (err error)
	GetById(id any, element *T) (data *T, err error)
  ...
}
```
##### Factory function signature
`func CreateRepository[T any](cxt context.Context, projectId, collectionName string) types.Repository[T]`
##### Factory function usage

`firestore.CreateRepository[Post](cxt, "hallix-blog", "Post")`

Code snippet:
```
package postrepository

import (
	"context"

	"github.com/hallix/go.io/pkg/firebase/firestore"
	"github.com/hallix/go.io/pkg/types"
)

var postRepository types.Repository[Post]

func GetRepository(cxt context.Context) types.Repository[Post] {

	if postRepository == nil {
		postRepository, _ = firestore.CreateRepository[Post](cxt, "hallix-blog", "Post")
	}

	return postRepository
}
```

2. Perform CRUD operation

##### Usage
```
postRepository = postrepository.GetRepository(cxt)
postRepository.Save(post)
```

Code Snippet: 
```
package postservice

import (
	"context"

	"github.com/hallix/blog-service/internal/repository/postrepository"
	"github.com/hallix/go.io/pkg/types"
)

var postRepository types.Repository[postrepository.Post]

type service struct{}

func GetService(cxt context.Context) PostService {

	if postRepository == nil {
		postRepository = postrepository.GetRepository(cxt)
	}

	return &service{}
}

func (s service) Save(post postrepository.Post) {

	postRepository.Save(post)
}
```
