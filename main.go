package main

import "fmt"

type Person struct {
	Name string
	age  int16
}

type PersonBuilder struct {
	Person
}

func (p *PersonBuilder) setAge(age int16) *PersonBuilder {
	p.age = age
	return p
}

func (p *PersonBuilder) setName(name string) *PersonBuilder {
	p.Name = name
	return p
}
func (p *PersonBuilder) Build() Person {
	return p.Person
}
func main() {

	person := Person{Name: "ali", age: 20}
	builder := PersonBuilder{Person: person}

	fmt.Println(builder)
	fmt.Println(builder.setAge(52).setName("meiti").Build())

}

// this was builder example
