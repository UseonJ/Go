package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

var (
	errorNotFound = errors.New("Not found")
	errCantUpdate = errors.New("Cant update non-existing word")
	errWordExists = errors.New("That Word already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if ok {
		return value, nil
	}
	return "", errorNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errorNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errorNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	delete(d, word)
	return nil
}
