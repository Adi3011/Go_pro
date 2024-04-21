package main

import "fmt"

func main() {
	top10 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, n := range top10 {
		fmt.Printf("%d --index, %d --id\n", i, n)
	}

	// if not using index , then use underscore
	for _, n := range top10 {
		fmt.Printf("%d\n", n)
	}

	// we can use range with arrays, slices and maps

	email := map[string]string{"aditya": "aditya@gmail.com", "anurag": "anurag@gmail.com", "amit": "amit@gmail.com"}
	for k, v := range email {
		fmt.Printf("%s --key, %s --value\n", k, v)
	}

	for k := range email {
		fmt.Printf("%s --key\n", k)
	}

}
