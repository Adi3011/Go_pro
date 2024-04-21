package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "aditya"
	names[1] = "aman"
	names[2] = "gaurisha"

	// names := [3]string{"aditya", "aman", "gaurisha}
	fmt.Println(names)
	fruits := []string{"orange", "apple", "mango"}
	fmt.Println(fruits[0:2])
}
