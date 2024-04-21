package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// firstname string
	// lastname  string
	// gender    string
	// age       int

	firstname, lastname, gender string
	age                         int
}

// sayHello() method ---(value receiver)
// this method can be called using person.sayHello() and it is now a method for the struct created above
func (person Person) sayHello() string {
	return "hello , my name is " + person.firstname + " " + person.lastname + " and my age is " + strconv.Itoa(person.age)
}

// pointer receiver method
func (person *Person) increaseAge() {
	person.age++
}
func (person *Person) changeName(name string) {
	person.firstname = name
}

func main() {
	// init person using struct
	// person1 := Person{firstname: "Aditya", lastname: "khare", gender: "male", age: 23}

	// alternative - keeping the order
	person1 := Person{"Aditya", "khare", "male", 23}

	// person1 := Person{"Aditya", "khare", "23", "male"} // not keeping the order will give error
	fmt.Println(person1)
	fmt.Println(person1.firstname)

	person1.age++
	fmt.Println(person1.age)

	// fmt.Println(person1.age++) // this doesn't works
	fmt.Println(person1.age + 1)

	fmt.Println(person1.sayHello())

	person1.increaseAge()
	person1.increaseAge()
	person1.increaseAge()
	person1.increaseAge()
	person1.increaseAge()

	fmt.Println(person1.age)
	person1.changeName("Shree")

	fmt.Println(person1.firstname)
}
