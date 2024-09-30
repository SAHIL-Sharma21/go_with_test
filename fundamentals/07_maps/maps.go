package main

import (
	"fmt"
)

func main() {

	fmt.Println("maps with test")
}

// first serach function
// func Search(dict map[string]string, word string) string {
// 	return dict[word]
// }

// making custom type
type Dictinory map[string]string

// var ErrNotFound = errors.New("could not find the word you were lookig for")
type DictinoryErr string

func (e DictinoryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound         = DictinoryErr("could not find the word you were looking for")
	ErrWordExists       = DictinoryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictinoryErr("cannot update word because it does not exist")
)

func (d Dictinory) Search(word string) (string, error) {

	defination, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return defination, nil
}

// func (d Dictinory) Add(word, defination string) error {
// 	d[word] = defination
// 	return nil
// }

func (d Dictinory) Add(word, defnation string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = defnation
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictinory) Update(word, newDefination string) error {
	// d[word] = newDefination
	// return nil
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefination
	default:
		return err
	}
	return nil
}

// delete method
func (d Dictinory) Delete(word string) {
	delete(d, word)
}
