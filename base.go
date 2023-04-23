package main

import "fmt"

type Animal interface {
	Sleep()
	Eat()
}

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Animal {
	p := &Person{
		Name: name,
		Age:  age,
	}

	fmt.Printf("NewPerson %p\n", p)
	return p
}

func (p *Person) Sleep() {
	fmt.Println("Hello, my name is", p.Name)
}

func (p *Person) Eat() {
	fmt.Println("I am eating")
}

func doWork(animal *Animal) {
	// print address of animal not in form &{PVC 23}
	fmt.Printf("doWork %p\n", animal)

	fmt.Printf("doWork %T\n", animal)
	(*animal).Sleep()
}

func main() {
	person := NewPerson("PVC", 23)
	// print address of person
	fmt.Println(&person)
	fmt.Printf("%p\n", &person)
	doWork(&person)
}
