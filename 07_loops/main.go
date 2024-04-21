package main

import "fmt"

func main() {
	fmt.Println("hello world")
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// or traditional method
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	for i := 3; i <= 30; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Printf("%d %s\n", i, "fizz")
		} else if i%3 != 0 && i%5 == 0 {
			fmt.Printf("%d %s\n", i, "buzz")
		} else if i%5 == 0 && i%3 == 0 {
			fmt.Printf("%d %s\n", i, "fizzbuzz")
		}
	}
}
