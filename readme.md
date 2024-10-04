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

### Dependency injection
- The Buffer type from the bytes package implements the Writer interface, because it has the method Write(p []byte) (n int, err error).

### Concurrency
- concurrency is the process of multiple goroutines executing at the same time.
- he purposes of the following, means "having more than one thing in progress." This is something that we do naturally everyday.

Normally in Go when we call a function **doSomething()** we wait for it to return (even if it has no value to return, we still wait for it to finish). We say that this operation is blocking - it makes us wait for it to finish. An operation that does not block in Go will run in a separate process called a **goroutine**. 

#### channels

- A channel is a data structure that can send and receive values.

We can solve this data race by coordinating our goroutines using channels. Channels are a Go data structure that can both receive and send values. These operations, along with their details, allow communication between different processes.

#### Waitgroups
A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.

- We are using ```sync.WaitGroup``` which is a convenient way of synchronising concurrent processes.

-  ```Mutex``` allows us to add locks to our data

- ```WaitGroup``` is a means of waiting for goroutines to finish jobs