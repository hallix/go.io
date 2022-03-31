package firestore

import "fmt"

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
