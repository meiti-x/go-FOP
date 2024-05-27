package main

import "fmt"

type Person struct {
	Name string
	age  int16
}

type PersonOptions func(person *Person)

func SetName(name string) PersonOptions {
	return func(person *Person) {
		person.Name = name
	}
}

func SetAge(age int16) PersonOptions {
	return func(person *Person) {
		person.age = age
	}
}

func newPerson(opts ...PersonOptions) *Person {
	person := &Person{}
	for _, opt := range opts {
		opt(person)
	}
	return person
}
func main() {
	person := newPerson(SetAge(13), SetName("abas"))
	fmt.Println(person)
}

// example with FOP
