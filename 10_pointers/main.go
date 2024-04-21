package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Println(*b) // *b = 5

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", *b)
	fmt.Printf("%T\n", &a)

	*b = 10 //change val using pointer where is stored (&a)
	fmt.Println(a)
}
