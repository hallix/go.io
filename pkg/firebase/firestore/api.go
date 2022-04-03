package firestore

import "fmt"

func (repo repository[T]) Save(element T) (err error) {

	_, _, err = repo.collection.Add(*repo.context, element)

	return err

}

func (repo repository[T]) GetById(id any, element *T) (data *T, err error) {

	stringId := fmt.Sprint(id)

	docSnapShot, err := repo.collection.Doc(stringId).Get(*repo.context)

	if err != nil {
		return element, err
	}

	err = docSnapShot.DataTo(element)

	return element, err
}
