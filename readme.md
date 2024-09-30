#Golang with testing


### pointers and error
 - when you call a function or a method the arguments are copied.
 - to solve this we will use pointer in this as it will give give the original value and change the original not the copy one

 ### maps
 - Maps allow you to store items in a manner similar to a dictionary. 
 - ou can think of the key as the word and the value as the definition.
 ```go
 func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}
```
 - It can return 2 values. The second value is a boolean which indicates if the key was found successfully.
 - Go has a built-in function delete that works on maps. It takes two arguments. The first is the map and the second is the key to be removed.